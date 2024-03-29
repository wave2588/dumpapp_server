// Code generated by "enumer -type=CertificateSource -json -sql -transform=snake -trimprefix=CertificateSource"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _CertificateSourceName = "v1v2v3"

var _CertificateSourceIndex = [...]uint8{0, 2, 4, 6}

const _CertificateSourceLowerName = "v1v2v3"

func (i CertificateSource) String() string {
	i -= 1
	if i < 0 || i >= CertificateSource(len(_CertificateSourceIndex)-1) {
		return fmt.Sprintf("CertificateSource(%d)", i+1)
	}
	return _CertificateSourceName[_CertificateSourceIndex[i]:_CertificateSourceIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _CertificateSourceNoOp() {
	var x [1]struct{}
	_ = x[CertificateSourceV1-(1)]
	_ = x[CertificateSourceV2-(2)]
	_ = x[CertificateSourceV3-(3)]
}

var _CertificateSourceValues = []CertificateSource{CertificateSourceV1, CertificateSourceV2, CertificateSourceV3}

var _CertificateSourceNameToValueMap = map[string]CertificateSource{
	_CertificateSourceName[0:2]:      CertificateSourceV1,
	_CertificateSourceLowerName[0:2]: CertificateSourceV1,
	_CertificateSourceName[2:4]:      CertificateSourceV2,
	_CertificateSourceLowerName[2:4]: CertificateSourceV2,
	_CertificateSourceName[4:6]:      CertificateSourceV3,
	_CertificateSourceLowerName[4:6]: CertificateSourceV3,
}

var _CertificateSourceNames = []string{
	_CertificateSourceName[0:2],
	_CertificateSourceName[2:4],
	_CertificateSourceName[4:6],
}

// CertificateSourceString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CertificateSourceString(s string) (CertificateSource, error) {
	if val, ok := _CertificateSourceNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _CertificateSourceNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CertificateSource values", s)
}

// CertificateSourceValues returns all values of the enum
func CertificateSourceValues() []CertificateSource {
	return _CertificateSourceValues
}

// CertificateSourceStrings returns a slice of all String values of the enum
func CertificateSourceStrings() []string {
	strs := make([]string, len(_CertificateSourceNames))
	copy(strs, _CertificateSourceNames)
	return strs
}

// IsACertificateSource returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CertificateSource) IsACertificateSource() bool {
	for _, v := range _CertificateSourceValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CertificateSource
func (i CertificateSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CertificateSource
func (i *CertificateSource) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CertificateSource should be a string, got %s", data)
	}

	var err error
	*i, err = CertificateSourceString(s)
	return err
}

func (i CertificateSource) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CertificateSource) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of CertificateSource: %[1]T(%[1]v)", value)
	}

	val, err := CertificateSourceString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
