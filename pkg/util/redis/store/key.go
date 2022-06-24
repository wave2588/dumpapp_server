package store

import (
	"bytes"
	"fmt"
)

func NewKey(funcName string, args ...interface{}) (cacheKey string) {
	buf := bytes.NewBufferString(funcName)
	for _, arg := range args {
		buf.WriteString(fmt.Sprintf("|%v", arg))
	}
	return buf.String()
}
