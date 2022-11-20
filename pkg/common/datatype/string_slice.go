package datatype

import (
	"fmt"
	"net/url"
	"strings"
)

type StringSlice []string

func (slice *StringSlice) Decode(text string) (err error) {
	cursorStr, err := url.QueryUnescape(text)
	if err != nil {
		return err
	}
	fmt.Print(111, cursorStr)
	value := strings.Split(cursorStr, ",")
	*slice = value
	return
}

func (slice StringSlice) Encode() (string, error) {
	fmt.Print(111, slice)
	return strings.Join(slice, ","), nil
}
