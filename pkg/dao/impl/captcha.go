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

func (d *CaptchaDAO) generateEmailCaptchaKey(email string) string {
	return fmt.Sprintf("dump:register:captcha:%s", email)
}

/// 5 分钟有效期
func (d *CaptchaDAO) SetEmailCaptcha(ctx context.Context, email, captcha string) error {
	key := d.generateEmailCaptchaKey(email)
	return d.redis.Set(ctx, key, captcha, 5*time.Minute).Err()
}

func (d *CaptchaDAO) GetEmailCaptcha(ctx context.Context, email string) (string, error) {
	key := d.generateEmailCaptchaKey(email)
	res, err := d.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return res, nil
}

func (d *CaptchaDAO) RemoveEmailCaptcha(ctx context.Context, email string) error {
	key := d.generateEmailCaptchaKey(email)
	_, err := d.redis.Del(ctx, key).Result()
	return err
}

func (d *CaptchaDAO) generatePhoneCaptchaKey(phone string) string {
	return fmt.Sprintf("dump:register:phone:captcha:%s", phone)
}

/// 15 分钟有效期
func (d *CaptchaDAO) SetPhoneCaptcha(ctx context.Context, phone, captcha string) error {
	key := d.generatePhoneCaptchaKey(phone)
	return d.redis.Set(ctx, key, captcha, 15*time.Minute).Err()
}

func (d *CaptchaDAO) GetPhoneCaptcha(ctx context.Context, phone string) (string, error) {
	key := d.generatePhoneCaptchaKey(phone)
	res, err := d.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return res, nil
}

func (d *CaptchaDAO) RemovePhoneCaptcha(ctx context.Context, phone string) error {
	key := d.generateResetPasswordCaptchaKey(phone)
	_, err := d.redis.Del(ctx, key).Result()
	return err
}

/// 重置密码
func (d *CaptchaDAO) generateResetPasswordCaptchaKey(email string) string {
	return fmt.Sprintf("dump:reset_password:captcha:%s", email)
}

func (d *CaptchaDAO) SetResetPassowordCaptcha(ctx context.Context, email, captcha string) error {
	key := d.generateResetPasswordCaptchaKey(email)
	return d.redis.Set(ctx, key, captcha, 15*time.Minute).Err()
}

func (d *CaptchaDAO) GetResetPassowordCaptcha(ctx context.Context, email string) (string, error) {
	key := d.generateResetPasswordCaptchaKey(email)
	res, err := d.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return res, nil
}

func (d *CaptchaDAO) RemoveResetPassowordCaptcha(ctx context.Context, email string) error {
	key := d.generateResetPasswordCaptchaKey(email)
	_, err := d.redis.Del(ctx, key).Result()
	return err
}
