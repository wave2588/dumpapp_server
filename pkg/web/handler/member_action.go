package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberActionHandler struct {
	memberDownloadIpaRecordDAO dao2.MemberDownloadIpaRecordDAO
	ipaDAO                     dao2.IpaDAO
}

func NewMemberActionHandler() *MemberActionHandler {
	return &MemberActionHandler{
		memberDownloadIpaRecordDAO: impl4.DefaultMemberDownloadIpaRecordDAO,
		ipaDAO:                     impl4.DefaultIpaDAO,
	}
}

type GetMemberActionItem struct {
	IpaID      int64          `json:"ipa_id,string"`
	IpaName    string         `json:"ipa_name"`
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

	ipaIDs := make([]int64, 0)
	for _, ipaRecord := range ipaRecords {
		if ipaRecord.IpaID.IsZero() {
			continue
		}
		ipaIDs = append(ipaIDs, ipaRecord.IpaID.Int64)
	}
	ipaIDs = util2.RemoveDuplicates(ipaIDs)
	ipaMap, err := h.ipaDAO.BatchGet(ctx, ipaIDs)
	util.PanicIf(err)

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
		ipa, ok := ipaMap[ipaRecord.IpaID.Int64]
		if !ok {
			continue
		}
		result = append(result, &GetMemberActionItem{
			IpaID:      ipaRecord.IpaID.Int64,
			IpaName:    ipa.Name,
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
