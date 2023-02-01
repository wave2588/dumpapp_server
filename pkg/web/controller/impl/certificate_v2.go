package impl

import (
	"context"

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
	"github.com/spf13/cast"
)

type CertificateV2WebController struct {
	memberDeviceDAO         dao.MemberDeviceDAO
	certificateDAO          dao.CertificateV2DAO
	certificateDeviceDAO    dao.CertificateDeviceDAO
	memberPayCountRecordDAO dao.MemberPayCountRecordDAO
	memberPayCountCtl       controller.MemberPayCountController
	certificatePriceCtl     controller.CertificatePriceController
	certificateBaseCtl      controller.CertificateBaseController
	certificateDeviceCtl    controller.CertificateDeviceController
	alterWebCtl             controller2.AlterWebController
}

var DefaultCertificateV2WebController *CertificateV2WebController

func init() {
	DefaultCertificateV2WebController = NewCertificateV2WebController()
}

func NewCertificateV2WebController() *CertificateV2WebController {
	return &CertificateV2WebController{
		memberDeviceDAO:         impl.DefaultMemberDeviceDAO,
		certificateDAO:          impl.DefaultCertificateV2DAO,
		certificateDeviceDAO:    impl.DefaultCertificateDeviceDAO,
		memberPayCountRecordDAO: impl.DefaultMemberPayCountRecordDAO,
		memberPayCountCtl:       impl2.DefaultMemberPayCountController,
		certificatePriceCtl:     impl2.DefaultCertificatePriceController,
		certificateBaseCtl:      impl2.DefaultCertificateBaseController,
		certificateDeviceCtl:    impl2.DefaultCertificateDeviceController,
		alterWebCtl:             NewAlterWebController(),
	}
}

func (c *CertificateV2WebController) Create(ctx context.Context, loginID int64, UDID, note string, priceID int64) (int64, error) {
	// 判断是候补还是正常购买
	isReplenish, err := c.certificateDeviceCtl.IsReplenish(ctx, loginID, UDID)
	if err != nil {
		return 0, err
	}

	// 绑定设备
	memberDevice, err := c.getMemberDevice(ctx, loginID, UDID)
	if memberDevice.MemberID != loginID {
		return 0, errors.ErrCreateCertificateFailV2
	}

	var cerID int64
	// 候补
	if isReplenish {
		cerID, err = c.replenish(ctx, loginID, UDID, note, priceID, memberDevice)
	} else {
		cerID, err = c.realCreate(ctx, loginID, UDID, note, priceID, memberDevice)
	}
	if err != nil {
		return 0, err
	}

	/// 发送消费成功通知
	c.alterWebCtl.SendCreateCertificateSuccessMsgV2(ctx, loginID, memberDevice.ID, cerID, isReplenish)

	return cerID, nil
}

func (c *CertificateV2WebController) realCreate(ctx context.Context, loginID int64, udid, note string, priceID int64, memberDevice *models.MemberDevice) (int64, error) {
	price, err := c.certificatePriceCtl.GetPriceByID(ctx, loginID, priceID)
	if err != nil {
		return 0, err
	}
	payCount := price.Price

	// 判断 D 币是否充足
	if err = c.memberPayCountCtl.CheckPayCount(ctx, loginID, payCount); err != nil {
		return 0, err
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	cerID := util2.MustGenerateID(ctx)

	// 扣币
	if err = c.memberPayCountCtl.DeductPayCount(ctx, loginID, payCount, enum.MemberPayCountStatusUsed, enum.MemberPayCountUseCertificate, datatype.MemberPayCountRecordBizExt{
		ObjectID:   cerID,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeCertificate,
	}); err != nil {
		return 0, err
	}

	// 生成证书
	if err = c.createCer(ctx, loginID, udid, note, priceID, memberDevice, false, cerID); err != nil {
		return 0, err
	}

	// 记录当前这本证书最新购买时间
	if err = c.certificateDeviceDAO.Insert(ctx, &models.CertificateDevice{
		DeviceID:      memberDevice.ID,
		CertificateID: cerID,
		BizExt:        datatype.CertificateDeviceBizExt{},
	}); err != nil {
		return 0, err
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return cerID, nil
}

func (c *CertificateV2WebController) replenish(ctx context.Context, loginID int64, udid, note string, priceID int64, memberDevice *models.MemberDevice) (int64, error) {
	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	cerID := util2.MustGenerateID(ctx)

	// 生成记录
	if err := c.memberPayCountRecordDAO.Insert(ctx, &models.MemberPayCountRecord{
		MemberID: loginID,
		Type:     enum.MemberPayCountRecordTypeReplenishCertificate,
		Count:    0,
		BizExt: datatype.MemberPayCountRecordBizExt{
			ObjectID:   cerID,
			ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeCertificate,
		},
	}); err != nil {
		return 0, nil
	}

	// 生成证书
	if err := c.createCer(ctx, loginID, udid, note, priceID, memberDevice, true, cerID); err != nil {
		return 0, err
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return cerID, nil
}

func (c *CertificateV2WebController) createCer(ctx context.Context, loginID int64, udid, note string, priceID int64, memberDevice *models.MemberDevice, isReplenish bool, cerID int64) error {
	resp, alterMsg, err := c.certificateBaseCtl.Create(ctx, udid)
	if alterMsg != "" {
		c.alterWebCtl.SendCreateCertificateFailMsg(ctx, loginID, memberDevice.ID, alterMsg)
	}
	if err != nil {
		return err
	}

	/// 记录证书等级, 方便后期候补
	resp.Response.BizExt.Level = cast.ToInt(priceID)

	/// 记录证书备注
	resp.Response.BizExt.Note = note
	/// 记录是否为售后证书
	resp.Response.BizExt.IsReplenish = isReplenish

	/// 计算证书 md5
	p12FileData := resp.Response.P12Data
	mpFileData := resp.Response.MobileProvisionData
	p12FileMd5 := util2.StringMd5(p12FileData)
	mpFileMd5 := util2.StringMd5(mpFileData)

	if err = c.certificateDAO.Insert(ctx, &models.CertificateV2{
		ID:                         cerID,
		DeviceID:                   memberDevice.ID,
		P12FileData:                p12FileData,
		P12FileDataMD5:             p12FileMd5,
		ModifiedP12FileDate:        resp.ModifiedP12FileData,
		MobileProvisionFileData:    mpFileData,
		MobileProvisionFileDataMD5: mpFileMd5,
		Source:                     resp.Response.Source,
		BizExt:                     resp.Response.BizExt,
	}); err != nil {
		return err
	}
	return nil
}

func (c *CertificateV2WebController) getMemberDevice(ctx context.Context, loginID int64, UDID string) (*models.MemberDevice, error) {
	// 绑定设备
	memberDevice, err := c.memberDeviceDAO.GetByMemberIDUdidSafe(ctx, loginID, UDID)
	if err != nil {
		return nil, err
	}

	if memberDevice != nil {
		return memberDevice, nil
	}
	id := util2.MustGenerateID(ctx)
	err = c.memberDeviceDAO.Insert(ctx, &models.MemberDevice{
		ID:       id,
		MemberID: loginID,
		Udid:     UDID,
		BizExt:   datatype.MemberDeviceBizExt{},
	})
	if err != nil {
		return nil, err
	}
	return c.memberDeviceDAO.Get(ctx, id)
}
