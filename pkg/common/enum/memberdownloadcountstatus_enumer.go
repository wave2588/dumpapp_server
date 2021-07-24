// Code generated by "enumer -type=MemberDownloadCountStatus -json -sql -transform=snake -trimprefix=MemberDownloadCountStatus"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _MemberDownloadCountStatusName = "normalused"

var _MemberDownloadCountStatusIndex = [...]uint8{0, 6, 10}

const _MemberDownloadCountStatusLowerName = "normalused"

func (i MemberDownloadCountStatus) String() string {
	i -= 1
	if i < 0 || i >= MemberDownloadCountStatus(len(_MemberDownloadCountStatusIndex)-1) {
		return fmt.Sprintf("MemberDownloadCountStatus(%d)", i+1)
	}
	return _MemberDownloadCountStatusName[_MemberDownloadCountStatusIndex[i]:_MemberDownloadCountStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _MemberDownloadCountStatusNoOp() {
	var x [1]struct{}
	_ = x[MemberDownloadCountStatusNormal-(1)]
	_ = x[MemberDownloadCountStatusUsed-(2)]
}

var _MemberDownloadCountStatusValues = []MemberDownloadCountStatus{MemberDownloadCountStatusNormal, MemberDownloadCountStatusUsed}

var _MemberDownloadCountStatusNameToValueMap = map[string]MemberDownloadCountStatus{
	_MemberDownloadCountStatusName[0:6]:       MemberDownloadCountStatusNormal,
	_MemberDownloadCountStatusLowerName[0:6]:  MemberDownloadCountStatusNormal,
	_MemberDownloadCountStatusName[6:10]:      MemberDownloadCountStatusUsed,
	_MemberDownloadCountStatusLowerName[6:10]: MemberDownloadCountStatusUsed,
}

var _MemberDownloadCountStatusNames = []string{
	_MemberDownloadCountStatusName[0:6],
	_MemberDownloadCountStatusName[6:10],
}

// MemberDownloadCountStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func MemberDownloadCountStatusString(s string) (MemberDownloadCountStatus, error) {
	if val, ok := _MemberDownloadCountStatusNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _MemberDownloadCountStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to MemberDownloadCountStatus values", s)
}

// MemberDownloadCountStatusValues returns all values of the enum
func MemberDownloadCountStatusValues() []MemberDownloadCountStatus {
	return _MemberDownloadCountStatusValues
}

// MemberDownloadCountStatusStrings returns a slice of all String values of the enum
func MemberDownloadCountStatusStrings() []string {
	strs := make([]string, len(_MemberDownloadCountStatusNames))
	copy(strs, _MemberDownloadCountStatusNames)
	return strs
}

// IsAMemberDownloadCountStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i MemberDownloadCountStatus) IsAMemberDownloadCountStatus() bool {
	for _, v := range _MemberDownloadCountStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for MemberDownloadCountStatus
func (i MemberDownloadCountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for MemberDownloadCountStatus
func (i *MemberDownloadCountStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MemberDownloadCountStatus should be a string, got %s", data)
	}

	var err error
	*i, err = MemberDownloadCountStatusString(s)
	return err
}

func (i MemberDownloadCountStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *MemberDownloadCountStatus) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of MemberDownloadCountStatus: %[1]T(%[1]v)", value)
	}

	val, err := MemberDownloadCountStatusString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
