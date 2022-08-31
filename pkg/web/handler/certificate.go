package handler

import (
	"dumpapp_server/pkg/common/constant"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

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

	cerID, err := h.certificateWebCtl.PayCertificate(ctx, loginID, args.UDID, payCount, payType)
	util.PanicIf(err)

	cerMap := render.NewCertificateRender([]int64{cerID}, loginID, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, cerMap[cerID])
}

type certificatePrice struct {
	ID          int    `json:"id"`
	Price       int    `json:"price"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *CertificateHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	data := []*certificatePrice{
		{
			ID:          1,
			Price:       constant.CertificatePriceL1,
			Title:       "普通版",
			Description: "理论 1 年，无质保。",
		},
		{
			ID:          2,
			Price:       constant.CertificatePriceL2,
			Title:       "稳定版",
			Description: "理论 1 年，售后半年，掉了无限补。",
		},
		{
			ID:          3,
			Price:       constant.CertificatePriceL3,
			Title:       "豪华版",
			Description: "理论 1 年，售后 1 年，掉了无限补。",
		},
	}
	util.RenderJSON(w, util.ListOutput{
		Data: data,
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
