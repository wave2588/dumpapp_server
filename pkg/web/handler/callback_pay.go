package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
)

type CallbackPayHandler struct {
	memberVipDAO      dao.MemberVipDAO
	memberVipOrderDAO dao.MemberVipOrderDAO

	alipayCtl controller.ALiPayController
}

func NewCallbackPayHandler() *CallbackPayHandler {
	return &CallbackPayHandler{
		memberVipDAO:      impl.DefaultMemberVipDAO,
		memberVipOrderDAO: impl.DefaultMemberVipOrderDAO,

		alipayCtl: impl2.DefaultALiPayController,
	}
}

type Tracks struct {
	out_trade_no string
}

func (h *CallbackPayHandler) ALiPayCallback(w http.ResponseWriter, r *http.Request) {
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

	duration := enum.MemberVipDurationTypeOneMonth

	order, err := h.memberVipOrderDAO.Get(ctx, orderID)
	util.PanicIf(err)

	if du, ok := constant.DurationToMemberVipDurationType[order.Duration.String]; ok {
		duration = du
	}

	util.PanicIf(h.alipayCtl.CheckPayStatus(ctx, orderID))

	order.Status = enum.MemberVipOrderStatusPaid

	account := GetAccountByLoginID(ctx, order.MemberID)
	memberVip, err := h.memberVipDAO.Get(ctx, order.MemberID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	util.PanicIf(h.memberVipOrderDAO.Update(ctx, order))

	endAt := constant.MemberVipDurationTypeToDays[duration]
	if memberVip == nil {
		util.PanicIf(h.memberVipDAO.Insert(ctx,
			&models.MemberVip{
				ID:      account.ID,
				StartAt: time.Now(),
				EndAt:   endAt,
			}))
	} else {
		days := util2.GetDiffDays(endAt, memberVip.EndAt)
		memberVip.EndAt = memberVip.EndAt.AddDate(0, 0, days)
		util.PanicIf(h.memberVipDAO.Update(ctx, memberVip))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)
}
