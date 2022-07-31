package util

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func HttpRequest(method, endpoint string, header, values map[string]string, timeout time.Duration) ([]byte, error) {
	data := url.Values{}
	for key, value := range values {
		data.Set(key, value)
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	r, err := http.NewRequest(method, endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
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

func HttpRequestV2(method, endpoint string, header map[string]string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	r, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		r.Header.Add(key, value)
	}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
