package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	http2 "dumpapp_server/pkg/http"
	impl3 "dumpapp_server/pkg/http/impl"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
	impl5 "dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
)

type CertificateHandler struct {
	alterWebCtl             controller2.AlterWebController
	certificateWebCtl       controller2.CertificateWebController
	memberDownloadNumberCtl controller.MemberDownloadController
	certificateV1Controller controller.CertificateController
	certificateV2Controller controller.CertificateController

	certificateServer http2.CertificateServer
	memberDeviceDAO   dao.MemberDeviceDAO
	certificateDAO    dao.CertificateV2DAO
}

func NewCertificateHandler() *CertificateHandler {
	return &CertificateHandler{
		alterWebCtl:             impl5.DefaultAlterWebController,
		certificateWebCtl:       impl5.DefaultCertificateWebController,
		memberDownloadNumberCtl: impl.DefaultMemberDownloadController,
		certificateV1Controller: impl.DefaultCertificateV1Controller,
		certificateV2Controller: impl.DefaultCertificateV2Controller,

		certificateServer: impl3.DefaultCertificateServer,
		memberDeviceDAO:   impl2.DefaultMemberDeviceDAO,
		certificateDAO:    impl2.DefaultCertificateV2DAO,
	}
}

type postCertificate struct {
	UDID string `json:"udid" validate:"required"`
	Type int    `json:"type" validate:"required"`
}

func (p *postCertificate) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.Type < 1 || p.Type > 3 {
		return errors.UnproccessableError("type 类型错误")
	}
	return nil
}

func (h *CertificateHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	args := &postCertificate{}
	util.PanicIf(util.JSONArgs(r, args))

	payCount := cast.ToInt64(30)
	payType := "private"
	switch args.Type {
	case 1: /// 30 售后七天，理论 1 年不掉签
		payCount = 30
		payType = "private"
	case 2: // 60 售后一年，等 1 - 7 天   public
		payCount = 78
		payType = "private"
	case 3: /// 90 售后一年，立即出   public
		payCount = 97
		payType = "private"
	}

	util.PanicIf(h.memberDownloadNumberCtl.CheckPayCount(ctx, loginID, payCount))

	memberDevice, err := h.memberDeviceDAO.GetByMemberIDUdidSafe(ctx, loginID, args.UDID)
	util.PanicIf(err)
	if memberDevice == nil {
		panic(errors.ErrDeviceNotFound)
	}
	if memberDevice.MemberID != loginID {
		panic(errors.ErrCreateCertificateFailV2)
	}

	/// 请求整数接口
	response := h.certificateV2Controller.CreateCer(ctx, args.UDID, payType)
	if response.ErrorMessage != nil {
		/// 创建失败推送
		h.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, *response.ErrorMessage)
		util.PanicIf(errors.ErrCreateCertificateFail)
	}
	if response.BizExt == nil {
		h.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, "response biz_ext is nil")
		util.PanicIf(errors.ErrCreateCertificateFail)
	}

	p12FileData := response.P12Data
	mpFileData := response.MobileProvisionData
	/// p12 文件修改内容
	modifiedP12FileData, err := h.certificateWebCtl.GetModifiedCertificateData(ctx, p12FileData, response.BizExt.OriginalP12Password, response.BizExt.NewP12Password)
	util.PanicIf(err)

	/// 计算证书 md5
	p12FileMd5 := util2.StringMd5(p12FileData)
	mpFileMd5 := util2.StringMd5(mpFileData)

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	cerID := util2.MustGenerateID(ctx)
	util.PanicIf(h.certificateDAO.Insert(ctx, &models.CertificateV2{
		ID:                         cerID,
		DeviceID:                   memberDevice.ID,
		P12FileData:                p12FileData,
		P12FileDataMD5:             p12FileMd5,
		ModifiedP12FileDate:        modifiedP12FileData,
		MobileProvisionFileData:    mpFileData,
		MobileProvisionFileDataMD5: mpFileMd5,
		Source:                     response.Source,
		BizExt:                     response.BizExt.String(),
	}))

	/// 扣除消费的 D 币
	util.PanicIf(h.memberDownloadNumberCtl.DeductPayCount(ctx, loginID, payCount, enum.MemberPayCountUseCertificate))

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	/// 发送消费成功通知
	h.alterWebCtl.SendCreateCertificateSuccessMsg(ctx, loginID, memberDevice.ID, cerID)

	memberMap := render.NewMemberRender([]int64{loginID}, loginID, render.MemberIncludes([]string{"Devices", "PayCount"})).RenderMap(ctx)
	util.RenderJSON(w, memberMap[loginID])
}

type downloadP12FileArgs struct {
	CerID string `form:"cer_id" validate:"required"`
}

func (args *downloadP12FileArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *CertificateHandler) DownloadP12File(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := downloadP12FileArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())
	cerID := cast.ToInt64(args.CerID)

	cer, err := h.certificateDAO.Get(ctx, cerID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if cer == nil {
		panic(errors.ErrNotFound)
	}
	uDec, err := base64.StdEncoding.DecodeString(cer.ModifiedP12FileDate)
	util.PanicIf(err)
	w.Header().Add("Content-Disposition", `attachment;filename="developer.p12`)
	w.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(int64(len(uDec)), 10))
	w.Write(uDec)
}

type downloadMobileprovisionFileArgs struct {
	CerID string `form:"cer_id" validate:"required"`
}

func (args *downloadMobileprovisionFileArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *CertificateHandler) DownloadMobileprovisionFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := downloadMobileprovisionFileArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())
	cerID := cast.ToInt64(args.CerID)

	cer, err := h.certificateDAO.Get(ctx, cerID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if cer == nil {
		panic(errors.ErrNotFound)
	}
	uDec, err := base64.StdEncoding.DecodeString(cer.MobileProvisionFileData)
	util.PanicIf(err)
	w.Header().Add("Content-Disposition", `attachment;filename="developer.mobileprovision"`)
	w.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(int64(len(uDec)), 10))
	w.Write(uDec)
}
