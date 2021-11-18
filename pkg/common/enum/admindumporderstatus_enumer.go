// Code generated by "enumer -type=AdminDumpOrderStatus -json -sql -transform=snake -trimprefix=AdminDumpOrderStatus"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _AdminDumpOrderStatusName = "progressingprogresseddeleted"

var _AdminDumpOrderStatusIndex = [...]uint8{0, 11, 21, 28}

const _AdminDumpOrderStatusLowerName = "progressingprogresseddeleted"

func (i AdminDumpOrderStatus) String() string {
	i -= 1
	if i < 0 || i >= AdminDumpOrderStatus(len(_AdminDumpOrderStatusIndex)-1) {
		return fmt.Sprintf("AdminDumpOrderStatus(%d)", i+1)
	}
	return _AdminDumpOrderStatusName[_AdminDumpOrderStatusIndex[i]:_AdminDumpOrderStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AdminDumpOrderStatusNoOp() {
	var x [1]struct{}
	_ = x[AdminDumpOrderStatusProgressing-(1)]
	_ = x[AdminDumpOrderStatusProgressed-(2)]
	_ = x[AdminDumpOrderStatusDeleted-(3)]
}

var _AdminDumpOrderStatusValues = []AdminDumpOrderStatus{AdminDumpOrderStatusProgressing, AdminDumpOrderStatusProgressed, AdminDumpOrderStatusDeleted}

var _AdminDumpOrderStatusNameToValueMap = map[string]AdminDumpOrderStatus{
	_AdminDumpOrderStatusName[0:11]:       AdminDumpOrderStatusProgressing,
	_AdminDumpOrderStatusLowerName[0:11]:  AdminDumpOrderStatusProgressing,
	_AdminDumpOrderStatusName[11:21]:      AdminDumpOrderStatusProgressed,
	_AdminDumpOrderStatusLowerName[11:21]: AdminDumpOrderStatusProgressed,
	_AdminDumpOrderStatusName[21:28]:      AdminDumpOrderStatusDeleted,
	_AdminDumpOrderStatusLowerName[21:28]: AdminDumpOrderStatusDeleted,
}

var _AdminDumpOrderStatusNames = []string{
	_AdminDumpOrderStatusName[0:11],
	_AdminDumpOrderStatusName[11:21],
	_AdminDumpOrderStatusName[21:28],
}

// AdminDumpOrderStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AdminDumpOrderStatusString(s string) (AdminDumpOrderStatus, error) {
	if val, ok := _AdminDumpOrderStatusNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _AdminDumpOrderStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AdminDumpOrderStatus values", s)
}

// AdminDumpOrderStatusValues returns all values of the enum
func AdminDumpOrderStatusValues() []AdminDumpOrderStatus {
	return _AdminDumpOrderStatusValues
}

// AdminDumpOrderStatusStrings returns a slice of all String values of the enum
func AdminDumpOrderStatusStrings() []string {
	strs := make([]string, len(_AdminDumpOrderStatusNames))
	copy(strs, _AdminDumpOrderStatusNames)
	return strs
}

// IsAAdminDumpOrderStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AdminDumpOrderStatus) IsAAdminDumpOrderStatus() bool {
	for _, v := range _AdminDumpOrderStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AdminDumpOrderStatus
func (i AdminDumpOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AdminDumpOrderStatus
func (i *AdminDumpOrderStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AdminDumpOrderStatus should be a string, got %s", data)
	}

	var err error
	*i, err = AdminDumpOrderStatusString(s)
	return err
}

func (i AdminDumpOrderStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *AdminDumpOrderStatus) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of AdminDumpOrderStatus: %[1]T(%[1]v)", value)
	}

	val, err := AdminDumpOrderStatusString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
