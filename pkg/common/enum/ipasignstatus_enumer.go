// Code generated by "enumer -type=IpaSignStatus -json -sql -transform=snake -trimprefix=IpaSignStatus"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _IpaSignStatusName = "unprocessedprocessingsuccessfail"

var _IpaSignStatusIndex = [...]uint8{0, 11, 21, 28, 32}

const _IpaSignStatusLowerName = "unprocessedprocessingsuccessfail"

func (i IpaSignStatus) String() string {
	i -= 1
	if i < 0 || i >= IpaSignStatus(len(_IpaSignStatusIndex)-1) {
		return fmt.Sprintf("IpaSignStatus(%d)", i+1)
	}
	return _IpaSignStatusName[_IpaSignStatusIndex[i]:_IpaSignStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _IpaSignStatusNoOp() {
	var x [1]struct{}
	_ = x[IpaSignStatusUnprocessed-(1)]
	_ = x[IpaSignStatusProcessing-(2)]
	_ = x[IpaSignStatusSuccess-(3)]
	_ = x[IpaSignStatusFail-(4)]
}

var _IpaSignStatusValues = []IpaSignStatus{IpaSignStatusUnprocessed, IpaSignStatusProcessing, IpaSignStatusSuccess, IpaSignStatusFail}

var _IpaSignStatusNameToValueMap = map[string]IpaSignStatus{
	_IpaSignStatusName[0:11]:       IpaSignStatusUnprocessed,
	_IpaSignStatusLowerName[0:11]:  IpaSignStatusUnprocessed,
	_IpaSignStatusName[11:21]:      IpaSignStatusProcessing,
	_IpaSignStatusLowerName[11:21]: IpaSignStatusProcessing,
	_IpaSignStatusName[21:28]:      IpaSignStatusSuccess,
	_IpaSignStatusLowerName[21:28]: IpaSignStatusSuccess,
	_IpaSignStatusName[28:32]:      IpaSignStatusFail,
	_IpaSignStatusLowerName[28:32]: IpaSignStatusFail,
}

var _IpaSignStatusNames = []string{
	_IpaSignStatusName[0:11],
	_IpaSignStatusName[11:21],
	_IpaSignStatusName[21:28],
	_IpaSignStatusName[28:32],
}

// IpaSignStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func IpaSignStatusString(s string) (IpaSignStatus, error) {
	if val, ok := _IpaSignStatusNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _IpaSignStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to IpaSignStatus values", s)
}

// IpaSignStatusValues returns all values of the enum
func IpaSignStatusValues() []IpaSignStatus {
	return _IpaSignStatusValues
}

// IpaSignStatusStrings returns a slice of all String values of the enum
func IpaSignStatusStrings() []string {
	strs := make([]string, len(_IpaSignStatusNames))
	copy(strs, _IpaSignStatusNames)
	return strs
}

// IsAIpaSignStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i IpaSignStatus) IsAIpaSignStatus() bool {
	for _, v := range _IpaSignStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for IpaSignStatus
func (i IpaSignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for IpaSignStatus
func (i *IpaSignStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("IpaSignStatus should be a string, got %s", data)
	}

	var err error
	*i, err = IpaSignStatusString(s)
	return err
}

func (i IpaSignStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *IpaSignStatus) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of IpaSignStatus: %[1]T(%[1]v)", value)
	}

	val, err := IpaSignStatusString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}