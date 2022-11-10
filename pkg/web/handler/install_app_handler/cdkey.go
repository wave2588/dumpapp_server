package install_app_handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	dao2 "dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render/install_app_render"
)

type CDKEYHandler struct {
	installAppCDKEYDAO     dao2.InstallAppCdkeyDAO
	installAppCDKEYCerDAO  dao2.InstallAppCertificateDAO
	installAppCKEYOrderDAO dao2.InstallAppCdkeyOrderDAO
	certificatePriceCtl    controller.CertificatePriceController
}

func NewCDKEYHandler() *CDKEYHandler {
	return &CDKEYHandler{
		installAppCDKEYDAO:     impl2.DefaultInstallAppCdkeyDAO,
		installAppCDKEYCerDAO:  impl2.DefaultInstallAppCertificateDAO,
		installAppCKEYOrderDAO: impl2.DefaultInstallAppCdkeyOrderDAO,
		certificatePriceCtl:    impl.DefaultCertificatePriceController,
	}
}

func (h *CDKEYHandler) GetCDKEYInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outID := util.URLParam(r, "out_id")

	res, err := h.installAppCDKEYDAO.BatchGetByOutID(ctx, []string{outID})
	util.PanicIf(err)

	cdkey, ok := res[outID]
	if !ok {
		util.PanicIf(errors.ErrInstallAppCDKeyNotFound)
	}
	data := install_app_render.NewCDKEYRender([]int64{cdkey.ID}, 0, install_app_render.CDKeyDefaultRenderFields...).RenderMap(ctx)

	util.RenderJSON(w, data[cdkey.ID])
}

type getOrderByContactWatResp struct {
	UDID   string                      `json:"udid"`
	CDKeys []*install_app_render.CDKEY `json:"cd_keys"`
}

func (h *CDKEYHandler) GetCDKEYInfoByUDID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	udid := util.URLParam(r, "udid")
	cerMap, err := h.installAppCDKEYCerDAO.BatchGetByUDIDs(ctx, []string{udid})
	util.PanicIf(err)

	cers, ok := cerMap[udid]
	if !ok {
		util.PanicIf(errors.UnproccessableError("未找到兑换码"))
	}

	cerIDs := make([]int64, 0)
	for _, cer := range cers {
		cerIDs = append(cerIDs, cer.ID)
	}

	cdkeyMap, err := h.installAppCDKEYDAO.BatchGetByCertificateIDs(ctx, cerIDs)
	util.PanicIf(err)

	cdkeyIDs := make([]int64, 0)
	for _, cdkeys := range cdkeyMap {
		for _, cdkey := range cdkeys {
			cdkeyIDs = append(cdkeyIDs, cdkey.ID)
		}
	}

	util.RenderJSON(w, &getOrderByContactWatResp{
		UDID:   udid,
		CDKeys: install_app_render.NewCDKEYRender(cdkeyIDs, 0, install_app_render.CDKeyDefaultRenderFields...).RenderSlice(ctx),
	})
}

type getOrderByContactResp struct {
	Contact string                      `json:"contact"`
	CDKeys  []*install_app_render.CDKEY `json:"cd_keys"`
}

func (h *CDKEYHandler) GetCDKEYInfoByContactWay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contact := util.URLParam(r, "contact")

	orderMap, err := h.installAppCKEYOrderDAO.BatchGetByContact(ctx, []string{contact})
	util.PanicIf(err)

	orderIDs := make([]int64, 0)
	for _, orders := range orderMap {
		for _, order := range orders {
			orderIDs = append(orderIDs, order.ID)
		}
	}

	cdkeys, err := h.installAppCDKEYDAO.BatchGetByOrderIDs(ctx, orderIDs)
	util.PanicIf(err)

	cdkeyIDs := make([]int64, 0)
	for _, appCdkeys := range cdkeys {
		for _, cdkey := range appCdkeys {
			cdkeyIDs = append(cdkeyIDs, cdkey.ID)
		}
	}

	util.RenderJSON(w, &getOrderByContactResp{
		Contact: contact,
		CDKeys:  install_app_render.NewCDKEYRender(cdkeyIDs, 0, install_app_render.CDKeyDefaultRenderFields...).RenderSlice(ctx),
	})
}

func (h *CDKEYHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	prices, err := h.certificatePriceCtl.GetPrices(ctx, 0)
	util.PanicIf(err)
	util.RenderJSON(w, util.ListOutput{
		Data: prices,
	})
}
