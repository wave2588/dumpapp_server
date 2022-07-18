package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
)

type MemberPayExpenseRecordBizExt struct {
	CountSource enum.MemberPayCountSource `json:"count_source"`
	OrderID     *int64                    `json:"order_id,omitempty"`

	/// 如果是管理员操作的话，会很有此 ID
	AdminMemberID *int64 `json:"admin_member_id,omitempty"`
}

func (i MemberPayExpenseRecordBizExt) String() string {
	data, err := json.Marshal(i)
	util.PanicIf(err)
	return string(data)
}

/// 写数据会走到这里
func (i MemberPayExpenseRecordBizExt) Value() (driver.Value, error) {
	return i.String(), nil
}

/// 读数据会走到这里
func (i *MemberPayExpenseRecordBizExt) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of MemberPayExpenseRecordBizExt: %[1]T(%[1]v)", value)
	}
	var data MemberPayExpenseRecordBizExt
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return err
	}
	*i = data
	return nil
}
