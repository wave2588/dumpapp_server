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
	content := fmt.Sprintf("感谢您使用 DumpAPP <br><br>您需要的 [%s] IPA 已更新，请访问 https://www.dumpapp.com/download?appid=%d 查看。 <br><br>有任何问题可打开官网，扫码联系微信 - https://www.dumpapp.com", name, ipaID)
	return h.emailCtl.SendEmail(ctx, title, content, email, []string{})
}
