package util

import (
	"net/http"
	"strconv"
)

func GetIntArgument(r *http.Request, key string, fallback int) int {
	if v, err := getIntArgument(r, key, 32); err != nil {
		return fallback
	} else {
		return int(v)
	}
}

func getIntArgument(r *http.Request, key string, bitSize int) (int64, error) {
	return strconv.ParseInt(r.URL.Query().Get(key), 10, bitSize)
}
