package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminOrderHandler struct {
	orderDAO dao.MemberDownloadOrderDAO
}

func NewAdminOrderHandler() *AdminOrderHandler {
	return &AdminOrderHandler{
		orderDAO: impl.DefaultMemberDownloadOrderDAO,
	}
}

type GetOrderArgs struct {
	StartAt int64 `form:"start_at"`
	EndAt   int64 `form:"end_at"`
}

func (args *GetOrderArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminOrderHandler) GetOrderCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := GetOrderArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	filter := make([]qm.QueryMod, 0)
	if args.StartAt != 0 {
		filter = append(filter, models.MemberDownloadOrderWhere.CreatedAt.GTE(cast.ToTime(args.StartAt)))
	}
	if args.EndAt != 0 {
		filter = append(filter, models.MemberDownloadOrderWhere.CreatedAt.LTE(cast.ToTime(args.EndAt)))
	}

	res, err := h.orderDAO.GetByFilters(ctx, filter, nil)
	util.PanicIf(err)

	paidCount := 0
	downloadCount := 0
	for _, re := range res {
		if re.Status == enum.MemberDownloadOrderStatusPaid {
			paidCount++
			downloadCount += cast.ToInt(re.Number)
		}
	}
	util.RenderJSON(w, map[string]int{
		"order_count":    len(res),
		"paid_count":     paidCount,
		"download_count": downloadCount,
	})
}
