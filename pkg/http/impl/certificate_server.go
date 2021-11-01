package impl

import (
	"context"
	"encoding/json"
	"io/ioutil"
	http2 "net/http"
	"net/url"
	"strconv"
	"strings"

	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/http"
)

type CertificateServer struct{}

var DefaultCertificateServer *CertificateServer

func init() {
	DefaultCertificateServer = NewCertificateServer()
}

func NewCertificateServer() *CertificateServer {
	return &CertificateServer{}
}

func (h *CertificateServer) CreateCer(ctx context.Context, udid string) (*http.CreateCerResponse, error) {
	endpoint := config.DumpConfig.AppConfig.CerCreateURL
	data := url.Values{}
	data.Set("token", config.DumpConfig.AppConfig.CerServerToken)
	data.Set("udid", udid)
	data.Set("udid_region_pool", "public")
	client := &http2.Client{}
	r, err := http2.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result http.CreateCerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
