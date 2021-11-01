package impl

import (
	"context"
	"dumpapp_server/pkg/common/errors"
	pkgErr "github.com/pkg/errors"

	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
)

type MemberIDEncryptionController struct {
	accountDAO            dao.AccountDAO
	memberIDEncryptionDAO dao.MemberIDEncryptionDAO
}

var DefaultMemberIDEncryptionController *MemberIDEncryptionController

func init() {
	DefaultMemberIDEncryptionController = NewMemberIDEncryptionController()
}

func NewMemberIDEncryptionController() *MemberIDEncryptionController {
	return &MemberIDEncryptionController{
		accountDAO:            impl.DefaultAccountDAO,
		memberIDEncryptionDAO: impl.DefaultMemberIDEncryptionDAO,
	}
}

func (c *MemberIDEncryptionController) GetCodeByMemberID(ctx context.Context, memberID int64) (string, error) {
	e, err := c.memberIDEncryptionDAO.GetByMemberID(ctx, memberID)
	if err != nil {
		return "", err
	}
	return e.Code, nil
}

func (c *MemberIDEncryptionController) GetMemberIDByCode(ctx context.Context, code string) (int64, error) {
	e, err := c.memberIDEncryptionDAO.GetByCode(ctx, code)
	if err != nil && pkgErr.Cause(err) != errors.ErrNotFound {
		return 0, err
	}
	if e == nil {
		return 0, errors.ErrNotFound
	}
	_, err = c.accountDAO.Get(ctx, e.MemberID)
	if err != nil {
		return 0, err
	}
	return e.MemberID, nil
}
