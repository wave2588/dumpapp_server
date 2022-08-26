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

type getUploadInfoArgs struct {
	Suffix string `json:"suffix" validate:"required"`
}

func (p *getUploadInfoArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.Suffix == "" {
		return errors.UnproccessableError("Suffix 格式错误")
	}
	return nil
}

func (h *LingshulianHandler) PostMultipartUploadInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &getUploadInfoArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.GetCreateMultipartUploadInfo(ctx, args.Suffix)
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

type getUploadPartInfoArgs struct {
	UploadID   string `json:"upload_id" validate:"required"`
	Key        string `json:"key" validate:"required"`
	Bucket     string `json:"bucket" validate:"required"`
	PartNumber int64  `json:"part_number" validate:"required"`
}

func (args *getUploadPartInfoArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if args.UploadID == "" || args.Key == "" || args.Bucket == "" {
		return errors.UnproccessableError("参数格式错误")
	}
	return nil
}

func (h *LingshulianHandler) PostMultipartUploadPartInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &getUploadPartInfoArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.GetMultipartUploadPartInfo(ctx, args.UploadID, args.Key, args.Bucket, args.PartNumber)
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

type getCompleteMultipartUploadInfoArgs struct {
	UploadID string `json:"upload_id" validate:"required"`
	Key      string `json:"key" validate:"required"`
	Bucket   string `json:"bucket" validate:"required"`
}

func (p *getCompleteMultipartUploadInfoArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.UploadID == "" || p.Key == "" || p.Bucket == "" {
		return errors.UnproccessableError("参数格式错误")
	}
	return nil
}

func (h *LingshulianHandler) PostCompleteMultipartUploadInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &getCompleteMultipartUploadInfoArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.GetCompleteMultipartUploadInfo(ctx, args.UploadID, args.Key, args.Bucket)
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}
