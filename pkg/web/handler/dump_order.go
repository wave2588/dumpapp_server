package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/errors"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
)

type DumpOrderHandler struct {
	adminDumpOrderCtl controller.AdminDumpOrderController
	alterWebCtl       controller2.AlterWebController
}

func NewDumpOrderHandler() *DumpOrderHandler {
	return &DumpOrderHandler{
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
