package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AppSourceHandler struct {
	appSourceCtl       controller.AppSourceController
	memberAppSourceDAO dao.MemberAppSourceDAO
}

func NewAppSourceHandler() *AppSourceHandler {
	return &AppSourceHandler{
		appSourceCtl:       impl.DefaultAppSourceController,
		memberAppSourceDAO: impl2.DefaultMemberAppSourceDAO,
	}
}

type postAppSourceArgs struct {
	URL string `json:"url" validate:"required"`
}

func (p *postAppSourceArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	err = util2.CheckURLValid(p.URL)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("链接格式错误: %s", err.Error()))
	}
	return nil
}

func (h *AppSourceHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)

	args := &postAppSourceArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// dumpapp 的源地址就不用重复添加了
	if defaultAppSource := h.getDefaultAppSource(ctx, loginID); defaultAppSource != nil {
		if args.URL == defaultAppSource.AppSource.URL {
			util.RenderJSON(w, defaultAppSource)
			return
		}
	}

	id, err := h.appSourceCtl.Insert(ctx, loginID, args.URL)
	util.PanicIf(err)

	data := render.NewMemberAppSourceRender([]int64{id}, loginID, render.MemberAppSourceDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[id])
}

func (h *AppSourceHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	memberAppSourceID := cast.ToInt64(util.URLParam(r, "id"))
	memberAppSourceMap, err := h.memberAppSourceDAO.BatchGet(ctx, []int64{memberAppSourceID})
	util.PanicIf(err)

	memberAppSource, ok := memberAppSourceMap[memberAppSourceID]
	if !ok {
		util.RenderJSON(w, h.getDefaultAppSource(ctx, loginID))
		return
	}
	if memberAppSource.MemberID != loginID {
		util.PanicIf(errors.ErrMemberAccessDenied)
	}

	data := render.NewMemberAppSourceRender([]int64{memberAppSourceID}, loginID, render.MemberAppSourceDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, data[memberAppSourceID])
}

func (h *AppSourceHandler) GetSelfList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	filter := []qm.QueryMod{
		models.MemberAppSourceWhere.MemberID.EQ(loginID),
	}
	ids, err := h.memberAppSourceDAO.ListIDs(ctx, offset, limit, filter, []string{"id desc"})
	util.PanicIf(err)

	totalCount, err := h.memberAppSourceDAO.Count(ctx, filter)
	util.PanicIf(err)

	resultData := make([]*render.MemberAppSource, 0)
	data := render.NewMemberAppSourceRender(ids, loginID, render.MemberAppSourceDefaultRenderFields...).RenderSlice(ctx)

	/// 第一页强插 dumpapp
	if offset == 0 {
		if d := h.getDefaultAppSource(ctx, loginID); d != nil {
			resultData = append(resultData, d)
		}
	}
	resultData = append(resultData, data...)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   resultData,
	})
}

func (h *AppSourceHandler) getDefaultAppSource(ctx context.Context, loginID int64) *render.MemberAppSource {
	defaultID := int64(1)
	appSourceMap := render.NewAppSourceRender([]int64{defaultID}, loginID, render.AppSourceDefaultRenderFields...).RenderMap(ctx)
	appSource, ok := appSourceMap[defaultID]
	if !ok {
		return nil
	}
	return &render.MemberAppSource{
		ID:            1,
		AppSource:     appSource,
		AppSourceMeta: appSource.AppSourceInfo,
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
	}
}

func (h *AppSourceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	memberAppSourceID := cast.ToInt64(util.URLParam(r, "id"))
	memberAppSourceMap, err := h.memberAppSourceDAO.BatchGet(ctx, []int64{memberAppSourceID})
	util.PanicIf(err)

	memberAppSource, ok := memberAppSourceMap[memberAppSourceID]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}
	if memberAppSource.MemberID != loginID {
		util.PanicIf(errors.ErrMemberAccessDenied)
	}

	util.PanicIf(h.memberAppSourceDAO.Delete(ctx, memberAppSourceID))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
