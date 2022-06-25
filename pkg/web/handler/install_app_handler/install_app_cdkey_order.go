package install_app_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/install_app"
	"dumpapp_server/pkg/controller/install_app/impl"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type InstallAppCDKEYOrderHandler struct {
	aliPayCtl install_app.ALiPayInstallAppController
}

func NewInstallAppCDKEYOrderHandler() *InstallAppCDKEYOrderHandler {
	return &InstallAppCDKEYOrderHandler{
		aliPayCtl: impl.DefaultALiPayInstallAppController,
	}
}

type getMemberPayOrderArgs struct {
	Number     int64  `form:"number" validate:"required"`
	ContactWay string `form:"contact_way" validate:"required"`
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

func (h *InstallAppCDKEYOrderHandler) GetOrderURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getMemberPayOrderArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	payURL, err := h.aliPayCtl.GetPayURLByInstallApp(ctx, args.Number, args.ContactWay)
	util.PanicIf(err)

	res := map[string]interface{}{
		"open_url": payURL,
	}
	util.RenderJSON(w, res)
}
