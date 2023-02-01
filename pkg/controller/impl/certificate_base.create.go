package impl

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
)

func (c *CertificateBaseController) Create(ctx context.Context, UDID string) (cerBase *controller.CerCreateResponse, alterMsg string, err error) {
	/// 请求整数接口
	config, err := c.adminConfigInfoDAO.GetConfig(ctx)
	if err != nil {
		return nil, "", err
	}

	// 生成证书
	// 这里做个分流, 后台可配置随意切换任何平台
	var response *controller.CertificateResponse
	switch config.BizExt.CerSource {
	case enum.CertificateSourceV2:
		response = c.certificateV2Ctl.CreateCer(ctx, UDID, "1")
	case enum.CertificateSourceV3:
		response = c.certificateV3Ctl.CreateCer(ctx, UDID, "1")
	default:
		return nil, "", errors.UnproccessableError("请联系管理员检查配置信息是否正确")
	}

	// 创建失败
	if response.ErrorMessage != nil {
		return nil, *response.ErrorMessage, errors.ErrCreateCertificateFail
	}

	p12FileData := response.P12Data
	/// p12 文件修改内容
	modifiedP12FileData, err := c.GetModifiedCertificateData(ctx, p12FileData, response.BizExt.OriginalP12Password, response.BizExt.NewP12Password)
	if err != nil {
		alterMsg = fmt.Sprintf("修改证书文件出错, err: %s", err.Error())
	}

	return &controller.CerCreateResponse{
		Response:            response,
		ModifiedP12FileData: modifiedP12FileData,
	}, "", nil
}

func (c *CertificateBaseController) GetModifiedCertificateData(ctx context.Context, p12Data, originalPassword, newPassword string) (string, error) {
	path, err := c.mobileconfigPath()
	if err != nil {
		return "", err
	}
	originFileID := util2.MustGenerateID(ctx)
	data, err := base64.StdEncoding.DecodeString(p12Data)
	util.PanicIf(err)

	originFilePath := fmt.Sprintf("%s/%d.p12", path, originFileID)
	err = ioutil.WriteFile(originFilePath, data, 0o644)
	if err != nil {
		return "", err
	}

	pemFileID := util2.MustGenerateID(ctx)
	pemFilePath := fmt.Sprintf("%s/%d.pem", path, pemFileID)
	cmd := fmt.Sprintf(`openssl pkcs12 -in %s -password pass:"%s" -passout pass:"%s" -name "www.dumpapp.com" -out %s`, originFilePath, originalPassword, newPassword, pemFilePath)
	err = util2.Cmd(cmd)
	if err != nil {
		return "", err
	}

	resultFileID := util2.MustGenerateID(ctx)
	resultFilePath := fmt.Sprintf("%s/%d.p12", path, resultFileID)
	cmd = fmt.Sprintf(`openssl pkcs12 -passin pass:"%s" -passout pass:"%s" -export -in %s  -name "www.dumpapp.com" -out %s`, newPassword, newPassword, pemFilePath, resultFilePath)
	err = util2.Cmd(cmd)
	if err != nil {
		return "", err
	}
	resultData, err := ioutil.ReadFile(resultFilePath)
	if err != nil {
		return "", err
	}

	_ = os.Remove(originFilePath)
	_ = os.Remove(pemFilePath)
	_ = os.Remove(resultFilePath)

	return base64.StdEncoding.EncodeToString(resultData), nil
}

func (c *CertificateBaseController) mobileconfigPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/templates/mobileconfig", path), nil
}
