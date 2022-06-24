package handler

import (
	"fmt"
	"net"
	"net/http"

	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/middleware/util"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberHandler struct {
	accountDAO      dao2.AccountDAO
	statisticsDAO   dao2.StatisticsDAO
	memberDeviceDAO dao2.MemberDeviceDAO
	certificateDAO  dao2.CertificateV2DAO
}

func NewMemberHandler() *MemberHandler {
	return &MemberHandler{
		accountDAO:      impl4.DefaultAccountDAO,
		statisticsDAO:   impl4.DefaultStatisticsDAO,
		memberDeviceDAO: impl4.DefaultMemberDeviceDAO,
		certificateDAO:  impl4.DefaultCertificateV2DAO,
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

func (h *MemberHandler) GetSelfCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	deviceMap, err := h.memberDeviceDAO.BatchGetByMemberIDs(ctx, []int64{loginID})
	util.PanicIf(err)

	deviceIDs := make([]int64, 0)
	for _, devices := range deviceMap {
		for _, memberDevice := range devices {
			deviceIDs = append(deviceIDs, memberDevice.ID)
		}
	}

	filters := []qm.QueryMod{
		models.CertificateV2Where.DeviceID.IN(deviceIDs),
	}
	ids, err := h.certificateDAO.ListIDs(ctx, offset, limit, filters, nil)
	util.PanicIf(err)
	count, err := h.certificateDAO.Count(ctx, filters)
	util.PanicIf(err)

	data := render.NewCertificateRender(ids, loginID, render.CertificateDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   data,
	})
}
