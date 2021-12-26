package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/web/render"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberDownloadRecordHandler struct {
	memberDownloadNumberDAO dao2.MemberDownloadNumberDAO
}

func NewMemberDownloadRecordHandler() *MemberDownloadRecordHandler {
	return &MemberDownloadRecordHandler{
		memberDownloadNumberDAO: impl4.DefaultMemberDownloadNumberDAO,
	}
}

func (h *MemberDownloadRecordHandler) GetSelfDownloadRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)
	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	account := GetAccountByLoginID(ctx, loginID)

	filter := []qm.QueryMod{
		models.MemberDownloadNumberWhere.MemberID.EQ(loginID),
		models.MemberDownloadNumberWhere.Status.EQ(enum.MemberDownloadNumberStatusUsed),
		models.MemberDownloadNumberWhere.Version.IsNotNull(),
	}
	ids, err := h.memberDownloadNumberDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)
	count, err := h.memberDownloadNumberDAO.Count(ctx, filter)
	util.PanicIf(err)

	data := render.NewMemberDownloadRecordRender(ids, account.ID, render.MemberDownloadRecordDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   data,
	})
}
