package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

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
	rpc "dumpapp_server/pkg/ice"
	impl4 "dumpapp_server/pkg/ice/impl"
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
	memberDownloadNumberCtl controller.MemberDownloadController
	certificateServer       http2.CertificateServer
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO
	memberDeviceDAO         dao.MemberDeviceDAO
	certificateDAO          dao.CertificateDAO
	certificateDeviceDAO    dao.CertificateDeviceDAO
	iceRPC                  rpc.IceRPC
}

func NewCertificateHandler() *CertificateHandler {
	return &CertificateHandler{
		alterWebCtl:             impl5.DefaultAlterWebController,
		memberDownloadNumberCtl: impl.DefaultMemberDownloadController,
		certificateServer:       impl3.DefaultCertificateServer,
		memberDownloadNumberDAO: impl2.DefaultMemberDownloadNumberDAO,
		memberDeviceDAO:         impl2.DefaultMemberDeviceDAO,
		certificateDAO:          impl2.DefaultCertificateDAO,
		certificateDeviceDAO:    impl2.DefaultCertificateDeviceDAO,
		iceRPC:                  impl4.DefaultIceRPC,
	}
}

type postCertificate struct {
	UDID string `json:"udid" validate:"required"`
}

func (p *postCertificate) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *CertificateHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	dns, err := h.memberDownloadNumberCtl.GetCertificateDownloadNumbers(ctx, loginID)
	util.PanicIf(err)

	args := &postCertificate{}
	util.PanicIf(util.JSONArgs(r, args))

	memberDevice, err := h.memberDeviceDAO.GetByUdid(ctx, args.UDID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if memberDevice == nil {
		panic(errors.ErrNotFound)
	}

	/// 请求整数接口
	result, err := h.certificateServer.CreateCer(ctx, args.UDID)
	util.PanicIf(err)

	if result.Data == nil || result.IsSuccess == false {
		/// 创建失败推送
		h.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, result.ErrorMessage)
		util.PanicIf(errors.ErrCreateCertificateFail)
	}
	cerData := result.Data

	/// 查看证书是否已经存在
	p12FileMd5 := util2.StringMd5(cerData.P12FileDate)
	mpFileMd5 := util2.StringMd5(cerData.MobileProvisionFileData)
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
		cerID := h.iceRPC.MustGenerateID(ctx)
		util.PanicIf(h.certificateDAO.Insert(ctx, &models.Certificate{
			ID:                         cerID,
			P12FileDate:                cerData.P12FileDate,
			P12FileDateMD5:             p12FileMd5,
			MobileProvisionFileData:    cerData.MobileProvisionFileData,
			MobileProvisionFileDataMD5: mpFileMd5,
			UdidBatchNo:                cerData.UdidBatchNo,
			CerAppleid:                 cerData.CerAppleid,
		}))
		/// 绑定证书和设备关系
		util.PanicIf(h.certificateDeviceDAO.Insert(ctx, &models.CertificateDevice{
			DeviceID:      memberDevice.ID,
			CertificateID: cerID,
		}))
		/// 消费 5 次, 这是因为完全新创建, 所以进行消费
		for _, dn := range dns {
			dn.Status = enum.MemberDownloadNumberStatusUsed
			util.PanicIf(h.memberDownloadNumberDAO.Update(ctx, dn))
		}
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
			/// 消费 5 次, 这是因为有证书了, 但是没绑定, 所以进行消费
			for _, dn := range dns {
				dn.Status = enum.MemberDownloadNumberStatusUsed
				util.PanicIf(h.memberDownloadNumberDAO.Update(ctx, dn))
			}
		}
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	memberMap := render.NewMemberRender([]int64{loginID}, loginID, render.MemberDefaultRenderFields...).RenderSlice(ctx)
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
	uDec, err := base64.StdEncoding.DecodeString(cer.P12FileDate)
	util.PanicIf(err)
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment;filename="%d.p12"`, cer.ID))
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
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment;filename="%d.mobileprovision"`, cer.ID))
	w.Write(uDec)
}
