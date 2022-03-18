package handler

import (
	"context"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminMemberHandler struct {
	accountDAO dao.AccountDAO
}

func NewAdminMemberHandler() *AdminMemberHandler {
	return &AdminMemberHandler{
		accountDAO: impl.DefaultAccountDAO,
	}
}

type ListMemberArgs struct {
	StartAt int64 `form:"start_at"`
	EndAt   int64 `form:"end_at"`
}

func (args *ListMemberArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminMemberHandler) ListMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := ListMemberArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	loginID := mustGetLoginID(ctx)

	filter := make([]qm.QueryMod, 0)
	if args.StartAt != 0 {
		filter = append(filter, models.AccountWhere.CreatedAt.GTE(cast.ToTime(args.StartAt)))
	}
	if args.EndAt != 0 {
		filter = append(filter, models.AccountWhere.CreatedAt.LTE(cast.ToTime(args.EndAt)))
	}

	h.getMembers(ctx, w, r, loginID, offset, limit, filter)
}

func (h *AdminMemberHandler) getMembers(ctx context.Context, w http.ResponseWriter, r *http.Request, loginID int64, offset, limit int, filter []qm.QueryMod) {
	ids, err := h.accountDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)
	totalCount, err := h.accountDAO.Count(ctx, filter)
	util.PanicIf(err)
	members := render.NewMemberRender(ids, loginID, render.MemberAdminRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   members,
	})
}
