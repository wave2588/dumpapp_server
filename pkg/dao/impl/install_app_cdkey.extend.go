package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *InstallAppCdkeyDAO) BatchGetByOrderIDs(ctx context.Context, orderIDs []int64) (map[int64][]*models.InstallAppCdkey, error) {
	qs := []qm.QueryMod{
		models.InstallAppCdkeyWhere.OrderID.IN(orderIDs),
	}

	data, err := models.InstallAppCdkeys(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}

	res := make(map[int64][]*models.InstallAppCdkey)
	for _, datum := range data {
		res[datum.OrderID] = append(res[datum.OrderID], datum)
	}
	return res, nil
}
