package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
)

type AdminMemberVipHandler struct {
	accountDAO   dao.AccountDAO
	memberVipDAO dao.MemberVipDAO
}

func NewAdminMemberVipHandler() *AdminMemberVipHandler {
	return &AdminMemberVipHandler{
		accountDAO:   impl.DefaultAccountDAO,
		memberVipDAO: impl.DefaultMemberVipDAO,
	}
}

type addDurationArgs struct {
	Email    string `json:"email" validate:"required"`
	Duration int    `json:"duration" validate:"required"`
}

func (p *addDurationArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminMemberVipHandler) AddDuration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &addDurationArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util.PanicIf(err)

	memberMap := render.NewMemberRender([]int64{account.ID}, loginID, render.MemberDefaultRenderFields...).RenderMap(ctx)
	member := memberMap[account.ID]
	if !member.Vip.IsVip {
		panic(errors.ErrUpgradeVip)
	}

	memberVip, err := h.memberVipDAO.Get(ctx, account.ID)
	util.PanicIf(err)

	memberVip.EndAt = memberVip.EndAt.AddDate(0, 0, args.Duration)
	util.PanicIf(h.memberVipDAO.Update(ctx, memberVip))

	util.RenderJSON(w, "ok")
}

type deleteMemberVipArgs struct {
	Email string `json:"email" validate:"required"`
}

func (p *deleteMemberVipArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminMemberVipHandler) DeleteMemberVip(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &deleteMemberVipArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util.PanicIf(err)

	memberVip, err := h.memberVipDAO.Get(ctx, account.ID)
	util.PanicIf(err)

	util.PanicIf(h.memberVipDAO.Delete(ctx, memberVip.ID))

	util.RenderJSON(w, "ok")
}
