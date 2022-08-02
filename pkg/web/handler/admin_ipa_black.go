package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
)

type AdminIpaBlackHandler struct {
	ipaBlackDAO dao.IpaBlackDAO
}

func NewAdminIpaBlackHandler() *AdminIpaBlackHandler {
	return &AdminIpaBlackHandler{
		ipaBlackDAO: impl.DefaultIpaBlackDAO,
	}
}

type postIpaBlackArgs struct {
	IpaID  int64  `json:"ipa_id,string" validate:"required"`
	Reason string `json:"reason" validate:"required"`
}

func (p *postIpaBlackArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminIpaBlackHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &postIpaBlackArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	util.PanicIf(h.ipaBlackDAO.Insert(ctx, &models.IpaBlack{
		IpaID:  args.IpaID,
		BizExt: datatype.IpaBlackBizExt{Reason: args.Reason},
	}))

	data := render.NewIpaRender([]int64{args.IpaID}, loginID, []enum.IpaType{enum.IpaTypeNormal, enum.IpaTypeCrack}, render.IpaAdminRenderFields...).RenderMap(ctx)

	util.RenderJSON(w, data[args.IpaID])
}

func (h *AdminIpaBlackHandler) GetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	ipaIDs, err := h.ipaBlackDAO.AdminListIpaIDs(ctx, offset, limit)
	util.PanicIf(err)

	paging := util.GenerateOffsetPaging(ctx, r, 0, offset, limit)
	paging.IsEnd = len(ipaIDs) != limit

	util.RenderJSON(w, util.ListOutput{
		Paging: paging,
		Data:   render.NewIpaRender(ipaIDs, loginID, []enum.IpaType{enum.IpaTypeNormal, enum.IpaTypeCrack}, render.IpaAdminRenderFields...).RenderSlice(ctx),
	})
}

func (h *AdminIpaBlackHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := cast.ToInt64(util.URLParam(r, "ipa_black_id"))

	dataMap, err := h.ipaBlackDAO.BatchGet(ctx, []int64{id})
	util.PanicIf(err)

	_, ok := dataMap[id]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}

	util.PanicIf(h.ipaBlackDAO.Delete(ctx, id))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
