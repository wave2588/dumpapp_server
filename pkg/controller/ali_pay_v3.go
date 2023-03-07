package controller

import (
	"context"
)

type ALiPayV3Controller interface {
	GetPayURLByNumber(ctx context.Context, loginID, number int64) (int64, string, error)
	GetPhoneWapPayURLByNumber(ctx context.Context, loginID, number int64) (int64, string, error)
	GetPhonePayURLByNumber(ctx context.Context, loginID, number int64) (int64, string, error)
	CheckPayStatus(ctx context.Context, orderID int64) error

	Test(ctx context.Context, amount int64) (int64, string, error)
}
