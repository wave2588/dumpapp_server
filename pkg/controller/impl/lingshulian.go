package impl

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	errors3 "errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	jsoniter "github.com/json-iterator/go"
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

func (c *LingshulianController) ss(ctx context.Context) {
}

func (c *LingshulianController) GetPutURL(ctx context.Context, bucket, key string) (*controller.GetPutURLResp, error) {
	resp, _ := c.Svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	expire := 60 * time.Minute
	URL, err := resp.Presign(expire)
	if err != nil {
		return nil, err
	}
	startAt := time.Now()
	expireAt := startAt.Add(expire)
	return &controller.GetPutURLResp{
		URL:      URL,
		ExpireTo: expireAt.Unix(),
		Key:      key,
	}, nil
}

var lingshulianAuthSecretURL = "https://api.lingshulian.com/api/auth/secret"

type postLingshulianAuthSecretResp struct {
	Status  string                             `json:"status"`
	Code    int64                              `json:"code"`
	Message string                             `json:"message"`
	Data    *postLingshulianAuthSecretDataResp `json:"data"`
	Error   map[string]interface{}             `json:"error"`
}

type postLingshulianAuthSecretDataResp struct {
	SecretID   string   `json:"secret_id"`
	SecretKey  string   `json:"secret_key"`
	BucketName string   `json:"bucket_name"`
	Prefix     string   `json:"prefix"`
	Key        string   `json:"key"`
	Policy     []string `json:"policy"`
	ExpireTo   int64    `json:"expire_to"`
}

