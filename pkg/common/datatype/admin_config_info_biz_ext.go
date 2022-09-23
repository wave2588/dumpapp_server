package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
)

type AdminConfigInfoBizExt struct {
	CerSource      enum.CertificateSource `json:"cer_source"`       /// 使用证书平台
	DailyFreeCount int                    `json:"daily_free_count"` /// 每日免费砸壳次数
}

func (i AdminConfigInfoBizExt) String() string {
	data, err := json.Marshal(i)
	util.PanicIf(err)
	return string(data)
}

/// 写数据会走到这里
func (i AdminConfigInfoBizExt) Value() (driver.Value, error) {
	return i.String(), nil
}

/// 读数据会走到这里
func (i *AdminConfigInfoBizExt) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of AdminConfigInfoBizExt: %[1]T(%[1]v)", value)
	}
	var data AdminConfigInfoBizExt
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return err
	}
	*i = data
	return nil
}
