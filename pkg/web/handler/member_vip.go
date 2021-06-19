package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type MemberVipHandler struct {
	ipaDAO        dao.IpaDAO
	ipaVersionDAO dao.IpaVersionDAO

	alipayCtl controller.ALiPayController
}

func NewMemberVipHandler() *MemberVipHandler {
	return &MemberVipHandler{
		ipaDAO:        impl.DefaultIpaDAO,
		ipaVersionDAO: impl.DefaultIpaVersionDAO,

		alipayCtl: impl2.DefaultALiPayController,
	}
}

type postMemberVipArgs struct {
	Type enum.MemberVipDurationType `json:"type" validate:"required"`
}

func (p *postMemberVipArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *MemberVipHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)
	_ = GetAccountByLoginID(ctx, loginID)

	args := &postMemberVipArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	url, err := h.alipayCtl.GetPayURL(ctx, loginID, args.Type)
	util.PanicIf(err)

	res := map[string]interface{}{
		"open_url": url,
	}
	util.RenderJSON(w, res)
}

func (h *MemberVipHandler) Get(w http.ResponseWriter, r *http.Request) {
	res := []map[string]interface{}{
		{
			"type":     enum.MemberVipDurationTypeOne,
			"title":    "10 天会员",
			"price":    constant.MemberVipDurationTypeToPrice[enum.MemberVipDurationTypeOne],
			"describe": fmt.Sprintf("10 天会员, 平均每天 %.2f 元", 18/10.0),
		},
		{
			"type":     enum.MemberVipDurationTypeTwo,
			"title":    "20 天会员",
			"price":    constant.MemberVipDurationTypeToPrice[enum.MemberVipDurationTypeTwo],
			"describe": fmt.Sprintf("20 天会员, 平均每天 %.2f 元", 28/20.0),
		},
		{
			"type":     enum.MemberVipDurationTypeThree,
			"title":    "30 天会员",
			"price":    constant.MemberVipDurationTypeToPrice[enum.MemberVipDurationTypeThree],
			"describe": fmt.Sprintf("30 天会员, 平均每天 %.2f 元", 38/30.0),
		},
	}
	util.RenderJSON(w, res)
}
