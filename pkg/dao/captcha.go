package dao

import "context"

type CaptchaDAO interface {
	/// 邮箱相关
	SetEmailCaptcha(ctx context.Context, email, captcha string) error
	GetEmailCaptcha(ctx context.Context, email string) (string, error)
	RemoveEmailCaptcha(ctx context.Context, email string) error

	/// 手机相关
	SetPhoneCaptcha(ctx context.Context, phone, captcha string) error
	GetPhoneCaptcha(ctx context.Context, phone string) (string, error)
	RemovePhoneCaptcha(ctx context.Context, phone string) error

	/// 重置密码
	SetResetPassowordCaptcha(ctx context.Context, phone, captcha string) error
	GetResetPassowordCaptcha(ctx context.Context, phone string) (string, error)
	RemoveResetPassowordCaptcha(ctx context.Context, phone string) error
}
