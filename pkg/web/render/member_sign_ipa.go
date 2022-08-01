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

type MemberSignIpa struct {
	Meta *models.MemberSignIpa `json:"-"`

	ID          int64  `json:"id,string"`
	IpaName     string `json:"ipa_name"`
	IpaBundleID string `json:"ipa_bundle_id"`

	DownloadURL string `json:"download_url" render:"method=RenderDownloadURL"`

	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"update_at"`
}

type MemberSignIpaRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberSignIpaMap map[int64]*MemberSignIpa

	memberSignIpaDAO dao.MemberSignIpaDAO
	tencentCtl       controller.TencentController
}

type MemberSignIpaOption func(*MemberSignIpaRender)

func MemberSignIpaIncludes(fields []string) MemberSignIpaOption {
	return func(render *MemberSignIpaRender) {
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

var MemberSignIpaDefaultRenderFields = []MemberSignIpaOption{
	MemberSignIpaIncludes([]string{
		"DownloadURL",
	}),
}

func NewMemberSignIpaRender(ids []int64, loginID int64, opts ...MemberSignIpaOption) *MemberSignIpaRender {
	f := &MemberSignIpaRender{
		ids:     ids,
		loginID: loginID,

		memberSignIpaDAO: impl.DefaultMemberSignIpaDAO,
		tencentCtl:       impl2.DefaultTencentController,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *MemberSignIpaRender) RenderSlice(ctx context.Context) []*MemberSignIpa {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*MemberSignIpa, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *MemberSignIpaRender) RenderMap(ctx context.Context) map[int64]*MemberSignIpa {
	if len(f.ids) == 0 {
		return f.memberSignIpaMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, MemberSignIpa{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.memberSignIpaMap
}

func (f *MemberSignIpaRender) fetch(ctx context.Context) {
	result := make(map[int64]*MemberSignIpa)

	metaMap, err := f.memberSignIpaDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)
	for _, id := range f.ids {
		meta, ok := metaMap[id]
		if !ok {
			continue
		}
		result[id] = &MemberSignIpa{
			Meta:        meta,
			ID:          meta.ID,
			IpaName:     meta.IpaName,
			IpaBundleID: meta.IpaBundleID,
			CreatedAt:   meta.CreatedAt.Unix(),
			UpdateAt:    meta.UpdatedAt.Unix(),
		}
	}

	f.memberSignIpaMap = result
}

func (f *MemberSignIpaRender) RenderDownloadURL(ctx context.Context) {
	for _, ipa := range f.memberSignIpaMap {
		url, err := f.tencentCtl.GetMemberSignIpa(ctx, ipa.Meta.IpaPlistFileToken)
		util.PanicIf(err)
		ipa.DownloadURL = url
	}
}
