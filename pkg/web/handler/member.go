package handler

import (
	"fmt"
	"net"
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

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

func (h *MemberHandler) GetSelf(w http.ResponseWriter, r *http.Request) {
	fmt.Println(111, r.Header)
	fmt.Println(111, r.Header.Get("X-FORWARDED-FOR"))

	fmt.Println(RemoteIp(r))

	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	account := GetAccountByLoginID(ctx, loginID)

	members := render.NewMemberRender([]int64{account.ID}, 0, render.MemberDefaultRenderFields...).RenderSlice(ctx)

	ticket, err := util2.GenerateRegisterTicket(account.ID)
	util.PanicIf(err)
	middleware.SetTicketCookie(w, r, ticket)

	util.RenderJSON(w, members[0])
}
