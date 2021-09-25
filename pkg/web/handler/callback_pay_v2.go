package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/spf13/cast"
)

type CallbackPayV2Handler struct {
	memberDownloadOrderDAO  dao.MemberDownloadOrderDAO
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO

	alipayCtl   controller.ALiPayController
	alterWebCtl controller2.AlterWebController
}

func NewCallbackPayV2Handler() *CallbackPayV2Handler {
	return &CallbackPayV2Handler{
		memberDownloadOrderDAO:  impl.DefaultMemberDownloadOrderDAO,
		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,

		alipayCtl:   impl2.DefaultALiPayController,
		alterWebCtl: impl3.DefaultAlterWebController,
	}
}

func (h *CallbackPayV2Handler) ALiPayCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	util.PanicIf(r.ParseForm())
	orderIDStrings := r.PostForm["out_trade_no"]
	var orderIDString string
	if len(orderIDStrings) >= 1 {
		orderIDString = orderIDStrings[0]
	}
	orderID := cast.ToInt64(orderIDString)
	if orderID == 0 {
		panic(errors.New(fmt.Sprintf("orderID 错啦 orderIDString: %s  orderID: %d  orderIDStrings: %v", orderIDString, orderID, orderIDStrings)))
	}

	util.PanicIf(h.alipayCtl.CheckPayStatus(ctx, orderID))

	order, err := h.memberDownloadOrderDAO.Get(ctx, orderID)
	util.PanicIf(err)

	/// 支付成功的订单即可忽略
	if order.Status == enum.MemberDownloadOrderStatusPaid {
		return
	}

	order.Status = enum.MemberDownloadOrderStatusPaid

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	util.PanicIf(h.memberDownloadOrderDAO.Update(ctx, order))

	number := cast.ToInt(order.Number)
	if number >= 3 && number < 5 { /// 买三送一
		number += 1
	} else if number >= 5 && number < 7 { /// 买五送二
		number += 2
	} else if number >= 7 { /// 买七送三
		number += 3
	}

	for i := 0; i < number; i++ {
		util.PanicIf(h.memberDownloadNumberDAO.Insert(ctx, &models.MemberDownloadNumber{
			MemberID: order.MemberID,
			Status:   enum.MemberDownloadNumberStatusNormal,
		}))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	h.alterWebCtl.SendPaidOrderMsg(ctx, orderID)
}
