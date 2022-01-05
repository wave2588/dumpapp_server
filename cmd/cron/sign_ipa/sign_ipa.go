package sign_ipa

import (
	"context"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Run() {
	fmt.Println("SignIpa")
	run()
}

func run() {
	ctx := context.Background()
	unprocessedIpaSigns, err := impl.DefaultIpaSignDAO.GetByStatus(ctx, enum.IpaSignStatusUnprocessed)
	util.PanicIf(err)
	if len(unprocessedIpaSigns) == 0 {
		return
	}

	/// 标记为处理中
	ipaSign := unprocessedIpaSigns[0]
	ipaSign.Status = enum.IpaSignStatusProcessing
	util.PanicIf(impl.DefaultIpaSignDAO.Update(ctx, ipaSign))

	var ipaSignBizExt constant.IpaSignBizExt
	util.PanicIf(json.Unmarshal([]byte(ipaSign.BizExt), &ipaSignBizExt))
	err = sign(ctx, ipaSign.MemberID, ipaSign.CertificateID, ipaSignBizExt.IpaVersionID)

	/// 签名失败
	if err != nil {
		ipaSign.Status = enum.IpaSignStatusFail
		util.PanicIf(impl.DefaultIpaSignDAO.Update(ctx, ipaSign))
		util.PanicIf(err)
		return
	}

	/// 签名成功
	ipaSign.Status = enum.IpaSignStatusSuccess
	util.PanicIf(impl.DefaultIpaSignDAO.Update(ctx, ipaSign))
}

func sign(ctx context.Context, memberID, certificateID, ipaVersionID int64) error {
	/// 检测 ipaVersionID
	ipaVersion, err := checkIpaVersionID(ctx, ipaVersionID)
	if err != nil {
		return err
	}
	/// 检测 certificateID
	cer, err := checkCertificateID(ctx, memberID, certificateID)
	if err != nil {
		return err
	}
	/// 获取 pem 文件路径
	pemFilePath, mpFilePath, err := generatePemAndMpFilePath(ctx, cer)
	if err != nil {
		return err
	}

	err = startSign(ctx, ipaVersion, pemFilePath, mpFilePath)
	if err != nil {
		return err
	}

	return err
}

func checkIpaVersionID(ctx context.Context, ipaVersionID int64) (*models.IpaVersion, error) {
	ipaVersionMap, err := impl.DefaultIpaVersionDAO.BatchGet(ctx, []int64{ipaVersionID})
	if err != nil {
		return nil, err
	}
	ipaVersion, ok := ipaVersionMap[ipaVersionID]
	if !ok {
		return nil, errors.ErrNotFoundIpaVersion
	}
	return ipaVersion, nil
}

func checkCertificateID(ctx context.Context, memberID, certificateID int64) (*models.Certificate, error) {
	cerMap := render.NewCertificateRender([]int64{certificateID}, memberID, render.CertificateDefaultRenderFields...).RenderMap(ctx)
	cer, ok := cerMap[certificateID]
	if !ok {
		return nil, errors.ErrNotFoundCertificate
	}
	if !cer.IsValidate || !cer.P12IsActive {
		return nil, errors.ErrCertificateInvalid
	}
	return cer.Meta, nil
}

func generatePemAndMpFilePath(ctx context.Context, cer *models.Certificate) (string, string, error) {
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

func startSign(ctx context.Context, ipaVersion *models.IpaVersion, pemFilePath, mpFilePath string) error {
	return nil
}
