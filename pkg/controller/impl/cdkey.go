package impl

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"github.com/volatiletech/strmangle"
)

type CdKeyController struct {
	memberDeviceDAO          dao.MemberDeviceDAO
	certificateDAO           dao.CertificateV2DAO
	installAppCdKeyDAO       dao.InstallAppCdkeyDAO
	installAppCertificateDAO dao.InstallAppCertificateDAO
}

var DefaultCdKeyController *CdKeyController

func init() {
	DefaultCdKeyController = NewCdKeyController()
}

func NewCdKeyController() *CdKeyController {
	return &CdKeyController{
		memberDeviceDAO:          impl.DefaultMemberDeviceDAO,
		certificateDAO:           impl.DefaultCertificateV2DAO,
		installAppCdKeyDAO:       impl.DefaultInstallAppCdkeyDAO,
		installAppCertificateDAO: impl.DefaultInstallAppCertificateDAO,
	}
}

func (c *CdKeyController) AddCdKeyByMemberBuyCertificate(ctx context.Context, certificateID int64) error {
	cerMap, err := c.certificateDAO.BatchGet(ctx, []int64{certificateID})
	if err != nil {
		return err
	}
	cer, ok := cerMap[certificateID]
	if !ok {
		return nil
	}

	deviceMap, err := c.memberDeviceDAO.BatchGet(ctx, []int64{cer.DeviceID})
	if err != nil {
		return err
	}
	device, ok := deviceMap[cer.DeviceID]
	if !ok {
		return nil
	}

	outIDs, err := c.GetOutIDs(ctx, 1, cer.BizExt.Level)
	if err != nil {
		return err
	}
	if len(outIDs) == 0 {
		return nil
	}
	outID := outIDs[0]
	cer.BizExt.CdKeyOutID = outID

	installAppCerID := util2.MustGenerateID(ctx)
	if err = c.installAppCertificateDAO.Insert(ctx, &models.InstallAppCertificate{
		ID:                         installAppCerID,
		Udid:                       device.Udid,
		P12FileData:                cer.P12FileData,
		P12FileDataMD5:             cer.P12FileDataMD5,
		ModifiedP12FileDate:        cer.ModifiedP12FileDate,
		MobileProvisionFileData:    cer.MobileProvisionFileData,
		MobileProvisionFileDataMD5: cer.MobileProvisionFileDataMD5,
		Source:                     cer.Source,
		BizExt:                     cer.BizExt,
	}); err != nil {
		return err
	}

	if err = c.installAppCdKeyDAO.Insert(ctx, &models.InstallAppCdkey{
		ID:            util2.MustGenerateID(ctx),
		OutID:         outID,
		Status:        enum.InstallAppCDKeyStatusUsed,
		CertificateID: installAppCerID,
		OrderID:       0, // 0 就代表是自动生成的
	}); err != nil {
		return err
	}

	return c.certificateDAO.Update(ctx, cer)
}

func (c *CdKeyController) GetOutIDs(ctx context.Context, number, cerLevel int) ([]string, error) {
	suffix := constant.GetInstallAppCDKeySuffix(cerLevel)
	outIDs := make([]string, 0)
	/// 生成 number * 10 的数量，以防重复
	for i := 0; i < number*10; i++ {
		outID := fmt.Sprintf("%s%s", util2.MustGenerateAppCDKEY(), suffix)
		outIDs = append(outIDs, outID)
	}
	outIDs = strmangle.RemoveDuplicates(outIDs)

	cMap, err := c.installAppCdKeyDAO.BatchGetByOutID(ctx, outIDs)
	if err != nil {
		return nil, err
	}

	resultOutIDs := make([]string, 0)
	for _, oID := range outIDs {
		if len(resultOutIDs) == number {
			return resultOutIDs, nil
		}
		if _, ok := cMap[oID]; !ok {
			resultOutIDs = append(resultOutIDs, oID)
		}
	}
	return nil, errors2.ErrInstallAppGenerateCDKeyFail
}
