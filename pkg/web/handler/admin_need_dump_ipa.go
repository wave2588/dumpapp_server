package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/web/render"
)

type AdminNeedDumpIpaHandler struct {
	adminNeedDumpIpaDAO dao.AdminNeedDumpIpaDAO
}

func NewAdminNeedDumpIpaHandler() *AdminNeedDumpIpaHandler {
	return &AdminNeedDumpIpaHandler{
		adminNeedDumpIpaDAO: impl.DefaultAdminNeedDumpIpaDAO,
	}
}

func (h *AdminNeedDumpIpaHandler) GetNeedDumpList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	ids, err := h.adminNeedDumpIpaDAO.ListIDs(ctx, offset, limit, nil, []string{"created_at desc"})
	util.PanicIf(err)
	totalCount, err := h.adminNeedDumpIpaDAO.Count(ctx, nil)
	util.PanicIf(err)
	needDumps, err := h.adminNeedDumpIpaDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	memberIDs := make([]int64, 0)
	for _, ndi := range needDumps {
		memberIDs = append(memberIDs, ndi.MemberID)
	}
	memberMap := render.NewMemberRender(memberIDs, loginID, render.MemberDefaultRenderFields...).RenderMap(ctx)

	result := make([]*NeedDumpResult, 0)
	for _, ndi := range needDumps {
		result = append(result, &NeedDumpResult{
			Member:     memberMap[ndi.MemberID],
			IpaID:      ndi.IpaID,
			IpaVersion: ndi.IpaVersion,
			IpaName:    ndi.IpaName,
			CreatedAt:  ndi.CreatedAt.Unix(),
			UpdatedAt:  ndi.UpdatedAt.Unix(),
		})
	}
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   result,
	})
}

type NeedDumpResult struct {
	Member     *render.Member `json:"member"`
	IpaID      int64          `json:"ipa_id"`
	IpaVersion string         `json:"ipa_version"`
	IpaName    string         `json:"ipa_name"`
	CreatedAt  int64          `json:"created_at"`
	UpdatedAt  int64          `json:"updated_at"`
}
