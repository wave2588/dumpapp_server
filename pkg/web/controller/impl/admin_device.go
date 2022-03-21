package impl

import (
	"context"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
)

type AdminDeviceWebController struct {
	accountDAO           dao.AccountDAO
	memberDeviceDAO      dao.MemberDeviceDAO
	certificateDeviceDAO dao.CertificateDeviceDAO
}

var DefaultAdminDeviceWebController *AdminDeviceWebController

func init() {
	DefaultAdminDeviceWebController = NewAdminDeviceWebController()
}

func NewAdminDeviceWebController() *AdminDeviceWebController {
	return &AdminDeviceWebController{
		accountDAO:           impl.DefaultAccountDAO,
		memberDeviceDAO:      impl.DefaultMemberDeviceDAO,
		certificateDeviceDAO: impl.DefaultCertificateDeviceDAO,
	}
}

func (c *AdminDeviceWebController) Unbind(ctx context.Context, email, udid string) error {
	accountMap, err := c.accountDAO.BatchGetByEmail(ctx, []string{email})
	if err != nil {
		return err
	}
	_, ok := accountMap[email]
	if !ok {
		return errors.ErrNotFoundMember
	}

	deviceMap, err := c.memberDeviceDAO.BatchGetByUdid(ctx, []string{udid})
	if err != nil {
		return err
	}
	device, ok := deviceMap[udid]
	if !ok {
		return errors.ErrDeviceNotFound
	}

	cs, err := c.certificateDeviceDAO.GetCertificateDeviceSliceByDeviceID(ctx, device.ID)
	if err != nil {
		return err
	}
	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	if err = c.memberDeviceDAO.Delete(ctx, device.ID); err != nil {
		return err
	}

	for _, certificateDevice := range cs {
		if err = c.certificateDeviceDAO.Delete(ctx, certificateDevice.ID); err != nil {
			return err
		}
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return nil
}
