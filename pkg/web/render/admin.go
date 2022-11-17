package render

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	util2 "dumpapp_server/pkg/util"
	"github.com/spf13/cast"
)

type Admin struct {
	OrderCount int64 `json:"order_count"` /// 提交的订单
	PaidCount  int64 `json:"paid_count"`  /// 支付成功的订单
}

type AdminRender struct {
	memberIDs     []int64
	loginID       int64
	includeFields []string

	adminMap map[int64]*Admin

	memberPayOrderDAO dao.MemberPayOrderDAO
}

func NewAdminRender(memberIDs []int64, loginID int64) *AdminRender {
	f := &AdminRender{
		memberIDs: memberIDs,
		loginID:   loginID,

		memberPayOrderDAO: impl.DefaultMemberPayOrderDAO,
	}
	return f
}

func (f *AdminRender) RenderSlice(ctx context.Context) []*Admin {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*Admin, len(f.memberIDs))
	for i, id := range f.memberIDs {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *AdminRender) RenderMap(ctx context.Context) map[int64]*Admin {
	if len(f.memberIDs) == 0 {
		return f.adminMap
	}

	f.fetch(ctx)

	err := util2.AutoRender(ctx, f, Admin{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.adminMap
}

func (f *AdminRender) fetch(ctx context.Context) {
	memberOrderMap, err := f.memberPayOrderDAO.BatchGetByMemberIDs(ctx, f.memberIDs)
	util.PanicIf(err)

	result := make(map[int64]*Admin)
	for _, memberID := range f.memberIDs {
		orders := memberOrderMap[memberID]
		paidCount := 0
		for _, order := range orders {
			if order.Status == enum.MemberPayOrderStatusPaid {
				paidCount++
			}
		}
		result[memberID] = &Admin{
			OrderCount: cast.ToInt64(len(orders)),
			PaidCount:  cast.ToInt64(paidCount),
		}
	}
	f.adminMap = result
}
