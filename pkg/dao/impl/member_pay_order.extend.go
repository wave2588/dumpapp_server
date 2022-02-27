package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberPayOrderDAO) GetByFilters(ctx context.Context, filters []qm.QueryMod, orderBys []string) ([]*models.MemberPayOrder, error) {
	qs := make([]qm.QueryMod, 0)
	qs = append(qs, filters...)
	if len(orderBys) > 0 {
		orderBys = append(orderBys, "id desc")
		for _, orderBy := range orderBys {
			qs = append(qs, qm.OrderBy(orderBy))
		}
	} else {
		qs = append(qs, qm.OrderBy("created_at DESC, id DESC"))
	}
	return models.MemberPayOrders(qs...).All(ctx, d.mysqlPool)
}

func (d *MemberPayOrderDAO) BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberPayOrder, error) {
	qs := []qm.QueryMod{
		models.MemberPayOrderWhere.MemberID.IN(memberIDs),
	}
	data, err := models.MemberPayOrders(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[int64][]*models.MemberPayOrder)
	for _, datum := range data {
		res[datum.MemberID] = append(res[datum.MemberID], datum)
	}
	return res, nil
}
