package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminOrderHandler struct {
	orderDAO dao.MemberPayOrderDAO
}

func NewAdminOrderHandler() *AdminOrderHandler {
	return &AdminOrderHandler{
		orderDAO: impl.DefaultMemberPayOrderDAO,
	}
}

type getListArgs struct {
	Status  enum.MemberPayOrderStatus `form:"status"`
	StartAt int64                     `form:"start_at"`
	EndAt   int64                     `form:"end_at"`
	Account string                    `form:"account"`
}

func (args *getListArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminOrderHandler) GetList(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
		loginID = mustGetLoginID(ctx)
	)

	args := getListArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	filters := make([]qm.QueryMod, 0)

	if args.Status.IsAMemberPayOrderStatus() {
		filters = append(filters, models.MemberPayOrderWhere.Status.EQ(args.Status))
	}
	if args.Account != "" {
		account := GetAccountByAccount(ctx, args.Account)
		filters = append(filters, models.MemberPayOrderWhere.MemberID.EQ(account.ID))
	}
	if args.StartAt != 0 {
		filters = append(filters, models.MemberPayOrderWhere.UpdatedAt.GTE(time.Unix(args.StartAt, 0)))
	}
	if args.EndAt != 0 {
		filters = append(filters, models.MemberPayOrderWhere.UpdatedAt.LTE(time.Unix(args.EndAt, 0)))
	}

	ids, err := h.orderDAO.ListIDs(ctx, offset, limit, filters, nil)
	util.PanicIf(err)

	count, err := h.orderDAO.Count(ctx, filters)
	util.PanicIf(err)

	data := render.NewMemberPayOrderRender(ids, loginID, render.MemberPayOrderAdminRenderFidles...).RenderSlice(ctx)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   data,
	})
}
