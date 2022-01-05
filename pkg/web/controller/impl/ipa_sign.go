package impl

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
)

type IpaSignWebController struct {
	ipaVersionDAO        dao.IpaVersionDAO
	memberDeviceDAO      dao.MemberDeviceDAO
	certificateDAO       dao.CertificateDAO
	certificateDeviceDAO dao.CertificateDeviceDAO
}

var DefaultIpaSignWebController *IpaSignWebController

func init() {
	DefaultIpaSignWebController = NewIpaSignWebController()
}

func NewIpaSignWebController() *IpaSignWebController {
	return &IpaSignWebController{
		ipaVersionDAO:        impl.DefaultIpaVersionDAO,
		memberDeviceDAO:      impl.DefaultMemberDeviceDAO,
		certificateDAO:       impl.DefaultCertificateDAO,
		certificateDeviceDAO: impl.DefaultCertificateDeviceDAO,
	}
}

func (c *IpaSignWebController) Sign(ctx context.Context, loginID, certificateID, ipaVersionID int64) error {
	/// 检测 ipaVersionID
	ipaVersion, err := c.checkIpaVersionID(ctx, ipaVersionID)
	if err != nil {
		return err
	}
	/// 检测 certificateID
	cer, err := c.checkCertificateID(ctx, loginID, certificateID)
	if err != nil {
		return err
	}
	/// 获取 pem 文件路径
	pemFilePath, mpFilePath, err := c.generatePemAndMpFilePath(ctx, cer)
	if err != nil {
		return err
	}

	fmt.Println(ipaVersion, cer)

	return nil
}

func (c *IpaSignWebController) checkIpaVersionID(ctx context.Context, ipaVersionID int64) (*models.IpaVersion, error) {
	ipaVersionMap, err := c.ipaVersionDAO.BatchGet(ctx, []int64{ipaVersionID})
	if err != nil {
		return nil, err
	}
	ipaVersion, ok := ipaVersionMap[ipaVersionID]
	if !ok {
		return nil, errors.ErrNotFoundIpaVersion
	}
	return ipaVersion, nil
}

func (c *IpaSignWebController) checkCertificateID(ctx context.Context, loginID, certificateID int64) (*models.Certificate, error) {
	cerMap := render.NewCertificateRender([]int64{certificateID}, loginID, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	cer, ok := cerMap[certificateID]
	if !ok {
		return nil, errors.ErrNotFoundCertificate
	}
	if !cer.IsValidate || !cer.P12IsActive {
		return nil, errors.ErrCertificateInvalid
	}
	return cer.Meta, nil
}

func (c *IpaSignWebController) generatePemAndMpFilePath(ctx context.Context, cer *models.Certificate) (string, string, error) {
	p12Data, err := base64.StdEncoding.DecodeString(cer.P12FileDate)
	if err != nil {
		return "", "", err
	}
	mpData, err := base64.StdEncoding.DecodeString(cer.MobileProvisionFileData)
	if err != nil {
		return "", "", err
	}

	path, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	/// 生成 p12 file
	p12FilePath := fmt.Sprintf("%s/templates/ipa_sign/%d.p12", path, cer.ID)
	err = ioutil.WriteFile(p12FilePath, p12Data, 0o644)
	if err != nil {
		return "", "", err
	}

	/// 生成 mp file
	mpFilePath := fmt.Sprintf("%s/templates/ipa_sign/%d.mobileprovision", path, cer.ID)
	err = ioutil.WriteFile(mpFilePath, mpData, 0o644)
	if err != nil {
		return "", "", err
	}

	/// p12 convert pem
	pemFilePath := fmt.Sprintf("%s/templates/ipa_sign/%d.pem", path, cer.ID)
	cmd := fmt.Sprintf(`openssl pkcs12 -in %s -password pass:"dumpapp" -passout pass:"123" -out %s`, p12FilePath, pemFilePath)
	err = util2.Cmd(cmd)
	if err != nil {
		return "", "", err
	}

	/// 删除 p12
	err = os.Remove(p12FilePath)
	if err != nil {
		return "", "", err
	}
	return pemFilePath, pemFilePath, nil
}

func (c *IpaSignWebController) startSign(ctx context.Context, ipaVersion *models.IpaVersion, pemFile string) error {
	return nil
}
