package handler

import (
	"errors"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/spf13/cast"
)

type CallbackPayV3Handler struct {
	memberPayOrderDAO    dao.MemberPayOrderDAO
	memberPayOrderWebCtl controller2.MemberPayOrderWebController
}

func NewCallbackPayV3Handler() *CallbackPayV3Handler {
	return &CallbackPayV3Handler{
		memberPayOrderDAO:    impl.DefaultMemberPayOrderDAO,
		memberPayOrderWebCtl: impl3.DefaultMemberPayOrderWebController,
	}
}

func (h *CallbackPayV3Handler) ALiPayCallback(w http.ResponseWriter, r *http.Request) {
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
	err := h.memberPayOrderWebCtl.AliPayCallbackOrder(ctx, orderID)
	util.PanicIf(err)
}
