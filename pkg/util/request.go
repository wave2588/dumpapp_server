package util

import (
	"io/ioutil"
	http2 "net/http"
	"net/url"
	"strconv"
	"strings"
)

func HttpRequest(method, endpoint string, header, values map[string]string) ([]byte, error) {
	data := url.Values{}
	for key, value := range values {
		data.Set(key, value)
	}
	client := &http2.Client{}
	r, err := http2.NewRequest(method, endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		r.Header.Add(key, value)
	}
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
