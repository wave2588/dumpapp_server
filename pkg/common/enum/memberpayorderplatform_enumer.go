// Code generated by "enumer -type=MemberPayOrderPlatform -json -sql -transform=snake -trimprefix=MemberPayOrderPlatform"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _MemberPayOrderPlatformName = "webios"

var _MemberPayOrderPlatformIndex = [...]uint8{0, 3, 6}

const _MemberPayOrderPlatformLowerName = "webios"

func (i MemberPayOrderPlatform) String() string {
	i -= 1
	if i < 0 || i >= MemberPayOrderPlatform(len(_MemberPayOrderPlatformIndex)-1) {
		return fmt.Sprintf("MemberPayOrderPlatform(%d)", i+1)
	}
	return _MemberPayOrderPlatformName[_MemberPayOrderPlatformIndex[i]:_MemberPayOrderPlatformIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _MemberPayOrderPlatformNoOp() {
	var x [1]struct{}
	_ = x[MemberPayOrderPlatformWeb-(1)]
	_ = x[MemberPayOrderPlatformIOS-(2)]
}

var _MemberPayOrderPlatformValues = []MemberPayOrderPlatform{MemberPayOrderPlatformWeb, MemberPayOrderPlatformIOS}

var _MemberPayOrderPlatformNameToValueMap = map[string]MemberPayOrderPlatform{
	_MemberPayOrderPlatformName[0:3]:      MemberPayOrderPlatformWeb,
	_MemberPayOrderPlatformLowerName[0:3]: MemberPayOrderPlatformWeb,
	_MemberPayOrderPlatformName[3:6]:      MemberPayOrderPlatformIOS,
	_MemberPayOrderPlatformLowerName[3:6]: MemberPayOrderPlatformIOS,
}

var _MemberPayOrderPlatformNames = []string{
	_MemberPayOrderPlatformName[0:3],
	_MemberPayOrderPlatformName[3:6],
}

// MemberPayOrderPlatformString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func MemberPayOrderPlatformString(s string) (MemberPayOrderPlatform, error) {
	if val, ok := _MemberPayOrderPlatformNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _MemberPayOrderPlatformNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to MemberPayOrderPlatform values", s)
}

// MemberPayOrderPlatformValues returns all values of the enum
func MemberPayOrderPlatformValues() []MemberPayOrderPlatform {
	return _MemberPayOrderPlatformValues
}

// MemberPayOrderPlatformStrings returns a slice of all String values of the enum
func MemberPayOrderPlatformStrings() []string {
	strs := make([]string, len(_MemberPayOrderPlatformNames))
	copy(strs, _MemberPayOrderPlatformNames)
	return strs
}

// IsAMemberPayOrderPlatform returns "true" if the value is listed in the enum definition. "false" otherwise
func (i MemberPayOrderPlatform) IsAMemberPayOrderPlatform() bool {
	for _, v := range _MemberPayOrderPlatformValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for MemberPayOrderPlatform
func (i MemberPayOrderPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for MemberPayOrderPlatform
func (i *MemberPayOrderPlatform) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MemberPayOrderPlatform should be a string, got %s", data)
	}

	var err error
	*i, err = MemberPayOrderPlatformString(s)
	return err
}

func (i MemberPayOrderPlatform) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *MemberPayOrderPlatform) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of MemberPayOrderPlatform: %[1]T(%[1]v)", value)
	}

	val, err := MemberPayOrderPlatformString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
