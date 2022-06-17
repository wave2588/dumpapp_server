package impl

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	util2 "dumpapp_server/pkg/util"
)

type SignMobileconfigWebController struct {
	memberIDEncryptionCtl controller.MemberIDEncryptionController
}

var DefaultSignMobileconfigWebController *SignMobileconfigWebController

func init() {
	DefaultSignMobileconfigWebController = NewSignMobileconfigWebController()
}

func NewSignMobileconfigWebController() *SignMobileconfigWebController {
	return &SignMobileconfigWebController{
		memberIDEncryptionCtl: impl.DefaultMemberIDEncryptionController,
	}
}

func (c *SignMobileconfigWebController) Sign(ctx context.Context, memberCode string) ([]byte, error) {
	memberID, err := c.memberIDEncryptionCtl.GetMemberIDByCode(ctx, memberCode)
	if err != nil {
		return nil, err
	}

	/// 日志
	NewAlterWebController().SendDeviceLog(ctx, "用户开始获取描述文件", memberID, map[string]string{
		"code": memberCode,
	})

	/// 签名过后文件路径
	mcPath, err := c.mobileconfigPath()
	if err != nil {
		return nil, err
	}

	/// 证书文件路径
	signFilePath, err := c.signPath()
	if err != nil {
		return nil, err
	}

	/// 写入为签名 mobileconfig
	unSignFilePath, signedFilePath, err := c.saveUnSignFile(ctx, mcPath, memberCode)
	if err != nil {
		return nil, err
	}

	/// 证书文件
	serverCrt := fmt.Sprintf("%s/server.crt", signFilePath)
	serverKey := fmt.Sprintf("%s/server.key", signFilePath)
	caCrt := fmt.Sprintf("%s/ca.crt", signFilePath)

	/// 进行签名
	cmdString := fmt.Sprintf("openssl smime -sign -in %s -out %s -signer %s -inkey %s  -certfile %s -outform der -nodetach", unSignFilePath, signedFilePath, serverCrt, serverKey, caCrt)
	err = util2.Cmd(cmdString)
	if err != nil {
		return nil, err
	}

	/// 读取签名过后的文件
	data, err := ioutil.ReadFile(signedFilePath)
	if err != nil {
		return nil, err
	}

	util.PanicIf(os.Remove(signedFilePath))
	util.PanicIf(os.Remove(unSignFilePath))
	return data, nil
}

func (c *SignMobileconfigWebController) mobileconfigPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/templates/mobileconfig", path), nil
}

func (c *SignMobileconfigWebController) signPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/templates/sign", path), nil
}

func (c *SignMobileconfigWebController) saveUnSignFile(ctx context.Context, mcPath, memberCode string) (string, string, error) {
	url := fmt.Sprintf("%s/device/bind/%s", constant.HOST, memberCode)
	configURL := strings.ReplaceAll(constant.DeviceMobileConfig, "%s", url)
	unSignFileData, err := ioutil.ReadAll(strings.NewReader(configURL))
	util.PanicIf(err)
	/// 未签名文件 mobileconfg 文件路径
	unSignFilePath := fmt.Sprintf("%s/un_sign_%s.mobileconfig", mcPath, memberCode)
	err = util2.WriteFile(unSignFilePath, unSignFileData, 0o644)
	if err != nil {
		return "", "", err
	}
	/// 签名过后的 mobileconfig 文件路径
	signedFilePath := fmt.Sprintf("%s/signed_%s.mobileconfig", mcPath, memberCode)
	return unSignFilePath, signedFilePath, nil
}
