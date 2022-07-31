package render

import (
	"context"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

type AppSource struct {
	meta          *models.AppSource         `json:"-"`
	AppSourceInfo *controller.AppSourceInfo `json:"-"`

	ID  int64  `json:"id,string"`
	URL string `json:"url"`

	Name     string `json:"name" render:"method=RenderName"`
	Icon     string `json:"icon"`
	IsActive bool   `json:"is_active"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type AppSourceRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	appSourceMap map[int64]*AppSource

	appSourceDAO dao.AppSourceDAO
	appSourceCtl controller.AppSourceController
}

type AppSourceOption func(*AppSourceRender)

func AppSourceIncludes(fields []string) AppSourceOption {
	return func(render *AppSourceRender) {
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

var AppSourceDefaultRenderFields = []AppSourceOption{
	AppSourceIncludes([]string{
		"Name",
	}),
}

func NewAppSourceRender(ids []int64, loginID int64, opts ...AppSourceOption) *AppSourceRender {
	f := &AppSourceRender{
		ids:     ids,
		loginID: loginID,

		appSourceDAO: impl.DefaultAppSourceDAO,
		appSourceCtl: impl2.DefaultAppSourceController,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *AppSourceRender) RenderSlice(ctx context.Context) []*AppSource {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*AppSource, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *AppSourceRender) RenderMap(ctx context.Context) map[int64]*AppSource {
	if len(f.ids) == 0 {
		return f.appSourceMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, AppSource{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.appSourceMap
}

func (f *AppSourceRender) fetch(ctx context.Context) {
	metaMap, err := f.appSourceDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*AppSource)
	for _, id := range f.ids {
		meta, ok := metaMap[id]
		if !ok {
			continue
		}
		result[id] = &AppSource{
			meta:      meta,
			ID:        meta.ID,
			URL:       meta.URL,
			CreatedAt: meta.CreatedAt.Unix(),
			UpdatedAt: meta.UpdatedAt.Unix(),
		}
	}
	f.appSourceMap = result
}

func (f *AppSourceRender) RenderName(ctx context.Context) {
	URLs := make([]string, 0)
	for _, source := range f.appSourceMap {
		URLs = append(URLs, source.URL)
	}

	infoMap, err := f.appSourceCtl.BatchGetAppSourceInfo(ctx, URLs)
	util.PanicIf(err)

	for _, source := range f.appSourceMap {
		sourceInfo, ok := infoMap[source.meta.URL]
		source.AppSourceInfo = sourceInfo
		if !ok || sourceInfo.Appstore != nil || len(sourceInfo.Apps) == 0 {
			source.Name = "已加密"
			continue
		}
		source.Name = sourceInfo.Name
		source.Icon = sourceInfo.Sourceicon
		source.IsActive = true
	}
}
