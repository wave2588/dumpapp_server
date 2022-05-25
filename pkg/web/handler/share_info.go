package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/web/render"
)

type ShareInfoHandler struct {
	memberRebateRecordDAO dao.MemberRebateRecordDAO
	memberPayOrderDAO     dao.MemberPayOrderDAO
}

func NewShareInfoHandler() *ShareInfoHandler {
	return &ShareInfoHandler{
		memberRebateRecordDAO: impl.DefaultMemberRebateRecordDAO,
		memberPayOrderDAO:     impl.DefaultMemberPayOrderDAO,
	}
}

func (h *ShareInfoHandler) Get(w http.ResponseWriter, r *http.Request) {
	util.RenderJSON(w, render.MustRenderShareInfo())
}
