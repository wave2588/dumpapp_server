package handler

import (
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
	accountDAO    dao2.AccountDAO
	statisticsDAO dao2.StatisticsDAO
}

func NewMemberHandler() *MemberHandler {
	return &MemberHandler{
		accountDAO:    impl4.DefaultAccountDAO,
		statisticsDAO: impl4.DefaultStatisticsDAO,
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
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	account := GetAccountByLoginID(ctx, loginID)

	members := render.NewMemberRender([]int64{account.ID}, loginID, render.MemberDefaultRenderFields...).RenderSlice(ctx)

	ticket, err := util2.GenerateRegisterTicket(account.ID)
	util.PanicIf(err)
	middleware.SetTicketCookie(w, r, ticket)

	_ = h.statisticsDAO.AddStatistics(ctx, loginID)

	util.RenderJSON(w, members[0])
}
