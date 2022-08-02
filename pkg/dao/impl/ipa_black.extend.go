package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *IpaBlackDAO) BatchGetByIpaIDs(ctx context.Context, ipaIDs []int64) (map[int64][]*models.IpaBlack, error) {
	qs := []qm.QueryMod{
		models.IpaBlackWhere.IpaID.IN(ipaIDs),
	}
	data, err := models.IpaBlacks(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	result := make(map[int64][]*models.IpaBlack)
	for _, datum := range data {
		result[datum.IpaID] = append(result[datum.IpaID], datum)
	}
	return result, nil
}

func (d *IpaBlackDAO) AdminListIpaIDs(ctx context.Context, offset, limit int) ([]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.IpaBlackColumns.IpaID),
		qm.Offset(offset),
		qm.Limit(limit),
		qm.OrderBy("id desc"),
		qm.GroupBy("ipa_id"),
	}

	datas, err := models.IpaBlacks(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}

	result := make([]int64, 0)
	for _, c := range datas {
		result = append(result, c.IpaID)
	}
	return result, nil
}
