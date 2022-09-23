package impl

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
)

type CertificateWebController struct {
	adminConfigInfoDAO dao.AdminConfigInfoDAO
	memberDeviceDAO    dao.MemberDeviceDAO
	certificateDAO     dao.CertificateV2DAO
	memberPayCountCtl  controller.MemberPayCountController
	certificateV2Ctl   controller.CertificateController
	certificateV3Ctl   controller.CertificateController
	alterWebCtl        controller2.AlterWebController
}

var DefaultCertificateWebController *CertificateWebController

func init() {
	DefaultCertificateWebController = NewCertificateWebController()
}

func NewCertificateWebController() *CertificateWebController {
	return &CertificateWebController{
		adminConfigInfoDAO: impl.DefaultAdminConfigInfoDAO,
		memberDeviceDAO:    impl.DefaultMemberDeviceDAO,
		certificateDAO:     impl.DefaultCertificateV2DAO,
		memberPayCountCtl:  impl2.DefaultMemberPayCountController,
		certificateV2Ctl:   impl2.DefaultCertificateV2Controller,
		certificateV3Ctl:   impl2.DefaultCertificateV3Controller,
		alterWebCtl:        NewAlterWebController(),
	}
}

func (c *CertificateWebController) PayCertificate(ctx context.Context, loginID int64, udid string, payCount int64, payType string) (int64, error) {
	util.PanicIf(c.memberPayCountCtl.CheckPayCount(ctx, loginID, payCount))

	memberDevice, err := c.memberDeviceDAO.GetByMemberIDUdidSafe(ctx, loginID, udid)
	util.PanicIf(err)
	if memberDevice == nil {
		return 0, errors.ErrDeviceNotFound
	}
	if memberDevice.MemberID != loginID {
		return 0, errors.ErrCreateCertificateFailV2
	}

	/// 发送用户开始购买证书日志
	c.alterWebCtl.SendBeganCreateCertificateMsg(ctx, loginID, udid)

	/// 请求整数接口
	config, err := c.adminConfigInfoDAO.GetConfig(ctx)
	if err != nil {
		return 0, err
	}

	/// 这里做个分流, 后台可配置随意切换任何平台
	var response *controller.CertificateResponse
	switch config.BizExt.CerSource {
	case enum.CertificateSourceV2:
		response = c.certificateV2Ctl.CreateCer(ctx, udid, "1")
	case enum.CertificateSourceV3:
		response = c.certificateV3Ctl.CreateCer(ctx, udid, "1")
	default:
		return 0, errors.UnproccessableError("请联系管理员检查配置信息是否正确")
	}

	if response.ErrorMessage != nil {
		/// 创建失败推送
		c.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, *response.ErrorMessage)
		util.PanicIf(errors.ErrCreateCertificateFail)
	}
	if response.BizExt == nil {
		c.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, "response biz_ext is nil")
		util.PanicIf(errors.ErrCreateCertificateFail)
	}

	p12FileData := response.P12Data
	mpFileData := response.MobileProvisionData
	/// p12 文件修改内容
	modifiedP12FileData, err := c.GetModifiedCertificateData(ctx, p12FileData, response.BizExt.OriginalP12Password, response.BizExt.NewP12Password)
	if err != nil {
		c.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, fmt.Sprintf("修改证书文件出错, err: %s", err.Error()))
	}
	// util.PanicIf(err)

	/// 记录证书等级, 方便后期候补
	if payCount == constant.CertificatePriceL1 {
		response.BizExt.Level = 1
	} else if payCount == constant.CertificatePriceL2 {
		response.BizExt.Level = 2
	} else if payCount == constant.CertificatePriceL3 {
		response.BizExt.Level = 3
	}

	/// 计算证书 md5
	p12FileMd5 := util2.StringMd5(p12FileData)
	mpFileMd5 := util2.StringMd5(mpFileData)

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	cerID := util2.MustGenerateID(ctx)
	util.PanicIf(c.certificateDAO.Insert(ctx, &models.CertificateV2{
		ID:                         cerID,
		DeviceID:                   memberDevice.ID,
		P12FileData:                p12FileData,
		P12FileDataMD5:             p12FileMd5,
		ModifiedP12FileDate:        modifiedP12FileData,
		MobileProvisionFileData:    mpFileData,
		MobileProvisionFileDataMD5: mpFileMd5,
		Source:                     response.Source,
		BizExt:                     response.BizExt.String(),
	}))

	/// 扣除消费的 D 币
	util.PanicIf(c.memberPayCountCtl.DeductPayCount(ctx, loginID, payCount, enum.MemberPayCountStatusUsed, enum.MemberPayCountUseCertificate, datatype.MemberPayCountRecordBizExt{
		ObjectID:   cerID,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeCertificate,
	}))

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	/// 发送消费成功通知
	c.alterWebCtl.SendCreateCertificateSuccessMsg(ctx, loginID, memberDevice.ID, cerID)

	return cerID, nil
}

func (c *CertificateWebController) GetModifiedCertificateData(ctx context.Context, p12Data, originalPassword, newPassword string) (string, error) {
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
