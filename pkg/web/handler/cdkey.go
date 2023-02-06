package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/web/controller"
	"dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render/install_app_render"
)

type CdkeyHandler struct {
	cdkeyWebCtl controller.CdkeyWebController
}

func NewCdkeyHandler() *CdkeyHandler {
	return &CdkeyHandler{
		cdkeyWebCtl: impl.DefaultCdkeyWebController,
	}
}

func (h *CdkeyHandler) Post(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
	)
	cdkeyID, err := h.cdkeyWebCtl.Create(ctx, loginID)
	util.PanicIf(err)

	cdkeyMap := install_app_render.NewCDKEYRender([]int64{cdkeyID}, loginID, install_app_render.CDKeyDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, cdkeyMap[cdkeyID])
}
