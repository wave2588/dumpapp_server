package impl

import (
	"context"
	"errors"
	"fmt"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"github.com/smartwalle/alipay/v3"
	"github.com/volatiletech/null/v8"
)

type ALiPayController struct {
	client *alipay.Client

	appID        string
	privateKey   string
	aliPublicKey string

	memberVipOrderDAO      dao.MemberVipOrderDAO
	memberDownloadOrderDAO dao.MemberDownloadOrderDAO
}

var DefaultALiPayController *ALiPayController

func init() {
	DefaultALiPayController = NewALiPayController()
}

func NewALiPayController() *ALiPayController {
	appID := config.DumpConfig.AppConfig.ALiPayDumpAppID
	privateKey := config.DumpConfig.AppConfig.ALiPayDumpPrivateKey
	aliPublicKey := config.DumpConfig.AppConfig.ALiPayPublicKey
	c, err := alipay.New(appID, privateKey, true)
	util.PanicIf(err)
	// 加载alipay公钥
	err = c.LoadAliPayPublicKey(aliPublicKey)
	util.PanicIf(err)
	return &ALiPayController{
		client: c,

		memberVipOrderDAO:      impl.DefaultMemberVipOrderDAO,
		memberDownloadOrderDAO: impl.DefaultMemberDownloadOrderDAO,
	}
}

func (c *ALiPayController) GetPayURLByNumber(ctx context.Context, loginID, number int64) (string, error) {
	id := util2.MustGenerateID(ctx)
	err := c.memberDownloadOrderDAO.Insert(ctx, &models.MemberDownloadOrder{
		ID:       id,
		MemberID: loginID,
		Status:   enum.MemberDownloadOrderStatusPending,
		Number:   number,
	})
	if err != nil {
		return "", err
	}

	p := alipay.TradePagePay{}
	p.NotifyURL = config.DumpConfig.AppConfig.ALiPayNotifyURLV2
	p.ReturnURL = "https://www.dumpapp.com"
	p.Subject = "Dumpapp"
	p.OutTradeNo = fmt.Sprintf("%d", id)
	p.TotalAmount = fmt.Sprintf("%d", number*constant.DownloadIpaPrice)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//p.ExtendParams = map[string]interface{}{
	//	"duration": duration.String(),
	//}
	p.TimeoutExpress = "15m"
	url, err := c.client.TradePagePay(p)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (c *ALiPayController) GetPayURL(ctx context.Context, loginID int64, duration enum.MemberVipDurationType) (string, error) {
	id := util2.MustGenerateID(ctx)
	err := c.memberVipOrderDAO.Insert(ctx, &models.MemberVipOrder{
		ID:       id,
		MemberID: loginID,
		Status:   enum.MemberVipOrderStatusPending,
		Duration: null.StringFrom(duration.String()),
	})
	if err != nil {
		return "", err
	}

	p := alipay.TradePagePay{}
	p.NotifyURL = config.DumpConfig.AppConfig.ALiPayNotifyURL
	p.ReturnURL = "https://www.dumpapp.com"
	p.Subject = constant.MemberVipDurationTypeToSubject[duration]
	p.OutTradeNo = fmt.Sprintf("%d", id)
	p.TotalAmount = fmt.Sprintf("%d", constant.MemberVipDurationTypeToPrice[duration])
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//p.ExtendParams = map[string]interface{}{
	//	"duration": duration.String(),
	//}
	p.TimeoutExpress = "15m"
	url, err := c.client.TradePagePay(p)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func (c *ALiPayController) CheckPayStatus(ctx context.Context, orderID int64) error {
	p := alipay.TradeQuery{}
	p.OutTradeNo = fmt.Sprintf("%d", orderID)
	p.QueryOptions = []string{"TRADE_SETTLE_INFO"}

	rsp, err := c.client.TradeQuery(p)
	if err != nil {
		return err
	}

	if rsp.Content.Code != alipay.CodeSuccess {
		return errors.New(fmt.Sprintf("订单支付未成功. order_id= %d", orderID))
	}

	return nil
}
