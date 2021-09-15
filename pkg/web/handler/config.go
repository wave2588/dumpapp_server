package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/web/render"
)

type ConfigHandler struct {
	configDAO dao.AdminConfigDAO
}

func NewConfigHandler() *ConfigHandler {
	return &ConfigHandler{
		configDAO: impl.DefaultAdminConfigDAO,
	}
}

func (h *ConfigHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	util.RenderJSON(w, render.NewConfigRender(0).Render(ctx))
}
