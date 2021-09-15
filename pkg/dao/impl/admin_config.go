package impl

import (
	"context"

	"dumpapp_server/pkg/common/clients"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

type AdminConfigDAO struct {
	redis redis.Client
}

var DefaultAdminConfigDAO *AdminConfigDAO

func init() {
	DefaultAdminConfigDAO = NewAdminConfigDAO()
}

func NewAdminConfigDAO() *AdminConfigDAO {
	d := &AdminConfigDAO{
		redis: clients.DumpRedis,
	}
	return d
}

func (d *AdminConfigDAO) generateAdminBusyKey() string {
	return "dump:admin_busy"
}

func (d *AdminConfigDAO) SetAdminBusy(ctx context.Context, busy bool) error {
	key := d.generateAdminBusyKey()
	_, err := d.redis.Set(ctx, key, busy, -1).Result()
	return err
}

func (d *AdminConfigDAO) GetAdminBusy(ctx context.Context) (bool, error) {
	key := d.generateAdminBusyKey()
	re, err := d.redis.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return cast.ToBool(re), nil
}
