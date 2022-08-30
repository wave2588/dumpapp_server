package datatype

import (
	"net/url"
	"strings"
)

type IncludeFields []string

func (c *IncludeFields) Decode(text string) (err error) {
	cursorStr, err := url.QueryUnescape(text)
	if err != nil {
		return err
	}
	value := strings.Split(cursorStr, ",")
	*c = value
	return
}

func (c IncludeFields) Encode() (string, error) {
	return strings.Join(c, ","), nil
}
