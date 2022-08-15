package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

type TencentCosHandler struct{}

func NewTencentCosHandler() *TencentCosHandler {
	return &TencentCosHandler{}
}

func (h *TencentCosHandler) Get(w http.ResponseWriter, r *http.Request) {
	appid := config.DumpConfig.AppConfig.TencentCOSAppID
	bucket := config.DumpConfig.AppConfig.TencentCosBucketName
	region := config.DumpConfig.AppConfig.TencentCosBucketRegion
	c := sts.NewClient(
		config.DumpConfig.AppConfig.TencentCosSecretID,
		config.DumpConfig.AppConfig.TencentCosSecretKey,
		nil,
	)

	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-beijing",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						// 简单上传
						"name/cos:PutObject",
						// 表单上传、小程序上传
						"name/cos:PostObject",
						// 分片上传
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/*",
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	util.PanicIf(err)

	util.RenderJSON(w, res)
}

func (h *TencentCosHandler) GetSignIpa(w http.ResponseWriter, r *http.Request) {
	appid := config.DumpConfig.AppConfig.TencentCOSAppID
	bucket := config.DumpConfig.AppConfig.TencentCosSignIpaBucketName
	region := config.DumpConfig.AppConfig.TencentCosBucketRegion
	c := sts.NewClient(
		config.DumpConfig.AppConfig.TencentCosSecretID,
		config.DumpConfig.AppConfig.TencentCosSecretKey,
		nil,
	)
	ss := "qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/*"
	fmt.Println(ss)

	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-beijing",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						// 简单上传
						"name/cos:PutObject",
						// 表单上传、小程序上传
						"name/cos:PostObject",
						// 分片上传
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/*",
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	util.PanicIf(err)

	util.RenderJSON(w, res)
}
