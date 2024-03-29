package impl

import (
	"context"
	"math"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	"github.com/spf13/cast"
)

type MemberPayOrderWebController struct {
	alipayCtl             controller.ALiPayV3Controller
	memberPayCountCtl     controller.MemberPayCountController
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
		memberPayCountCtl:     impl.DefaultMemberPayCountController,
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

	memberID := order.MemberID
	accountMap, err := c.accountDAO.BatchGet(ctx, []int64{memberID})
	if err != nil {
		return err
	}
	account, ok := accountMap[memberID]
	if !ok {
		return errors2.ErrNotFoundMember
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	order.Status = enum.MemberPayOrderStatusPaid
	util.PanicIf(c.memberPayOrderDAO.Update(ctx, order))

	number := cast.ToInt64(order.Amount)
	util.PanicIf(c.memberPayCountCtl.AddCount(ctx, order.MemberID, number, enum.MemberPayCountSourceNormal, datatype.MemberPayCountRecordBizExt{
		ObjectID:   orderID,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeOrder,
	}))

	/// 多买多送，冲 500 送 30, 冲 1000 送 60
	/*
		500 - 15
		1000 - 70
		2000 - 260
		5000 - 1290
		10000 - 3330
	*/
	freeNumber := int64(0)
	if number >= 500 && number < 1500 {
		freeNumber = 15
	} else if number >= 1500 && number < 3000 {
		freeNumber = 100
	} else if number >= 3000 && number < 5000 {
		freeNumber = 300
	} else if number >= 5000 {
		freeNumber = 800
	}

	/// 充值大于等于 1000 则自动升级为代理商
	if number >= 1000 {
		account.Role = enum.AccountRoleAgent
		if err = c.accountDAO.Update(ctx, account); err != nil {
			return err
		}
	}

	util.PanicIf(c.memberPayCountCtl.AddCount(ctx, order.MemberID, freeNumber, enum.MemberPayCountSourcePayForFree, datatype.MemberPayCountRecordBizExt{
		ObjectID:   orderID,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeOrder,
	}))

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	/// 赠送失败了先不处理
	_ = c.rebaseRecord(ctx, order)

	return nil
}

func (c *MemberPayOrderWebController) rebaseRecord(ctx context.Context, order *models.MemberPayOrder) error {
	if order.Amount < 10 {
		return nil
	}

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

	/// 正常是返还 0.05%，如果是大 V 则返还 10%
	ratio := 0.05
	if account.Role == enum.AccountRoleInfluential {
		ratio = 0.1
	}

	/// 写入返还次数
	count := cast.ToInt(math.Ceil(order.Amount * ratio))

	err = c.memberPayCountCtl.AddCount(ctx, inviterID, int64(count), enum.MemberPayCountSourceRebate, datatype.MemberPayCountRecordBizExt{
		ObjectID:   order.ID,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeOrder,
	})
	if err != nil {
		return err
	}

	return c.memberRebateRecordDAO.Insert(ctx, &models.MemberRebateRecord{
		OrderID:          order.ID,
		ReceiverMemberID: inviterID,
		Count:            count,
	})
}
