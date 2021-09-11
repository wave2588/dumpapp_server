package controller

import "context"

type AlterWebController interface {
	SendPendingOrderMsg(ctx context.Context, orderID int64)
	SendPaidOrderMsg(ctx context.Context, orderID int64)
}
