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

func (h *EmailWebController) SendUpdateIpaEmail(ctx context.Context, ipaID int64, email, name string) error {
	title := "ipa 已更新~"
	content := fmt.Sprintf("DumpAPP - 您需要的 [%s] IPA 已更新，请访问 https://www.dumpapp.com/download?appid=%d 查看。", name, ipaID)
	return h.emailCtl.SendEmail(ctx, title, content, email, []string{})
}
