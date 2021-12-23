package render

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	util2 "dumpapp_server/pkg/util"
)

type MemberDownloadRecord struct {
	ID      int64                           `json:"id,string"`
	Version string                          `json:"version"`
	Status  enum.MemberDownloadNumberStatus `json:"status"`
	IpaType enum.IpaType                    `json:"ipa_type"`
	Ipa     *Ipa                            `json:"ipa"`

	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"update_at"`
}

type MemberDownloadRecordRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberDownloadRecordMap map[int64]*MemberDownloadRecord

	memberDownloadNumberDAO dao.MemberDownloadNumberDAO
}

type MemberDownloadRecordOption func(*MemberDownloadRecordRender)

func MemberDownloadRecordIncludes(fields []string) MemberDownloadRecordOption {
	return func(render *MemberDownloadRecordRender) {
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

var MemberDownloadRecordDefaultRenderFields = []MemberDownloadRecordOption{
	MemberDownloadRecordIncludes([]string{}),
}

func NewMemberDownloadRecordRender(ids []int64, loginID int64, opts ...MemberDownloadRecordOption) *MemberDownloadRecordRender {
	f := &MemberDownloadRecordRender{
		ids:     ids,
		loginID: loginID,

		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *MemberDownloadRecordRender) RenderSlice(ctx context.Context) []*MemberDownloadRecord {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*MemberDownloadRecord, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *MemberDownloadRecordRender) RenderMap(ctx context.Context) map[int64]*MemberDownloadRecord {
	if len(f.ids) == 0 {
		return f.memberDownloadRecordMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, MemberDownloadRecord{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.memberDownloadRecordMap
}

func (f *MemberDownloadRecordRender) fetch(ctx context.Context) {
	memberDownloadMap, err := f.memberDownloadNumberDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	ipaIDs := make([]int64, 0)
	for _, id := range f.ids {
		d, ok := memberDownloadMap[id]
		if !ok {
			continue
		}
		if d.IpaID.Int64 == 0 {
			continue
		}
		ipaIDs = append(ipaIDs, d.IpaID.Int64)
	}
	ipaMap := NewIpaRender(ipaIDs, f.loginID, []enum.IpaType{enum.IpaTypeNormal, enum.IpaTypeCrack}, IpaDefaultRenderFields...).RenderMap(ctx)

	result := make(map[int64]*MemberDownloadRecord)
	for _, id := range f.ids {
		d, ok := memberDownloadMap[id]
		if !ok {
			continue
		}
		if d.IpaID.Int64 == 0 {
			continue
		}
		/// version 没有说明是证书记录，则忽略
		if !d.Version.Valid || !d.IpaType.Valid {
			continue
		}
		/// 小于 2021-12-09 号之前的记录不下发
		if d.CreatedAt.Unix() <= 1638979200 {
			continue
		}
		ipaType, _ := enum.IpaTypeString(d.IpaType.String)
		result[id] = &MemberDownloadRecord{
			ID:        d.ID,
			Version:   d.Version.String,
			Status:    d.Status,
			IpaType:   ipaType,
			Ipa:       ipaMap[d.IpaID.Int64],
			CreatedAt: d.CreatedAt.Unix(),
			UpdateAt:  d.UpdatedAt.Unix(),
		}
	}

	f.memberDownloadRecordMap = result
}
