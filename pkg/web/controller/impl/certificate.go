package impl

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"dumpapp_server/pkg/common/util"
	util2 "dumpapp_server/pkg/util"
)

type CertificateWebController struct{}

var DefaultCertificateWebController *CertificateWebController

func init() {
	DefaultCertificateWebController = NewCertificateWebController()
}

func NewCertificateWebController() *CertificateWebController {
	return &CertificateWebController{}
}

func (c *CertificateWebController) GetModifiedCertificateData(ctx context.Context, p12Data string) (string, error) {
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
		panic(err)
	}

	pemFileID := util2.MustGenerateID(ctx)
	pemFilePath := fmt.Sprintf("%s/%d.pem", path, pemFileID)
	cmd := fmt.Sprintf(`openssl pkcs12 -in %s -password pass:"1" -passout pass:"123" -name "www.dumpapp.com" -out %s`, originFilePath, pemFilePath)
	err = util2.Cmd(cmd)
	if err != nil {
		return "", err
	}

	resultFileID := util2.MustGenerateID(ctx)
	resultFilePath := fmt.Sprintf("%s/%d.p12", path, resultFileID)
	cmd = fmt.Sprintf(`openssl pkcs12 -passin pass:"dumpapp" -passout pass:"dumpapp" -export -in %s  -name "www.dumpapp.com" -out %s`, pemFilePath, resultFilePath)
	err = util2.Cmd(cmd)
	if err != nil {
		return "", err
	}
	resultData, err := ioutil.ReadFile(resultFilePath)
	if err != nil {
		return "", err
	}

	util.PanicIf(os.Remove(originFilePath))
	util.PanicIf(os.Remove(pemFilePath))
	util.PanicIf(os.Remove(resultFilePath))

	return base64.StdEncoding.EncodeToString(resultData), nil
}

func (c *CertificateWebController) mobileconfigPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/templates/mobileconfig", path), nil
}
