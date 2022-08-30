package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *DispenseCountDAO) BatchGetMemberNormalCount(ctx context.Context, memberIDs []int64) (map[int64]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.DispenseCountColumns.MemberID, "count(*) as count"),
		qm.From("dispense_count"),
		models.DispenseCountWhere.MemberID.IN(memberIDs),
		models.DispenseCountWhere.Status.EQ(enum.DispenseCountStatusNormal),
		qm.GroupBy(models.DispenseCountColumns.MemberID),
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
