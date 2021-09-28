package impl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	errors2 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentController struct {
	client *cos.Client

	credential *common.Credential
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
		credential: common.NewCredential(config.DumpConfig.AppConfig.TencentCosSecretID, config.DumpConfig.AppConfig.TencentCosSecretKey),
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
	res, err := c.client.Object.GetPresignedURL(ctx, http.MethodGet, name, config.DumpConfig.AppConfig.TencentCosSecretID, config.DumpConfig.AppConfig.TencentCosSecretKey, 5*time.Hour, nil)
	util.PanicIf(err)
	return res.String(), nil
}

func (c *TencentController) SendPhoneRegisterCaptcha(ctx context.Context, captcha, phone string) error {
	client, err := sms.NewClient(c.credential, regions.Beijing, profile.NewClientProfile())
	if err != nil {
		return err
	}
	/// 发送短信实例
	request := sms.NewSendSmsRequest()
	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名，例如：腾讯云，签名信息可登录 [短信控制台](https://console.cloud.tencent.com/smsv2) 查看。
	// <dx-alert infotype="notice" title="注意">国内短信为必填参数。</dx-alert>
	request.SignName = util.StringPtr(config.DumpConfig.AppConfig.TencentSMSSignName)
	/// 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666
	request.SmsSdkAppId = util.StringPtr(config.DumpConfig.AppConfig.TencentSMSAppSDKID)
	/// 模板 ID: 必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台] 查看
	request.TemplateId = util.StringPtr(config.DumpConfig.AppConfig.TencentSMSTemplateID)
	/// 模板参数: 若无模板参数，则设置为空
	request.TemplateParamSet = common.StringPtrs([]string{captcha, "15"})
	request.PhoneNumberSet = common.StringPtrs([]string{fmt.Sprintf("+86%s", phone)})
	_, err = client.SendSms(request)
	if _, ok := err.(*errors2.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
