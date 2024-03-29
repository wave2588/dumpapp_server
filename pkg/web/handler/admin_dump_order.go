package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminDumpOrderHandler struct {
	adminDumpOrderCtl controller.AdminDumpOrderController
	adminDumpOrderDAO dao.AdminDumpOrderDAO
}

func NewAdminDumpOrderHandler() *AdminDumpOrderHandler {
	return &AdminDumpOrderHandler{
		adminDumpOrderCtl: impl2.DefaultAdminDumpOrderController,
		adminDumpOrderDAO: impl.DefaultAdminDumpOrderDAO,
	}
}

type DumpOrderResult struct {
	ID                  int64                     `json:"id,string"`
	DemanderMember      *render.Member            `json:"demander_member"`
	OtherDemanderMember []*render.Member          `json:"other_demander_member"`
	OperatorMember      *render.Member            `json:"operator_member"`
	Status              enum.AdminDumpOrderStatus `json:"status"`
	IpaID               int64                     `json:"ipa_id,string"`
	IpaVersion          string                    `json:"ipa_version"`
	IpaName             string                    `json:"ipa_name"`
	IpaBundleID         string                    `json:"ipa_bundle_id"`
	IpaAppStoreLink     string                    `json:"ipa_app_store_link"`
	IsOld               bool                      `json:"is_old"`
	CreatedAt           int64                     `json:"created_at"`
	UpdatedAt           int64                     `json:"updated_at"`
}

func (h *AdminDumpOrderHandler) GetDumpOrderList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	filter := []qm.QueryMod{
		qm.Where("(status=?) or (status=?)", enum.AdminDumpOrderStatusUnprocessed, enum.AdminDumpOrderStatusProgressing),
	}
	ids, err := h.adminDumpOrderDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)
	totalCount, err := h.adminDumpOrderDAO.Count(ctx, filter)
	util.PanicIf(err)
	dumpOrderMap, err := h.adminDumpOrderDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	memberIDs := make([]int64, 0)
	for _, do := range dumpOrderMap {
		var bizExt constant.AdminDumpOrderBizExt
		util.PanicIf(json.Unmarshal([]byte(do.IpaBizExt), &bizExt))
		memberIDs = append(memberIDs, do.DemanderID)
		if do.OperatorID != 0 {
			memberIDs = append(memberIDs, do.OperatorID)
		}
		memberIDs = append(memberIDs, bizExt.DemanderIDs...)
	}
	memberIDs = util2.RemoveDuplicates(memberIDs)
	memberMap := render.NewMemberRender(memberIDs, loginID, render.MemberAdminRenderFields...).RenderMap(ctx)

	result := make([]*DumpOrderResult, 0)
	for _, orderID := range ids {
		do, ok := dumpOrderMap[orderID]
		if !ok {
			continue
		}
		var bizExt constant.AdminDumpOrderBizExt
		util.PanicIf(json.Unmarshal([]byte(do.IpaBizExt), &bizExt))
		otherDemanderMembers := make([]*render.Member, 0)
		for _, otherDemanderMemberID := range bizExt.DemanderIDs {
			otherDemanderMembers = append(otherDemanderMembers, memberMap[otherDemanderMemberID])
		}
		res := &DumpOrderResult{
			ID:                  orderID,
			DemanderMember:      memberMap[do.DemanderID],
			OtherDemanderMember: otherDemanderMembers,
			OperatorMember:      memberMap[do.OperatorID],
			Status:              do.Status,
			IpaID:               do.IpaID,
			IpaVersion:          do.IpaVersion,
			IpaName:             bizExt.IpaName,
			IpaBundleID:         bizExt.IpaBundleID,
			IpaAppStoreLink:     bizExt.IpaAppStoreLink,
			IsOld:               bizExt.IsOld,
			CreatedAt:           do.CreatedAt.Unix(),
			UpdatedAt:           do.UpdatedAt.Unix(),
		}
		result = append(result, res)
	}
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   result,
	})
}

type putAdminDumpOrderArgs struct {
	Status enum.AdminDumpOrderStatus `json:"status" validate:"required"`
}

func (p *putAdminDumpOrderArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		fmt.Println()
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminDumpOrderHandler) PutDumpOrderList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	dumpOrderID := cast.ToInt64(util.URLParam(r, "dump_order_id"))

	args := &putAdminDumpOrderArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	orderMap, err := h.adminDumpOrderDAO.BatchGet(ctx, []int64{dumpOrderID})
	util.PanicIf(err)

	order, ok := orderMap[dumpOrderID]
	if !ok {
		util.RenderJSON(w, DefaultSuccessBody(ctx))
		return
	}
	order.OperatorID = loginID
	order.Status = args.Status
	util.PanicIf(h.adminDumpOrderDAO.Update(ctx, order))
	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type deleteAdminDumpOrderArgs struct {
	IpaID      string `json:"ipa_id" validate:"required"`
	IpaVersion string `json:"ipa_version" validate:"required"`
}

func (p *deleteAdminDumpOrderArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		fmt.Println()
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminDumpOrderHandler) DeleteDumpOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)

	args := &deleteAdminDumpOrderArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	util.PanicIf(h.adminDumpOrderCtl.Delete(ctx, loginID, cast.ToInt64(args.IpaID), args.IpaVersion))
}
