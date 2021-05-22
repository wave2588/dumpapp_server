package controller

import (
	"context"

	"dumpapp_server/pkg/common/enum"
)

type ALiPayController interface {
	GetPayURL(ctx context.Context, loginID int64, duration enum.MemberVipDurationType) (string, error)
	CheckPayStatus(ctx context.Context, orderID int64) error
}
