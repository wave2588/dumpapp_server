package controller

import (
	"context"
)

type ALiPayController interface {
	/// 根据次数计算价钱
	GetPayURLByNumber(ctx context.Context, loginID, number int64) (string, error)
	CheckPayStatus(ctx context.Context, orderID int64) error
}
