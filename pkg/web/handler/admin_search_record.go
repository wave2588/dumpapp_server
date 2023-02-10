package handler

import (
	"fmt"
	"net/http"

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
	StartAt int64 `form:"start_at"`
	EndAt   int64 `form:"end_at"`
}

func (args *GetMemberSearchRecordArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

type record struct {
	IpaID int64  `json:"ipa_id,string"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

func (h *AdminSearchRecordHandler) GetMemberSearchRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := GetMemberSearchRecordArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	filter := make([]qm.QueryMod, 0)
	if args.StartAt != 0 {
		filter = append(filter, models.SearchRecordV2Where.CreatedAt.GTE(cast.ToTime(args.StartAt)))
	}
	if args.EndAt != 0 {
		filter = append(filter, models.SearchRecordV2Where.CreatedAt.LTE(cast.ToTime(args.EndAt)))
	}

	data, err := h.searchRecordV2DAO.GetOrderBySearchCount(ctx, offset, limit, filter)
	util.PanicIf(err)

	totalCount, err := h.searchRecordV2DAO.CountOrderBySearchCount(ctx, filter)
	util.PanicIf(err)

	result := make([]*record, 0)
	for _, datum := range data {
		result = append(result, &record{
			IpaID: datum.IpaID,
			//Name:  datum.Name,
			Count: datum.Count,
		})
	}

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   result,
	})
}
