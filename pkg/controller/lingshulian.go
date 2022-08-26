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

	/// 分片上传
	GetCreateMultipartUploadInfo(ctx context.Context, suffix string) (*GetCreateMultipartUploadInfoResp, error)
	GetMultipartUploadPartInfo(ctx context.Context, uploadID, key, bucket string, partNumber int64) (*GetMultipartUploadPartInfoResp, error)
	GetCompleteMultipartUploadInfo(ctx context.Context, uploadID, key, bucket string) (*GetCompleteMultipartUploadInfoResp, error)
}

type GetPutURLResp struct {
	URL      string `json:"url"`
	ExpireTo int64  `json:"expire_to"`
	Key      string `json:"key"`
}

type GetTempSecretKeyResp struct {
	SecretID   string   `json:"secret_id"`
	SecretKey  string   `json:"secret_key"`
	BucketName string   `json:"bucket_name"`
	Prefix     string   `json:"prefix"`
	Key        string   `json:"key"`
	Policy     []string `json:"policy"`
	ExpireTo   int64    `json:"expire_to"`
}

type GetHeaderSignResp struct {
	Sign string `json:"sign"`
	Body string `json:"body"`
	TTL  int64  `json:"ttl"`
}

type GetCreateMultipartUploadInfoResp struct {
	UploadID string `json:"upload_id"`
	Key      string `json:"key"`
	Bucket   string `json:"bucket"`
	ExpireTo int64  `json:"expire_to"`
}

type GetMultipartUploadPartInfoResp struct {
	URL      string `json:"url"`
	ExpireTo int64  `json:"expire_to"`
}

type GetCompleteMultipartUploadInfoResp struct {
	Key string `json:"key"`
}
