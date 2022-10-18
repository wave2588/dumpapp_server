package render

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"github.com/spf13/cast"
)

type MemberSignIpa struct {
	Meta *models.MemberSignIpa `json:"-"`

	ID              int64  `json:"id,string"`
	ExpenseID       string `json:"expense_id"`
	IpaName         string `json:"ipa_name"`
	IpaBundleID     string `json:"ipa_bundle_id"`
	IpaVersion      string `json:"ipa_version"`
	IpaSize         int64  `json:"ipa_size"`
	CertificateName string `json:"certificate_name"`
	IsDelete        bool   `json:"is_delete"`

	DownloadURL string  `json:"download_url" render:"method=RenderDownloadURL"`
	PlistURL    *string `json:"plist_url,omitempty" render:"method=RenderPlistURL"`

	/// 分发相关的控制
	Dispense *Dispense `json:"dispense" render:"method=RenderDispense"`

	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"update_at"`
}

type Dispense struct {
	Count     int64 `json:"count"`      /// 用户设置的可分发次数
	UsedCount int64 `json:"used_count"` /// 使用过的分发次数
	IsValid   bool  `json:"is_valid"`   /// 是否还能使用
}

type MemberSignIpaRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberSignIpaMap map[int64]*MemberSignIpa

	memberSignIpaDAO       dao.MemberSignIpaDAO
	dispenseCountRecordDAO dao.DispenseCountRecordDAO
	lingshulianCtl         controller.LingshulianController
	fileCtl                controller.FileController
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

var DefaultMemberSignIpaFields = []string{
	"DownloadURL",
	"Dispense",
}

var MemberSignIpaDefaultRenderFields = []MemberSignIpaOption{
	MemberSignIpaIncludes(DefaultMemberSignIpaFields),
}

func NewMemberSignIpaRender(ids []int64, loginID int64, opts ...MemberSignIpaOption) *MemberSignIpaRender {
	f := &MemberSignIpaRender{
		ids:     ids,
		loginID: loginID,

		memberSignIpaDAO:       impl.DefaultMemberSignIpaDAO,
		dispenseCountRecordDAO: impl.DefaultDispenseCountRecordDAO,
		lingshulianCtl:         impl2.DefaultLingshulianController,
		fileCtl:                impl2.DefaultFileController,
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
			Meta:            meta,
			ID:              meta.ID,
			ExpenseID:       meta.ExpenseID,
			IpaName:         meta.BizExt.IpaName,
			IpaBundleID:     meta.BizExt.IpaBundleID,
			IpaVersion:      meta.BizExt.IpaVersion,
			IpaSize:         meta.BizExt.IpaSize,
			CertificateName: meta.BizExt.CertificateName,
			IsDelete:        meta.IsDelete,
			CreatedAt:       meta.CreatedAt.Unix(),
			UpdateAt:        meta.UpdatedAt.Unix(),
		}
	}

	f.memberSignIpaMap = result
}

func (f *MemberSignIpaRender) RenderDownloadURL(ctx context.Context) {
	for _, ipa := range f.memberSignIpaMap {
		ipa.DownloadURL = fmt.Sprintf("https://www.dumpapp.com/installipa?id=%d", ipa.Meta.ID)
	}
}

func (f *MemberSignIpaRender) RenderPlistURL(ctx context.Context) {
	for _, ipa := range f.memberSignIpaMap {
		ipa.PlistURL = util.StringPtr(f.fileCtl.GetPlistFileURL(ctx, ipa.Meta.IpaPlistFileToken))
	}
}

func (f *MemberSignIpaRender) RenderDispense(ctx context.Context) {
	signIpaIDs := make([]int64, 0)
	for _, ipa := range f.memberSignIpaMap {
		if ipa.Meta.BizExt.DispenseCount == 0 {
			continue
		}
		signIpaIDs = append(signIpaIDs, ipa.ID)
	}

	res, err := f.dispenseCountRecordDAO.BatchGetByObjectIDsAndRecordType(ctx, signIpaIDs, enum.DispenseCountRecordTypeInstallSignIpa)
	util.PanicIf(err)

	for _, ipa := range f.memberSignIpaMap {
		usedCount := cast.ToInt64(len(res[ipa.ID]))
		/// 等于 0 相当于没设置, 按照不需要拦截处理
		dispenseCount := ipa.Meta.BizExt.DispenseCount
		if dispenseCount == 0 {
			ipa.Dispense = &Dispense{
				Count:     0,
				UsedCount: usedCount,
				IsValid:   true,
			}
			continue
		}
		ipa.Dispense = &Dispense{
			Count:     dispenseCount,
			UsedCount: usedCount,
			IsValid:   usedCount < dispenseCount,
		}
	}
}
