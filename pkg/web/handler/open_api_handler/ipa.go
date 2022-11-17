package open_api_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type OpenIpaHandler struct {
	memberDeviceDAO     dao.MemberDeviceDAO
	certificateDAO      dao.CertificateV2DAO
	certificateWebCtl   controller2.CertificateWebController
	certificatePriceCtl controller.CertificatePriceController
}

func NewOpenIpaHandler() *OpenIpaHandler {
	return &OpenIpaHandler{
		memberDeviceDAO:     impl.DefaultMemberDeviceDAO,
		certificateDAO:      impl.DefaultCertificateV2DAO,
		certificateWebCtl:   impl3.DefaultCertificateWebController,
		certificatePriceCtl: impl2.DefaultCertificatePriceController,
	}
}

type getIpaArgs struct {
	IpaID      string `form:"ipa_id" validate:"required"`
	IpaVersion string `form:"ipa_version" validate:"required"`
}

func (p *getIpaArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *OpenIpaHandler) Get(w http.ResponseWriter, r *http.Request) {
}
