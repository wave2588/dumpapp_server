package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberPayCountDAO) BatchGetMemberNormalCount(ctx context.Context, memberIDs []int64) (map[int64]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.MemberPayCountColumns.MemberID, "count(*) as count"),
		qm.From("member_pay_count"),
		models.MemberPayCountWhere.MemberID.IN(memberIDs),
		models.MemberPayCountWhere.Status.EQ(enum.MemberPayCountStatusNormal),
		qm.GroupBy(models.MemberPayCountColumns.MemberID),
	}

	var data []struct {
		MemberID int64 `boil:"member_id"`
		Count    int64 `boil:"count"`
	}
	err := models.NewQuery(qs...).Bind(ctx, d.mysqlPool, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]int64)
	for _, datum := range data {
		result[datum.MemberID] = datum.Count
	}
	return result, nil
}

func (d *MemberPayCountDAO) BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberPayCount, error) {
	qs := []qm.QueryMod{
		models.MemberPayCountWhere.MemberID.IN(memberIDs),
	}
	data, err := models.MemberPayCounts(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[int64][]*models.MemberPayCount)
	for _, datum := range data {
		res[datum.MemberID] = append(res[datum.MemberID], datum)
	}
	return res, nil
}
