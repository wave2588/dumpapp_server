package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"dumpapp_server/pkg/common/util"
)

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

	CdKeyID int64 `json:"cd_key_id"` // 兑换码 id
}

func (i CertificateBizExt) String() string {
	data, err := json.Marshal(i)
	util.PanicIf(err)
	return string(data)
}

/// 写数据会走到这里
func (i CertificateBizExt) Value() (driver.Value, error) {
	return i.String(), nil
}

/// 读数据会走到这里
func (i *CertificateBizExt) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CertificateBizExt: %[1]T(%[1]v)", value)
	}
	var data CertificateBizExt
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return err
	}
	*i = data
	return nil
}
