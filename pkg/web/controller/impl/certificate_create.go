package impl

import (
	"context"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/web/controller"
)

type CertificateCreateWebController struct {
	adminConfigInfoDAO  dao.AdminConfigInfoDAO
	certificateV1WebCtl controller.CertificateWebController
	certificateV2WebCtl controller.CertificateV2WebController
}

var DefaultCertificateCreateWebController *CertificateCreateWebController

func init() {
	DefaultCertificateCreateWebController = NewCertificateCreateWebController()
}

func NewCertificateCreateWebController() *CertificateCreateWebController {
	return &CertificateCreateWebController{
		adminConfigInfoDAO:  impl.DefaultAdminConfigInfoDAO,
		certificateV1WebCtl: NewCertificateWebController(),
		certificateV2WebCtl: NewCertificateV2WebController(),
	}
}

func (c *CertificateCreateWebController) Create(ctx context.Context, loginID int64, udid, note string, priceID int64, isReplenish bool, payType string) (int64, error) {
	config, err := c.adminConfigInfoDAO.GetConfig(ctx)
	if err != nil {
		return c.certificateV1WebCtl.PayCertificate(ctx, loginID, udid, note, priceID, isReplenish, payType)
	}
	if !config.BizExt.OpenReplenish {
		return c.certificateV1WebCtl.PayCertificate(ctx, loginID, udid, note, priceID, isReplenish, payType)
	}

	// 打开开关后则走新版创建证书
	return c.certificateV2WebCtl.Create(ctx, loginID, udid, note, priceID)
}
