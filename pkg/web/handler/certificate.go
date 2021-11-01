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
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
)

type CertificateHandler struct {
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

	md, err := h.memberDeviceDAO.GetByUdid(ctx, args.UDID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if md == nil {
		panic(errors.ErrNotFound)
	}

	/// 请求整数接口
	result, err := h.certificateServer.CreateCer(ctx, args.UDID)
	util.PanicIf(err)
	if result.Data == nil {
		util.PanicIf(errors.ErrHttpFail)
	}
	cerData := result.Data

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	/// 创建证书
	cerID := h.iceRPC.MustGenerateID(ctx)
	util.PanicIf(h.certificateDAO.Insert(ctx, &models.Certificate{
		ID:                      cerID,
		P12FileDate:             cerData.P12FileDate,
		MobileProvisionFileData: cerData.MobileProvisionFileData,
		UdidBatchNo:             cerData.UdidBatchNo,
		CerAppleid:              cerData.CerAppleid,
	}))
	/// 绑定证书和设备关系
	util.PanicIf(h.certificateDeviceDAO.Insert(ctx, &models.CertificateDevice{
		DeviceID:      md.ID,
		CertificateID: cerID,
	}))
	/// 消费 5 次
	for _, dn := range dns {
		dn.Status = enum.MemberDownloadNumberStatusUsed
		util.PanicIf(h.memberDownloadNumberDAO.Update(ctx, dn))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)
}

type downloadP12FileArgs struct {
	DeviceID string `form:"cer_id" validate:"required"`
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
	DeviceID string `form:"cer_id" validate:"required"`
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
	uDec, err := base64.StdEncoding.DecodeString(cer.MobileProvisionFileData)
	util.PanicIf(err)
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment;filename="%d.p12"`, cer.ID))
	w.Write(uDec)
}
