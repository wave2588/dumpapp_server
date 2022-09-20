package install_app_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	impl2 "dumpapp_server/pkg/controller/install_app/impl"
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
	Number   int `json:"number" validate:"required"`
	CerLevel int `json:"cer_level" validate:"required"`
	Price    int `json:"price" validate:"required"`
}

func (p *postCDKeyArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.Number <= 0 {
		return errors.UnproccessableError("number > 0")
	}
	if p.CerLevel > 3 || p.CerLevel < 1 {
		return errors.UnproccessableError(fmt.Sprintf("检查 cer_level 是否符合要求: %d", p.CerLevel))
	}
	return nil
}

func (h *AdminCDKeyHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &postCDKeyArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	outIDs, err := impl2.DefaultALiPayInstallAppController.GetOutIDs(ctx, args.Number, args.CerLevel)
	util.PanicIf(err)

	orderID := util2.MustGenerateID(ctx)
	util.PanicIf(h.cdkeyOrderDAO.Insert(ctx, &models.InstallAppCdkeyOrder{
		ID:     orderID,
		Status: enum.MemberPayOrderStatusPaid,
		Number: cast.ToInt64(args.Number),
		Amount: cast.ToFloat64(args.Price),
		BizExt: datatype.InstallAppCdkeyOrderBizExt{
			CerLevel: args.CerLevel,
		},
	}))

	cdkeyIDs := make([]int64, 0)
	for _, outID := range outIDs {
		cdkeyID := util2.MustGenerateID(ctx)
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
