package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/controller"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type AdminCertificateHandler struct {
	accountDAO        dao.AccountDAO
	certificateWebCtl controller.CertificateWebController
}

func NewAdminCertificateHandler() *AdminCertificateHandler {
	return &AdminCertificateHandler{
		accountDAO:        impl.DefaultAccountDAO,
		certificateWebCtl: impl2.DefaultCertificateWebController,
	}
}

type replenishCertificateArgs struct {
	Email string `json:"email" validate:"required"`
	UDID  string `json:"udid" validate:"required"`
}

func (p *replenishCertificateArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if !util2.CheckUDIDValid(p.UDID) {
		return errors.UnproccessableError(fmt.Sprintf("无效的 UDID: %s", p.UDID))
	}
	return nil
}

func (h *AdminCertificateHandler) Replenish(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &replenishCertificateArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util.PanicIf(err)

	_, err = h.certificateWebCtl.PayCertificate(ctx, account.ID, args.UDID, "售后证书", constant.CertificateIDL1, true, "")
	util.PanicIf(err)

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
