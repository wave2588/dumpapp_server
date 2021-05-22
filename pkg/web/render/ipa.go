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

type Ipa struct {
	meta *models.Ipa

	ID   int64  `json:"id,string"`
	Name string `json:"name"`

	Versions []*Version `json:"versions,omitempty" render:"method=RenderVersions"`
}

type Version struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

type IpaRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	IpaMap map[int64]*Ipa

	ipaDAO        dao.IpaDAO
	ipaVersionDAO dao.IpaVersionDAO

	tencentCtl controller.TencentController
}

type IpaOption func(*IpaRender)

var defaultFields = []string{}

func IpaIncludes(fields []string) IpaOption {
	return func(render *IpaRender) {
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

var IpaDefaultRenderFields = []IpaOption{
	IpaIncludes([]string{
		"Versions",
	}),
}

func NewIpaRender(ids []int64, loginID int64, opts ...IpaOption) *IpaRender {
	f := &IpaRender{
		ids:     ids,
		loginID: loginID,

		ipaDAO:        impl.DefaultIpaDAO,
		ipaVersionDAO: impl.DefaultIpaVersionDAO,

		tencentCtl: impl2.DefaultTencentController,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *IpaRender) RenderSlice(ctx context.Context) []*Ipa {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*Ipa, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *IpaRender) RenderMap(ctx context.Context) map[int64]*Ipa {
	if len(f.ids) == 0 {
		return f.IpaMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, Ipa{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.IpaMap
}

func (f *IpaRender) fetch(ctx context.Context) {
	aMap, err := f.ipaDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	res := make(map[int64]*Ipa)
	for _, a := range aMap {
		res[a.ID] = &Ipa{
			meta: a,
			ID:   a.ID,
			Name: a.Name,
		}
	}
	f.IpaMap = res
}

func (f *IpaRender) RenderVersions(ctx context.Context) {
	totalVersionMap, err := f.ipaVersionDAO.BatchGetIpaVersions(ctx, f.ids)
	util.PanicIf(err)

	versionMap, err := f.ipaVersionDAO.BatchGetLatestVersion(ctx, f.ids)
	util.PanicIf(err)

	memberMap := NewMemberRender([]int64{f.loginID}, f.loginID, MemberDefaultRenderFields...).RenderMap(ctx)
	member := memberMap[f.loginID]

	for _, ipa := range f.IpaMap {
		res := make([]*Version, 0)
		/// 如果是 vip 返回所有 ipa
		if member.Vip.IsVip {
			if vs, ok := totalVersionMap[ipa.ID]; ok {
				for _, v := range vs {
					url, err := f.tencentCtl.GetSignatureURL(ctx, v.TokenPath)
					util.PanicIf(err)
					res = append(res, &Version{
						Version: v.Version,
						URL:     url,
					})
				}
			}
		} else {
			if v, ok := versionMap[ipa.ID]; ok {
				url, err := f.tencentCtl.GetSignatureURL(ctx, v.TokenPath)
				util.PanicIf(err)
				res = append(res, &Version{
					Version: v.Version,
					URL:     url,
				})
			}
		}
		ipa.Versions = res
	}
}
