package controller

import "context"

type AlterWebController interface {
	/// 支付成功推送
	SendPaidOrderMsg(ctx context.Context, orderID int64)
	/// 砸壳订单推送
	SendDumpOrderMsg(ctx context.Context, loginID, ipaID int64, ipaName string)
}
