package config

import (
	"fmt"
	"log"
	"os"
)

var DumpConfig *dumpConfig

func init() {
	DumpConfig = getDumpConfig()
}

func getDumpConfig() *dumpConfig {
	log.Println("get dump config...")
	cfg := &dumpConfig{}

	currentENV := os.Getenv("DumpEnv")
	fmt.Println("currentENV---> ", currentENV)
	cfg.AppConfig = &config
	return cfg
}

type dumpConfig struct {
	AppConfig *appConfig
}

var config = appConfig{
	MySQL: &MySQL{
		Master:   os.Getenv("MYSQL_MASTER"),
		Slaves:   []string{os.Getenv("MYSQL_SLAVE")},
		Offlines: []string{os.Getenv("MYSQL_OFFLINE")},
	},
	Redis: &Redis{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	},
	Cookie: &Cookie{HMACKey: os.Getenv("COOKIE_HMAC_KEY"), AESKey: os.Getenv("COOKIE_AES_KEY")},

	Env: DumpEnv(os.Getenv("DumpEnv")),

	SentryDSN: os.Getenv("SENTRY_DSN"),

	TencentCOSAppID:             os.Getenv("TENCENT_COS_APP_ID"),
	TencentCosBucketName:        os.Getenv("TENCENT_COS_BUCKET_NAME"),
	TencentCosBucketRegion:      os.Getenv("TENCENT_COS_BUCKET_REGION"),
	TencentCosSecretID:          os.Getenv("TENCENT_COS_SECRET_ID"),
	TencentCosSecretKey:         os.Getenv("TENCENT_COS_SECRET_KEY"),
	TencentCosIpaHost:           os.Getenv("TENCENT_COS_IPA_HOST"),
	TencentCosSignIpaHost:       os.Getenv("TENCENT_COS_SIGN_IPA_HOST"),
	TencentCosSignIpaBucketName: os.Getenv("TENCENT_COS_SIGN_IPA_BUCKET_NAME"),

	TencentSMSAppSDKID:   os.Getenv("TENCENT_SMS_APP_SKD_ID"),
	TencentSMSTemplateID: os.Getenv("TENCENT_SMS_TEMPLATE_ID"),
	TencentSMSSignName:   os.Getenv("TENCENT_SMS_SIGN_NAME"),

	TencentGroupKey: os.Getenv("TENCENT_GROUP_KEY"),

	ALiPayDumpAppID:             os.Getenv("ALIPAY_DUMP_APP_ID"),
	ALiPayDumpPublicKey:         os.Getenv("ALIPAY_DUMP_PUBLIC_KEY"),
	ALiPayDumpPrivateKey:        os.Getenv("ALIPAY_DUMP_PRIVATE_KEY"),
	ALiPayPublicKey:             os.Getenv("ALIPAY_PUBLIC_KEY"),
	ALiPayNotifyURLV3:           os.Getenv("ALIPAY_NOTIFY_URL_v3"),
	ALiPayNotifyURLByInstallApp: os.Getenv("ALIPAY_NOTIFY_URL_BY_INSTALL_APP"),

	DumpAppEmail:             os.Getenv("DUMP_APP_EMAIL"),
	DumpAppFromEmail:         os.Getenv("DUMP_APP_FROM_EMAIL"),
	DumpAppFromEmailPassword: os.Getenv("DUMP_APP_FROM_EMAIL_PASSWORD"),

	DumpAppRegisterFromEmail:         os.Getenv("DUMP_APP_REGISTER_FROM_EMAIL"),
	DumpAppRegisterFromEmailPassword: os.Getenv("DUMP_APP_REGISTER_FROM_EMAIL_PASSWORD"),

	/// cer server
	CerCreateURL:        os.Getenv("CER_CREATE_URL"),
	CerCheckP12URL:      os.Getenv("CER_CHECK_P12_URL"),
	CerCheckValidateURL: os.Getenv("CER_CHECK_VALIDATE"),
	CerServerToken:      os.Getenv("CER_SERVER_TOKEN"),

	/// cer v2 server
	CerServerTokenV2:   os.Getenv("CER_SERVER_TOKEN_V2"),
	CerCreateURLV2:     os.Getenv("CER_CREATE_URL_V2"),
	CerGetDeviceListV2: os.Getenv("CER_GET_DEVICE_LIST_V2"),
	CerGetV2:           os.Getenv("CER_GET_CER_V2"),
}

type appConfig struct {
	MySQL  *MySQL
	Redis  *Redis
	Cookie *Cookie

	Env DumpEnv

	SentryDSN string

	TencentCOSAppID             string
	TencentCosBucketName        string
	TencentCosBucketRegion      string
	TencentCosSecretID          string
	TencentCosSecretKey         string
	TencentCosIpaHost           string
	TencentCosSignIpaHost       string
	TencentCosSignIpaBucketName string

	TencentSMSAppSDKID   string
	TencentSMSTemplateID string
	TencentSMSSignName   string

	TencentGroupKey string

	ALiPayDumpAppID      string
	ALiPayDumpPublicKey  string
	ALiPayDumpPrivateKey string
	ALiPayPublicKey      string
	ALiPayNotifyURLV3    string

	/// 安装 app 支付订单回调
	ALiPayNotifyURLByInstallApp string

	DumpAppEmail             string
	DumpAppFromEmail         string
	DumpAppFromEmailPassword string

	/// 注册专有邮箱
	DumpAppRegisterFromEmail         string
	DumpAppRegisterFromEmailPassword string

	/// cer v1
	CerCreateURL        string
	CerCheckP12URL      string
	CerCheckValidateURL string
	CerServerToken      string

	/// cer v2
	CerServerTokenV2   string
	CerCreateURLV2     string
	CerGetDeviceListV2 string
	CerGetV2           string
}

type MySQL struct {
	Master   string
	Slaves   []string
	Offlines []string
}

type Redis struct {
	Addr     string /// redis 地址, 带端口, localhost:6379
	Password string
}

type Cookie struct {
	HMACKey string
	AESKey  string
}

type DumpEnv string

const (
	DumpEnvProduction = "production"
	DumpEnvDev        = "dev"
)
