package impl

import (
	"context"
	"fmt"
	"time"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

type CribberDAO struct {
	redis redis.Client
}

var DefaultCribberDAO *CribberDAO

func init() {
	DefaultCribberDAO = NewCribberDAO()
}

func NewCribberDAO() *CribberDAO {
	d := &CribberDAO{
		redis: clients.DumpRedis,
	}
	return d
}

func (d *CribberDAO) generateIncrMemberIP(memberID int64, ip string) string {
	return fmt.Sprintf("dump:incr:memberID:%d:ip:%s", memberID, ip)
}

func (d *CribberDAO) IncrMemberIP(ctx context.Context, memberID int64, ip string) error {
	key := d.generateIncrMemberIP(memberID, ip)
	incrCmd := d.redis.Incr(ctx, key)
	_, err := incrCmd.Result()
	if err != nil {
		return err
	}
	expireCmd := d.redis.Expire(ctx, key, time.Minute)
	expireResult, err := expireCmd.Result()
	if err != nil {
		return err
	}
	if !expireResult {
		return errors.ErrRedisFail
	}
	return nil
}

func (d *CribberDAO) GetMemberIPIncrCount(ctx context.Context, memberID int64, ip string) (int, error) {
	key := d.generateIncrMemberIP(memberID, ip)
	cmd := d.redis.Get(ctx, key)
	res, err := cmd.Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	return cast.ToInt(res), nil
}

func (d *CribberDAO) generateMemberBlacklist(memberID int64) string {
	return fmt.Sprintf("dump:blacklist:memberID:%d", memberID)
}

func (d *CribberDAO) SetMemberIDToBlacklist(ctx context.Context, memberID int64) error {
	key := d.generateMemberBlacklist(memberID)
	return d.redis.Set(ctx, key, true, 20*time.Minute).Err()
}

func (d *CribberDAO) GetBlacklistByMemberID(ctx context.Context, memberID int64) (bool, error) {
	key := d.generateMemberBlacklist(memberID)
	res, err := d.redis.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return cast.ToBool(res), nil
}
