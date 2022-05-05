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
	certificateServer       http2.CertificateServer
	memberDeviceDAO         dao.MemberDeviceDAO
	certificateDAO          dao.CertificateDAO
	certificateDeviceDAO    dao.CertificateDeviceDAO
}

func NewCertificateHandler() *CertificateHandler {
	return &CertificateHandler{
		alterWebCtl:             impl5.DefaultAlterWebController,
		certificateWebCtl:       impl5.DefaultCertificateWebController,
		memberDownloadNumberCtl: impl.DefaultMemberDownloadController,
		certificateServer:       impl3.DefaultCertificateServer,
		memberDeviceDAO:         impl2.DefaultMemberDeviceDAO,
		certificateDAO:          impl2.DefaultCertificateDAO,
		certificateDeviceDAO:    impl2.DefaultCertificateDeviceDAO,
	}
}

type postCertificate struct {
	UDID string `json:"udid" validate:"required"`
	Type int    `json:"type"`
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
		payType = "public"
	case 2: // 60 售后一年，等 1 - 7 天
		payCount = 60
		payType = "private"
	case 3: /// 90 售后一年，立即出
		payCount = 90
		payType = "private"
	}

	util.PanicIf(h.memberDownloadNumberCtl.CheckPayCount(ctx, loginID, payCount))

	memberDevice, err := h.memberDeviceDAO.GetByUdid(ctx, args.UDID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if memberDevice == nil {
		panic(errors.ErrNotFound)
	}

	/// 请求整数接口
	result, err := h.certificateServer.CreateCer(ctx, args.UDID, payType)
	util.PanicIf(err)

	if result.Data == nil || result.IsSuccess == false {
		/// 创建失败推送
		h.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, result.ErrorMessage)
		util.PanicIf(errors.ErrCreateCertificateFail)
	}
	cerData := result.Data

	p12FileData := cerData.P12FileDate
	mpFileData := cerData.MobileProvisionFileData
	/// p12 文件修改内容
	modifiedP12FileData, err := h.certificateWebCtl.GetModifiedCertificateData(ctx, cerData.P12FileDate)
	util.PanicIf(err)

	/// 查看证书是否已经存在, p12 文件还是按照元数据计算
	p12FileMd5 := util2.StringMd5(p12FileData)
	mpFileMd5 := util2.StringMd5(mpFileData)
	cer, err := h.certificateDAO.GetByP12FileDateMD5MobileProvisionFileDataMD5(ctx, p12FileMd5, mpFileMd5)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		panic(err)
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	if cer == nil {
		/// 创建证书
		cerID := util2.MustGenerateID(ctx)
		util.PanicIf(h.certificateDAO.Insert(ctx, &models.Certificate{
			ID:                         cerID,
			P12FileDate:                p12FileData,
			P12FileDateMD5:             p12FileMd5,
			ModifiedP12FileDate:        modifiedP12FileData,
			MobileProvisionFileData:    mpFileData,
			MobileProvisionFileDataMD5: mpFileMd5,
			UdidBatchNo:                cerData.UdidBatchNo,
			CerAppleid:                 cerData.CerAppleid,
		}))
		/// 绑定证书和设备关系
		util.PanicIf(h.certificateDeviceDAO.Insert(ctx, &models.CertificateDevice{
			DeviceID:      memberDevice.ID,
			CertificateID: cerID,
		}))
		/// 消费 6 次, 这是因为完全新创建, 所以进行消费
		util.PanicIf(h.memberDownloadNumberCtl.DeductPayCount(ctx, loginID, payCount, enum.MemberPayCountUseCertificate))

		h.alterWebCtl.SendCreateCertificateSuccessMsg(ctx, loginID, memberDevice.ID, cerID)
	} else {
		/// 发现设备和此证书没绑定过, 则进行绑定
		mc, err := h.certificateDeviceDAO.GetByDeviceIDCertificateID(ctx, memberDevice.ID, cer.ID)
		if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
			util.PanicIf(err)
		}
		if mc == nil {
			util.PanicIf(h.certificateDeviceDAO.Insert(ctx, &models.CertificateDevice{
				DeviceID:      memberDevice.ID,
				CertificateID: cer.ID,
			}))
			/// 消费 6 次, 这是因为有证书了, 但是没绑定, 所以进行消费
			util.PanicIf(h.memberDownloadNumberCtl.DeductPayCount(ctx, loginID, payCount, enum.MemberPayCountUseCertificate))
			h.alterWebCtl.SendCreateCertificateSuccessMsg(ctx, loginID, memberDevice.ID, cer.ID)
		}
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	memberMap := render.NewMemberRender([]int64{loginID}, loginID, render.MemberDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, memberMap[loginID])
}

type downloadP12FileArgs struct {
	DeviceID string `form:"device_id" validate:"required"`
	CerID    string `form:"cer_id" validate:"required"`
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
	DeviceID string `form:"device_id" validate:"required"`
	CerID    string `form:"cer_id" validate:"required"`
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
