package sign_ipa

import (
	"context"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	impl2 "dumpapp_server/pkg/controller/impl"
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
	fmt.Println("sign ipa")
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
	/// ipa sign folder path
	ipaSignPath := ipaSignFolderPath()

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
	pemFilePath, mpFilePath, err := generatePemAndMpFilePath(ctx, cer, ipaSignPath)
	if err != nil {
		return err
	}

	/// 获取签名工具
	zsignPath, err := getZsignPath()
	if err != nil {
		return err
	}

	/// 下载原始 ipa
	originIpaID := util2.MustGenerateCode(ctx, 8)
	originIpaName := fmt.Sprintf("%s.ipa", originIpaID)
	originIpaPath := fmt.Sprintf("%s/%d.ipa", ipaSignPath, originIpaName)
	err = impl2.DefaultTencentController.GetToFile(ctx, ipaVersion.TokenPath, originIpaPath)
	if err != nil {
		return err
	}

	/// 生成签名过后 ipa 路径
	signIpaID := util2.MustGenerateCode(ctx, 8)
	signIpaName := fmt.Sprintf("%s.ipa", signIpaID)
	signedIpaPath := fmt.Sprintf("%s/%d.ipa", ipaSignPath, signIpaName)

	/// 开始签名
	//./zsign/build/zsign -k developer.pem -p 123 -m developer.mobileprovision -o output.ipa -z 9 test.ipa
	cmd := fmt.Sprintf("%s -k %s -p dumpapp -m %s -o %s -z 9 %s", zsignPath, pemFilePath, mpFilePath, signedIpaPath, originIpaPath)
	err = util2.Cmd(cmd)
	if err != nil {
		return err
	}

	/// 上传签名 ipa
	//err = impl2.DefaultTencentController.PutSignIpaByFile(ctx, signIpaName, signedIpaPath)
	//if err != nil {
	//	return err
	//}

	/// 签名结束 删除所有文件
	os.Remove(pemFilePath)
	os.Remove(mpFilePath)
	os.Remove(originIpaPath)
	os.Remove(signedIpaPath)

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
	if !cer.P12IsActive {
		return nil, errors.ErrCertificateInvalid
	}
	return cer.Meta, nil
}

func generatePemAndMpFilePath(ctx context.Context, cer *models.Certificate, ipaSignPath string) (string, string, error) {
	p12Data, err := base64.StdEncoding.DecodeString(cer.ModifiedP12FileDate)
	if err != nil {
		return "", "", err
	}
	mpData, err := base64.StdEncoding.DecodeString(cer.MobileProvisionFileData)
	if err != nil {
		return "", "", err
	}

	/// 生成 p12 file
	p12FilePath := fmt.Sprintf("%s/%d.p12", ipaSignPath, cer.ID)
	err = ioutil.WriteFile(p12FilePath, p12Data, 0o644)
	if err != nil {
		return "", "", err
	}

	/// 生成 mp file
	mpFilePath := fmt.Sprintf("%s/%d.mobileprovision", ipaSignPath, cer.ID)
	err = ioutil.WriteFile(mpFilePath, mpData, 0o644)
	if err != nil {
		return "", "", err
	}

	/// p12 convert pem
	pemFilePath := fmt.Sprintf("%s/%d.pem", ipaSignPath, cer.ID)
	cmd := fmt.Sprintf(`openssl pkcs12 -in %s -password pass:"dumpapp" -passout pass:"dumpapp" -out %s`, p12FilePath, pemFilePath)
	err = util2.Cmd(cmd)
	if err != nil {
		return "", "", err
	}

	/// 删除 p12
	err = os.Remove(p12FilePath)
	if err != nil {
		return "", "", err
	}
	return pemFilePath, mpFilePath, nil
}

func getZsignPath() (string, error) {
	path := rootPath()
	zsignPath := ""
	switch config.DumpConfig.AppConfig.Env {
	case config.DumpEnvProduction:
		zsignPath = fmt.Sprintf("%s/tools/bin/linux/zsign", path)
	case config.DumpEnvDev:
		zsignPath = fmt.Sprintf("%s/tools/bin/macos/zsign", path)
	default:
		return "", errors2.NewError(1000, "env fail", "为识别的 env")
	}
	return zsignPath, nil
}

func ipaSignFolderPath() string {
	path := rootPath()
	return fmt.Sprintf("%s/templates/ipa_sign", path)
}

func rootPath() string {
	path, err := os.Getwd()
	util.PanicIf(err)
	return path
}
