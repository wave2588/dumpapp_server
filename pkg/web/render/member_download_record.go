package render

import (
	"context"
	"dumpapp_server/pkg/common/enum"
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
	ids, err := f.memberDownloadNumberDAO.batch
}
