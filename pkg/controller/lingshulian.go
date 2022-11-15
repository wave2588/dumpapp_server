package controller

import (
	"context"
	"fmt"
	"io"
	"time"

	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/errors"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-playground/validator/v10"
)

type LingshulianController interface {
	GetPutURL(ctx context.Context, bucket, key string) (*GetPutURLResp, error)
	GetTempSecretKey(ctx context.Context) (*GetTempSecretKeyResp, error)

	PutFile(ctx context.Context, url string, data io.Reader) error
	GetURL(ctx context.Context, bucket, key string) (string, error)
	Put(ctx context.Context, bucket, key string, body io.ReadSeeker) error
	Delete(ctx context.Context, bucket, key string) error
	List(ctx context.Context, bucket string, marker *string, limit int64) (*s3.ListObjectsOutput, error)

	/// 分片上传
	PostCreateMultipartUploadInfo(ctx context.Context, request *PostCreateMultipartUploadInfoRequest) (*PostCreateMultipartUploadInfoResp, error)
	PostMultipartUploadPartInfo(ctx context.Context, request *PostMultipartUploadPartInfoRequest) (*PostMultipartUploadPartInfoResp, error)
	PostCompleteMultipartUploadInfo(ctx context.Context, request *PostCompleteMultipartUploadInfoRequest) (*PostCompleteMultipartUploadInfoResp, error)
	PostAbortMultipartUploadInfo(ctx context.Context, request *PostAbortMultipartUploadPartInfoRequest) (*PostAbortMultipartUploadPartInfoResponse, error)

	GetSignatureURL(ctx context.Context, bucket, key string, expired time.Duration) (string, error)
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

/// 开始上传
type PostCreateMultipartUploadInfoRequest struct {
	Suffix *string `json:"suffix"`
	Key    *string `json:"key"`

	Bucket string `json:"bucket"`
}

func (p *PostCreateMultipartUploadInfoRequest) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.Suffix == nil && p.Key == nil {
		return errors.UnproccessableError("Suffix or Key 不能为空")
	}
	if p.Suffix != nil && p.Key != nil {
		return errors.UnproccessableError("Suffix or Key 不能同时使用")
	}
	/// 如果 bucket 传空，则默认给 membersignipa bucket
	if p.Bucket == "" {
		p.Bucket = config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket
	}
	return nil
}

type PostCreateMultipartUploadInfoResp struct {
	UploadID string `json:"upload_id"`
	Key      string `json:"key"`
	Bucket   string `json:"bucket"`
	ExpireTo int64  `json:"expire_to"`
}

/// 上传分片
type PostMultipartUploadPartInfoRequest struct {
	UploadID   string `json:"upload_id" validate:"required"`
	Key        string `json:"key" validate:"required"`
	Bucket     string `json:"bucket" validate:"required"`
	PartNumber int64  `json:"part_number" validate:"required"`
}

func (args *PostMultipartUploadPartInfoRequest) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if args.UploadID == "" || args.Key == "" || args.Bucket == "" || args.PartNumber <= 0 {
		return errors.UnproccessableError("参数格式错误")
	}
	return nil
}

type PostMultipartUploadPartInfoResp struct {
	URLData  []string `json:"url_data"`
	ExpireTo int64    `json:"expire_to"`
}

/// 完成上传
type PostCompleteMultipartUploadInfoRequest struct {
	UploadID string                                        `json:"upload_id" validate:"required"`
	Key      string                                        `json:"key" validate:"required"`
	Bucket   string                                        `json:"bucket" validate:"required"`
	Parts    []*PostCompleteMultipartUploadPartInfoRequest `json:"parts" validate:"required"`
}

type PostCompleteMultipartUploadPartInfoRequest struct {
	PartNumber int64  `json:"part_number" validate:"required"`
	ETag       string `json:"e_tag" validate:"required"`
}

func (p *PostCompleteMultipartUploadInfoRequest) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.UploadID == "" || p.Key == "" || p.Bucket == "" || len(p.Parts) == 0 {
		return errors.UnproccessableError("参数格式错误")
	}
	return nil
}

type PostCompleteMultipartUploadInfoResp struct {
	Key string `json:"key"`
}

/// 取消上传
type PostAbortMultipartUploadPartInfoRequest struct {
	UploadID string `json:"upload_id" validate:"required"`
	Key      string `json:"key" validate:"required"`
	Bucket   string `json:"bucket" validate:"required"`
}

func (p *PostAbortMultipartUploadPartInfoRequest) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.UploadID == "" || p.Key == "" || p.Bucket == "" {
		return errors.UnproccessableError("参数格式错误")
	}
	return nil
}

type PostAbortMultipartUploadPartInfoResponse struct {
	RequestCharged *string `json:"request_charged"`
}
