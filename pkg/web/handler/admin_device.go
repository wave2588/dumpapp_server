package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/controller"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type AdminDeviceHandler struct {
	adminDeviceWebCtl controller.AdminDeviceController
}

func NewAdminDeviceHandler() *AdminDeviceHandler {
	return &AdminDeviceHandler{
		adminDeviceWebCtl: impl2.DefaultAdminDeviceWebController,
	}
}

type unbindArgs struct {
	UDID  string `json:"udid" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func (p *unbindArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminDeviceHandler) Unbind(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &unbindArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	util.PanicIf(h.adminDeviceWebCtl.Unbind(ctx, args.Email, args.UDID))
}
