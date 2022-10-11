package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/errors"
	util3 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type LingshulianHandler struct {
	lingshulianCtl controller.LingshulianController
	alterWebCtl    controller2.AlterWebController
}

func NewLingshulianHandler() *LingshulianHandler {
	return &LingshulianHandler{
		lingshulianCtl: impl.DefaultLingshulianController,
		alterWebCtl:    impl2.DefaultAlterWebController,
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

	loginID := mustGetLoginID(ctx)

	args := &controller.PostCreateMultipartUploadInfoRequest{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.PostCreateMultipartUploadInfo(ctx, args)
	if err != nil {
		h.sendMsg(ctx, "获取上传信息失败", loginID, args, err.Error())
	}
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) PostMultipartUploadPartInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &controller.PostMultipartUploadPartInfoRequest{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.PostMultipartUploadPartInfo(ctx, args)
	if err != nil {
		h.sendMsg(ctx, "获取上传分片信息失败", loginID, args, err.Error())
	}
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) PostCompleteMultipartUploadInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &controller.PostCompleteMultipartUploadInfoRequest{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.PostCompleteMultipartUploadInfo(ctx, args)
	if err != nil {
		h.sendMsg(ctx, "合并文件失败", loginID, args, err.Error())
	}
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) PostAbortMultipartUploadInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &controller.PostAbortMultipartUploadPartInfoRequest{}
	util.PanicIf(util.JSONArgs(r, args))

	resp, err := h.lingshulianCtl.PostAbortMultipartUploadInfo(ctx, args)
	if err != nil {
		h.sendMsg(ctx, "取消上传失败", loginID, args, err.Error())
	}
	util.PanicIf(err)

	util.RenderJSON(w, resp)
}

func (h *LingshulianHandler) sendMsg(ctx context.Context, title string, loginID int64, requestBody interface{}, errMsg string) {
	appVersion, _ := ctx.Value(constant.CtxKeyAppVersion).(string)
	jsonData, _ := json.Marshal(requestBody)
	titleString := fmt.Sprintf("<font color=\"warning\">%s</font>\n>", title)
	loginString := fmt.Sprintf("用户 ID：<font color=\"comment\">%d</font>\n", loginID)
	jsonString := fmt.Sprintf("request body：<font color=\"comment\">%s</font>\n", string(jsonData))
	errString := fmt.Sprintf("错误信息：<font color=\"comment\">%s</font>\n", errMsg)
	appVersionString := fmt.Sprintf("版本：<font color=\"comment\">%s</font>\n", appVersion)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	h.alterWebCtl.SendCustomMsg(ctx, "16a2bd1b-a03a-4a46-bbec-f218cbcfe17d", titleString+loginString+jsonString+errString+appVersionString+timeStr)
}
