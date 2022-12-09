package open_api_handler

import (
	"net/http"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
)

type OpenMemberHandler struct{}

func NewOpenMemberHandler() *OpenMemberHandler {
	return &OpenMemberHandler{}
}

type Member struct {
	ID        int64            `json:"id,string"`
	Status    string           `json:"status"`
	Role      enum.AccountRole `json:"role"`
	PayCount  *int64           `json:"pay_count,omitempty"`
	CreatedAt int64            `json:"created_at"`
	UpdatedAt int64            `json:"updated_at"`
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
	util.RenderJSON(w, &Member{
		ID:        member.ID,
		Status:    member.Status,
		Role:      member.Role,
		PayCount:  member.PayCount,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
	})
}
