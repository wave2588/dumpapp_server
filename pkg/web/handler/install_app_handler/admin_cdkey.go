package install_app_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render/install_app_render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
)

type AdminCDKeyHandler struct {
	cdkeyDAO      dao.InstallAppCdkeyDAO
	cdkeyOrderDAO dao.InstallAppCdkeyOrderDAO
}

func NewAdminCDKeyHandler() *AdminCDKeyHandler {
	return &AdminCDKeyHandler{
		cdkeyDAO:      impl.DefaultInstallAppCdkeyDAO,
		cdkeyOrderDAO: impl.DefaultInstallAppCdkeyOrderDAO,
	}
}

type postCDKeyArgs struct {
	Number int `json:"number" validate:"required"`
}

func (p *postCDKeyArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.Number <= 0 {
		return errors.UnproccessableError("number > 0")
	}
	return nil
}

func (h *AdminCDKeyHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &postCDKeyArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	cdkeyIDs := make([]int64, 0)
	for i := 0; i < args.Number; i++ {
		orderID := util2.MustGenerateID(ctx)
		bizExt := constant.InstallAppCDKEYOrderBizExt{
			IsTest: true,
		}
		util.PanicIf(h.cdkeyOrderDAO.Insert(ctx, &models.InstallAppCdkeyOrder{
			ID:     orderID,
			Status: enum.MemberPayOrderStatusPaid,
			Number: 1,
			Amount: 0,
			BizExt: bizExt.String(),
		}))

		cdkeyID := util2.MustGenerateID(ctx)
		outID := util2.MustGenerateAppCDKEY()
		util.PanicIf(h.cdkeyDAO.Insert(ctx, &models.InstallAppCdkey{
			ID:      cdkeyID,
			OutID:   outID,
			Status:  enum.InstallAppCDKeyStatusNormal,
			OrderID: orderID,
		}))

		cdkeyIDs = append(cdkeyIDs, cdkeyID)
	}

	util.RenderJSON(w, util.ListOutput{
		Paging: nil,
		Data:   install_app_render.NewCDKEYRender(cdkeyIDs, 0, install_app_render.CDKeyDefaultRenderFields...).RenderSlice(ctx),
	})
}

func (h *AdminCDKeyHandler) GetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	ids, err := h.cdkeyDAO.ListIDs(ctx, offset, limit, nil, nil)
	util.PanicIf(err)
	count, err := h.cdkeyDAO.Count(ctx, nil)
	util.PanicIf(err)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   install_app_render.NewCDKEYRender(ids, 0, install_app_render.CDKeyDefaultRenderFields...).RenderSlice(ctx),
	})
}

func (h *AdminCDKeyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cdkeyID := cast.ToInt64(util.URLParam(r, "cdkey_id"))

	cdkeyMap, err := h.cdkeyDAO.BatchGet(ctx, []int64{cdkeyID})
	util.PanicIf(err)

	cdkey, ok := cdkeyMap[cdkeyID]
	if !ok {
		util.PanicIf(errors.ErrInstallAppCDKeyNotFound)
	}
	cdkey.Status = enum.InstallAppCDKeyStatusAdminDelete
	util.PanicIf(h.cdkeyDAO.Update(ctx, cdkey))

	data := install_app_render.NewCDKEYRender([]int64{cdkeyID}, 0, install_app_render.CDKeyDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[cdkeyID])
}
