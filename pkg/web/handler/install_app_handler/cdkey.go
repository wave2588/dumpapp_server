package install_app_handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render/install_app_render"
)

type CDKEYHandler struct {
	installAppCDKEYDAO dao2.InstallAppCdkeyDAO
}

func NewCDKEYHandler() *CDKEYHandler {
	return &CDKEYHandler{
		installAppCDKEYDAO: impl2.DefaultInstallAppCdkeyDAO,
	}
}

func (h *CDKEYHandler) GetCDKEYInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outID := util.URLParam(r, "out_id")

	res, err := h.installAppCDKEYDAO.BatchGetByOutID(ctx, []string{outID})
	util.PanicIf(err)

	cdkey, ok := res[outID]
	if !ok {
		util.PanicIf(errors.ErrInstallAppGenerateCDKeyNotFound)
	}
	data := install_app_render.NewCDKEYRender([]int64{cdkey.ID}, 0).RenderMap(ctx)

	util.RenderJSON(w, data[cdkey.ID])
}
