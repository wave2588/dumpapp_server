package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type WeComHandler struct{}

func NewWeComHandler() *WeComHandler {
	return &WeComHandler{}
}

type postNotificationArgs struct {
	Content string `json:"content" validate:"required"`
}

func (args *postNotificationArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *WeComHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &postNotificationArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	impl.DefaultAlterWebController.SendCustomMsg(ctx, "2ff8e2b8-1098-4418-8bde-97c0f5e15ab5", args.Content)

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
