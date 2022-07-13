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
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DumpOrderHandler struct {
	dumpOrderDAO      dao.AdminDumpOrderDAO
	adminDumpOrderCtl controller.AdminDumpOrderController
	alterWebCtl       controller2.AlterWebController
}

func NewDumpOrderHandler() *DumpOrderHandler {
	return &DumpOrderHandler{
		dumpOrderDAO:      impl.DefaultAdminDumpOrderDAO,
		adminDumpOrderCtl: impl2.DefaultAdminDumpOrderController,
		alterWebCtl:       impl3.DefaultAlterWebController,
	}
}

type postDumpOrderArgs struct {
	IpaID        string `json:"ipa_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	BundleID     string `json:"bundle_id" validate:"required"`
	Version      string `json:"version" validate:"required"`
	AppStoreLink string `json:"app_store_link" validate:"required"`
	IsOld        bool   `json:"is_old"`
}

func (p *postDumpOrderArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *DumpOrderHandler) Post(w http.ResponseWriter, r *http.Request) {
	args := &postDumpOrderArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	ctx := r.Context()
	loginID := mustGetLoginID(ctx)

	ipaID := cast.ToInt64(args.IpaID)

	/// 写入记录库里
	util.PanicIf(h.adminDumpOrderCtl.Upsert(ctx, loginID, ipaID, args.Name, args.Version, args.BundleID, args.AppStoreLink, args.IsOld))

	/// 库内没有找到对应的砸壳信息，需要发送推送给负责人进行砸壳。
	h.alterWebCtl.SendDumpOrderMsg(ctx, loginID, ipaID, args.BundleID, args.Name, args.Version)

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type dumpOrderItem struct {
	ID         int64                     `json:"id,string"`
	IpaID      int64                     `json:"ipa_id,string"`
	IpaName    string                    `json:"ipa_name"`
	IpaVersion string                    `json:"ipa_version"`
	Status     enum.AdminDumpOrderStatus `json:"status"`
	CreatedAt  int64                     `json:"created_at"`
	UpdatedAt  int64                     `json:"updated_at"`
}

func (h *DumpOrderHandler) GetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		offset = GetIntArgument(r, "offset", 0)
		limit  = GetIntArgument(r, "limit", 10)
	)

	filter := []qm.QueryMod{
		qm.Where("(status=?) or (status=?)", enum.AdminDumpOrderStatusUnprocessed, enum.AdminDumpOrderStatusProgressing),
	}
	ids, err := h.dumpOrderDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)
	totalCount, err := h.dumpOrderDAO.Count(ctx, filter)
	util.PanicIf(err)

	dumpOrderMap, err := h.dumpOrderDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	res := make([]*dumpOrderItem, 0)

	for _, id := range ids {
		do, ok := dumpOrderMap[id]
		if !ok {
			continue
		}
		var bizExt constant.AdminDumpOrderBizExt
		util.PanicIf(json.Unmarshal([]byte(do.IpaBizExt), &bizExt))

		res = append(res, &dumpOrderItem{
			ID:         id,
			IpaID:      do.IpaID,
			IpaName:    bizExt.IpaName,
			IpaVersion: do.IpaVersion,
			Status:     do.Status,
			CreatedAt:  do.CreatedAt.Unix(),
			UpdatedAt:  do.UpdatedAt.Unix(),
		})
	}

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   res,
	})
}
