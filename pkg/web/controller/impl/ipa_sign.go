package impl

import (
	"context"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
)

type IpaSignWebController struct {
	ipaVersionDAO dao.IpaVersionDAO
	ipaSignDAO    dao.IpaSignDAO
}

var DefaultIpaSignWebController *IpaSignWebController

func init() {
	DefaultIpaSignWebController = NewIpaSignWebController()
}

func NewIpaSignWebController() *IpaSignWebController {
	return &IpaSignWebController{
		ipaVersionDAO: impl.DefaultIpaVersionDAO,
		ipaSignDAO:    impl.DefaultIpaSignDAO,
	}
}

func (c *IpaSignWebController) AddSignTask(ctx context.Context, loginID, certificateID, ipaVersionID int64) error {
	/// fixme: 检查次数

	/// 检测 ipaVersionID
	ipaVersion, err := c.checkIpaVersionID(ctx, ipaVersionID)
	if err != nil {
		return err
	}
	/// 检测 certificateID
	_, err = c.checkCertificateID(ctx, loginID, certificateID)
	if err != nil {
		return err
	}

	bizExt := constant.IpaSignBizExt{
		IpaVersionID: ipaVersionID,
		IpaVersion:   ipaVersion.Version,
		IpaType:      ipaVersion.IpaType,
		TokenPath:    ipaVersion.TokenPath,
	}
	return c.ipaSignDAO.Insert(ctx, &models.IpaSign{
		IpaID:         ipaVersion.IpaID,
		CertificateID: certificateID,
		MemberID:      loginID,
		Status:        enum.IpaSignStatusUnprocessed,
		BizExt:        bizExt.String(),
	})
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
