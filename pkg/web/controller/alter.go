package controller

import "context"

type AlterWebController interface {
	/// 支付成功推送
	SendPaidOrderMsg(ctx context.Context, orderID int64)
	/// 砸壳订单推送
	SendDumpOrderMsg(ctx context.Context, loginID, ipaID int64, bundleID, ipaName, version string)
	/// feedback 推送
	SendFeedbackMsg(ctx context.Context, loginID int64, content string)
}
