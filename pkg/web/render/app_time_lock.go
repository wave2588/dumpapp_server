package render

import (
	"context"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

type AppTimeLock struct {
	meta *models.AppTimeLock `json:"-"`

	ID          int64  `json:"id,string"`
	IsDelete    bool   `json:"is_delete"`
	IsStop      bool   `json:"is_stop"`
	Description string `json:"description"` /// 一些额外的描述，由客户端直接上报 json string
	StartAt     int64  `json:"start_at"`
	EndAt       int64  `json:"end_at"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`

	Member *Member `json:"member" render:"method=RenderMember"`
}

type AppTimeLockRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	appTimeLockMap map[int64]*AppTimeLock

	appTimeLockDAO dao.AppTimeLockDAO
}

type AppTimeLockOption func(*AppTimeLockRender)

func AppTimeLockIncludes(fields []string) AppTimeLockOption {
	return func(render *AppTimeLockRender) {
		fields = append(fields, defaultFields...)
		uniqFields := make([]string, 0)
		fieldSet := util2.NewSet()
		for _, field := range fields {
			if fieldSet.Exists(field) {
				continue
			}
			fieldSet.Add(field)
			uniqFields = append(uniqFields, field)
		}
		render.includeFields = uniqFields
	}
}

var AppTimeLockDefaultRenderFields = []AppTimeLockOption{
	AppTimeLockIncludes([]string{
		"Member",
	}),
}

func NewAppTimeLockRender(ids []int64, loginID int64, opts ...AppTimeLockOption) *AppTimeLockRender {
	f := &AppTimeLockRender{
		ids:     ids,
		loginID: loginID,

		appTimeLockDAO: impl.DefaultAppTimeLockDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *AppTimeLockRender) RenderSlice(ctx context.Context) []*AppTimeLock {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*AppTimeLock, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *AppTimeLockRender) RenderMap(ctx context.Context) map[int64]*AppTimeLock {
	if len(f.ids) == 0 {
		return f.appTimeLockMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, AppTimeLock{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.appTimeLockMap
}

func (f *AppTimeLockRender) fetch(ctx context.Context) {
	metaMap, err := f.appTimeLockDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*AppTimeLock)
	for _, id := range f.ids {
		meta, ok := metaMap[id]
		if !ok {
			continue
		}
		data := &AppTimeLock{
			meta:        meta,
			ID:          meta.ID,
			IsDelete:    meta.IsDelete,
			IsStop:      meta.IsStop,
			Description: meta.BizExt.Description,
			StartAt:     meta.StartAt.Unix(),
			EndAt:       meta.EndAt.Unix(),
			CreatedAt:   meta.CreatedAt.Unix(),
			UpdatedAt:   meta.UpdatedAt.Unix(),
		}
		if meta.IsStop {
			data.IsStop = meta.IsStop
		} else {
			/// 判断是否还在有效时间内
			now := time.Now().Unix()
			if meta.StartAt.Unix() < now && now < meta.EndAt.Unix() {
				data.IsStop = true
			} else {
				data.IsStop = false
			}
		}
		result[id] = data
	}

	f.appTimeLockMap = result
}

func (f *AppTimeLockRender) RenderMember(ctx context.Context) {
	memberIDs := make([]int64, 0)
	for _, lock := range f.appTimeLockMap {
		memberIDs = append(memberIDs, lock.meta.MemberID)
	}
	memberIDs = util2.RemoveDuplicates(memberIDs)

	memberMap := NewMemberRender(memberIDs, f.loginID, MemberDefaultRenderFields...).RenderMap(ctx)
	for _, lock := range f.appTimeLockMap {
		lock.Member = memberMap[lock.meta.MemberID]
	}
}
