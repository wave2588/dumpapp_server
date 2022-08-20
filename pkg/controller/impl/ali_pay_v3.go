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

func (c *ALiPayV3Controller) GetPhoneWapPayURLByNumber(ctx context.Context, loginID, number int64) (int64, string, error) {
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

	p := alipay.TradeWapPay{}
	p.OutTradeNo = fmt.Sprintf("%d", id)
	p.TotalAmount = fmt.Sprintf("%d.00", totalAmount)
	p.NotifyURL = config.DumpConfig.AppConfig.ALiPayNotifyURLV3
	p.Subject = "Dumpapp"
	p.ProductCode = "QUICK_WAP_WAY"
	p.TimeoutExpress = "15m"

	url, err := c.client.TradeWapPay(p)
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
	p.TotalAmount = fmt.Sprintf("%d.00", totalAmount)
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
	if rsp.Content.TradeStatus != alipay.TradeStatusSuccess {
		return errors.New(fmt.Sprintf("订单支付未成功. order_id= %d", orderID))
	}

	return nil
}
