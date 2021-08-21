// Code generated by "enumer -type=MemberDownloadOrderStatus -json -sql -transform=snake -trimprefix=MemberDownloadOrderStatus"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _MemberDownloadOrderStatusName = "pendingpaid"

var _MemberDownloadOrderStatusIndex = [...]uint8{0, 7, 11}

const _MemberDownloadOrderStatusLowerName = "pendingpaid"

func (i MemberDownloadOrderStatus) String() string {
	i -= 1
	if i < 0 || i >= MemberDownloadOrderStatus(len(_MemberDownloadOrderStatusIndex)-1) {
		return fmt.Sprintf("MemberDownloadOrderStatus(%d)", i+1)
	}
	return _MemberDownloadOrderStatusName[_MemberDownloadOrderStatusIndex[i]:_MemberDownloadOrderStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _MemberDownloadOrderStatusNoOp() {
	var x [1]struct{}
	_ = x[MemberDownloadOrderStatusPending-(1)]
	_ = x[MemberDownloadOrderStatusPaid-(2)]
}

var _MemberDownloadOrderStatusValues = []MemberDownloadOrderStatus{MemberDownloadOrderStatusPending, MemberDownloadOrderStatusPaid}

var _MemberDownloadOrderStatusNameToValueMap = map[string]MemberDownloadOrderStatus{
	_MemberDownloadOrderStatusName[0:7]:       MemberDownloadOrderStatusPending,
	_MemberDownloadOrderStatusLowerName[0:7]:  MemberDownloadOrderStatusPending,
	_MemberDownloadOrderStatusName[7:11]:      MemberDownloadOrderStatusPaid,
	_MemberDownloadOrderStatusLowerName[7:11]: MemberDownloadOrderStatusPaid,
}

var _MemberDownloadOrderStatusNames = []string{
	_MemberDownloadOrderStatusName[0:7],
	_MemberDownloadOrderStatusName[7:11],
}

// MemberDownloadOrderStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func MemberDownloadOrderStatusString(s string) (MemberDownloadOrderStatus, error) {
	if val, ok := _MemberDownloadOrderStatusNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _MemberDownloadOrderStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to MemberDownloadOrderStatus values", s)
}

// MemberDownloadOrderStatusValues returns all values of the enum
func MemberDownloadOrderStatusValues() []MemberDownloadOrderStatus {
	return _MemberDownloadOrderStatusValues
}

// MemberDownloadOrderStatusStrings returns a slice of all String values of the enum
func MemberDownloadOrderStatusStrings() []string {
	strs := make([]string, len(_MemberDownloadOrderStatusNames))
	copy(strs, _MemberDownloadOrderStatusNames)
	return strs
}

// IsAMemberDownloadOrderStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i MemberDownloadOrderStatus) IsAMemberDownloadOrderStatus() bool {
	for _, v := range _MemberDownloadOrderStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for MemberDownloadOrderStatus
func (i MemberDownloadOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for MemberDownloadOrderStatus
func (i *MemberDownloadOrderStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MemberDownloadOrderStatus should be a string, got %s", data)
	}

	var err error
	*i, err = MemberDownloadOrderStatusString(s)
	return err
}

func (i MemberDownloadOrderStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *MemberDownloadOrderStatus) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of MemberDownloadOrderStatus: %[1]T(%[1]v)", value)
	}

	val, err := MemberDownloadOrderStatusString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
