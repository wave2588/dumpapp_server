package impl

import (
	"context"
	"fmt"
	"time"

	"dumpapp_server/pkg/common/clients"
	"github.com/go-redis/redis/v8"
)

type CaptchaDAO struct {
	redis redis.Client
}

var DefaultCaptchaDAO *CaptchaDAO

func init() {
	DefaultCaptchaDAO = NewCaptchaDAO()
}

func NewCaptchaDAO() *CaptchaDAO {
	d := &CaptchaDAO{
		redis: clients.DumpRedis,
	}
	return d
}

func (d *CaptchaDAO) generateCaptchaKey(email string) string {
	return fmt.Sprintf("dump:register:captcha:%s", email)
}

/// 5 分钟有效期
func (d *CaptchaDAO) SetCaptcha(ctx context.Context, email, captcha string) error {
	key := d.generateCaptchaKey(email)
	return d.redis.Set(ctx, key, captcha, 300000*time.Millisecond).Err()
}

func (d *CaptchaDAO) GetCaptcha(ctx context.Context, email string) (string, error) {
	key := d.generateCaptchaKey(email)
	res, err := d.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return res, nil
}

func (d *CaptchaDAO) RemoveCaptcha(ctx context.Context, email string) error {
	key := d.generateCaptchaKey(email)
	_, err := d.redis.Del(ctx, key).Result()
	return err
}
