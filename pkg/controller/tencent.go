package controller

import "context"

type TencentController interface {
	DeleteFile(ctx context.Context, TokenPath string) error
	GetSignatureURL(ctx context.Context, name string) (string, error)
}
