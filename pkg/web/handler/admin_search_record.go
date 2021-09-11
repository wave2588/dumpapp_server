package handler

import (
	"fmt"
	"net/http"
	"sort"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminSearchRecordHandler struct {
	searchRecordV2DAO dao.SearchRecordV2DAO
}

func NewAdminSearchRecordHandler() *AdminSearchRecordHandler {
	return &AdminSearchRecordHandler{
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,
	}
}

type GetMemberSearchRecordArgs struct {
	StartAt int64 `form:"start_at" validate:"required"`
	EndAt   int64 `form:"end_at" validate:"required"`
}

func (args *GetMemberSearchRecordArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

type record struct {
	IpaID    int64  `json:"ipa_id,string"`
	Name     string `json:"name"`
	Count    int64  `json:"count"`
	LatestAt int64  `json:"latest_at"`
}

func (h *AdminSearchRecordHandler) GetMemberSearchRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := GetMemberSearchRecordArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	filter := []qm.QueryMod{
		models.SearchRecordV2Where.CreatedAt.GT(cast.ToTime(args.StartAt)),
		models.SearchRecordV2Where.CreatedAt.LT(cast.ToTime(args.EndAt)),
	}
	ids, err := h.searchRecordV2DAO.ListIDs(ctx, 0, 10000, filter, nil)
	util.PanicIf(err)

	recordMap, err := h.searchRecordV2DAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	recordIpaMap := make(map[int64][]*models.SearchRecordV2, 0)
	for _, record := range recordMap {
		recordIpaMap[record.IpaID] = append(recordIpaMap[record.IpaID], record)
	}

	result := make([]*record, 0)
	for ipaID, records := range recordIpaMap {
		if len(records) < 1 {
			continue
		}
		result = append(result, &record{
			IpaID:    ipaID,
			Name:     records[0].Name,
			Count:    cast.ToInt64(len(records)),
			LatestAt: records[len(records)-1].CreatedAt.Unix(),
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	util.RenderJSON(w, util.ListOutput{
		Paging: nil,
		Data:   result,
	})
}
