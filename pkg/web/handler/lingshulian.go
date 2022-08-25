package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	util3 "dumpapp_server/pkg/util"
)

type LingshulianHandler struct {
	lingshulianCtl controller.LingshulianController
}

func NewLingshulianHandler() *LingshulianHandler {
	return &LingshulianHandler{
		lingshulianCtl: impl.DefaultLingshulianController,
	}
}

func (h *LingshulianHandler) GetSignIpaPutURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := util3.MustGenerateID(ctx)
	key := fmt.Sprintf("%d.ipa", id)
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

func (h *LingshulianHandler) Get(w http.ResponseWriter, r *http.Request) {
}
