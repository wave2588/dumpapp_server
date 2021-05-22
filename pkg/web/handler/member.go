package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/middleware/util"
	"dumpapp_server/pkg/web/render"
)

type MemberHandler struct {
	accountDAO dao2.AccountDAO
}

func NewMemberHandler() *MemberHandler {
	return &MemberHandler{
		accountDAO: impl4.DefaultAccountDAO,
	}
}

func (h *MemberHandler) GetSelf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	account := GetAccountByLoginID(ctx, loginID)

	members := render.NewMemberRender([]int64{account.ID}, 0, render.MemberDefaultRenderFields...).RenderSlice(ctx)

	ticket, err := util2.GenerateRegisterTicket(account.ID)
	util.PanicIf(err)
	middleware.SetTicketCookie(w, r, ticket)

	util.RenderJSON(w, members[0])
}
