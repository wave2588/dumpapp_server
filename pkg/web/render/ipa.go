package render

import (
	"context"
	"sort"

	"dumpapp_server/pkg/common/enum"
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

	ID        int64  `json:"id,string"`
	Name      string `json:"name"`
	BundleID  string `json:"bundle_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`

	Counter *Counter `json:"counter,omitempty" render:"method=RenderCounter"`

	Versions []*Version `json:"versions,omitempty" render:"method=RenderVersions"`

	/// 是否在黑名单中
	BlackInfo BlackInfo `json:"black_info" render:"method=RenderBlack"`
}

type Version struct {
	ID          int64        `json:"id,string"`
	Version     string       `json:"version"`
	IpaType     enum.IpaType `json:"ipa_type"`
	DescribeURL *string      `json:"describe_url,omitempty"`
	Describe    *string      `json:"describe,omitempty"`
	Size        int64        `json:"size"`
	Country     string       `json:"country"`
	CreatedAt   int64        `json:"created_at"`
	UpdatedAt   int64        `json:"updated_at"`
}

type Counter struct {
	DownloadCount    int64 `json:"download_count"`
	LastDownloadTime int64 `json:"last_download_time"`
}

type BlackInfo struct {
	IsBlack bool               `json:"is_black"`
	Reasons []*BlackInfoReason `json:"reasons"`
}

type BlackInfoReason struct {
	ID     int64  `json:"id,string"`
	Reason string `json:"reason"`
}

type IpaRender struct {
	ids               []int64
	loginID           int64
	supportIpaTypeMap map[enum.IpaType]struct{}
	includeFields     []string

	IpaMap map[int64]*Ipa

	ipaDAO                     dao.IpaDAO
	ipaVersionDAO              dao.IpaVersionDAO
	ipaBlackDAO                dao.IpaBlackDAO
	memberDownloadIpaRecordDAO dao.MemberDownloadIpaRecordDAO

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

var IpaAdminRenderFields = []IpaOption{
	IpaIncludes([]string{
		"Versions",
		"Counter",
		"BlackInfo",
	}),
}

var IpaDefaultRenderFields = []IpaOption{
	IpaIncludes([]string{
		"Versions",
		"BlackInfo",
	}),
}

func NewIpaRender(ids []int64, loginID int64, ipaTypes []enum.IpaType, opts ...IpaOption) *IpaRender {
	ipaTypeMap := make(map[enum.IpaType]struct{})
	for _, ipaType := range ipaTypes {
		ipaTypeMap[ipaType] = struct{}{}
	}

	f := &IpaRender{
		ids:               ids,
		loginID:           loginID,
		supportIpaTypeMap: ipaTypeMap,

		ipaDAO:                     impl.DefaultIpaDAO,
		ipaVersionDAO:              impl.DefaultIpaVersionDAO,
		ipaBlackDAO:                impl.DefaultIpaBlackDAO,
		memberDownloadIpaRecordDAO: impl.DefaultMemberDownloadIpaRecordDAO,

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

	err := util2.AutoRender(ctx, f, Ipa{}, f.includeFields)
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
			meta:      a,
			ID:        a.ID,
			Name:      a.Name,
			BundleID:  a.BundleID,
			CreatedAt: a.CreatedAt.Unix(),
			UpdatedAt: a.UpdatedAt.Unix(),
		}
	}
	f.IpaMap = res
}

func (f *IpaRender) RenderVersions(ctx context.Context) {
	totalVersionMap, err := f.ipaVersionDAO.BatchGetIpaVersions(ctx, f.ids)
	util.PanicIf(err)

	for _, ipa := range f.IpaMap {
		vs := totalVersionMap[ipa.ID]
		if vs == nil || len(vs) == 0 {
			continue
		}

		sort.Slice(vs, func(i, j int) bool {
			re1 := vs[i].Version
			re2 := vs[j].Version
			// return util2.CompareLittleVer(re1, re2) == util2.VersionCompareResBig
			return util2.Compare(re1, re2) == util2.VersionCompareResBig
		})

		res := make([]*Version, 0)
		for _, v := range vs {
			_, ok := f.supportIpaTypeMap[v.IpaType]
			if !ok {
				continue
			}
			version := &Version{
				ID:        v.ID,
				Version:   v.Version,
				IpaType:   v.IpaType,
				Size:      v.BizExt.Size,
				Country:   v.BizExt.Country,
				CreatedAt: v.CreatedAt.Unix(),
				UpdatedAt: v.UpdatedAt.Unix(),
			}
			if v.BizExt.DescribeURL != nil && *v.BizExt.DescribeURL != "" {
				version.DescribeURL = v.BizExt.DescribeURL
			}
			if v.BizExt.Describe != nil && *v.BizExt.Describe != "" {
				version.Describe = v.BizExt.Describe
			}
			res = append(res, version)
		}
		ipa.Versions = res
	}
}

func (f *IpaRender) RenderCounter(ctx context.Context) {
	for _, ipa := range f.IpaMap {
		count, updatedAt, err := f.memberDownloadIpaRecordDAO.GetIpaDownloadCount(ctx, ipa.ID)
		util.PanicIf(err)
		if count != 0 {
			ipa.Counter = &Counter{
				DownloadCount:    count,
				LastDownloadTime: updatedAt,
			}
		}
	}
}

func (f *IpaRender) RenderBlack(ctx context.Context) {
	blackMap, err := f.ipaBlackDAO.BatchGetByIpaIDs(ctx, f.ids)
	util.PanicIf(err)

	for _, ipa := range f.IpaMap {
		blacks := blackMap[ipa.meta.ID]
		reasons := make([]*BlackInfoReason, 0)
		for _, black := range blacks {
			reasons = append(reasons, &BlackInfoReason{
				ID:     black.ID,
				Reason: black.BizExt.Reason,
			})
		}
		ipa.BlackInfo = BlackInfo{
			IsBlack: len(reasons) != 0,
			Reasons: reasons,
		}
	}
}
