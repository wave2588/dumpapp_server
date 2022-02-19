package controller

import "context"

type MemberPayOrderWebController interface {
	AliPayCallbackOrder(ctx context.Context, orderID int64) error
}
