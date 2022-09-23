package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
)

type AdminConfigHandler struct {
	adminConfigDAO dao.AdminConfigInfoDAO
}

func NewAdminConfigHandler() *AdminConfigHandler {
	return &AdminConfigHandler{
		adminConfigDAO: impl.DefaultAdminConfigInfoDAO,
	}
}

type postConfigArgs struct {
	AdminBusy      *bool                   `json:"admin_busy"`
	DailyFreeCount *int64                  `json:"daily_free_count"`
	CerSource      *enum.CertificateSource `json:"cer_source"`
}

func (p *postConfigArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminConfigHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	args := &postConfigArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	config, err := h.adminConfigDAO.GetConfig(ctx)
	util.PanicIf(err)

	if args.AdminBusy != nil {
		config.BizExt.AdminBusy = *args.AdminBusy
	}
	if args.DailyFreeCount != nil {
		config.BizExt.DailyFreeCount = *args.DailyFreeCount
	}
	if args.CerSource != nil {
		config.BizExt.CerSource = *args.CerSource
	}

	util.PanicIf(h.adminConfigDAO.Update(ctx, config))

	util.RenderJSON(w, render.NewConfigRender(loginID).Render(ctx))
}

func (h *AdminConfigHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	util.RenderJSON(w, render.NewConfigRender(loginID).Render(ctx))
}
