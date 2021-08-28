package impl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentController struct {
	client *cos.Client
}

var DefaultTencentController *TencentController

func init() {
	DefaultTencentController = NewTencentController()
}

func NewTencentController() *TencentController {
	//u, err := url.Parse(config.DumpConfig.AppConfig.TencentCosIpaHost)
	/// 删除暂时用这个地址
	u, err := url.Parse(config.DumpConfig.AppConfig.TencentCosIpaHost)
	util.PanicIf(err)
	b := &cos.BaseURL{BucketURL: u}

	return &TencentController{
		client: cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.DumpConfig.AppConfig.TencentCosSecretID,
				SecretKey: config.DumpConfig.AppConfig.TencentCosSecretKey,
			},
		}),
	}
}

func (c *TencentController) DeleteFile(ctx context.Context, TokenPath string) error {
	response, err := c.client.Object.Delete(ctx, TokenPath)
	if err != nil {
		return err
	}
	if response.StatusCode >= 300 {
		return errors.UnproccessableError(fmt.Sprintf("服务器删除文件错误. 错误码: %s", response.Status))
	}
	return nil
}

func (c *TencentController) ListFile(ctx context.Context, marker *string, limit int) (*cos.BucketGetResult, error) {
	options := &cos.BucketGetOptions{
		MaxKeys: limit,
	}
	if marker != nil {
		options.Marker = *marker
	}
	result, response, err := c.client.Bucket.Get(ctx, options)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.UnproccessableError(fmt.Sprintf("服务器删除文件错误. 错误码: %s", response.Status))
	}
	return result, nil
}

func (c *TencentController) GetSignatureURL(ctx context.Context, name string) (string, error) {
	return util.GetImageUrl(name), nil

	//url, err := c.client.Object.GetPresignedURL(ctx, http.MethodGet, name, config.DumpConfig.AppConfig.TencentCosSecretID, config.DumpConfig.AppConfig.TencentCosSecretKey, time.Hour*24, nil)
	//if err != nil {
	//	return "", err
	//}
	//
	//s := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000))
	//i, err := strconv.ParseInt(s, 10, 64)
	//util.PanicIf(err)
	//
	//appID := config.DumpConfig.AppConfig.TencentCOSAppID
	//bucket := config.DumpConfig.AppConfig.TencentCosBucketName
	//secretID := config.DumpConfig.AppConfig.TencentCosSecretID
	//secretKey := config.DumpConfig.AppConfig.TencentCosSecretKey
	////expired := time.Now().Add(60).Unix()
	//current := time.Now().Unix()
	//onceExpired := 0
	//rdm := i
	////fileid := "/200001/newbucket/tencent_test.jpg"
	//fileid := fmt.Sprintf("%s/%s/%s", appID, bucket, name)
	//
	////$once_signature=
	////'a='.$appid.'&b='.$bucket.'&k='.$secret_id.'&e='.$onceExpired.'&t='.$current.'&r='.$rdm.'&f='.$fileid;
	//
	//onceSignature := fmt.Sprintf("a=%s&./src/nocodeb=%s&k=%s&e=%d&t=%d&r=%d&f=%s", appID, bucket, secretID, onceExpired, current, rdm, fileid)
	//
	////$once_signature = base64_encode(hash_hmac('SHA1',$once_signature,$secret_key, true).$once_signature);
	//
	////$once_signature = base64_encode(hash_hmac('SHA1',$once_signature,$secret_key, true).$once_signature);
	//
	//h := hmac.New(sha1.New, []byte(secretKey))
	//h.Write([]byte(onceSignature))
	//// Get result and encode as hexadecimal string
	//sha := hex.EncodeToString(h.Sum(nil))
	//
	//fmt.Println("sha--->: ", sha)
	////expectedMAC := hmac.
	////return hmac.Equal(messageMAC, expectedMAC)
	//
	//return url.String(), nil
}
