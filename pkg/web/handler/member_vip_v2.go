package handler

import (
	"dumpapp_server/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
)

type MemberVipV2Handler struct {
	ipaDAO        dao.IpaDAO
	ipaVersionDAO dao.IpaVersionDAO

	alipayCtl controller.ALiPayController
}

func NewMemberVipV2Handler() *MemberVipV2Handler {
	return &MemberVipV2Handler{
		ipaDAO:        impl.DefaultIpaDAO,
		ipaVersionDAO: impl.DefaultIpaVersionDAO,

		alipayCtl: impl2.DefaultALiPayController,
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

	payURL, err := h.alipayCtl.GetPayURLByNumber(ctx, loginID, args.Number)
	util.PanicIf(err)

	res := map[string]interface{}{
		"open_url": payURL,
	}
	util.RenderJSON(w, res)
}