package controller

import (
	"context"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentController interface {
	DeleteFile(ctx context.Context, TokenPath string) error
	PutByFile(ctx context.Context, name, path string) error
	GetToFile(ctx context.Context, name, path string) error
	GetSignatureURL(ctx context.Context, name string) (string, error)
	ListFile(ctx context.Context, marker *string, limit int) (*cos.BucketGetResult, error)

	SendPhoneRegisterCaptcha(ctx context.Context, captcha, phone string) error
}
