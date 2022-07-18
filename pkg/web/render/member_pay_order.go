package render

import (
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"github.com/spf13/cast"
	"golang.org/x/net/context"
)

type MemberPayOrder struct {
	meta *models.MemberPayOrder `json:"-"`

	ID     int64                     `json:"id,string"`
	Status enum.MemberPayOrderStatus `json:"status"`
	Amount int64                     `json:"amount"`

	Member *Member `json:"member" render:"method=RenderMember"`

	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"update_at"`
}

type MemberPayOrderRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberPayOrderMap map[int64]*MemberPayOrder

	memberPayOrderDAO dao.MemberPayOrderDAO
}

type MemberPayOrderOption func(*MemberPayOrderRender)

func MemberPayOrderIncludes(fields []string) MemberPayOrderOption {
	return func(render *MemberPayOrderRender) {
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

var MemberPayOrderDefaultRenderFields = []MemberPayOrderOption{
	MemberPayOrderIncludes([]string{
		"Member",
	}),
}

func NewMemberPayOrderRender(ids []int64, loginID int64, opts ...MemberPayOrderOption) *MemberPayOrderRender {
	f := &MemberPayOrderRender{
		ids:     ids,
		loginID: loginID,

		memberPayOrderDAO: impl.DefaultMemberPayOrderDAO,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *MemberPayOrderRender) RenderSlice(ctx context.Context) []*MemberPayOrder {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*MemberPayOrder, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *MemberPayOrderRender) RenderMap(ctx context.Context) map[int64]*MemberPayOrder {
	if len(f.ids) == 0 {
		return f.memberPayOrderMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, MemberPayOrder{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.memberPayOrderMap
}

func (f *MemberPayOrderRender) fetch(ctx context.Context) {
	result := make(map[int64]*MemberPayOrder)

	orderMap, err := f.memberPayOrderDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)
	for _, id := range f.ids {
		order, ok := orderMap[id]
		if !ok {
			continue
		}
		result[id] = &MemberPayOrder{
			meta:      order,
			ID:        order.ID,
			Status:    order.Status,
			Amount:    cast.ToInt64(order.Amount),
			CreatedAt: order.CreatedAt.Unix(),
			UpdateAt:  order.UpdatedAt.Unix(),
		}
	}

	f.memberPayOrderMap = result
}

func (f *MemberPayOrderRender) RenderMember(ctx context.Context) {
	memberIDs := make([]int64, 0)
	for _, order := range f.memberPayOrderMap {
		memberIDs = append(memberIDs, order.meta.MemberID)
	}
	memberIDs = util2.RemoveDuplicates(memberIDs)
	memberMap := NewMemberRender(memberIDs, f.loginID, MemberDefaultRenderFields...).RenderMap(ctx)
	for _, order := range f.memberPayOrderMap {
		order.Member = memberMap[order.meta.MemberID]
	}
}
