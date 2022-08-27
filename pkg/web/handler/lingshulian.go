package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/errors"
	util3 "dumpapp_server/pkg/util"
	"github.com/go-playground/validator/v10"
)

type LingshulianHandler struct {
	lingshulianCtl controller.LingshulianController
}

func NewLingshulianHandler() *LingshulianHandler {
	return &LingshulianHandler{
		lingshulianCtl: impl.DefaultLingshulianController,
	}
}

type postPutURLArgs struct {
	Suffix string `json:"suffix" validate:"required"`
}

func (p *postPutURLArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.Suffix == "" {
		return errors.UnproccessableError("Suffix 格式错误")
	}
	return nil
}

func (h *LingshulianHandler) PostUploadInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &postPutURLArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	id := util3.MustGenerateID(ctx)
	key := fmt.Sprintf("%d.%s", id, args.Suffix)
	resp, err := h.lingshulianCtl.GetPutURL(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, key)
	util.PanicIf(err)
	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) GetTempSecretKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp, err := h.lingshulianCtl.GetTempSecretKey(ctx)
	util.PanicIf(err)
	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) PostMultipartUploadInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &controller.PostCreateMultipartUploadInfoRequest{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.PostCreateMultipartUploadInfo(ctx, args)
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) PostMultipartUploadPartInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &controller.PostMultipartUploadPartInfoRequest{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.PostMultipartUploadPartInfo(ctx, args)
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) PostCompleteMultipartUploadInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &controller.PostCompleteMultipartUploadInfoRequest{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.PostCompleteMultipartUploadInfo(ctx, args)
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}
