package impl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type LingshulianController struct {
	Session *session.Session
	Svc     *s3.S3
}

var DefaultLingshulianController *LingshulianController

func init() {
	DefaultLingshulianController = NewLingshulianController()
}

func NewLingshulianController() *LingshulianController {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			config.DumpConfig.AppConfig.LingshulianSecretID,
			config.DumpConfig.AppConfig.LingshulianSecretKey,
			"",
		),
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("s3-us-east-1.ossfiles.com"),
	})
	util.PanicIf(err)
	return &LingshulianController{
		Session: sess,
		Svc:     s3.New(sess),
	}
}

func (c *LingshulianController) GetPutURL(ctx context.Context, bucket, key string) (*controller.GetPutURLResp, error) {
	resp, _ := c.Svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	expire := 60 * time.Minute
	url, err := resp.Presign(expire)
	if err != nil {
		return nil, err
	}
	startAt := time.Now()
	expireAt := startAt.Add(expire)
	return &controller.GetPutURLResp{
		URL:      url,
		StartAt:  startAt.Unix(),
		ExpireAt: expireAt.Unix(),
		Token:    key,
	}, nil
}

func (c *LingshulianController) PutFile(ctx context.Context, url string, data io.Reader) error {
	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *LingshulianController) GetURL(ctx context.Context, bucket, key string) (string, error) {
	return fmt.Sprintf("%s/%s/%s", config.DumpConfig.AppConfig.LingshulianMemberSignIpaHost, bucket, key), nil
}

func (c *LingshulianController) Put(ctx context.Context, bucket, key string, body io.ReadSeeker) error {
	_, err := c.Svc.PutObject(&s3.PutObjectInput{
		Bucket: util.StringPtr(bucket),
		Key:    util.StringPtr(key),
		Body:   body,
	})
	return err
}

func (c *LingshulianController) Delete(ctx context.Context, bucket, key string) error {
	_, err := c.Svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: util.StringPtr(bucket),
		Key:    util.StringPtr(key),
	})
	return err
}
