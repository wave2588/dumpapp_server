package render

import (
	"context"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

type MemberAppSource struct {
	Meta *models.MemberAppSource `json:"-"`

	ID int64 `json:"id,string"`

	AppSource     *AppSource                `json:"app_source" render:"method=RenderAppSource"`
	AppSourceMeta *controller.AppSourceInfo `json:"app_source_meta"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type MemberAppSourceRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberAppSourceMap map[int64]*MemberAppSource

	memberAppSourceDAO dao.MemberAppSourceDAO
}

type MemberAppSourceOption func(*MemberAppSourceRender)

func MemberAppSourceIncludes(fields []string) MemberAppSourceOption {
	return func(render *MemberAppSourceRender) {
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

var MemberAppSourceDefaultRenderFields = []MemberAppSourceOption{
	MemberAppSourceIncludes([]string{
		"AppSource",
	}),
}

func NewMemberAppSourceRender(ids []int64, loginID int64, opts ...MemberAppSourceOption) *MemberAppSourceRender {
	f := &MemberAppSourceRender{
		ids:     ids,
		loginID: loginID,

		memberAppSourceDAO: impl.DefaultMemberAppSourceDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *MemberAppSourceRender) RenderSlice(ctx context.Context) []*MemberAppSource {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*MemberAppSource, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *MemberAppSourceRender) RenderMap(ctx context.Context) map[int64]*MemberAppSource {
	if len(f.ids) == 0 {
		return f.memberAppSourceMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, MemberAppSource{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.memberAppSourceMap
}

func (f *MemberAppSourceRender) fetch(ctx context.Context) {
	metaMap, err := f.memberAppSourceDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*MemberAppSource)
	for _, id := range f.ids {
		meta, ok := metaMap[id]
		if !ok {
			continue
		}
		result[id] = &MemberAppSource{
			Meta:      meta,
			ID:        meta.ID,
			CreatedAt: meta.CreatedAt.Unix(),
			UpdatedAt: meta.UpdatedAt.Unix(),
		}
	}
	f.memberAppSourceMap = result
}

func (f *MemberAppSourceRender) RenderAppSource(ctx context.Context) {
	appSourceIDs := make([]int64, 0)
	for _, source := range f.memberAppSourceMap {
		appSourceIDs = append(appSourceIDs, source.Meta.AppSourceID)
	}
	appSourceIDs = util2.RemoveDuplicates(appSourceIDs)
	appSourceMap := NewAppSourceRender(appSourceIDs, f.loginID, AppSourceDefaultRenderFields...).RenderMap(ctx)
	for _, source := range f.memberAppSourceMap {
		appSource, ok := appSourceMap[source.Meta.AppSourceID]
		if !ok {
			continue
		}
		source.AppSource = appSource
		source.AppSourceMeta = appSource.AppSourceInfo
	}
}
