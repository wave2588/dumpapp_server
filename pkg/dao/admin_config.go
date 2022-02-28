package dao

import "context"

type AdminConfigDAO interface {
	SetAdminBusy(ctx context.Context, busy bool) error
	GetAdminBusy(ctx context.Context) (bool, error)

	/// 每日限额
	SetDailyFreeCount(ctx context.Context, count int64) error
	GetDailyFreeCount(ctx context.Context) (int64, error)
}
