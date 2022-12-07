package open_api_handler

import (
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"net/http"
)

type OpenMemberHandler struct {
}

func NewOpenMemberHandler() *OpenMemberHandler {
	return &OpenMemberHandler{}
}

func (h *OpenMemberHandler) GetMember(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx, r)
	)

	memberMap := render.NewMemberRender([]int64{loginID}, loginID, render.MemberDefaultRenderFields...).RenderMap(ctx)
	member, ok := memberMap[loginID]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
	}
	util.RenderJSON(w, member)
}
