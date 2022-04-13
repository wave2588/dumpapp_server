package handler

import (
	"dumpapp_server/pkg/web/render"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberRebateRecordHandler struct {
	memberRebateRecordDAO dao.MemberRebateRecordDAO
	memberPayOrderDAO     dao.MemberPayOrderDAO
}

func NewMemberRebateRecordHandler() *MemberRebateRecordHandler {
	return &MemberRebateRecordHandler{
		memberRebateRecordDAO: impl.DefaultMemberRebateRecordDAO,
		memberPayOrderDAO:     impl.DefaultMemberPayOrderDAO,
	}
}

type rebateRecord struct {
	ID        int64          `json:"id,string"`
	Count     int            `json:"count"`
	Member    *render.Member `json:"member"`
	CreatedAt int64          `json:"created_at"`
}

func (h *MemberRebateRecordHandler) GetRebateRecords(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)
	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	filter := []qm.QueryMod{
		models.MemberRebateRecordWhere.ReceiverMemberID.EQ(loginID),
	}
	ids, err := h.memberRebateRecordDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

	totalCount, err := h.memberRebateRecordDAO.Count(ctx, filter)
	util.PanicIf(err)

	rebateRecordMap, err := h.memberRebateRecordDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	orderIDs := make([]int64, 0)
	for _, memberRebateRecord := range rebateRecordMap {
		orderIDs = append(orderIDs, memberRebateRecord.OrderID)
	}

	orderMap, err := h.memberPayOrderDAO.BatchGet(ctx, orderIDs)
	util.PanicIf(err)
	memberIDs := make([]int64, 0)
	for _, order := range orderMap {
		memberIDs = append(memberIDs, order.MemberID)
	}
	memberMap := render.NewMemberRender(memberIDs, loginID, render.MemberDefaultRenderFields...).RenderMap(ctx)

	result := make([]*rebateRecord, 0)
	for _, id := range ids {
		record, ok := rebateRecordMap[id]
		if !ok {
			continue
		}
		order, ok := orderMap[record.OrderID]
		if !ok {
			continue
		}
		member, ok := memberMap[order.MemberID]
		if !ok {
			continue
		}
		result = append(result, &rebateRecord{
			ID:        record.ID,
			Count:     record.Count,
			Member:    member,
			CreatedAt: record.CreatedAt.Unix(),
		})
	}
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   result,
	})
}
