package impl

import (
	"context"
	"fmt"
	"time"

	"dumpapp_server/pkg/common/clients"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
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
	cmd := d.redis.ZIncrBy(ctx, key, 1, cast.ToString(memberID))
	_, err := cmd.Result()
	return err
}

func (d *StatisticsDAO) GetPageView(ctx context.Context, time time.Time) (int64, error) {
	key := d.generateStatisticsKey(time)
	data, err := d.redis.ZRevRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		return 0, err
	}
	var count int64 = 0
	for _, datum := range data {
		count += cast.ToInt64(datum.Score)
	}
	return count, nil
}

func (d *StatisticsDAO) GetUserView(ctx context.Context, time time.Time) (int64, error) {
	key := d.generateStatisticsKey(time)
	return d.redis.ZCard(ctx, key).Result()
}
