package handler

import (
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/render"
	"github.com/spf13/cast"
)

type IpaHandler struct {
	ipaDAO            dao.IpaDAO
	ipaVersionDAO     dao.IpaVersionDAO
	searchRecordV2DAO dao.SearchRecordV2DAO

	memberDownloadCtl controller.MemberDownloadController
}

func NewIpaHandler() *IpaHandler {
	return &IpaHandler{
		ipaDAO:            impl.DefaultIpaDAO,
		ipaVersionDAO:     impl.DefaultIpaVersionDAO,
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,

		memberDownloadCtl: impl2.DefaultMemberDownloadController,
	}
}

func (h *IpaHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	offset := GetIntArgument(r, "offset", 0)
	limit := GetIntArgument(r, "limit", 10)

	count, err := h.ipaDAO.Count(ctx, nil)
	util.PanicIf(err)

	ids, err := h.ipaDAO.ListIDs(ctx, offset, limit, nil, nil)
	util.PanicIf(err)

	data := render.NewIpaRender(ids, loginID, render.IpaDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(count), offset, limit),
		Data:   data,
	})
}

type getIpaArgs struct {
	Name string `form:"name" validate:"required"`
}

func (p *getIpaArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *IpaHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	args := getIpaArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	/// 记录用户获取的记录
	util.PanicIf(h.searchRecordV2DAO.Insert(ctx, &models.SearchRecordV2{
		MemberID: loginID,
		IpaID:    ipaID,
		Name:     args.Name,
	}))

	_, err := h.memberDownloadCtl.GetDownloadNumber(ctx, loginID)
	util.PanicIf(err)

	data := render.NewIpaRender([]int64{ipaID}, loginID, render.IpaDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[ipaID])
}
