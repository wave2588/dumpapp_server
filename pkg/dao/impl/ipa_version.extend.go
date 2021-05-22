package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *IpaVersionDAO) BatchGetIpaVersions(ctx context.Context, ipaIDs []int64) (map[int64][]*models.IpaVersion, error) {
	qs := []qm.QueryMod{
		models.IpaVersionWhere.IpaID.IN(ipaIDs),
		qm.OrderBy("created_at desc"),
	}
	data, err := models.IpaVersions(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[int64][]*models.IpaVersion)
	for _, datum := range data {
		res[datum.IpaID] = append(res[datum.IpaID], datum)
	}
	return res, nil
}

func (d *IpaVersionDAO) BatchGetLatestVersion(ctx context.Context, ipaIDs []int64) (map[int64]*models.IpaVersion, error) {
	qs := []qm.QueryMod{
		models.IpaVersionWhere.IpaID.IN(ipaIDs),
		qm.OrderBy("created_at desc"),
	}
	data, err := models.IpaVersions(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[int64]*models.IpaVersion)
	for _, datum := range data {
		res[datum.IpaID] = datum
	}
	return res, nil
}
