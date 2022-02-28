package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
)

type AdminConfigHandler struct {
	configDAO dao.AdminConfigDAO
}

func NewAdminConfigHandler() *AdminConfigHandler {
	return &AdminConfigHandler{
		configDAO: impl.DefaultAdminConfigDAO,
	}
}

type postConfigArgs struct {
	AdminBusy      *bool  `json:"admin_busy"`
	DailyFreeCount *int64 `json:"daily_free_count"`
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

	if args.AdminBusy != nil {
		util.PanicIf(h.configDAO.SetAdminBusy(ctx, *args.AdminBusy))
	}
	if args.DailyFreeCount != nil {
		util.PanicIf(h.configDAO.SetDailyFreeCount(ctx, *args.DailyFreeCount))
	}

	util.RenderJSON(w, render.NewConfigRender(loginID).Render(ctx))
}

func (h *AdminConfigHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	util.RenderJSON(w, render.NewConfigRender(loginID).Render(ctx))
}
