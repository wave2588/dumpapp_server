package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/web/render"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberActionHandler struct {
	memberDownloadIpaRecordDAO dao2.MemberDownloadIpaRecordDAO
}

func NewMemberActionHandler() *MemberActionHandler {
	return &MemberActionHandler{
		memberDownloadIpaRecordDAO: impl4.DefaultMemberDownloadIpaRecordDAO,
	}
}

type GetMemberActionItem struct {
	IpaID      int64          `json:"ipa_id,string"`
	IpaVersion string         `json:"ipa_version"`
	Type       string         `json:"action"` /// download
	Member     *render.Member `json:"member"`
}

func (h *MemberActionHandler) GetMemberActions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	filter := []qm.QueryMod{
		models.MemberDownloadIpaRecordWhere.CreatedAt.GT(time.Now().AddDate(0, 0, -1)),
	}
	ids, err := h.memberDownloadIpaRecordDAO.ListIDs(ctx, 0, 100, filter, nil)
	util.PanicIf(err)

	ipaRecords, err := h.memberDownloadIpaRecordDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	memberIDs := make([]int64, 0)
	for _, ipaRecord := range ipaRecords {
		memberIDs = append(memberIDs, ipaRecord.MemberID)
	}
	options := []render.MemberOption{
		render.MemberIncludes([]string{}),
	}
	memberMap := render.NewMemberRender(memberIDs, 0, options...).RenderMap(ctx)
	for _, member := range memberMap {
		email := member.Email
		member.Email = fmt.Sprintf("***%s", email[3:])
		member.Phone = nil
	}

	result := make([]*GetMemberActionItem, 0)
	for _, recordID := range ids {
		ipaRecord, ok := ipaRecords[recordID]
		if !ok {
			continue
		}
		member, ok := memberMap[ipaRecord.MemberID]
		if !ok {
			continue
		}
		result = append(result, &GetMemberActionItem{
			IpaID:      ipaRecord.IpaID.Int64,
			IpaVersion: ipaRecord.Version.String,
			Type:       "download",
			Member:     member,
		})
	}

	util.RenderJSON(w, util.ListOutput{
		Paging: nil,
		Data:   result,
	})
}
