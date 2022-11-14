package constant

import (
	"encoding/json"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
)

type TransactionKey = string

const TransactionKeyTxn TransactionKey = "txn"

type AdminDumpOrderBizExt struct {
	IpaName         string  `json:"ipa_name"`
	IpaBundleID     string  `json:"ipa_bundle_id"`
	IpaAppStoreLink string  `json:"ipa_app_store_link"`
	DemanderIDs     []int64 `json:"demander_ids"`
	IsOld           bool    `json:"is_old"`
}

type SearchCount struct {
	IpaID int64
	Name  string
	Count int64
}

type IpaVersionBizExt struct {
	DescribeURL *string `json:"describe_url,omitempty"`
	Describe    *string `json:"describe,omitempty"`
	Storage     string  `json:"storage"` /// cos lingshulian
}

func (d *IpaVersionBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}

type IpaSignBizExt struct {
	IpaVersionID int64        `json:"ipa_version_id"`
	IpaVersion   string       `json:"ipa_version"`
	IpaType      enum.IpaType `json:"ipa_type"`
	TokenPath    string       `json:"token_path"`
}

func (d *IpaSignBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}

type CertificateBizExt struct {
	/// v1 扩展字段
	V1UDIDBatchNo string `json:"v1_udid_batch_no,omitempty"`
	V1CerAppleID  string `json:"v1_cer_apple_id,omitempty"`

	/// v2 扩展字段
	V2DeviceID string `json:"v2_device_id,omitempty"`

	/// v3 扩展字段
	V3DeviceID string `json:"v3_device_id,omitempty"`

	/// 公共字段
	OriginalP12Password string `json:"original_p12_password"` /// 原本的密码
	NewP12Password      string `json:"new_p12_password"`      /// 新密码
	IsReplenish         bool   `json:"is_replenish"`          /// 是否是补充证书
	Level               int    `json:"level"`                 /// 1: 普通版   2: 高级版  3: 豪华版
	Note                string `json:"note"`                  /// 证书备注
}

func (d *CertificateBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}
