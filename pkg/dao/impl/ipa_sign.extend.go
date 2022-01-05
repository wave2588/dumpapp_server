package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *IpaSignDAO) GetByStatus(ctx context.Context, status enum.IpaSignStatus) ([]*models.IpaSign, error) {
	qs := []qm.QueryMod{
		models.IpaSignWhere.Status.EQ(status),
	}
	return models.IpaSigns(qs...).All(ctx, d.mysqlPool)
}

func (d *IpaSignDAO) BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.IpaSign, error) {
	qs := []qm.QueryMod{
		models.IpaSignWhere.MemberID.IN(memberIDs),
	}
	data, err := models.IpaSigns(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	result := make(map[int64][]*models.IpaSign)
	for _, datum := range data {
		result[datum.MemberID] = append(result[datum.MemberID], datum)
	}
	return result, nil
}
