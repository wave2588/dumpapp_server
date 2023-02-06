package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
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
	certificateCreateCtl controller2.CertificateCreateWebController
	certificateWebCtl    controller2.CertificateWebController
	certificateV2WebCtl  controller2.CertificateV2WebController
	certificatePriceCtl  controller.CertificatePriceController

	certificateDAO dao.CertificateV2DAO
}

func NewCertificateHandler() *CertificateHandler {
	return &CertificateHandler{
		certificateCreateCtl: impl5.DefaultCertificateCreateWebController,
		certificateWebCtl:    impl5.DefaultCertificateWebController,
		certificateV2WebCtl:  impl5.DefaultCertificateV2WebController,
		certificatePriceCtl:  impl.DefaultCertificatePriceController,

		certificateDAO: impl2.DefaultCertificateV2DAO,
	}
}

type postCertificate struct {
	UDID string `json:"udid" validate:"required"`
	Type int    `json:"type" validate:"required"` // 就是 price_id
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
	if !util2.CheckUDIDValid(p.UDID) {
		return errors.UnproccessableError(fmt.Sprintf("无效的 UDID: %s", p.UDID))
	}
	return nil
}

func (h *CertificateHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	args := &postCertificate{}
	util.PanicIf(util.JSONArgs(r, args))

	payType := "private"
	cerID, err := h.certificateCreateCtl.Create(ctx, loginID, args.UDID, args.Note, cast.ToInt64(args.Type), false, payType)
	util.PanicIf(err)

	cerMap := render.NewCertificateRender([]int64{cerID}, loginID, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, cerMap[cerID])
}

func (h *CertificateHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)
	prices, err := h.certificatePriceCtl.GetPrices(ctx, loginID)
	util.PanicIf(err)
	util.RenderJSON(w, map[string]interface{}{
		"data":  prices,
		"title": "网站所有证书已开启推送权限，证书掉签后非代理用户 7 天(代理用户 30 天)内使用掉签 UDID 进行购买证书不会消耗 D 币。",
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
