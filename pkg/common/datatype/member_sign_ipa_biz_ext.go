package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"dumpapp_server/pkg/common/util"
)

type MemberSignIpaBizExt struct {
	IpaName         string `json:"ipa_name"`
	IpaBundleID     string `json:"ipa_bundle_id"`
	IpaVersion      string `json:"ipa_version"`
	IpaSize         int64  `json:"ipa_size"`
	CertificateName string `json:"certificate_name"`
	DispenseCount   int64  `json:"dispense_count"`
}

func (i MemberSignIpaBizExt) String() string {
	data, err := json.Marshal(i)
	util.PanicIf(err)
	return string(data)
}

/// 写数据会走到这里
func (i MemberSignIpaBizExt) Value() (driver.Value, error) {
	return i.String(), nil
}

/// 读数据会走到这里
func (i *MemberSignIpaBizExt) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of MemberSignIpaBizExt: %[1]T(%[1]v)", value)
	}
	var data MemberSignIpaBizExt
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return err
	}
	*i = data
	return nil
}
