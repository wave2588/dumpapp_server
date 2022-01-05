package render

import (
	"context"
	"encoding/json"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

type IpaSign struct {
	meta *models.IpaSign

	ID        int64              `json:"id,string"`
	Status    enum.IpaSignStatus `json:"status"`
	URL       string             `json:"url"`
	CreatedAt int64              `json:"created_at"`
	UpdateAt  int64              `json:"update_at"`

	CurrentIpaVersion string       `json:"current_ipa_version"`
	CurrentIpaType    enum.IpaType `json:"current_ipa_type"`
	Ipa               *Ipa         `json:"ipa,omitempty" render:"method=RenderIpa"`
}

type IpaSignRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	ipaSignMap map[int64]*IpaSign

	ipaSignDAO dao.IpaSignDAO
}

type IpaSignOption func(*IpaSignRender)

var defaultIpaSignFields = []string{}

func IpaSignIncludes(fields []string) IpaSignOption {
	return func(render *IpaSignRender) {
		fields = append(fields, defaultIpaSignFields...)
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

var IpaSignDefaultRenderFields = []IpaSignOption{
	IpaSignIncludes([]string{
		"Ipa",
	}),
}

func NewIpaSignRender(ids []int64, loginID int64, opts ...IpaSignOption) *IpaSignRender {
	f := &IpaSignRender{
		ids:     ids,
		loginID: loginID,

		ipaSignDAO: impl.DefaultIpaSignDAO,
	}

	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *IpaSignRender) RenderSlice(ctx context.Context) []*IpaSign {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*IpaSign, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *IpaSignRender) RenderMap(ctx context.Context) map[int64]*IpaSign {
	if len(f.ids) == 0 {
		return f.ipaSignMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, IpaSign{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.ipaSignMap
}

func (f *IpaSignRender) fetch(ctx context.Context) {
	data, err := f.ipaSignDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*IpaSign)
	for ipaSignID, ipaSign := range data {
		var ipaSignBizExt constant.IpaSignBizExt
		util.PanicIf(json.Unmarshal([]byte(ipaSign.BizExt), &ipaSignBizExt))

		result[ipaSignID] = &IpaSign{
			meta:              ipaSign,
			ID:                ipaSign.ID,
			Status:            ipaSign.Status,
			URL:               "",
			CreatedAt:         ipaSign.CreatedAt.Unix(),
			UpdateAt:          ipaSign.UpdatedAt.Unix(),
			CurrentIpaVersion: ipaSignBizExt.IpaVersion,
			CurrentIpaType:    ipaSignBizExt.IpaType,
		}
	}

	f.ipaSignMap = result
}

func (f *IpaSignRender) RenderIpa(ctx context.Context) {
	ipaIDs := make([]int64, 0)
	for _, sign := range f.ipaSignMap {
		ipaIDs = append(ipaIDs, sign.meta.IpaID)
	}
	ipaMap := NewIpaRender(ipaIDs, f.loginID, []enum.IpaType{enum.IpaTypeNormal, enum.IpaTypeCrack}, IpaDefaultRenderFields...).RenderMap(ctx)

	for _, sign := range f.ipaSignMap {
		sign.Ipa = ipaMap[sign.meta.IpaID]
	}
}
