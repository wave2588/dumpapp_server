package impl

import (
	"context"
	"encoding/json"

	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/http"
	"dumpapp_server/pkg/util"
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
	body, err := util.HttpRequest("POST", endpoint, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}, map[string]string{
		"token":            config.DumpConfig.AppConfig.CerServerToken,
		"udid":             udid,
		"udid_region_pool": "private",
	})
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

func (h *CertificateServer) CheckP12File(ctx context.Context, p12FileData, p12Password string) (*http.CheckCerResponse, error) {
	endpoint := config.DumpConfig.AppConfig.CerCheckP12URL
	body, err := util.HttpRequest("POST", endpoint, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}, map[string]string{
		"p12_file_data": p12FileData,
		"p12_password":  p12Password,
	})
	if err != nil {
		return nil, err
	}
	var result http.CheckCerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (h *CertificateServer) CheckCerByUDIDBatchNo(ctx context.Context, udidBatchNo string) (*http.CheckCerResponse, error) {
	endpoint := config.DumpConfig.AppConfig.CerCheckValidateURL
	body, err := util.HttpRequest("POST", endpoint, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}, map[string]string{
		"token":      config.DumpConfig.AppConfig.CerServerToken,
		"udid_batch": udidBatchNo,
	})
	if err != nil {
		return nil, err
	}
	var result http.CheckCerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
