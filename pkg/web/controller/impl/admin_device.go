package impl

import (
	"context"

	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
)

type AdminDeviceWebController struct {
	accountDAO      dao.AccountDAO
	memberDeviceDAO dao.MemberDeviceDAO
}

var DefaultAdminDeviceWebController *AdminDeviceWebController

func init() {
	DefaultAdminDeviceWebController = NewAdminDeviceWebController()
}

func NewAdminDeviceWebController() *AdminDeviceWebController {
	return &AdminDeviceWebController{
		accountDAO:      impl.DefaultAccountDAO,
		memberDeviceDAO: impl.DefaultMemberDeviceDAO,
	}
}

func (c *AdminDeviceWebController) Unbind(ctx context.Context, email, udid string) error {
	//accountMap, err := c.accountDAO.BatchGetByEmail(ctx, []string{email})
	//if err != nil {
	//	return err
	//}
	//_, ok := accountMap[email]
	//if !ok {
	//	return errors.ErrNotFoundMember
	//}
	//
	//deviceMap, err := c.memberDeviceDAO.BatchGetByUdid(ctx, []string{udid})
	//if err != nil {
	//	return err
	//}
	//device, ok := deviceMap[udid]
	//if !ok {
	//	return errors.ErrDeviceNotFound
	//}
	//
	//if err = c.memberDeviceDAO.Delete(ctx, device.ID); err != nil {
	//	return err
	//}

	return nil
}
