package handler

import (
	"fmt"
	"net"
	"net/http"

	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/middleware/util"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
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

type getSelfDeviceArgs struct {
	IsFilterEmptyCer int `form:"is_filter_empty_cer"`
}

func (p *getSelfDeviceArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *MemberHandler) GetSelfDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getSelfDeviceArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	loginID := mustGetLoginID(ctx)

	account := GetAccountByLoginID(ctx, loginID)

	members := render.NewMemberRender([]int64{account.ID}, loginID, render.MemberIncludes([]string{"Devices", "PayCount"})).RenderSlice(ctx)

	/// 过滤掉没有证书的设备信息
	if cast.ToBool(args.IsFilterEmptyCer) {
		for _, member := range members {
			devices := member.Devices
			if len(devices) == 0 {
				continue
			}
			resDevice := make([]*render.Device, 0)
			for _, device := range devices {
				if len(device.Certificates) != 0 {
					resDevice = append(resDevice, device)
				}
			}
			member.Devices = resDevice
		}
	}

	util.RenderJSON(w, members[0])
}
