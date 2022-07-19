package render

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

type MemberPayCountRecord struct {
	meta *models.MemberPayCountRecord

	ID int64 `json:"id,string"`

	Count int64 `json:"count"`

	Type        string `json:"type"`
	Description string `json:"description" render:"method=RenderDescription"`

	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"update_at"`
}

type MemberPayCountRecordRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberPayCountRecordMap map[int64]*MemberPayCountRecord

	memberPayCountRecordDAO dao.MemberPayCountRecordDAO
}

type MemberPayCountRecordOption func(*MemberPayCountRecordRender)

func MemberPayCountRecordIncludes(fields []string) MemberPayCountRecordOption {
	return func(render *MemberPayCountRecordRender) {
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

var MemberPayCountRecordDefaultRenderFields = []MemberPayCountRecordOption{
	MemberPayCountRecordIncludes([]string{
		"Description",
	}),
}

func NewMemberPayCountRecordRender(ids []int64, loginID int64, opts ...MemberPayCountRecordOption) *MemberPayCountRecordRender {
	f := &MemberPayCountRecordRender{
		ids:     ids,
		loginID: loginID,

		memberPayCountRecordDAO: impl.DefaultMemberPayCountRecordDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *MemberPayCountRecordRender) RenderSlice(ctx context.Context) []*MemberPayCountRecord {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*MemberPayCountRecord, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *MemberPayCountRecordRender) RenderMap(ctx context.Context) map[int64]*MemberPayCountRecord {
	if len(f.ids) == 0 {
		return f.memberPayCountRecordMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, MemberPayCountRecord{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.memberPayCountRecordMap
}

func (f *MemberPayCountRecordRender) fetch(ctx context.Context) {
	result := make(map[int64]*MemberPayCountRecord)

	data, err := f.memberPayCountRecordDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	for _, id := range f.ids {
		meta, ok := data[id]
		if !ok {
			continue
		}
		result[id] = &MemberPayCountRecord{
			meta:      meta,
			ID:        meta.ID,
			Count:     meta.Count,
			CreatedAt: meta.CreatedAt.Unix(),
			UpdateAt:  meta.UpdatedAt.Unix(),
		}
	}

	f.memberPayCountRecordMap = result
}

func (f *MemberPayCountRecordRender) RenderDescription(ctx context.Context) {
	for _, record := range f.memberPayCountRecordMap {
		switch record.meta.Type {
		case enum.MemberPayCountRecordTypePay:
			record.Type = "add"
			record.Description = fmt.Sprintf("充值了 %d 个 D 币", record.meta.Count)
		case enum.MemberPayCountRecordTypePayForFree:
			record.Type = "add"
			record.Description = fmt.Sprintf("充值赠送了 %d 个 D 币", record.meta.Count)
		case enum.MemberPayCountRecordTypeAdminPresented:
			record.Type = "add"
			record.Description = fmt.Sprintf("系统添加了 %d 个 D 币", record.meta.Count)
		case enum.MemberPayCountRecordTypeInvitedPresented:
			record.Type = "add"
			record.Description = fmt.Sprintf("邀请新用户赠送了 %d 个 D 币", record.meta.Count)
		case enum.MemberPayCountRecordTypeRebate:
			record.Type = "add"
			record.Description = fmt.Sprintf("邀请用户充值返利赠送 %d 个 D 币", record.meta.Count)
		case enum.MemberPayCountRecordTypeBuyIpa:
			record.Type = "deduct"
			record.Description = fmt.Sprintf("购买 ipa 消费了 %d 个 D 币", record.meta.Count)
		case enum.MemberPayCountRecordTypeBuyCertificate:
			record.Type = "deduct"
			record.Description = fmt.Sprintf("购买证书消费了 %d 个 D 币", record.meta.Count)
		case enum.MemberPayCountRecordTypeAdminDelete:
			record.Type = "deduct"
			record.Description = fmt.Sprintf("系统删除了 %d 个 D 币", record.meta.Count)
		}
	}
}
