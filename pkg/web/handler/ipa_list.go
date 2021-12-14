package handler

import (
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
	"net/http"
)

type IpaListHandler struct {
	ipaVersionDAO dao.IpaVersionDAO
}

func NewIpaListHandler() *IpaListHandler {
	return &IpaListHandler{
		ipaVersionDAO: impl.DefaultIpaVersionDAO,
	}
}

func (h *IpaListHandler) GetByIpaType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ipaType, err := enum.IpaTypeString(util.URLParam(r, "ipa_type"))
	util.PanicIf(err)

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	loginID := mustGetLoginID(ctx)

	ivs, err := h.ipaVersionDAO.GetByIpaType(ctx, ipaType)
	util.PanicIf(err)
	ipaIDs := make([]int64, 0)
	for _, iv := range ivs {
		ipaIDs = append(ipaIDs, iv.IpaID)
	}
	ipaIDs = util2.RemoveDuplicates(ipaIDs)

	paging := util.GenerateOffsetPaging(ctx, r, len(ipaIDs), offset, limit)
	paging.IsEnd = true

	ipa := render.NewIpaRender(ipaIDs, loginID, []enum.IpaType{ipaType}, render.IpaDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: paging,
		Data:   ipa,
	})
}
