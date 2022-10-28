package go_utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"

	"dumpapp_server/pkg/common/util"
	"github.com/pkg/errors"
)

func MarshalToString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func MustMarshalToString(v interface{}) string {
	b, err := json.Marshal(v)
	util.PanicIf(err)
	return string(b)
}

// JSONUnmarshal 为 JSON 为大整数提供无损的反序列化
// 反序列化目标 v 必须为指针类型
func JSONUnmarshal(data []byte, v interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buffer)
	decoder.UseNumber()
	return decoder.Decode(&v)
}

// JSONUnmarshal 为 JSON 为大整数提供无损的反序列化
// 反序列化目标 v 必须为指针类型
func MustJSONUnmarshal(data []byte, v interface{}) {
	util.PanicIf(JSONUnmarshal(data, v))
}

func JSON2String(i interface{}) string {
	if i == nil {
		return ""
	}
	str, ok := i.(string)
	if !ok {
		return ""
	}
	return str
}

func JSON2Int64(i interface{}) int64 {
	if i == nil {
		return 0
	}
	switch i := i.(type) {
	case json.Number:
		var num int64
		var err error
		num, err = i.Int64()
		if err != nil {
			num2, err2 := i.Float64()
			if err2 != nil {
				return 0
			}
			num = int64(num2)
		}
		return num
	case int64:
		return i
	case float64:
		return int64(i)
	case string:
		res, err := strconv.ParseInt(i, 10, 64)
		util.PanicIf(err)
		return res
	default:
		panic(errors.New("Unknown Type" + reflect.TypeOf(i).Name()))
	}
}
