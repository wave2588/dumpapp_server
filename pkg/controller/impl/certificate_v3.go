package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	util2 "dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/util"
	"github.com/pkg/errors"
)

type CertificateV3Controller struct {
	certificateDAO  dao.CertificateV2DAO
	memberDeviceDAO dao.MemberDeviceDAO
}

var DefaultCertificateV3Controller controller.CertificateController

func init() {
	DefaultCertificateV3Controller = NewCertificateV3Controller()
}

func NewCertificateV3Controller() *CertificateV3Controller {
	return &CertificateV3Controller{
		certificateDAO:  impl.DefaultCertificateV2DAO,
		memberDeviceDAO: impl.DefaultMemberDeviceDAO,
	}
}

type cerResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Time string           `json:"time"`
	Data *cerDataResponse `json:"data"`

	ErrorMessage *string `json:"error_message"` /// 业务自己加的字段
}

type cerDataResponse struct {
	ID              string `json:"id"`
	Mobileprovision string `json:"mobileprovision"`
	P12             string `json:"p12"`
	DeviceID        string `json:"device_id"`
	State           bool   `json:"state"`
}

func (c *CertificateV3Controller) CreateCer(ctx context.Context, UDID, regionPool string) *controller.CertificateResponse {
	res := &controller.CertificateResponse{}

	endpoint := "https://developer.52tzs.com/api/adddevice"
	requestBodyMap := map[string]interface{}{
		"token": config.DumpConfig.AppConfig.CerServerTokenV3,
		"udid":  UDID,
	}
	requestBody, _ := json.Marshal(requestBodyMap)
	body, err := util.HttpRequestV2("POST", endpoint, map[string]string{
		"Content-Type": "application/json",
	}, bytes.NewBuffer(requestBody))
	if err != nil {
		res.ErrorMessage = util2.StringPtr("v3 cer server fail")
		return res
	}

	var response *cerResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v3 response body json.unmarshal fail. json: %s", string(body)))
		return res
	}

	if response.Code != 1 {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v3 cer server error：%s", response.Msg))
		return res
	}

	if response.Data == nil || response.Data.P12 == "" || response.Data.Mobileprovision == "" {
		data, _ := json.Marshal(response)
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v3 cer server 返回数据错误：%s", string(data)))
		return res
	}

	return &controller.CertificateResponse{
		P12Data:             response.Data.P12,
		MobileProvisionData: response.Data.Mobileprovision,
		Source:              enum.CertificateSourceV3,
		BizExt: datatype.CertificateBizExt{
			V3DeviceID:          response.Data.ID,
			OriginalP12Password: "1",
			NewP12Password:      "1",
		},
	}
}

func (c *CertificateV3Controller) getCerByServer(ctx context.Context, id string) *cerResponse {
	res := &cerResponse{}
	endpoint := "https://developer.52tzs.com/api/getCertificate"
	requestBodyMap := map[string]interface{}{
		"token": config.DumpConfig.AppConfig.CerServerTokenV3,
		"id":    id,
	}
	requestBody, _ := json.Marshal(requestBodyMap)
	body, err := util.HttpRequestV2("POST", endpoint, map[string]string{
		"Content-Type": "application/json",
	}, bytes.NewBuffer(requestBody))
	if err != nil {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v3 get_cer_by_server fail.  id: %s", id))
		return res
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v3 get_cer_by_server json.Unmarshal fail.  id: %s   json: %s", id, string(body)))
		return res
	}
	if res.Code != 1 {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v3 get_cer_by_server code != 1.  id: %s   json: %s", id, string(body)))
		return res
	}
	return res
}

func (c *CertificateV3Controller) CheckCerIsActive(ctx context.Context, certificateID int64) (bool, error) {
	cerMap, err := c.certificateDAO.BatchGet(ctx, []int64{certificateID})
	if err != nil {
		return false, err
	}
	cer, ok := cerMap[certificateID]
	if !ok {
		return false, nil
	}

	resp := c.getCerByServer(ctx, cer.BizExt.V3DeviceID)
	if resp.ErrorMessage != nil {
		return false, nil
	}

	if resp.Code != 1 || resp.Data == nil {
		return false, nil
	}

	return resp.Data.State, nil
}

type cerBalanceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
	Data struct {
		Score   int64  `json:"score"`
		Balance string `json:"balance"`
	} `json:"data"`
}

func (c *CertificateV3Controller) GetBalance(ctx context.Context) (*controller.CertificateBalance, error) {
	res := &cerBalanceResponse{}
	endpoint := "https://developer.52tzs.com/api/getbalance"
	requestBodyMap := map[string]interface{}{
		"token": config.DumpConfig.AppConfig.CerServerTokenV3,
	}
	requestBody, _ := json.Marshal(requestBodyMap)
	body, err := util.HttpRequestV2("POST", endpoint, map[string]string{
		"Content-Type": "application/json",
	}, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("v3 get_balance_by_server fail. err: %s", err.Error()))
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("v3 get_balance_by_server json.Unmarshal fail.  %s", string(body)))
	}
	return &controller.CertificateBalance{
		Count: res.Data.Score,
	}, nil
}
