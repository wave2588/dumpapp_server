package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/render"
	"github.com/spf13/cast"
)

type IpaHandler struct {
	ipaDAO        dao.IpaDAO
	ipaVersionDAO dao.IpaVersionDAO
}

func NewIpaHandler() *IpaHandler {
	return &IpaHandler{
		ipaDAO:        impl.DefaultIpaDAO,
		ipaVersionDAO: impl.DefaultIpaVersionDAO,
	}
}

func (h *IpaHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	count, err := h.ipaDAO.Count(ctx, nil)
	util.PanicIf(err)

	ids, err := h.ipaDAO.ListIDs(ctx, offset, limit, nil, nil)
	util.PanicIf(err)

	data := render.NewIpaRender(ids, loginID, render.IpaDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   data,
	})
}

func (h *IpaHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	data := render.NewIpaRender([]int64{ipaID}, loginID, render.IpaDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[ipaID])
}
