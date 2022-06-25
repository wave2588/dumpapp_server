package install_app_render

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	util2 "dumpapp_server/pkg/util"
)

type CDKEY struct {
	ID     int64                      `json:"id,string"`
	OutID  string                     `json:"out_id"`
	Status enum.InstallAppCDKeyStatus `json:"status"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type CDKEYRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	cKeyMap map[int64]*CDKEY

	installAppCDKeyDAO dao.InstallAppCdkeyDAO
}

type DeviceOption func(*CDKEYRender)

var defaultFields = []string{}

func DeviceIncludes(fields []string) DeviceOption {
	return func(render *CDKEYRender) {
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

func NewCDKEYRender(ids []int64, loginID int64, opts ...DeviceOption) *CDKEYRender {
	f := &CDKEYRender{
		ids:     ids,
		loginID: loginID,

		installAppCDKeyDAO: impl.DefaultInstallAppCdkeyDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *CDKEYRender) RenderSlice(ctx context.Context) []*CDKEY {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*CDKEY, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *CDKEYRender) RenderMap(ctx context.Context) map[int64]*CDKEY {
	if len(f.ids) == 0 {
		return f.cKeyMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, CDKEY{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.cKeyMap
}

func (f *CDKEYRender) fetch(ctx context.Context) {
	cMap, err := f.installAppCDKeyDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*CDKEY)
	for _, id := range f.ids {
		c, ok := cMap[id]
		if !ok {
			continue
		}
		result[id] = &CDKEY{
			ID:        c.ID,
			OutID:     c.OutID,
			Status:    c.Status,
			CreatedAt: c.CreatedAt.Unix(),
			UpdatedAt: c.UpdatedAt.Unix(),
		}
	}
	f.cKeyMap = result
}
