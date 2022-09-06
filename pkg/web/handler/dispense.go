package handler

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DispenseHandler struct {
	memberPayCountCtl      controller.MemberPayCountController
	dispenseCountCtl       controller.DispenseCountController
	memberSignIpaDAO       dao.MemberSignIpaDAO
	dispenseCountRecordDAO dao.DispenseCountRecordDAO
}

func NewDispenseHandler() *DispenseHandler {
	return &DispenseHandler{
		memberPayCountCtl:      impl.DefaultMemberPayCountController,
		dispenseCountCtl:       impl.DefaultDispenseCountController,
		memberSignIpaDAO:       impl2.DefaultMemberSignIpaDAO,
		dispenseCountRecordDAO: impl2.DefaultDispenseCountRecordDAO,
	}
}

type postDispenseArgs struct {
	Count int64 `json:"count" validate:"required"`
}

func (p *postDispenseArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *DispenseHandler) Post(w http.ResponseWriter, r *http.Request) {
	args := &postDispenseArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
	)

	/// 扣费流程, 里边有检查 D 币是否足够
	util.PanicIf(h.memberPayCountCtl.DeductPayCount(ctx, loginID, args.Count, enum.MemberPayCountStatusUsed, enum.MemberPayCountUseDispense, datatype.MemberPayCountRecordBizExt{}))

	/// fixme: 活动期间分发卷双倍
	util.PanicIf(h.dispenseCountCtl.AddCount(ctx, loginID, args.Count*constant.DispenseRatioByPayCount, enum.DispenseCountRecordTypePay))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type expenseArgs struct {
	ID          string `json:"id" validate:"required"`
	ExpenseType string `json:"expense_type" validate:"required"`
}

func (p *expenseArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.ExpenseType != "sign_ipa" {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *DispenseHandler) Expense(w http.ResponseWriter, r *http.Request) {
	args := &expenseArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	ctx := r.Context()

	signIpaMap, err := h.memberSignIpaDAO.BatchGetByExpenseID(ctx, []string{args.ID})
	util.PanicIf(err)

	signIpa, ok := signIpaMap[args.ID]
	if !ok {
		util.PanicIf(errors.NewDefaultAPIError(404, 404, "NotFound", fmt.Sprintf("记录未找到, id:%s type:%s", cast.ToString(args.ID), args.ExpenseType)))
	}

	///  根据大小同步进行扣费 signIpa.BizExt.IpaSize / 1024 / 1024
	dCount := h.dispenseCountCtl.CalculateMemberSignIpaDispenseCount(ctx, signIpa.BizExt.IpaSize)
	util.PanicIf(h.dispenseCountCtl.DeductCount(ctx, signIpa.MemberID, dCount, enum.DispenseCountRecordTypeInstallSignIpa))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

func (h *DispenseHandler) Records(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	filter := []qm.QueryMod{
		models.DispenseCountRecordWhere.MemberID.EQ(loginID),
	}
	ids, err := h.dispenseCountRecordDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

	count, err := h.dispenseCountRecordDAO.Count(ctx, filter)
	util.PanicIf(err)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   render.NewDispenseCountRecordRender(ids, loginID, render.DispenseCountRecordDefaultRenderFields...).RenderSlice(ctx),
	})
}
