package dao

import (
	"context"
	"time"
)

type StatisticsDAO interface {
	AddStatistics(ctx context.Context, memberID int64) error
	GetPageView(ctx context.Context, time time.Time) (int64, error)
	GetUserView(ctx context.Context, time time.Time) (int64, error)
}