func (c *LingshulianController) GetTempSecretKey(ctx context.Context) (*controller.GetTempSecretKeyResp, error) {
	sign, err := c.GetHeaderSign(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := util2.HttpRequestV2("POST", lingshulianAuthSecretURL, map[string]string{
		"x-lingshulian-sign": sign.Sign,
		"Content-Type":       "application/json",
	}, strings.NewReader(sign.Body))
	util.PanicIf(err)

	var re *postLingshulianAuthSecretResp
	err = json.Unmarshal(resp, &re)
	if err != nil {
		return nil, err
	}
	if re == nil || re.Code != 200 || re.Data == nil {
		return nil, errors.NewDefaultAPIError(401, 401, "ErrGetTempSecretFail", "获取秘钥失败")
	}

	return &controller.GetTempSecretKeyResp{
		SecretID:   re.Data.SecretID,
		SecretKey:  re.Data.SecretKey,
		BucketName: re.Data.BucketName,
		Prefix:     re.Data.Prefix,
		Key:        re.Data.Key,
		Policy:     re.Data.Policy,
		ExpireTo:   re.Data.ExpireTo,
	}, nil
}

func (c *LingshulianController) GetHeaderSign(ctx context.Context) (*controller.GetHeaderSignResp, error) {
	accessID := config.DumpConfig.AppConfig.LingshulianSecretID
	accessKey := config.DumpConfig.AppConfig.LingshulianSecretKey

	ttl := int64(900)
	bodyInfo := map[string]interface{}{
		"ttl": ttl,
		"policy": []string{
			"full_control",
		},
		"bucket_name": config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket,
	}
	model := "POST"

	urlInfo, err := url.Parse(lingshulianAuthSecretURL)
	util.PanicIf(err)

	host := urlInfo.Host
	urlPath := urlInfo.Path
	bodyData, err := json.Marshal(bodyInfo)
	util.PanicIf(err)
	body := string(bodyData)

	startAt := time.Now()
	expiryTo := startAt.Unix() + ttl
	signAccessSecret := fmt.Sprintf("%s-%s", accessID, accessKey)
	signString := fmt.Sprintf("%s\n%s\n%s\n%s\n%d", model, host, urlPath, body, expiryTo)
	sign := fmt.Sprintf("%s-%d-%s", accessID, expiryTo, c.hmac(signString, signAccessSecret))
	return &controller.GetHeaderSignResp{
		Sign: sign,
		Body: body,
		TTL:  ttl,
	}, nil
}

func (c *LingshulianController) hmac(key, data string) string {
	mac := hmac.New(sha1.New, []byte(data))
	mac.Write([]byte(key))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
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

func (c *LingshulianController) List(ctx context.Context, bucket string, marker *string, limit int64) (*s3.ListObjectsOutput, error) {
	return c.Svc.ListObjects(&s3.ListObjectsInput{
		Bucket:  util.StringPtr(bucket),
		Marker:  marker,
		MaxKeys: util2.Int64Ptr(limit),
	})
}

func (c *LingshulianController) PostCreateMultipartUploadInfo(ctx context.Context, request *controller.PostCreateMultipartUploadInfoRequest) (*controller.PostCreateMultipartUploadInfoResp, error) {
	id := util2.MustGenerateID(ctx)
	var key string
	if request.Suffix != nil { /// 如果有后缀, 则使用后缀, 并且自动生成 id
		key = fmt.Sprintf("%d.%s", id, *request.Suffix)
	} else if request.Key != nil { /// 如果指定了 key, 则使用指定的 key
		key = *request.Key
	}

	if key == "" {
		data, err := jsoniter.MarshalToString(request)
		if err != nil {
			return nil, err
		}
		return nil, errors3.New(fmt.Sprintf("PostCreateMultipartUploadInfo key is nil. request-->: %s", data))
	}

	expireTo := time.Now().Add(time.Hour)
	output, err := c.Svc.CreateMultipartUpload(&s3.CreateMultipartUploadInput{
		Bucket:  util.StringPtr(request.Bucket),
		Key:     util.StringPtr(key),
		Expires: aws.Time(expireTo),
	})
	if err != nil {
		return nil, err
	}
	if output.UploadId == nil {
		return nil, errors.NewDefaultAPIError(401, 401, "ErrGetUploadIDFail", "UploadID 获取失败")
	}
	return &controller.PostCreateMultipartUploadInfoResp{
		UploadID: *output.UploadId,
		Key:      key,
		Bucket:   request.Bucket,
		ExpireTo: expireTo.Unix(),
	}, nil
}

func (c *LingshulianController) PostMultipartUploadPartInfo(ctx context.Context, request *controller.PostMultipartUploadPartInfoRequest) (*controller.PostMultipartUploadPartInfoResp, error) {
	expire := 60 * time.Minute
	urlData := make([]string, request.PartNumber)
	for i := 0; i < int(request.PartNumber); i++ {
		req, _ := c.Svc.UploadPartRequest(&s3.UploadPartInput{
			UploadId:   util.StringPtr(request.UploadID),
			Bucket:     util.StringPtr(request.Bucket),
			Key:        util.StringPtr(request.Key),
			PartNumber: util2.Int64Ptr(int64(i + 1)),
		})
		URL, err := req.Presign(expire)
		if err != nil {
			return nil, err
		}
		urlData[i] = URL
	}
	return &controller.PostMultipartUploadPartInfoResp{
		URLData:  urlData,
		ExpireTo: time.Now().Add(expire).Unix(),
	}, nil
}

func (c *LingshulianController) PostCompleteMultipartUploadInfo(ctx context.Context, request *controller.PostCompleteMultipartUploadInfoRequest) (*controller.PostCompleteMultipartUploadInfoResp, error) {
	parts := make([]*s3.CompletedPart, 0)
	for _, part := range request.Parts {
		parts = append(parts, &s3.CompletedPart{
			ETag:       util.StringPtr(part.ETag),
			PartNumber: util2.Int64Ptr(part.PartNumber),
		})
	}
	output, err := c.Svc.CompleteMultipartUpload(&s3.CompleteMultipartUploadInput{
		UploadId: util.StringPtr(request.UploadID),
		Key:      util.StringPtr(request.Key),
		Bucket:   util.StringPtr(request.Bucket),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: parts,
		},
	})
	if err != nil {
		return nil, err
	}
	if output.Key == nil {
		return nil, errors.NewDefaultAPIError(401, 401, "ErrGetKeyFail", "Key 获取失败")
	}
	return &controller.PostCompleteMultipartUploadInfoResp{
		Key: request.Key,
	}, nil
}

func (c *LingshulianController) PostAbortMultipartUploadInfo(ctx context.Context, request *controller.PostAbortMultipartUploadPartInfoRequest) (*controller.PostAbortMultipartUploadPartInfoResponse, error) {
	output, err := c.Svc.AbortMultipartUpload(&s3.AbortMultipartUploadInput{
		Bucket:   util.StringPtr(request.Bucket),
		Key:      util.StringPtr(request.Key),
		UploadId: util.StringPtr(request.UploadID),
	})
	if err != nil {
		return nil, err
	}
	return &controller.PostAbortMultipartUploadPartInfoResponse{
		RequestCharged: output.RequestCharged,
	}, nil
}

func (c *LingshulianController) GetSignatureURL(ctx context.Context, bucket, key string, expired time.Duration) (string, error) {
	request, _ := c.Svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: util.StringPtr(bucket),
		Key:    util.StringPtr(key),
	})
	return request.Presign(expired)
}
