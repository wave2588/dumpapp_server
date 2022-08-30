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

type DispenseCountRecord struct {
	meta *models.DispenseCountRecord

	ID int64 `json:"id,string"`

	Count int64 `json:"count"`

	Type        string `json:"type"`
	Description string `json:"description" render:"method=RenderDescription"`

	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"update_at"`
}

type DispenseCountRecordRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	dispenseCountRecordMap map[int64]*DispenseCountRecord

	dispenseCountRecordDAO dao.DispenseCountRecordDAO
}

type DispenseCountRecordOption func(*DispenseCountRecordRender)

func DispenseCountRecordIncludes(fields []string) DispenseCountRecordOption {
	return func(render *DispenseCountRecordRender) {
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

var DispenseCountRecordDefaultRenderFields = []DispenseCountRecordOption{
	DispenseCountRecordIncludes([]string{
		"Description",
	}),
}

func NewDispenseCountRecordRender(ids []int64, loginID int64, opts ...DispenseCountRecordOption) *DispenseCountRecordRender {
	f := &DispenseCountRecordRender{
		ids:     ids,
		loginID: loginID,

		dispenseCountRecordDAO: impl.DefaultDispenseCountRecordDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *DispenseCountRecordRender) RenderSlice(ctx context.Context) []*DispenseCountRecord {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*DispenseCountRecord, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *DispenseCountRecordRender) RenderMap(ctx context.Context) map[int64]*DispenseCountRecord {
	if len(f.ids) == 0 {
		return f.dispenseCountRecordMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, DispenseCountRecord{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.dispenseCountRecordMap
}

func (f *DispenseCountRecordRender) fetch(ctx context.Context) {
	result := make(map[int64]*DispenseCountRecord)

	data, err := f.dispenseCountRecordDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	for _, id := range f.ids {
		meta, ok := data[id]
		if !ok {
			continue
		}
		result[id] = &DispenseCountRecord{
			meta:      meta,
			ID:        meta.ID,
			Count:     meta.Count,
			CreatedAt: meta.CreatedAt.Unix(),
			UpdateAt:  meta.UpdatedAt.Unix(),
		}
	}

	f.dispenseCountRecordMap = result
}

func (f *DispenseCountRecordRender) RenderDescription(ctx context.Context) {
	for _, record := range f.dispenseCountRecordMap {
		switch record.meta.Type {
		case enum.DispenseCountRecordTypePay:
			record.Type = "add"
			record.Description = fmt.Sprintf("兑换了 %d 个下载次数", record.meta.Count)
		case enum.DispenseCountRecordTypeInstallSignIpa:
			record.Type = "deduct"
			record.Description = fmt.Sprintf("安装消费了 %d 个下载次数", record.meta.Count)
		case enum.DispenseCountRecordTypeAdminPresented:
			record.Type = "add"
			record.Description = fmt.Sprintf("系统添加了 %d 个下载次数", record.meta.Count)
		case enum.DispenseCountRecordTypeAdminDeleted:
			record.Type = "deduct"
			record.Description = fmt.Sprintf("系统删除了 %d 个下载次数", record.meta.Count)
		}
	}
}
