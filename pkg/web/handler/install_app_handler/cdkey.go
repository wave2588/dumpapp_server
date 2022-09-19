package install_app_handler

import (
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render/install_app_render"
)

type CDKEYHandler struct {
	installAppCDKEYDAO    dao2.InstallAppCdkeyDAO
	installAppCDKEYCerDAO dao2.InstallAppCertificateDAO
}

func NewCDKEYHandler() *CDKEYHandler {
	return &CDKEYHandler{
		installAppCDKEYDAO:    impl2.DefaultInstallAppCdkeyDAO,
		installAppCDKEYCerDAO: impl2.DefaultInstallAppCertificateDAO,
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

type cdkeyPrice struct {
	ID          int64  `json:"id,string"`
	Price       int64  `json:"price"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *CDKEYHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	data := []*cdkeyPrice{
		{
			ID:          1,
			Price:       constant.CertificatePriceL1,
			Title:       "普通版",
			Description: "理论 1 年，无质保。",
		},
		{
			ID:          2,
			Price:       constant.CertificatePriceL2,
			Title:       "稳定版",
			Description: "理论 1 年，售后半年，掉了无限补。",
		},
		{
			ID:          3,
			Price:       constant.CertificatePriceL3,
			Title:       "豪华版",
			Description: "理论 1 年，售后 1 年，掉了无限补。",
		},
	}
	util.RenderJSON(w, util.ListOutput{
		Data: data,
	})
}
