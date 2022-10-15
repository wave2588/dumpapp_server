package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"dumpapp_server/pkg/common/constant"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	controller2 "dumpapp_server/pkg/web/controller"
	impl5 "dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
)

type CertificateHandler struct {
	certificateWebCtl controller2.CertificateWebController

	certificateDAO dao.CertificateV2DAO
}

func NewCertificateHandler() *CertificateHandler {
	return &CertificateHandler{
		certificateWebCtl: impl5.DefaultCertificateWebController,

		certificateDAO: impl2.DefaultCertificateV2DAO,
	}
}

type postCertificate struct {
	UDID string `json:"udid" validate:"required"`
	Type int    `json:"type" validate:"required"`
	Note string `json:"note"`
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

	payCount := cast.ToInt64(constant.CertificatePriceL1)
	payType := "private"
	switch args.Type {
	case 1: /// 30 售后七天，理论 1 年不掉签
		payCount = constant.CertificatePriceL1
		payType = "private"
	case 2: // 60 售后一年，等 1 - 7 天   public
		payCount = constant.CertificatePriceL2
		payType = "private"
	case 3: /// 90 售后一年，立即出   public
		payCount = constant.CertificatePriceL3
		payType = "private"
	}

	cerID, err := h.certificateWebCtl.PayCertificate(ctx, loginID, args.UDID, args.Note, payCount, payType)
	util.PanicIf(err)

	cerMap := render.NewCertificateRender([]int64{cerID}, loginID, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, cerMap[cerID])
}

func (h *CertificateHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	util.RenderJSON(w, util.ListOutput{
		Data: constant.GetCertificatePrices(),
	})
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

	p12Data := cer.ModifiedP12FileDate
	if p12Data == "" {
		p12Data = cer.P12FileData
	}
	uDec, err := base64.StdEncoding.DecodeString(p12Data)
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
