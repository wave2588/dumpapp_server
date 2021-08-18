package controller

import (
	"context"

	"dumpapp_server/pkg/common/enum"
)

type ALiPayController interface {
	GetPayURL(ctx context.Context, loginID int64, duration enum.MemberVipDurationType) (string, error)
	/// 根据次数计算价钱
	GetPayURLByNumber(ctx context.Context, loginID, number int64) (string, error)
	CheckPayStatus(ctx context.Context, orderID int64) error
}
