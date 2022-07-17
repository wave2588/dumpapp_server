package impl

import (
	"context"
	"errors"
	"fmt"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/cast"
)

type ALiPayV3Controller struct {
	client *alipay.Client

	appID        string
	privateKey   string
	aliPublicKey string

	memberPayOrderDAO dao.MemberPayOrderDAO
}

var DefaultALiPayV3Controller *ALiPayV3Controller

func init() {
	DefaultALiPayV3Controller = NewALiPayV3Controller()
}

func NewALiPayV3Controller() *ALiPayV3Controller {
	appID := config.DumpConfig.AppConfig.ALiPayDumpAppID
	privateKey := config.DumpConfig.AppConfig.ALiPayDumpPrivateKey
	aliPublicKey := config.DumpConfig.AppConfig.ALiPayPublicKey
	c, err := alipay.New(appID, privateKey, true)
	util.PanicIf(err)
	// 加载alipay公钥
	err = c.LoadAliPayPublicKey(aliPublicKey)
	util.PanicIf(err)
	return &ALiPayV3Controller{
		client:            c,
		memberPayOrderDAO: impl.DefaultMemberPayOrderDAO,
	}
}

func (c *ALiPayV3Controller) GetPayURLByNumber(ctx context.Context, loginID, number int64) (int64, string, error) {
	id := util2.MustGenerateID(ctx)
	totalAmount := number
	err := c.memberPayOrderDAO.Insert(ctx, &models.MemberPayOrder{
		ID:       id,
		MemberID: loginID,
		Status:   enum.MemberPayOrderStatusPending,
		Amount:   cast.ToFloat64(totalAmount),
		BizExt: datatype.MemberPayOrderBizExt{
			Platform: enum.MemberPayOrderPlatformWeb,
		},
	})
	if err != nil {
		return 0, "", err
	}

	p := alipay.TradePagePay{}
	p.NotifyURL = config.DumpConfig.AppConfig.ALiPayNotifyURLV3
	p.ReturnURL = "https://www.dumpapp.com"
	p.Subject = "Dumpapp"
	p.OutTradeNo = fmt.Sprintf("%d", id)
	p.TotalAmount = fmt.Sprintf("%d", totalAmount)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//p.ExtendParams = map[string]interface{}{
	//	"duration": duration.String(),
	//}
	p.TimeoutExpress = "15m"
	url, err := c.client.TradePagePay(p)
	if err != nil {
		return 0, "", err
	}
	return id, url.String(), nil
}

func (c *ALiPayV3Controller) GetPhonePayURLByNumber(ctx context.Context, loginID, number int64) (int64, string, error) {
	id := util2.MustGenerateID(ctx)
	totalAmount := number
	err := c.memberPayOrderDAO.Insert(ctx, &models.MemberPayOrder{
		ID:       id,
		MemberID: loginID,
		Status:   enum.MemberPayOrderStatusPending,
		Amount:   cast.ToFloat64(totalAmount),
		BizExt: datatype.MemberPayOrderBizExt{
			Platform: enum.MemberPayOrderPlatformIOS,
		},
	})
	if err != nil {
		return 0, "", err
	}

	p := alipay.TradeAppPay{}
	p.OutTradeNo = fmt.Sprintf("%d", id)
	p.TotalAmount = fmt.Sprintf("%d", totalAmount)
	p.NotifyURL = config.DumpConfig.AppConfig.ALiPayNotifyURLV3
	p.Subject = "Dumpapp"
	p.ProductCode = "QUICK_MSECURITY_PAY"
	p.TimeoutExpress = "15m"

	url, err := c.client.TradeAppPay(p)
	if err != nil {
		return 0, "", err
	}

	return id, url, nil
}

func (c *ALiPayV3Controller) CheckPayStatus(ctx context.Context, orderID int64) error {
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

// web-url-->;  https://openapi.alipay.com/gateway.do?app_id=2021002145649331&biz_content=%7B%22subject%22%3A%22Dumpapp%22%2C%22out_trade_no%22%3A%221548736399497236480%22%2C%22total_amount%22%3A%221%22%2C%22product_code%22%3A%22FAST_INSTANT_TRADE_PAY%22%2C%22timeout_express%22%3A%2215m%22%7D&charset=utf-8&format=JSON&method=alipay.trade.page.pay&notify_url=&return_url=https%3A%2F%2Fwww.dumpapp.com&sign=MMEsQ7bSntbdZvbXsuPf5179COGPZjYEC%2B8RY7clMNGG7N286BBbegRS96AQDXPpbGPr85Lqv1N%2BSaB13zeEfh8K1UKqF8%2FSChG0nP2shvb48XRNMfdnNXjbhzB51pkL%2BBd5hSV8nzbjglMPoycEXhZsyPFZKFTXIDttmjvxJIDOQJzCBX1vbNkxk%2Fk1EFm14Wkc%2Fzvld6EJKLC4Cevp%2B24rTSrwrMPcgGqNFtdqcMg7gPy0ueSCG5ws2nKMU4OCZl3VfGIZd%2FACnxwMBAOpc%2B8wcPvd%2BsdTRU8gaMyb5F7DzUF2S%2FX6%2BfDuxmC42qWeOmcHCq2vCX9Eq2uuKR1BMw%3D%3D&sign_type=RSA2&timestamp=2022-07-18+02%3A28%3A26&version=1.0

// url-->:  app_id=2021002145649331&biz_content=%7B%22subject%22%3A%22Dumpapp%22%2C%22out_trade_no%22%3A%221548737020552024064%22%2C%22total_amount%22%3A%221%22%2C%22product_code%22%3A%22QUICK_MSECURITY_PAY%22%2C%22timeout_express%22%3A%2215m%22%7D&charset=utf-8&format=JSON&method=alipay.trade.app.pay&notify_url=&sign=EAd1HcuuYU7t5%2B4RRqIaEbl19hJQL68JoFDAw6HxWTiKi%2BVTAAWX4JUPWel01Yj51UWTqrHe6kIKjAhcSXSLdQLkEJWgCDe7psBI%2BpKt0DIgUecd3PTingnk88KjBidqTRLwYKFr6kpfJaIDXc6KWwHcA25WlUMHi0LFUYZZfyTxgvl6jCm78DRZ5mf548Ypudb6cK9L8O7t%2BjtOxpWdwdIdiwOXTbbVn8jFqNIQoUfyOF26HgAqxF2YUiXSXfboMFihsczvKG71OyeEc%2FShjxlFJW0xcnghUpdYiJT56jdcz4epHH%2Fhw2EH%2Fns9WHYRjFyZi0C63GIJ3WCtkwws0A%3D%3D&sign_type=RSA2&timestamp=2022-07-18+02%3A30%3A54&version=1.0
