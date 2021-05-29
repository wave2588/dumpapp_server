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
func (h *EmailWebController) SendEmailToMaster(ctx context.Context, appName, version, memberEmail string) error {
	title := fmt.Sprintf("Hi~ 订单来了, 应用名称:「%s」", appName)
	content := fmt.Sprintf("应用名称: %s <br>应用版本: %s <br>用户邮箱: %s", appName, version, memberEmail)
	return h.emailCtl.SendEmail(ctx, title, content, "", []string{})
}

func (h *EmailWebController) SendVipEmailToMaster(ctx context.Context, appName, version, memberEmail string) error {
	title := fmt.Sprintf("Hi~ Vip 订单来了, 应用名称:「%s」", appName)
	content := fmt.Sprintf("应用名称: %s <br>应用版本: %s <br>用户邮箱: %s", appName, version, memberEmail)
	return h.emailCtl.SendEmail(ctx, title, content, "", []string{})
}

func (h *EmailWebController) SendUpdateIpaEmail(ctx context.Context, email, name string) error {
	title := "ipa 已更新~"
	content := fmt.Sprintf("您搜索的「%s」ipa 已更新，请访问 https://dumpapp.com 查看。", name)
	return h.emailCtl.SendEmail(ctx, title, content, email, []string{})
}
