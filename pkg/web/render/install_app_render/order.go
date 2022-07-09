package install_app_render

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

type CDKeyOrder struct {
	Meta *models.InstallAppCdkeyOrder `json:"-"`

	ID        int64                     `json:"id,string"`
	Status    enum.MemberPayOrderStatus `json:"status"`
	Number    int64                     `json:"number"`
	CreatedAt int64                     `json:"created_at"`
	UpdatedAt int64                     `json:"updated_at"`

	Source     string `json:"source"`
	ContactWat string `json:"contact_wat"`
}

type CDKeyOrderRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	cdkeyOrderMap map[int64]*CDKeyOrder

	cdkeyOrderDAO dao.InstallAppCdkeyOrderDAO
}

type CDKeyOrderOption func(*CDKeyOrderRender)

func CDKeyOrderIncludes(fields []string) CDKeyOrderOption {
	return func(render *CDKeyOrderRender) {
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

var CDKeyOrderDefaultRenderFields = []CDKeyOrderOption{
	CDKeyOrderIncludes([]string{}),
}

func NewCDKeyOrderRender(ids []int64, loginID int64, opts ...CDKeyOrderOption) *CDKeyOrderRender {
	f := &CDKeyOrderRender{
		ids:     ids,
		loginID: loginID,

		cdkeyOrderDAO: impl.DefaultInstallAppCdkeyOrderDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *CDKeyOrderRender) RenderSlice(ctx context.Context) []*CDKeyOrder {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*CDKeyOrder, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *CDKeyOrderRender) RenderMap(ctx context.Context) map[int64]*CDKeyOrder {
	if len(f.ids) == 0 {
		return f.cdkeyOrderMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, CDKeyOrder{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.cdkeyOrderMap
}

func (f *CDKeyOrderRender) fetch(ctx context.Context) {
	orderMap, err := f.cdkeyOrderDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*CDKeyOrder)
	for _, id := range f.ids {
		order, ok := orderMap[id]
		if !ok {
			continue
		}

		var bizExt constant.InstallAppCDKEYOrderBizExt
		util.PanicIf(json.Unmarshal([]byte(order.BizExt), &bizExt))

		source := "normal"
		if bizExt.IsTest {
			source = "admin"
		}
		result[id] = &CDKeyOrder{
			Meta:       order,
			ID:         order.ID,
			Status:     order.Status,
			Number:     order.Number,
			CreatedAt:  order.CreatedAt.Unix(),
			UpdatedAt:  order.UpdatedAt.Unix(),
			ContactWat: bizExt.ContactWay,
			Source:     source,
		}
	}

	f.cdkeyOrderMap = result
}
