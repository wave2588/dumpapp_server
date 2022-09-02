package impl

import (
	"context"
	"encoding/json"
	"time"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/dao"
	"github.com/go-redis/redis/v8"
)

type IpaRankingDAO struct {
	redis redis.Client
}

var DefaultIpaRankingDAO *IpaRankingDAO

func init() {
	DefaultIpaRankingDAO = NewIpaRankingDAO()
}

func NewIpaRankingDAO() *IpaRankingDAO {
	d := &IpaRankingDAO{
		redis: clients.DumpRedis,
	}
	return d
}

func (d *IpaRankingDAO) generateIpaRankingKey() string {
	return "dump:ipa:ranking:v0"
}

func (d *IpaRankingDAO) SetIpaRankingData(ctx context.Context, data *dao.IpaRanking) error {
	key := d.generateIpaRankingKey()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return d.redis.Set(ctx, key, jsonData, 6*time.Hour).Err()
}

func (d *IpaRankingDAO) GetIpaRankingData(ctx context.Context) (*dao.IpaRanking, error) {
	key := d.generateIpaRankingKey()
	data, err := d.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var res *dao.IpaRanking
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *IpaRankingDAO) RemoveIpaRankingData(ctx context.Context) error {
	key := d.generateIpaRankingKey()
	_, err := d.redis.Del(ctx, key).Result()
	return err
}
