package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"dumpapp_server/pkg/common/util"
)

type IpaVersionBizExt struct {
	DescribeURL *string `json:"describe_url,omitempty"`
	Describe    *string `json:"describe,omitempty"`
	Storage     string  `json:"storage"` // cos lingshulian
	Size        int64   `json:"size"`    // ipa 大小
	Country     string  `json:"country"` // 国家
}

func (i IpaVersionBizExt) String() string {
	data, err := json.Marshal(i)
	util.PanicIf(err)
	return string(data)
}

/// 写数据会走到这里
func (i IpaVersionBizExt) Value() (driver.Value, error) {
	return i.String(), nil
}

/// 读数据会走到这里
func (i *IpaVersionBizExt) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of IpaVersionBizExt: %[1]T(%[1]v)", value)
	}
	var data IpaVersionBizExt
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return err
	}
	*i = data
	return nil
}
