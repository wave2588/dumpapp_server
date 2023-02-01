package impl

import (
	"context"
	"time"

	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
)

type CertificateDeviceController struct {
	accountDAO           dao.AccountDAO
	certificateDeviceDAO dao.CertificateDeviceDAO
	memberDeviceDAO      dao.MemberDeviceDAO
	certificateDAO       dao.CertificateV2DAO
	certificateBaseCtl   controller.CertificateBaseController
}

var DefaultCertificateDeviceController controller.CertificateDeviceController

func init() {
	DefaultCertificateDeviceController = NewCertificateDeviceController()
}

func NewCertificateDeviceController() *CertificateDeviceController {
	return &CertificateDeviceController{
		accountDAO:           impl.DefaultAccountDAO,
		certificateDeviceDAO: impl.DefaultCertificateDeviceDAO,
		memberDeviceDAO:      impl.DefaultMemberDeviceDAO,
		certificateDAO:       impl.DefaultCertificateV2DAO,
		certificateBaseCtl:   NewCertificateBaseController(),
	}
}

func (c *CertificateDeviceController) IsReplenish(ctx context.Context, memberID int64, UDID string) (bool, error) {
	mDevice, err := c.memberDeviceDAO.GetByMemberIDUdidSafe(ctx, memberID, UDID)
	if err != nil {
		return false, err
	}
	// 没找到这个设备，说明就要扣币购买证书
	if mDevice == nil {
		return false, nil
	}
	deviceID := mDevice.ID

	// 这个设备最新的一本证书
	cd, err := c.certificateDeviceDAO.GetLastByDeviceID(ctx, deviceID)
	if err != nil {
		return false, err
	}
	// 说明此设备没有购买过证书，则进行扣费操作
	if cd == nil {
		return false, nil
	}

	certificateID := cd.CertificateID

	// 获取证书
	cerMap, err := c.certificateDAO.BatchGet(ctx, []int64{certificateID})
	if err != nil {
		return false, err
	}
	cer, ok := cerMap[certificateID]
	if !ok {
		return false, nil
	}

	// 检测证书候补时间
	cerReplenishExpireAtMap, err := c.certificateBaseCtl.GetCertificateReplenishExpireAtByModels(ctx, []*models.CertificateV2{cer})
	if err != nil {
		return false, err
	}

	cerReplenishExpireAt := cerReplenishExpireAtMap[cer.ID]

	// 如果证书候补时间大于当前时间，则不能候补
	if cerReplenishExpireAt.Unix() >= time.Now().Unix() {
		return false, nil
	}

	// 检测证书是否有效
	isActiveMap, _ := c.certificateBaseCtl.CheckCertificateIsActiveByModels(ctx, []*models.CertificateV2{cer})
	// 如果证书还有效，则不进行候补，进行扣币操作
	if isActive := isActiveMap[cer.ID]; isActive {
		return false, nil
	}

	// 允许候补
	return true, nil
}
