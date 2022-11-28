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

	LingshulianSecretID:            os.Getenv("LINGSHULIAN_SECRET_ID"),
	LingshulianSecretKey:           os.Getenv("LINGSHULIAN_SECRET_KEY"),
	LingshulianMemberSignIpaBucket: os.Getenv("LINGSHULIAN_MEMBER_SIGN_IPA_BUCKET"),
	LingshulianMemberSignIpaHost:   os.Getenv("LINGSHULIAN_MEMBER_SIGN_IPA_HOST"),
	LingshulianShareIpaBucket:      os.Getenv("LINGSHULIAN_SHARE_IPA_BUCKET"),

	// ALiPayDumpAppID:             os.Getenv("ALIPAY_DUMP_APP_ID"),
	ALiPayDumpAppID: "2021002145649331",
	// ALiPayDumpPrivateKey: os.Getenv("ALIPAY_DUMP_PRIVATE_KEY"),
	ALiPayDumpPrivateKey: "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCIl7pUz/A9zStduW7pZpI11+FbRPdebf4Vesof50AwDfxx3RJbWPp2d0QdLKdkKx94V+A9g+ufrAFeNrThIXGcAqKUsLQG3Zqqut3AXes4B5FLgY/WEPQmo8ULVvVwqitBB5AbRqKNOhEQ7ztXDvnF80itDK4OMtPmiSkPqeuQQxiPvY9onqcA1W1JKIM5sfwfHnR6ErlFxqeGpM7ye0htvHAhiEjnKIpgZepl2e1CBEx95l1mPpzV4LPWVUItOHPETNfjvF3yf+Pgj+oFZsA7PiCiSNXMMSKnhFFXB2eF6RW3XpMeuF+pkGn59e3+A/hoJh4ioeZSfTCh8BPtrLqxAgMBAAECggEAdCYRG+11rBe6uIfl+DJwQbjAuIt8jZ+aX0l6doZ93l8GOwxxf5u53uKr2OMPs+23ZO3UkHqM8cmhoCuEw6tcn/zdovJfIzdtPaCIz/sM+Sf2NO6HENB5zCGWfH4FVzhcb3+u2oLF1rF5sZy8KNtjKSAmvH/7wbcW2QHpGJi05oXk5EoecFYHuzAH7RtF9uFeUByT1XVUpYS4hzge4uWy479KhQxwEe/fzThs5ihF9WivwEcrUcmKr8jJl/kn7MWZTZS5RjYD6k13k/pgWbWt565UioxMcsnZepKkO0diZKPyx0y8GfKqTlXXTNxwtFrbb33lUKbOieC6l+7GOd7MwQKBgQDdQgwXoPjigkyffpzi8VjfCkPHxROzPIOwfptwMec/tIVIm9BWixyrQhGlEGR+XLkbOFQN7i4YdAiUgnv4Sx9pQ/GrY5jtZeaJzgR+IkMFrrlcawjsVnWUX+L9V4bagDyIwcxp23gd1iQpwLdmHo3+DKaGXPGBYlsgA+QpiBQQBwKBgQCeCl81gittPo3JnWbcsImGsaciE7IF1t8DnldAfkKkpZS6SBf8zdMMVZAzQmYEHGzCGUr2WJWRbB5Fi0FKl0lNcfWQxxGt/l7WWqOVP7jAXhWBSkhCQH1kXVSZidlZcYrtVuMRo3fNsQY1CgpIlfs78fu2SAWoaGHXnX0v9iXBhwKBgHgS1dkk5KyYJdkQrzeB3sb8HRiW3UASAS2RJ+3VRzgNUZ+73254BFD0g/reUIs66sHY6dS9g4qFvfpKbdirfBp2YvquDFoZSOlUQp/pHBJDZhi/hZIswZaKCveNvoNpwHA/LB3umtsmUW3PRjhHMKvEVcLpQa8Dn4xaUEIxtSGVAoGAVFps3v6Pd0vAGjCtSnXfscj40DN6/armdeP55i5+G6tVaug4BXNGhnrU8Mcr9F3HnwIpBLvbeTcgITZmrw14zqFY1OGsChaPQBI45dyRG/wbtlqTnukVBJDcKuds31S/Nlb989gwhdVK3txxCLUk16YdF/nxKyYrsw4YV5UmKdkCgYEAgsftrXlcVETR7Tov4e2I7A4ef8njUWmdu4WG1ZTjyfiRmB0b6Xi5g0zVQRRsj9ZU8KPHYxREdxtMfZONGYwmTOiPuysU+sf3MD/uJrbOsNXtvjtyUxJVTU4KChYs1JGOZCc39MA8EvlOvmkC0C7e5nNqH/tV5uPzOE2k3hdkIxA=",
	// ALiPayPublicKey:             os.Getenv("ALIPAY_PUBLIC_KEY"),
	ALiPayPublicKey:             "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiJe6VM/wPc0rXblu6WaSNdfhW0T3Xm3+FXrKH+dAMA38cd0SW1j6dndEHSynZCsfeFfgPYPrn6wBXja04SFxnAKilLC0Bt2aqrrdwF3rOAeRS4GP1hD0JqPFC1b1cKorQQeQG0aijToREO87Vw75xfNIrQyuDjLT5okpD6nrkEMYj72PaJ6nANVtSSiDObH8Hx50ehK5RcanhqTO8ntIbbxwIYhI5yiKYGXqZdntQgRMfeZdZj6c1eCz1lVCLThzxEzX47xd8n/j4I/qBWbAOz4gokjVzDEip4RRVwdnhekVt16THrhfqZBp+fXt/gP4aCYeIqHmUn0wofAT7ay6sQIDAQAB",
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

	/// cer v3 server
	CerServerTokenV3: os.Getenv("CER_SERVER_TOKEN_V3"),
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

	/// lingshulian
	LingshulianSecretID            string
	LingshulianSecretKey           string
	LingshulianMemberSignIpaBucket string
	LingshulianMemberSignIpaHost   string
	LingshulianShareIpaBucket      string

	ALiPayDumpAppID      string
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

	/// cer v3
	CerServerTokenV3 string
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
