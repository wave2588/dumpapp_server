package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/controller"
	"dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type IpaSignHandler struct {
	ipaSignWebCtl controller.IpaSignWebController
	ipaSignDAO    dao.IpaSignDAO
}

func NewIpaSignHandler() *IpaSignHandler {
	return &IpaSignHandler{
		ipaSignWebCtl: impl.DefaultIpaSignWebController,
		ipaSignDAO:    impl2.DefaultIpaSignDAO,
	}
}

type postSignArgs struct {
	CertificateID int64 `json:"certificate_id,string" validate:"required"`
	IpaVersionID  int64 `json:"ipa_version_id,string" validate:"required"`
}

func (p *postSignArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *IpaSignHandler) PostSign(w http.ResponseWriter, r *http.Request) {
	args := &postSignArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	ctx := r.Context()

	loginID := mustGetLoginID(ctx)
	util.PanicIf(h.ipaSignWebCtl.AddSignTask(ctx, loginID, args.CertificateID, args.IpaVersionID))

	util.RenderJSON(w, "ok")
}

func (h *IpaSignHandler) GetMemberSignList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	filter := []qm.QueryMod{
		models.IpaSignWhere.MemberID.EQ(loginID),
	}
	ids, err := h.ipaSignDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

	count, err := h.ipaSignDAO.Count(ctx, filter)
	util.PanicIf(err)

	data := render.NewIpaSignRender(ids, loginID, render.IpaSignDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   data,
	})
}
