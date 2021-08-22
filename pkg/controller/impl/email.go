package impl

import (
	"context"
	"crypto/tls"

	"dumpapp_server/pkg/config"
	"gopkg.in/gomail.v2"
)

type EmailController struct{}

var DefaultEmailController *EmailController

func init() {
	DefaultEmailController = NewEmailController()
}

func NewEmailController() *EmailController {
	return &EmailController{}
}

func (c *EmailController) SendEmail(ctx context.Context, title, content, receiveEmail string, images []string) error {
	/// re == "" 说明是要发给自己的, 订单邮件
	re := receiveEmail
	if re == "" {
		re = config.DumpConfig.AppConfig.DumpAppEmail
	}

	fromEmail := config.DumpConfig.AppConfig.DumpAppFromEmail
	fromEmailPassword := config.DumpConfig.AppConfig.DumpAppFromEmailPassword
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail) // 发件人
	m.SetHeader("To", re)          // 收件人
	m.SetHeader("Subject", title)  // 邮件标题
	// m.Attach("E:\\IMGP0814.JPG")       //邮件附件
	for _, image := range images {
		m.Embed(image)
	}
	m.SetBody("text/html", content) // 邮件内容

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, fromEmail, fromEmailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	/// 邮件发送服务器信息,使用授权码而非密码
	return d.DialAndSend(m)
}
