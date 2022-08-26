package impl

import (
	"context"
	"errors"
	"fmt"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/cast"
	"github.com/volatiletech/strmangle"
)

type ALiPayInstallAppController struct {
	client *alipay.Client

	appID        string
	privateKey   string
	aliPublicKey string

	installAppCDKEYOrderDAO dao.InstallAppCdkeyOrderDAO
	installAppCDKEYDAO      dao.InstallAppCdkeyDAO
}

var DefaultALiPayInstallAppController *ALiPayInstallAppController

func init() {
	DefaultALiPayInstallAppController = NewALiPayInstallAppController()
}

func NewALiPayInstallAppController() *ALiPayInstallAppController {
	appID := config.DumpConfig.AppConfig.ALiPayDumpAppID
	privateKey := config.DumpConfig.AppConfig.ALiPayDumpPrivateKey
	aliPublicKey := config.DumpConfig.AppConfig.ALiPayPublicKey
	c, err := alipay.New(appID, privateKey, true)
	util.PanicIf(err)
	// 加载alipay公钥
	err = c.LoadAliPayPublicKey(aliPublicKey)
	util.PanicIf(err)
	return &ALiPayInstallAppController{
		client:                  c,
		installAppCDKEYOrderDAO: impl.DefaultInstallAppCdkeyOrderDAO,
		installAppCDKEYDAO:      impl.DefaultInstallAppCdkeyDAO,
	}
}

func (c *ALiPayInstallAppController) GetPayURLByInstallApp(ctx context.Context, number int64, contactWay string) (int64, string, error) {
	id := util2.MustGenerateID(ctx)
	totalAmount := number * 30
	bizExt := constant.InstallAppCDKEYOrderBizExt{
		ContactWay: contactWay,
	}
	err := c.installAppCDKEYOrderDAO.Insert(ctx, &models.InstallAppCdkeyOrder{
		ID:     id,
		Status: enum.MemberPayOrderStatusPending,
		Amount: cast.ToFloat64(totalAmount),
		Number: number,
		BizExt: bizExt.String(),
	})
	if err != nil {
		return 0, "", err
	}

	p := alipay.TradePagePay{}
	p.NotifyURL = config.DumpConfig.AppConfig.ALiPayNotifyURLByInstallApp
	p.ReturnURL = fmt.Sprintf("https://www.dumpapp.com/installbuy?order_id=%d", id)
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

func (c *ALiPayInstallAppController) AliPayCallbackOrder(ctx context.Context, orderID int64) error {
	util.PanicIf(c.checkPayStatus(ctx, orderID))

	order, err := c.installAppCDKEYOrderDAO.Get(ctx, orderID)
	if err != nil {
		return err
	}

	/// 支付成功的订单即可忽略
	if order.Status == enum.MemberPayOrderStatusPaid {
		return nil
	}

	number := int(order.Number)
	outIDs, err := c.getOutIDs(ctx, number)
	if err != nil {
		return err
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	order.Status = enum.MemberPayOrderStatusPaid
	util.PanicIf(c.installAppCDKEYOrderDAO.Update(ctx, order))

	for _, oID := range outIDs {
		id := util2.MustGenerateID(ctx)
		err = c.installAppCDKEYDAO.Insert(ctx, &models.InstallAppCdkey{
			ID:      id,
			OutID:   oID,
			Status:  enum.InstallAppCDKeyStatusNormal,
			OrderID: orderID,
		})
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return nil
}

func (c *ALiPayInstallAppController) checkPayStatus(ctx context.Context, orderID int64) error {
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

func (c *ALiPayInstallAppController) getOutIDs(ctx context.Context, number int) ([]string, error) {
	outIDs := make([]string, 0)
	/// 生成 number * 10 的数量，以防重复
	for i := 0; i < number*10; i++ {
		outIDs = append(outIDs, util2.MustGenerateAppCDKEY())
	}
	outIDs = strmangle.RemoveDuplicates(outIDs)

	cMap, err := c.installAppCDKEYDAO.BatchGetByOutID(ctx, outIDs)
	if err != nil {
		return nil, err
	}

	resultOutIDs := make([]string, 0)
	for _, oID := range outIDs {
		if len(resultOutIDs) == number {
			return resultOutIDs, nil
		}
		if _, ok := cMap[oID]; !ok {
			resultOutIDs = append(resultOutIDs, oID)
		}
	}
	return nil, errors2.ErrInstallAppGenerateCDKeyFail
}
