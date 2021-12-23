package handler

import (
	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
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
	}
	ids, err := h.memberDownloadNumberDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

}
