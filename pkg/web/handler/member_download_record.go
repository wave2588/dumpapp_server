package handler

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/web/render"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberDownloadRecordHandler struct {
	memberDownloadIpaRecordDAO dao2.MemberDownloadIpaRecordDAO
}

func NewMemberDownloadRecordHandler() *MemberDownloadRecordHandler {
	return &MemberDownloadRecordHandler{
		memberDownloadIpaRecordDAO: impl4.DefaultMemberDownloadIpaRecordDAO,
	}
}

func (h *MemberDownloadRecordHandler) GetSelfDownloadRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)
	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	account := GetAccountByLoginID(ctx, loginID)

	filter := []qm.QueryMod{
		models.MemberDownloadIpaRecordWhere.MemberID.EQ(loginID),
		models.MemberDownloadIpaRecordWhere.Status.EQ("used"),
	}
	ids, err := h.memberDownloadIpaRecordDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

	count, err := h.memberDownloadIpaRecordDAO.Count(ctx, filter)
	util.PanicIf(err)

	data := render.NewMemberDownloadRecordRender(ids, account.ID, render.MemberDownloadRecordDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   data,
	})
}
