package datatype

import (
	"net/url"
	"strings"
)

type StringSlice []string

func (slice *StringSlice) Decode(text string) (err error) {
	cursorStr, err := url.QueryUnescape(text)
	if err != nil {
		return err
	}
	value := strings.Split(cursorStr, ",")
	*slice = value
	return
}

func (slice StringSlice) Encode() (string, error) {
	return strings.Join(slice, ","), nil
}
