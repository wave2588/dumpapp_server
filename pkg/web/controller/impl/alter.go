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
	emailCtl          controller3.EmailController
	accountDAO        dao.AccountDAO
	orderDAO          dao.MemberDownloadOrderDAO
	downloadNumberDAO dao.MemberDownloadNumberDAO
}

var DefaultAlterWebController *AlterWebController

func init() {
	DefaultAlterWebController = NewAlterWebController()
}

func NewAlterWebController() *AlterWebController {
	return &AlterWebController{
		emailCtl:          impl.DefaultEmailController,
		accountDAO:        impl2.DefaultAccountDAO,
		orderDAO:          impl2.DefaultMemberDownloadOrderDAO,
		downloadNumberDAO: impl2.DefaultMemberDownloadNumberDAO,
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

func (c *AlterWebController) SendPaidOrderMsg(ctx context.Context, orderID int64) {
	order, err := c.orderDAO.Get(ctx, orderID)
	if err != nil {
		return
	}
	account, err := c.accountDAO.Get(ctx, order.MemberID)
	if err != nil {
		return
	}
	countMap, err := c.downloadNumberDAO.BatchGetMemberNormalCount(ctx, []int64{account.ID})
	if err != nil {
		return
	}
	email := fmt.Sprintf("邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	number := fmt.Sprintf("充值次数：<font color=\"comment\">%d</font>\n", order.Number)
	number2 := fmt.Sprintf("剩余次数：<font color=\"comment\">%d</font>\n", countMap[account.ID])
	amount := fmt.Sprintf("充值金额：<font color=\"comment\">%v</font>\n", order.Amount.Float64)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">支付成功</font>\n>" +
				email + number + amount + number2 + timeStr,
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})
}

func (c *AlterWebController) SendDumpOrderMsg(ctx context.Context, loginID, ipaID int64, bundleID, ipaName, version string) {
	account, err := c.accountDAO.Get(ctx, loginID)
	if err != nil {
		return
	}
	countMap, err := c.downloadNumberDAO.BatchGetMemberNormalCount(ctx, []int64{account.ID})
	if err != nil {
		return
	}

	ipaIDStr := fmt.Sprintf("应用 ID：<font color=\"comment\">%d</font>\n", ipaID)
	ipaNameStr := fmt.Sprintf("应用名称：<font color=\"comment\">%s</font>\n", ipaName)
	versionStr := fmt.Sprintf("应用版本：<font color=\"comment\">%s</font>\n", version)
	bundleIDStr := fmt.Sprintf("BundleID：<font color=\"comment\">%s</font>\n", bundleID)
	numberStr := fmt.Sprintf("剩余次数：<font color=\"comment\">%d</font>\n", countMap[account.ID])
	emailStr := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">需求来了</font>\n>" +
				ipaIDStr + ipaNameStr + versionStr + bundleIDStr + numberStr + emailStr + timeStr,
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
