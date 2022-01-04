package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/controller"
	"dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type IpaSignHandler struct {
	ipaSignWebCtl controller.IpaSignWebController
}

func NewIpaSignHandler() *IpaSignHandler {
	return &IpaSignHandler{
		ipaSignWebCtl: impl.DefaultIpaSignWebController,
	}
}

type postSignArgs struct {
	CertificateID int64 `json:"certificate_id,string" validate:"required"`
	IpaVersionID  int64 `json:"ipa_version_id,string" validate:"required"`
}

func (p *postSignArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *IpaSignHandler) PostSign(w http.ResponseWriter, r *http.Request) {
	args := &postSignArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	h.ipaSignWebCtl.Sign(ctx, loginID, args.CertificateID, args.IpaVersionID)
}
