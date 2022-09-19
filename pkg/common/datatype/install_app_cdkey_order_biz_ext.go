package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"dumpapp_server/pkg/common/util"
)

type InstallAppCdkeyOrderBizExt struct {
	ContactWay string `json:"contact_way"`
	IsTest     bool   `json:"is_test"`  /// 是否是测试或者后台添加生成的订单
	IsAgent    bool   `json:"is_agent"` /// 是否是代理商
}

func (i InstallAppCdkeyOrderBizExt) String() string {
	data, err := json.Marshal(i)
	util.PanicIf(err)
	return string(data)
}

/// 写数据会走到这里
func (i InstallAppCdkeyOrderBizExt) Value() (driver.Value, error) {
	return i.String(), nil
}

/// 读数据会走到这里
func (i *InstallAppCdkeyOrderBizExt) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of InstallAppCdkeyOrderBizExt: %[1]T(%[1]v)", value)
	}
	var data InstallAppCdkeyOrderBizExt
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return err
	}
	*i = data
	return nil
}
