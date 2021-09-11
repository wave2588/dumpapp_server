package impl

import (
	"context"
	"fmt"

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

func (c *AlterWebController) SendPendingOrderMsg(ctx context.Context, orderID int64) {
	order, err := c.orderDAO.Get(ctx, orderID)
	if err != nil {
		return
	}
	account, err := c.accountDAO.Get(ctx, order.MemberID)
	if err != nil {
		return
	}

	email := fmt.Sprintf("邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	number := fmt.Sprintf("充值次数：：<font color=\"comment\">%d</font>\n", order.Number)
	amount := fmt.Sprintf("充值金额：<font color=\"comment\">%v</font>\n", order.Amount.Float64)
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "生成订单了，暂未支付。\n>" +
				email + number + amount,
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})
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
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">支付成功</font>\n>" +
				email + number + amount + number2,
		},
	}
	util.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{})
}
