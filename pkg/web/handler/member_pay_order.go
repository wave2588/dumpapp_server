package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
)

type MemberPayOrderHandler struct {
	aliPayCtl         controller.ALiPayV3Controller
	memberPayOrderCtl controller.MemberPayOrderController
}

func NewMemberPayOrderHandler() *MemberPayOrderHandler {
	return &MemberPayOrderHandler{
		aliPayCtl:         impl.DefaultALiPayV3Controller,
		memberPayOrderCtl: impl.DefaultMemberPayOrderController,
	}
}

type getMemberPayOrderArgs struct {
	Number int64 `form:"number" validate:"required"`
}

func (args *getMemberPayOrderArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if args.Number <= 0 {
		return errors.UnproccessableError("number > 0")
	}
	if args.Number > 99999 {
		return errors.UnproccessableError("创建订单失败")
	}
	return nil
}

func (h *MemberPayOrderHandler) GetPayOrderURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getMemberPayOrderArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	loginID := mustGetLoginID(ctx)

	var (
		orderID int64
		payURL  string
		err     error
	)
	platform := ctx.Value(constant.CtxKeyAppPlatform).(string)
	if platform == "ios" {
		orderID, payURL, err = h.aliPayCtl.GetPhoneWapPayURLByNumber(ctx, loginID, args.Number)
	} else {
		orderID, payURL, err = h.aliPayCtl.GetPayURLByNumber(ctx, loginID, args.Number)
	}
	util.PanicIf(err)

	res := map[string]interface{}{
		"order_id": cast.ToString(orderID),
		"open_url": payURL,
	}
	util.RenderJSON(w, res)
}

func (h *MemberPayOrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		loginID = mustGetLoginID(ctx)
		orderID = cast.ToInt64(util.URLParam(r, "order_id"))
	)

	orderMap := render.NewMemberPayOrderRender([]int64{orderID}, loginID, render.MemberPayOrderDefaultRenderFields...).RenderMap(ctx)
	order, ok := orderMap[orderID]
	if !ok {
		util.PanicIf(errors.ErrMemberPayOrderNotFound)
	}
	util.RenderJSON(w, order)
}

func (h *MemberPayOrderHandler) GetOrderRule(w http.ResponseWriter, r *http.Request) {
	util.RenderJSON(w, h.memberPayOrderCtl.GetPayCampaignRule())
}
