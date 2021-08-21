package handler

import (
	"context"
	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"net/http"
)

type CallbackPayV2Handler struct {
	memberDownloadOrderDAO  dao.MemberDownloadOrderDAO
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO

	alipayCtl controller.ALiPayController
}

func NewCallbackPayV2Handler() *CallbackPayV2Handler {
	return &CallbackPayV2Handler{
		memberDownloadOrderDAO:  impl.DefaultMemberDownloadOrderDAO,
		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,

		alipayCtl: impl2.DefaultALiPayController,
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

	order.Status = enum.MemberDownloadOrderStatusPaid

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	util.PanicIf(h.memberDownloadOrderDAO.Update(ctx, order))

	for i := 0; i <= cast.ToInt(order.Number); i++ {
		util.PanicIf(h.memberDownloadNumberDAO.Insert(ctx, &models.MemberDownloadNumber{
			MemberID: order.MemberID,
			Status:   enum.MemberDownloadNumberStatusNormal,
		}))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)
}
