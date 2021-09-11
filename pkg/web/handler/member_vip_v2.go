package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type MemberVipV2Handler struct {
	ipaDAO        dao.IpaDAO
	ipaVersionDAO dao.IpaVersionDAO

	alipayCtl   controller.ALiPayController
	alertWebCtl controller2.AlterWebController
}

func NewMemberVipV2Handler() *MemberVipV2Handler {
	return &MemberVipV2Handler{
		ipaDAO:        impl.DefaultIpaDAO,
		ipaVersionDAO: impl.DefaultIpaVersionDAO,

		alipayCtl:   impl2.DefaultALiPayController,
		alertWebCtl: impl3.DefaultAlterWebController,
	}
}

func (h *MemberVipV2Handler) GetV2(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"price": constant.DownloadIpaPrice,
	}
	util.RenderJSON(w, res)
}

type getNumber struct {
	Number int64 `form:"number" validate:"required"`
}

func (p *getNumber) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *MemberVipV2Handler) GetPayURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &getNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	loginID := mustGetLoginID(ctx)

	orderID, payURL, err := h.alipayCtl.GetPayURLByNumber(ctx, loginID, args.Number)
	util.PanicIf(err)

	h.alertWebCtl.SendPendingOrderMsg(ctx, orderID)

	res := map[string]interface{}{
		"open_url": payURL,
	}
	util.RenderJSON(w, res)
}
