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

	SentryDSN: os.Getenv("SENTRY_DSN"),

	TencentCOSAppID:        os.Getenv("TENCENT_COS_APP_ID"),
	TencentCosBucketName:   os.Getenv("TENCENT_COS_BUCKET_NAME"),
	TencentCosBucketRegion: os.Getenv("TENCENT_COS_BUCKET_REGION"),
	TencentCosSecretID:     os.Getenv("TENCENT_COS_SECRET_ID"),
	TencentCosSecretKey:    os.Getenv("TENCENT_COS_SECRET_KEY"),
	TencentCosIpaHost:      os.Getenv("TENCENT_COS_IPA_HOST"),

	TencentSMSAppSDKID:   os.Getenv("TENCENT_SMS_APP_SKD_ID"),
	TencentSMSTemplateID: os.Getenv("TENCENT_SMS_TEMPLATE_ID"),
	TencentSMSSignName:   os.Getenv("TENCENT_SMS_SIGN_NAME"),

	ALiPayDumpAppID:      os.Getenv("ALIPAY_DUMP_APP_ID"),
	ALiPayDumpPublicKey:  os.Getenv("ALIPAY_DUMP_PUBLIC_KEY"),
	ALiPayDumpPrivateKey: os.Getenv("ALIPAY_DUMP_PRIVATE_KEY"),
	ALiPayPublicKey:      os.Getenv("ALIPAY_PUBLIC_KEY"),
	ALiPayNotifyURL:      os.Getenv("ALIPAY_NOTIFY_URL"),
	ALiPayNotifyURLV2:    os.Getenv("ALIPAY_NOTIFY_URL_v2"),

	DumpAppEmail:             os.Getenv("DUMP_APP_EMAIL"),
	DumpAppFromEmail:         os.Getenv("DUMP_APP_FROM_EMAIL"),
	DumpAppFromEmailPassword: os.Getenv("DUMP_APP_FROM_EMAIL_PASSWORD"),
}

type appConfig struct {
	MySQL  *MySQL
	Redis  *Redis
	Cookie *Cookie

	SentryDSN string

	TencentCOSAppID        string
	TencentCosBucketName   string
	TencentCosBucketRegion string
	TencentCosSecretID     string
	TencentCosSecretKey    string
	TencentCosIpaHost      string

	TencentSMSAppSDKID   string
	TencentSMSTemplateID string
	TencentSMSSignName   string

	ALiPayDumpAppID      string
	ALiPayDumpPublicKey  string
	ALiPayDumpPrivateKey string
	ALiPayPublicKey      string
	ALiPayNotifyURL      string
	ALiPayNotifyURLV2    string

	DumpAppEmail             string
	DumpAppFromEmail         string
	DumpAppFromEmailPassword string
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
