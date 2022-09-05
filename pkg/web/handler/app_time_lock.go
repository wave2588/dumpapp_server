package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AppTimeLockHandler struct {
	appTimeLockDAO dao.AppTimeLockDAO
}

func NewAppTimeLockHandler() *AppTimeLockHandler {
	return &AppTimeLockHandler{
		appTimeLockDAO: impl.DefaultAppTimeLockDAO,
	}
}

type postAppTimeLockArgs struct {
	StartAt     int64  `json:"start_at" validate:"required"`
	EndAt       int64  `json:"end_at" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (p *postAppTimeLockArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.StartAt >= p.EndAt {
		return errors.UnproccessableError("开始时间不能大于或等于结束时间")
	}
	return nil
}

func (h *AppTimeLockHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	/// 时间锁暂时先不限制了
	//count, err := h.appTimeLockDAO.Count(ctx, []qm.QueryMod{
	//	models.AppTimeLockWhere.MemberID.EQ(loginID),
	//	models.AppTimeLockWhere.IsDelete.EQ(false),
	//})
	//util.PanicIf(err)
	//if count >= 5 {
	//	util.PanicIf(errors.UnproccessableError("时间锁当前仅开放最多创建 5 个"))
	//}

	args := &postAppTimeLockArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	id := util2.MustGenerateID(ctx)
	util.PanicIf(h.appTimeLockDAO.Insert(ctx, &models.AppTimeLock{
		ID:       id,
		MemberID: loginID,
		IsDelete: false,
		IsStop:   false,
		StartAt:  time.Unix(args.StartAt, 0),
		EndAt:    time.Unix(args.EndAt, 0),
		BizExt: datatype.AppTimeLockBizExt{
			Description: args.Description,
		},
	}))

	respMap := render.NewAppTimeLockRender([]int64{id}, loginID, render.AppTimeLockDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, respMap[id])
}

type putAppTimeLockArgs struct {
	StartAt     int64  `json:"start_at" validate:"required"`
	EndAt       int64  `json:"end_at" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsStop      bool   `json:"is_stop"` /// 犹豫不能传 0 和 false，所以 1 表示为停止，2 标识为不停止
}

func (p *putAppTimeLockArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.StartAt >= p.EndAt {
		return errors.UnproccessableError("开始时间不能大于或等于结束时间")
	}
	return nil
}

func (h *AppTimeLockHandler) Put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &putAppTimeLockArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	timeLockID := cast.ToInt64(util.URLParam(r, "id"))
	timeLockMap, err := h.appTimeLockDAO.BatchGet(ctx, []int64{timeLockID})
	util.PanicIf(err)
	timeLock, ok := timeLockMap[timeLockID]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}

	/// 只能自己修改
	if timeLock.MemberID != loginID {
		util.PanicIf(errors.ErrMemberAccessDenied)
	}

	timeLock.StartAt = time.Unix(args.StartAt, 0)
	timeLock.EndAt = time.Unix(args.EndAt, 0)
	timeLock.BizExt.Description = args.Description
	timeLock.IsStop = args.IsStop
	util.PanicIf(h.appTimeLockDAO.Update(ctx, timeLock))

	respMap := render.NewAppTimeLockRender([]int64{timeLockID}, loginID, render.AppTimeLockDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, respMap[timeLockID])
}

func (h *AppTimeLockHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)
	timeLockID := cast.ToInt64(util.URLParam(r, "id"))

	timeLockMap, err := h.appTimeLockDAO.BatchGet(ctx, []int64{timeLockID})
	util.PanicIf(err)

	timeLock, ok := timeLockMap[timeLockID]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}

	/// 只能自己删除
	if timeLock.MemberID != loginID {
		util.PanicIf(errors.ErrMemberAccessDenied)
	}

	timeLock.IsDelete = true
	timeLock.IsStop = true
	util.PanicIf(h.appTimeLockDAO.Update(ctx, timeLock))

	respMap := render.NewAppTimeLockRender([]int64{timeLockID}, loginID, render.AppTimeLockDefaultRenderFields...).RenderMap(ctx)
	util.RenderJSON(w, respMap[timeLockID])
}

func (h *AppTimeLockHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	timeLockID := cast.ToInt64(util.URLParam(r, "id"))

	respMap := render.NewAppTimeLockRender([]int64{timeLockID}, 0, render.AppTimeLockDefaultRenderFields...).RenderMap(ctx)
	resp, ok := respMap[timeLockID]
	if !ok {
		util.PanicIf(errors.ErrNotFound)
	}
	util.RenderJSON(w, resp)
}

func (h *AppTimeLockHandler) GetList(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	filter := []qm.QueryMod{
		models.AppTimeLockWhere.MemberID.EQ(loginID),
		models.AppTimeLockWhere.IsDelete.EQ(false),
	}
	ids, err := h.appTimeLockDAO.ListIDs(ctx, offset, limit, filter, nil)
	util.PanicIf(err)

	totalCount, err := h.appTimeLockDAO.Count(ctx, filter)
	util.PanicIf(err)

	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   render.NewAppTimeLockRender(ids, loginID, render.AppTimeLockDefaultRenderFields...).RenderSlice(ctx),
	})
}
