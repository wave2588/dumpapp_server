package impl

import (
	"context"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberVipOrderDAO) BatchGetOrdersByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberVipOrder, error) {
	qs := []qm.QueryMod{
		models.MemberVipOrderWhere.MemberID.IN(memberIDs),
	}
	data, err := models.MemberVipOrders(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[int64][]*models.MemberVipOrder)
	for _, datum := range data {
		res[datum.MemberID] = append(res[datum.MemberID], datum)
	}
	return res, nil
}
