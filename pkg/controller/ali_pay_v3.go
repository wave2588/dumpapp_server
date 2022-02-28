package controller

import (
	"context"
)

type ALiPayV3Controller interface {
	GetPayURLByNumber(ctx context.Context, loginID, number int64) (int64, string, error)
	CheckPayStatus(ctx context.Context, orderID int64) error
}