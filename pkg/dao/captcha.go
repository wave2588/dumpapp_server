package dao

import "context"

type CaptchaDAO interface {
	SetCaptcha(ctx context.Context, email, captcha string) error
	GetCaptcha(ctx context.Context, email string) (string, error)
	RemoveCaptcha(ctx context.Context, email string) error
}
