package render

import (
	"context"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	util2 "dumpapp_server/pkg/util"
)

type Device struct {
	ID           int64          `json:"id,string"`
	UDID         string         `json:"udid"`
	Product      string         `json:"product"`
	Certificates []*Certificate `json:"certificates,omitempty" render:"method=RenderCertificates"` /// 证书列表
}

type DeviceRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	DeviceMap map[int64]*Device

	memberDeviceDAO      dao.MemberDeviceDAO
	certificateDeviceDAO dao.CertificateDeviceDAO
}

type DeviceOption func(*DeviceRender)

func DeviceIncludes(fields []string) DeviceOption {
	return func(render *DeviceRender) {
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

var DeviceDefaultRenderFields = []DeviceOption{
	DeviceIncludes([]string{
		"Certificates",
	}),
}

func NewDeviceRender(ids []int64, loginID int64, opts ...DeviceOption) *DeviceRender {
	f := &DeviceRender{
		ids:     ids,
		loginID: loginID,

		memberDeviceDAO:      impl.DefaultMemberDeviceDAO,
		certificateDeviceDAO: impl.DefaultCertificateDeviceDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *DeviceRender) RenderSlice(ctx context.Context) []*Device {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*Device, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *DeviceRender) RenderMap(ctx context.Context) map[int64]*Device {
	if len(f.ids) == 0 {
		return f.DeviceMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, Device{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.DeviceMap
}

func (f *DeviceRender) fetch(ctx context.Context) {
	deviceMap, err := f.memberDeviceDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*Device)
	for _, id := range f.ids {
		device, ok := deviceMap[id]
		if !ok {
			continue
		}
		result[device.ID] = &Device{
			ID:      device.ID,
			UDID:    device.Udid,
			Product: device.Product,
		}
	}
	f.DeviceMap = result
}

func (f *DeviceRender) RenderCertificates(ctx context.Context) {
	deviceCerMap, err := f.certificateDeviceDAO.BatchGetByDeviceIDs(ctx, f.ids)
	util.PanicIf(err)
	cerIDs := make([]int64, 0)
	for _, devices := range deviceCerMap {
		for _, device := range devices {
			cerIDs = append(cerIDs, device.CertificateID)
		}
	}
	cerMap := NewCertificateRender(cerIDs, f.loginID, CertificateDefaultRenderFields...).RenderMap(ctx)
	for _, device := range f.DeviceMap {
		deviceCers, ok := deviceCerMap[device.ID]
		if !ok {
			continue
		}
		certificates := make([]*Certificate, 0)
		for _, dc := range deviceCers {
			cer, ok := cerMap[dc.CertificateID]
			if !ok {
				continue
			}
			certificates = append(certificates, cer)
		}
		device.Certificates = certificates
	}
}