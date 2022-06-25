package install_app_handler

import (
	"errors"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/install_app"
	"dumpapp_server/pkg/controller/install_app/impl"
	"github.com/spf13/cast"
)

type CallbackPayHandler struct {
	aliPayCtl install_app.ALiPayInstallAppController
}

func NewCallbackPayHandler() *CallbackPayHandler {
	return &CallbackPayHandler{
		aliPayCtl: impl.DefaultALiPayInstallAppController,
	}
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
	err := h.aliPayCtl.AliPayCallbackOrder(ctx, orderID)
	util.PanicIf(err)
}
