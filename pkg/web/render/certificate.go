package render

import (
	"context"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl3 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/render/install_app_render"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type Certificate struct {
	Meta *models.CertificateV2 `json:"-"`

	ID        int64 `json:"id,string"`
	CreatedAt int64 `json:"created_at"`
	ExpireAt  int64 `json:"expire_at"`
	UpdatedAt int64 `json:"updated_at"`

	/// 备注
	Note string `json:"note"`

	/// p12 文件密码
	P12Password     string `json:"p12_password"`
	P12             string `json:"p12"`
	Mobileprovision string `json:"mobileprovision"`
	Level           int    `json:"level"` /// 0: 未知   1: 普通版   2: 高级版  3: 豪华版

	IsReplenish *bool `json:"is_replenish,omitempty" render:"method=RenderIsReplenish"` // 是否是候补证书

	/// p12 文件是否有效
	P12IsActive bool `json:"p12_is_active" render:"method=RenderP12IsActive"`
	/// 证书对应绑定的设备
	Device *Device `json:"device" render:"method=RenderDevice"`

	CdKey *install_app_render.CDKEY `json:"cd_key,omitempty" render:"method=RenderCdKey"`
}

type CertificateRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	certificateMap map[int64]*Certificate

	certificateDAO       dao.CertificateV2DAO
	certificateDeviceDAO dao.CertificateDeviceDAO
	certificateBaseCtl   controller.CertificateBaseController
}

type CertificateOption func(*CertificateRender)

func CertificateIncludes(fields []string) CertificateOption {
	return func(render *CertificateRender) {
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

var CertificateDefaultRenderFields = []CertificateOption{
	CertificateIncludes([]string{
		"P12IsActive",
		"Device",
		"IsReplenish",
		"CdKey",
	}),
}

func NewCertificateRender(ids []int64, loginID int64, opts ...CertificateOption) *CertificateRender {
	f := &CertificateRender{
		ids:     ids,
		loginID: loginID,

		certificateDAO:       impl.DefaultCertificateV2DAO,
		certificateDeviceDAO: impl.DefaultCertificateDeviceDAO,
		certificateBaseCtl:   impl3.DefaultCertificateBaseController,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *CertificateRender) RenderSlice(ctx context.Context) []*Certificate {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*Certificate, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *CertificateRender) RenderMap(ctx context.Context) map[int64]*Certificate {
	if len(f.ids) == 0 {
		return f.certificateMap
	}

	f.fetch(ctx)

	err := util2.AutoRender(ctx, f, Certificate{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.certificateMap
}

func (f *CertificateRender) fetch(ctx context.Context) {
	cerMap, err := f.certificateDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*Certificate)
	for _, id := range f.ids {
		meta, ok := cerMap[id]
		if !ok {
			continue
		}

		cer := &Certificate{
			Meta:            meta,
			ID:              meta.ID,
			CreatedAt:       meta.CreatedAt.Unix(),
			ExpireAt:        meta.CreatedAt.AddDate(1, 0, 0).Unix(),
			UpdatedAt:       meta.UpdatedAt.Unix(),
			Note:            meta.BizExt.Note,
			P12Password:     meta.BizExt.NewP12Password,
			P12:             meta.ModifiedP12FileDate,
			Mobileprovision: meta.MobileProvisionFileData,
			Level:           meta.BizExt.Level,
		}

		/// fixme: 做个兜底策略, 防止 read |0: file already closed 错误再次出现
		if meta.ModifiedP12FileDate == "" {
			cer.P12 = meta.P12FileData
			cer.P12Password = meta.BizExt.OriginalP12Password
		}

		result[meta.ID] = cer
	}
	f.certificateMap = result
}

func (f *CertificateRender) RenderP12IsActive(ctx context.Context) {
	cerMetas := make([]*models.CertificateV2, 0)
	for _, certificate := range f.certificateMap {
		cerMetas = append(cerMetas, certificate.Meta)
	}

	isActiveMap, err := f.certificateBaseCtl.CheckCertificateIsActiveByModels(ctx, cerMetas)
	util.PanicIf(err)

	for _, certificate := range f.certificateMap {
		if isActive, ok := isActiveMap[certificate.ID]; ok {
			certificate.P12IsActive = isActive
		} else {
			certificate.P12IsActive = true // todo:  如果没获取到, 默认展示有效
		}
	}
}

func (f *CertificateRender) RenderDevice(ctx context.Context) {
	deviceIDs := make([]int64, 0)
	for _, certificate := range f.certificateMap {
		deviceIDs = append(deviceIDs, certificate.Meta.DeviceID)
	}
	deviceIDs = util2.RemoveDuplicates(deviceIDs)

	deviceMap := NewDeviceRender(deviceIDs, f.loginID).RenderMap(ctx)
	for _, certificate := range f.certificateMap {
		certificate.Device = deviceMap[certificate.Meta.DeviceID]
	}
}

func (f *CertificateRender) RenderIsReplenish(ctx context.Context) {
	cdMap, err := f.certificateDeviceDAO.BatchGetByCertificateID(ctx, f.ids)
	util.PanicIf(err)

	endAt := time.Date(2023, 2, 3, 15, 0, 0, 0, time.Now().Location())
	for _, certificate := range f.certificateMap {
		if certificate.CreatedAt < endAt.Unix() {
			continue
		}
		_, ok := cdMap[certificate.ID]
		certificate.IsReplenish = common.BoolPtr(!ok) // 如果存在说明不是候补证书
	}
}

func (f *CertificateRender) RenderCdKey(ctx context.Context) {
	cdKeyIDs := make([]int64, 0)
	for _, certificate := range f.certificateMap {
		cdKeyID := certificate.Meta.BizExt.CdKeyID
		if cdKeyID == 0 {
			continue
		}
		cdKeyIDs = append(cdKeyIDs, cdKeyID)
	}

	cdKeyMap := install_app_render.NewCDKEYRender(cdKeyIDs, f.loginID, install_app_render.CDKeyDefaultRenderFields...).RenderMap(ctx)
	for _, certificate := range f.certificateMap {
		certificate.CdKey = cdKeyMap[certificate.Meta.BizExt.CdKeyID]
	}
}
