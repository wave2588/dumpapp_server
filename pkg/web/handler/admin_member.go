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

	loginID := mustGetLoginID(ctx)

	filter := make([]qm.QueryMod, 0)
	if args.StartAt != 0 {
		filter = append(filter, models.AccountWhere.CreatedAt.GT(cast.ToTime(args.StartAt)))
	}
	if args.EndAt != 0 {
		filter = append(filter, models.AccountWhere.CreatedAt.LT(cast.ToTime(args.EndAt)))
	}
	ids, err := h.accountDAO.ListIDs(ctx, 0, 10000, filter, nil)
	util.PanicIf(err)

	members := render.NewMemberRender(ids, loginID, render.MemberAdminRenderFields...).RenderSlice(ctx)
	sort.Slice(members, func(i, j int) bool {
		return members[i].Admin.PaidCount > members[j].Admin.PaidCount
	})

	util.RenderJSON(w, util.ListOutput{
		Paging: nil,
		Data:   members,
	})
}
