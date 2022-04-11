package impl

import (
	"context"
	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"github.com/spf13/cast"
)

type MemberPayOrderWebController struct {
	alipayCtl             controller.ALiPayV3Controller
	accountDAO            dao.AccountDAO
	memberPayOrderDAO     dao.MemberPayOrderDAO
	memberPayCountDAO     dao.MemberPayCountDAO
	memberInviteDAO       dao.MemberInviteDAO
	memberRebateRecordDAO dao.MemberRebateRecordDAO
}

var DefaultMemberPayOrderWebController *MemberPayOrderWebController

func init() {
	DefaultMemberPayOrderWebController = NewMemberPayOrderWebController()
}

func NewMemberPayOrderWebController() *MemberPayOrderWebController {
	return &MemberPayOrderWebController{
		alipayCtl:             impl.DefaultALiPayV3Controller,
		accountDAO:            impl2.DefaultAccountDAO,
		memberPayOrderDAO:     impl2.DefaultMemberPayOrderDAO,
		memberPayCountDAO:     impl2.DefaultMemberPayCountDAO,
		memberInviteDAO:       impl2.DefaultMemberInviteDAO,
		memberRebateRecordDAO: impl2.DefaultMemberRebateRecordDAO,
	}
}

func (c *MemberPayOrderWebController) AliPayCallbackOrder(ctx context.Context, orderID int64) error {
	util.PanicIf(c.alipayCtl.CheckPayStatus(ctx, orderID))

	order, err := c.memberPayOrderDAO.Get(ctx, orderID)
	if err != nil {
		return err
	}

	/// 支付成功的订单即可忽略
	if order.Status == enum.MemberPayOrderStatusPaid {
		return nil
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	order.Status = enum.MemberPayOrderStatusPaid
	util.PanicIf(c.memberPayOrderDAO.Update(ctx, order))

	number := cast.ToInt(order.Amount)
	for i := 0; i < number; i++ {
		err := c.memberPayCountDAO.Insert(ctx, &models.MemberPayCount{
			MemberID: order.MemberID,
			Status:   enum.MemberPayCountStatusNormal,
			Source:   enum.MemberPayCountSourceNormal,
		})
		if err != nil {
			return err
		}
	}

	/// 多买多送，买 27 送 9，买 45 送 18，买 63 送 27。
	freeNumber := 0
	if number >= 27 && number < 45 {
		freeNumber = 9
	} else if number >= 45 && number < 63 {
		freeNumber = 18
	} else if number >= 63 {
		freeNumber = 27
	}
	for i := 0; i < freeNumber; i++ {
		err := c.memberPayCountDAO.Insert(ctx, &models.MemberPayCount{
			MemberID: order.MemberID,
			Status:   enum.MemberPayCountStatusNormal,
			Source:   enum.MemberPayCountSourcePayForFree,
		})
		if err != nil {
			return err
		}
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	/// 赠送失败了先不处理
	_ = c.rebaseRecord(ctx, order)

	return nil
}

func (c *MemberPayOrderWebController) rebaseRecord(ctx context.Context, order *models.MemberPayOrder) error {

	inviteeID := order.MemberID

	/// 支付成功后要送邀请者
	inviteMap, err := c.memberInviteDAO.BatchGetByInviteeID(ctx, []int64{inviteeID})
	if err != nil {
		return err
	}

	invite, ok := inviteMap[inviteeID]
	if !ok {
		return nil
	}

	inviterID := invite.InviterID
	accountMap, err := c.accountDAO.BatchGet(ctx, []int64{inviterID})
	if err != nil {
		return nil
	}

	account, ok := accountMap[inviterID]
	if !ok {
		return nil
	}

	/// 正常是返还 30%，如果是大 V 则返还 60%
	ratio := 0.3
	if account.Role == enum.AccountRoleInfluential {
		ratio = 0.6
	}

	count := cast.ToUint(order.Amount * ratio)
	return c.memberRebateRecordDAO.Insert(ctx, &models.MemberRebateRecord{
		OrderID:          order.ID,
		ReceiverMemberID: inviterID,
		Count:            count,
	})
}
