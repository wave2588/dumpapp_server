package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
)

type MemberPayOrderHandler struct {
	aliPayCtl controller.ALiPayV3Controller
}

func NewMemberPayOrderHandler() *MemberPayOrderHandler {
	return &MemberPayOrderHandler{
		aliPayCtl: impl.DefaultALiPayV3Controller,
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
		orderID, payURL, err = h.aliPayCtl.GetPhonePayURLByNumber(ctx, loginID, args.Number)
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
