package controller

import (
	"context"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentController interface {
	DeleteFile(ctx context.Context, TokenPath string) error
	GetToFile(ctx context.Context, name, path string) error
	GetSignatureURL(ctx context.Context, name string, expired time.Duration) (string, error)
	ListFile(ctx context.Context, marker *string, limit int) (*cos.BucketGetResult, error)

	/// sign_ipa bucket
	PutMemberSignIpa(ctx context.Context, name string, data string) error
	GetMemberSignIpa(ctx context.Context, ipaToken string) (string, error)
	DeleteMemberSignIpa(ctx context.Context, tokenPath string) error

	/// 验证码相关
	SendPhoneRegisterCaptcha(ctx context.Context, captcha, phone string) error
}
