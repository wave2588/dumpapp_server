package impl

import (
	"context"
	"dumpapp_server/pkg/common/clients"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type StatisticsDAO struct {
	redis redis.Client
}

var DefaultStatisticsDAO *StatisticsDAO

func init() {
	DefaultStatisticsDAO = NewStatisticsDAO()
}

func NewStatisticsDAO() *StatisticsDAO {
	d := &StatisticsDAO{
		redis: clients.DumpRedis,
	}
	return d
}

func (d *StatisticsDAO) generateStatisticsKey(time time.Time) string {
	return fmt.Sprintf("dumpapp:page_view:%d-%d-%d", time.Year(), time.Month(), time.Day())
}

func (d *StatisticsDAO) AddStatistics(ctx context.Context, memberID int64) error {
	key := d.generateStatisticsKey(time.Now())
	cmd := d.redis.ZIncr(ctx, key, &redis.Z{
		Score:  1,
		Member: memberID,
	})
	_, err := cmd.Result()
	return err
}

func (d *StatisticsDAO) GetPageView(ctx context.Context, time time.Time) (int64, error) {
	return 0, nil
}

func (d *StatisticsDAO) GetUserView(ctx context.Context, time time.Time) (int64, error) {
	key := d.generateStatisticsKey(time)
	return d.redis.ZCard(ctx, key).Result()
}
