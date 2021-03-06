package impl

import (
	"context"
	"fmt"
	"time"

	"dumpapp_server/pkg/config"
	controller3 "dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/util"
)

type AlterWebController struct {
	emailCtl                   controller3.EmailController
	accountDAO                 dao.AccountDAO
	memberDownloadIpaRecordDAO dao.MemberDownloadIpaRecordDAO
	memberDeviceDAO            dao.MemberDeviceDAO
	certificateDAO             dao.CertificateV2DAO
}

var DefaultAlterWebController *AlterWebController

func init() {
	DefaultAlterWebController = NewAlterWebController()
}

func NewAlterWebController() *AlterWebController {
	return &AlterWebController{
		emailCtl:                   impl.DefaultEmailController,
		accountDAO:                 impl2.DefaultAccountDAO,
		memberDownloadIpaRecordDAO: impl2.DefaultMemberDownloadIpaRecordDAO,
		memberDeviceDAO:            impl2.DefaultMemberDeviceDAO,
		certificateDAO:             impl2.DefaultCertificateV2DAO,
	}
}

func (c *AlterWebController) SendMsg(ctx context.Context, memberID int64, name, version, bundleID string) {
	//data := map[string]interface{}{
	//	"msgtype": "markdown",
	//	"markdown": map[string]interface{}{
	//		"content": "DumpApp - 应用名称：<font color=\"info\">微信</font>\n>" +
	//			"应用名称:<font color=\"comment\">微信</font>\n" +
	//			"应用版本:<font color=\"comment\">微信</font>\n" +
	//			"bundleID:<font color=\"comment\">11111</font>\n" +
	//			"邮箱:<font color=\"comment\">zhanghaibo</font>\n" +
	//			"手机号:<font color=\"comment\">15711367321</font>",
	//	},
	//}
	//SendWeiXinBot(ctx, keyID, data, receivers)
}

func (c *AlterWebController) SendCustomMsg(ctx context.Context, content string) {
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": content,
		},
	}
	util.SendWeiXinBot(ctx, "2ff8e2b8-1098-4418-8bde-97c0f5e15ab5", data, []string{})

	data = map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":        "",
			"mentioned_list": []string{"@all"},
		},
	}
	util.SendWeiXinBot(ctx, "2ff8e2b8-1098-4418-8bde-97c0f5e15ab5", data, []string{})
}

func (c *AlterWebController) SendDumpOrderMsg(ctx context.Context, loginID, ipaID int64, bundleID, ipaName, version string) {
	account, err := c.accountDAO.Get(ctx, loginID)
	if err != nil {
		return
	}

	ipaIDStr := fmt.Sprintf("应用 ID：<font color=\"comment\">%d</font>\n", ipaID)
	ipaNameStr := fmt.Sprintf("应用名称：<font color=\"comment\">%s</font>\n", ipaName)
	versionStr := fmt.Sprintf("应用版本：<font color=\"comment\">%s</font>\n", version)
	bundleIDStr := fmt.Sprintf("BundleID：<font color=\"comment\">%s</font>\n", bundleID)
	emailStr := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">需求来了</font>\n>" +
				ipaIDStr + ipaNameStr + versionStr + bundleIDStr + emailStr + timeStr,
		},
	}
	util.SendWeiXinBot(ctx, "2ff8e2b8-1098-4418-8bde-97c0f5e15ab5", data, []string{})

	data = map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":        "",
			"mentioned_list": []string{"@all"},
		},
	}
	util.SendWeiXinBot(ctx, "2ff8e2b8-1098-4418-8bde-97c0f5e15ab5", data, []string{})
}

func (c *AlterWebController) SendFeedbackMsg(ctx context.Context, loginID int64, content string) {
	account, err := c.accountDAO.Get(ctx, loginID)
	if err != nil {
		return
	}

	emailStr := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	contentStr := fmt.Sprintf("反馈：<font color=\"comment\">%s</font>\n", content)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"comment\">反馈来了</font>\n>" +
				emailStr + contentStr + timeStr,
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})

	data = map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":        "",
			"mentioned_list": []string{"@all"},
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})
}

func (c *AlterWebController) SendCreateCertificateFailMsg(ctx context.Context, loginID, deviceID int64, errorMessage string) {
	account, err := c.accountDAO.Get(ctx, loginID)
	if err != nil {
		return
	}

	errorStr := fmt.Sprintf("错误信息：<font color=\"comment\">%s</font>\n", errorMessage)
	deviceStr := fmt.Sprintf("设备 ID：<font color=\"comment\">%d</font>\n", deviceID)
	emailStr := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">证书服务报错了!</font>\n>" +
				errorStr + deviceStr + emailStr + timeStr,
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})

	data = map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":        "",
			"mentioned_list": []string{"@all"},
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})
}

func (c *AlterWebController) SendCreateCertificateSuccessMsg(ctx context.Context, loginID, deviceID, cerID int64) {
	account, err := c.accountDAO.Get(ctx, loginID)
	if err != nil {
		return
	}
	device, err := c.memberDeviceDAO.Get(ctx, deviceID)
	if err != nil {
		return
	}
	cer, err := c.certificateDAO.Get(ctx, cerID)
	if err != nil {
		return
	}

	cerIDStr := fmt.Sprintf("证书 ID：<font color=\"comment\">%d</font>\n", cer.ID)
	deviceIDStr := fmt.Sprintf("设备 ID：<font color=\"comment\">%d</font>\n", device.ID)
	udidStr := fmt.Sprintf("UDID：<font color=\"comment\">%s</font>\n", device.Udid)
	emailStr := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"info\">证书购买成功</font>\n>" +
				cerIDStr + deviceIDStr + udidStr + emailStr + timeStr,
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})
}

func (c *AlterWebController) SendAccountMsg(ctx context.Context) {
	count, err := c.accountDAO.Count(ctx, nil)
	if err != nil {
		return
	}
	message := fmt.Sprintf("当前注册用户总数：%d", count)

	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">注册用户实时同步</font>\n>" +
				message,
		},
	}
	util.SendWeiXinBot(ctx, "2ff8e2b8-1098-4418-8bde-97c0f5e15ab5", data, []string{})
}

func (c *AlterWebController) SendDeviceLog(ctx context.Context, title string, memberID int64, values map[string]string) {
	account, err := c.accountDAO.Get(ctx, memberID)
	if err != nil {
		return
	}
	message := fmt.Sprintf("邮箱：<font color=\"comment\">%s</font>\n member_id: <font color=\"comment\">%d</font>\n", account.Email, memberID)
	for key, value := range values {
		msg := fmt.Sprintf("%s：<font color=\"comment\">%s</font>\n", key, value)
		message += msg
	}
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": fmt.Sprintf("<font color=\"warning\">%s</font>\n>", title) + message,
		},
	}
	util.SendWeiXinBot(ctx, "c8e8a862-4d4a-44ea-8250-de9297b5e8bc", data, []string{})
}

func (c *AlterWebController) SendInstallAppCreateCertificateFailMsg(ctx context.Context, cdkey, udid string, errorMessage string) {
	errorStr := fmt.Sprintf("错误信息：<font color=\"comment\">%s</font>\n", errorMessage)
	deviceStr := fmt.Sprintf("udid：<font color=\"comment\">%s</font>\n", udid)
	emailStr := fmt.Sprintf("兑换码：<font color=\"comment\">%s</font>\n", cdkey)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">证书服务报错了!</font>\n>" +
				errorStr + deviceStr + emailStr + timeStr,
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})

	data = map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":        "",
			"mentioned_list": []string{"@all"},
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})
}
