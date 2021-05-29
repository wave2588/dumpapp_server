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
			"type":     enum.MemberVipDurationTypeOneMonth,
			"title":    "月度会员",
			"price":    constant.MemberVipDurationTypeToPrice[enum.MemberVipDurationTypeOneMonth],
			"describe": fmt.Sprintf("月度会员每月 %d 元。", constant.MemberVipDurationTypeToPrice[enum.MemberVipDurationTypeOneMonth]),
		},
		{
			"type":     enum.MemberVipDurationTypeThreeMonth,
			"title":    "季度会员",
			"price":    constant.MemberVipDurationTypeToPrice[enum.MemberVipDurationTypeThreeMonth],
			"describe": fmt.Sprintf("季度会员享受 85 折优惠，每月仅需 57 元。"),
		},
		{
			"type":     enum.MemberVipDurationTypeSixMonth,
			"title":    "年度会员",
			"price":    constant.MemberVipDurationTypeToPrice[enum.MemberVipDurationTypeSixMonth],
			"describe": fmt.Sprintf("年度会员享受 7 折优惠，每月仅需 47 元。"),
		},
	}
	util.RenderJSON(w, res)
}
