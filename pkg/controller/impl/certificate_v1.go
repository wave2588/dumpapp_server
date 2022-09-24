package impl

import (
	"context"
	"encoding/json"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	util2 "dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/util"
)

type CertificateV1Controller struct {
	certificateDAO dao.CertificateV2DAO
}

var DefaultCertificateV1Controller controller.CertificateController

func init() {
	DefaultCertificateV1Controller = NewCertificateV1Controller()
}

func NewCertificateV1Controller() *CertificateV1Controller {
	return &CertificateV1Controller{
		certificateDAO: impl.DefaultCertificateV2DAO,
	}
}

type createCerResponse struct {
	IsSuccess    bool                   `json:"IsSuccess"`
	Data         *createCerResponseData `json:"Data"`
	ErrorCode    int                    `json:"ErrorCode"`
	ErrorMessage string                 `json:"ErrorMessage"`
}

type createCerResponseData struct {
	P12FileDate             string `json:"p12_file_date"`
	MobileProvisionFileData string `json:"mobile_provision_file_data"`
	UdidBatchNo             string `json:"udid_batch_no"`
	CerAppleid              string `json:"cer_appleid"`
}

func (c *CertificateV1Controller) CreateCer(ctx context.Context, UDID, regionPool string) *controller.CertificateResponse {
	res := &controller.CertificateResponse{}

	endpoint := config.DumpConfig.AppConfig.CerCreateURL
	body, err := util.HttpRequest("POST", endpoint, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}, map[string]string{
		"token":            config.DumpConfig.AppConfig.CerServerToken,
		"udid":             UDID,
		"udid_region_pool": regionPool,
	}, 0)
	if err != nil {
		res.ErrorMessage = util2.StringPtr("请求 v1 接口失败")
		return res
	}

	var result createCerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		res.ErrorMessage = util2.StringPtr("v1 接口 response body json.unmarshal 失败")
		return res
	}
	if result.Data == nil || result.IsSuccess == false {
		res.ErrorMessage = util2.StringPtr(result.ErrorMessage)
		return res
	}

	return &controller.CertificateResponse{
		P12Data:             result.Data.P12FileDate,
		MobileProvisionData: result.Data.MobileProvisionFileData,
		Source:              enum.CertificateSourceV1,
		BizExt: &constant.CertificateBizExt{
			V1UDIDBatchNo:       result.Data.UdidBatchNo,
			V1CerAppleID:        result.Data.CerAppleid,
			OriginalP12Password: "1",
			NewP12Password:      "1",
		},
	}
}

type checkCerResponse struct {
	IsSuccess    bool   `json:"IsSuccess"`
	Data         bool   `json:"Data"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

func (c *CertificateV1Controller) CheckCerIsActive(ctx context.Context, certificateID int64) (bool, error) {
	certificateMap, err := c.certificateDAO.BatchGet(ctx, []int64{certificateID})
	if err != nil {
		return false, err
	}
	cer, ok := certificateMap[certificateID]
	if !ok {
		return false, nil
	}
	endpoint := config.DumpConfig.AppConfig.CerCheckP12URL
	body, err := util.HttpRequest("POST", endpoint, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}, map[string]string{
		"p12_file_data": cer.P12FileData,
		"p12_password":  "1",
	}, 0)
	if err != nil {
		return false, err
	}
	var result checkCerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}
	return result.Data, nil
}
