package impl

import (
	"context"
	"fmt"

	controller3 "dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
)

type EmailWebController struct {
	emailCtl controller3.EmailController
}

var DefaultEmailWebController *EmailWebController

func init() {
	DefaultEmailWebController = NewEmailWebController()
}

func NewEmailWebController() *EmailWebController {
	return &EmailWebController{
		emailCtl: impl.DefaultEmailController,
	}
}

/// 发给自己
func (h *EmailWebController) SendEmailToMaster(ctx context.Context, appName, receiveEmail string) error {
	title := fmt.Sprintf("Hi~ 订单来了, 应用名称:「%s」", appName)
	content := fmt.Sprintf("应用名称: %s <br>接收邮箱: %s", appName, receiveEmail)
	return h.emailCtl.SendEmail(ctx, title, content, "", []string{})
}
