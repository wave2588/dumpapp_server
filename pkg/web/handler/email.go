package handler

import (
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type EmailHandler struct {
	emailCtl controller.EmailController
}

func NewEmailHandler() *EmailHandler {
	return &EmailHandler{
		emailCtl: impl.DefaultEmailController,
	}
}

type postEmailArgs struct {
	Emails  []string `json:"emails"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
}

func (p *postEmailArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *EmailHandler) PostEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &postEmailArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	for _, email := range args.Emails {
		util.PanicIf(h.emailCtl.SendEmail(ctx, args.Title, args.Content, email, []string{}))
	}

	util.RenderJSON(w, "ok")
}
