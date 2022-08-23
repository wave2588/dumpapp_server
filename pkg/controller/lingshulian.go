package controller

import (
	"context"
	"io"
)

type LingshulianController interface {
	GetPutURL(ctx context.Context, bucket, key string) (*GetPutURLResp, error)
	GetTempSecretKey(ctx context.Context) (*GetTempSecretKeyResp, error)

	PutFile(ctx context.Context, url string, data io.Reader) error
	GetURL(ctx context.Context, bucket, key string) (string, error)
	Put(ctx context.Context, bucket, key string, body io.ReadSeeker) error
	Delete(ctx context.Context, bucket, key string) error
}

type GetPutURLResp struct {
	URL      string `json:"url"`
	StartAt  int64  `json:"start_at"`
	ExpireAt int64  `json:"expire_at"`
	Token    string `json:"token"`
}

type GetTempSecretKeyResp struct {
	TempSecretKey string `json:"temp_secret_key"`
	StartAt       int64  `json:"start_at"`
	ExpireAt      int64  `json:"expire_at"`
}
