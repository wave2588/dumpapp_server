package impl

import (
	"context"
	"time"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

type CertificateBaseController struct {
	certificateDAO     dao.CertificateV2DAO
	memberDeviceDAO    dao.MemberDeviceDAO
	accountDAO         dao.AccountDAO
	adminConfigInfoDAO dao.AdminConfigInfoDAO
	certificateV1Ctl   controller.CertificateController
	certificateV2Ctl   controller.CertificateController
	certificateV3Ctl   controller.CertificateController
}

var DefaultCertificateBaseController controller.CertificateBaseController

func init() {
	DefaultCertificateBaseController = NewCertificateBaseController()
}

func NewCertificateBaseController() *CertificateBaseController {
	return &CertificateBaseController{
		certificateDAO:     impl.DefaultCertificateV2DAO,
		memberDeviceDAO:    impl.DefaultMemberDeviceDAO,
		accountDAO:         impl.DefaultAccountDAO,
		adminConfigInfoDAO: impl.DefaultAdminConfigInfoDAO,
		certificateV1Ctl:   NewCertificateV1Controller(),
		certificateV2Ctl:   NewCertificateV2Controller(),
		certificateV3Ctl:   NewCertificateV3Controller(),
	}
}

func (c *CertificateBaseController) CheckCertificateIsActive(ctx context.Context, ids []int64) (map[int64]bool, error) {
	certificateMap, err := c.certificateDAO.BatchGet(ctx, ids)
	if err != nil {
		return nil, err
	}
	certificates := make([]*models.CertificateV2, 0)
	for _, id := range ids {
		cer, ok := certificateMap[id]
		if !ok {
			continue
		}
		certificates = append(certificates, cer)
	}
	return c.CheckCertificateIsActiveByModels(ctx, certificates)
}

func (c *CertificateBaseController) CheckCertificateIsActiveByModels(ctx context.Context, certificates models.CertificateV2Slice) (map[int64]bool, error) {
	isActiveMap := make(map[int64]bool)
	batch := util2.NewBatch(ctx)
	for _, certificate := range certificates {
		batch.Append(func(cer *models.CertificateV2) util2.FutureFunc {
			return func() error {
				switch cer.Source {
				case enum.CertificateSourceV1:
					response, err := c.certificateV1Ctl.CheckCerIsActive(ctx, cer.ID)
					if err != nil {
						return err
					}
					isActiveMap[cer.ID] = response
				case enum.CertificateSourceV2:
					response, err := c.certificateV2Ctl.CheckCerIsActive(ctx, cer.ID)
					if err != nil {
						return err
					}
					isActiveMap[cer.ID] = response
				case enum.CertificateSourceV3:
					response, err := c.certificateV3Ctl.CheckCerIsActive(ctx, cer.ID)
					if err != nil {
						return err
					}
					isActiveMap[cer.ID] = response
				}
				return nil
			}
		}(certificate))
	}
	batch.Wait()

	result := make(map[int64]bool)
	for _, cer := range certificates {
		if res, ok := isActiveMap[cer.ID]; ok {
			result[cer.ID] = res
		} else {
			result[cer.ID] = true // todo:  如果没获取到, 默认展示有效
		}
	}
	return result, nil
}

func (c *CertificateBaseController) GetCertificateReplenishExpireAt(ctx context.Context, ids []int64) (map[int64]time.Time, error) {
	certificateMap, err := c.certificateDAO.BatchGet(ctx, ids)
	if err != nil {
		return nil, err
	}
	certificates := make([]*models.CertificateV2, 0)
	for _, id := range ids {
		cer, ok := certificateMap[id]
		if !ok {
			continue
		}
		certificates = append(certificates, cer)
	}
	return c.GetCertificateReplenishExpireAtByModels(ctx, certificates)
}

func (c *CertificateBaseController) GetCertificateReplenishExpireAtByModels(ctx context.Context, certificates models.CertificateV2Slice) (map[int64]time.Time, error) {
	// 获取设备
	deviceIDs := make([]int64, 0)
	for _, certificate := range certificates {
		deviceIDs = append(deviceIDs, certificate.DeviceID)
	}
	deviceIDs = util2.RemoveDuplicates(deviceIDs)
	mDeviceMap, err := c.memberDeviceDAO.BatchGet(ctx, deviceIDs)
	if err != nil {
		return nil, err
	}
	// 获取 member_ids
	memberIDs := make([]int64, 0)
	for _, device := range mDeviceMap {
		memberIDs = append(memberIDs, device.MemberID)
	}
	memberIDs = util2.RemoveDuplicates(memberIDs)

	accountMap, err := c.accountDAO.BatchGet(ctx, memberIDs)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]time.Time)
	for _, certificate := range certificates {
		device, ok := mDeviceMap[certificate.DeviceID]
		if !ok {
			continue
		}
		account, ok := accountMap[device.MemberID]
		if !ok {
			continue
		}
		replenishExpireAt := certificate.CreatedAt
		switch certificate.BizExt.Level {
		case 1:
			if account.Role == enum.AccountRoleNone {
				replenishExpireAt = replenishExpireAt.AddDate(0, 0, 7) // 7 天售后
			} else {
				replenishExpireAt = replenishExpireAt.AddDate(0, 0, 30) // 如果是代理商普通版给 30 天售后
			}
		case 2:
			replenishExpireAt = replenishExpireAt.AddDate(0, 0, 180) // 180 天售后
		case 3:
			replenishExpireAt = replenishExpireAt.AddDate(0, 0, 365) // 365 天售后
		}
		result[certificate.ID] = replenishExpireAt
	}
	return result, nil
}
