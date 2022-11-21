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
)

type CertificateV2Controller struct {
	certificateDAO  dao.CertificateV2DAO
	memberDeviceDAO dao.MemberDeviceDAO
}

var DefaultCertificateV2Controller controller.CertificateController

func init() {
	DefaultCertificateV2Controller = NewCertificateV2Controller()
}

func NewCertificateV2Controller() *CertificateV2Controller {
	return &CertificateV2Controller{
		certificateDAO:  impl.DefaultCertificateV2DAO,
		memberDeviceDAO: impl.DefaultMemberDeviceDAO,
	}
}

type createResponse struct {
	Code int64       `json:"code"` ///
	Msg  string      `json:"msg"`  /// 添加成功
	Data interface{} `json:"data"`
}

type createDataResponse struct {
	ID                 string `json:"id"` /// 设备 id
	CertificateContent string `json:"certificateContent"`
	ProfileContent     string `json:"profileContent"`
}

func (c *CertificateV2Controller) CreateCer(ctx context.Context, UDID, regionPool string) *controller.CertificateResponse {
	/// regionPool 1 是私有 2 和 3 是共有
	res := &controller.CertificateResponse{}

	endpoint := config.DumpConfig.AppConfig.CerCreateURLV2
	requestBodyMap := map[string]interface{}{
		"token": config.DumpConfig.AppConfig.CerServerTokenV2,
		"udid":  UDID,
		"type":  1,
	}
	requestBody, _ := json.Marshal(requestBodyMap)
	body, err := util.HttpRequestV2("POST", endpoint, map[string]string{
		"Content-Type": "application/json",
	}, bytes.NewBuffer(requestBody))
	if err != nil {
		res.ErrorMessage = util2.StringPtr("v2 cer server fail")
		return res
	}

	var response *createResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		res.ErrorMessage = util2.StringPtr("v2 response body json.unmarshal fail")
		return res
	}
	if response.Code != 1 {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v2 cer server error：%s", response.Msg))
		return res
	}
	d, err := json.Marshal(response.Data)
	if err != nil {
		res.ErrorMessage = util2.StringPtr("v2 json.Marshal(response.Data) fail")
		return res
	}
	var responseData *createDataResponse
	err = json.Unmarshal(d, &responseData)
	if err != nil {
		res.ErrorMessage = util2.StringPtr("v2 responseData json.Unmarshal fail")
		return res
	}

	cerData := c.getCerByServer(ctx, responseData.ID)
	if cerData.ErrorMessage != nil {
		res.ErrorMessage = cerData.ErrorMessage
		return res
	}
	return &controller.CertificateResponse{
		P12Data:             cerData.Data.P12,
		MobileProvisionData: cerData.Data.Mobileprovision,
		Source:              enum.CertificateSourceV2,
		BizExt: datatype.CertificateBizExt{
			V2DeviceID:          responseData.ID,
			OriginalP12Password: cerData.Data.Password,
			NewP12Password:      "1",
		},
	}
}

type getCerResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data *getCerDataResponse `json:"data"`

	ErrorMessage *string `json:"error_message"`
}

type getCerDataResponse struct {
	Mobileprovision string `json:"mobileprovision"` /// 描述文件 base64 位编码
	P12             string `json:"p12"`             /// p12文件 base64 位编码内容
	Password        string `json:"password"`        /// 证书密码
	AppleDeviceID   string `json:"deviceId"`        /// 苹果平台设备 ID
}

func (c *CertificateV2Controller) getCerByServer(ctx context.Context, id string) *getCerResponse {
	res := &getCerResponse{}
	endpoint := config.DumpConfig.AppConfig.CerGetV2
	requestBodyMap := map[string]interface{}{
		"token": config.DumpConfig.AppConfig.CerServerTokenV2,
		"page":  1,
		"id":    id,
	}
	requestBody, _ := json.Marshal(requestBodyMap)
	body, err := util.HttpRequestV2("POST", endpoint, map[string]string{
		"Content-Type": "application/json",
	}, bytes.NewBuffer(requestBody))
	if err != nil {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v2 get_cer_by_server fail.  id: %s", id))
		return res
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v2 get_cer_by_server json.Unmarshal fail.  id: %s   json: %s", id, string(body)))
		return res
	}
	if res.Code != 1 {
		res.ErrorMessage = util2.StringPtr(fmt.Sprintf("v2 get_cer_by_server code != 1.  id: %s   json: %s", id, string(body)))
		return res
	}
	return res
}

func (c *CertificateV2Controller) CheckCerIsActive(ctx context.Context, certificateID int64) (bool, error) {
	cerMap, err := c.certificateDAO.BatchGet(ctx, []int64{certificateID})
	if err != nil {
		return false, err
	}
	cer, ok := cerMap[certificateID]
	if !ok {
		return false, nil
	}

	resp := c.getCerByServer(ctx, cer.BizExt.V2DeviceID)
	if resp.ErrorMessage != nil {
		return false, nil
	}

	if resp.Code != 1 || resp.Data == nil || resp.Data.P12 == "" {
		return false, nil
	}

	return true, nil
}

func (c *CertificateV2Controller) GetBalance(ctx context.Context) (*controller.CertificateBalance, error) {
	return &controller.CertificateBalance{}, nil
}
