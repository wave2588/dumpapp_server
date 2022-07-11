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

	/// 公共字段
	OriginalP12Password string `json:"original_p12_password"` /// 原本的密码
	NewP12Password      string `json:"new_p12_password"`      /// 新密码
}

func (d *CertificateBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}

/// install_app
type InstallAppCDKEYOrderBizExt struct {
	ContactWay string `json:"contact_way"`
	IsTest     bool   `json:"is_test"` /// 是否是测试或者后台添加生成的订单
}

func (d *InstallAppCDKEYOrderBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}

/// member_pay_order
type MemberPayOrderBizExt struct {
	Platform enum.MemberPayOrderPlatform `json:"platform"`
}

func (d *MemberPayOrderBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}
