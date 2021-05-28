package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *SearchRecordV2DAO) BatchGetByIpaIDs(ctx context.Context, ipaIDs []int64, filters []qm.QueryMod) ([]*models.SearchRecordV2, error) {
	qs := []qm.QueryMod{
		models.SearchRecordV2Where.IpaID.IN(ipaIDs),
	}
	qs = append(qs, filters...)
	data, err := models.SearchRecordV2S(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make([]*models.SearchRecordV2, 0)
	for _, datum := range data {
		res = append(res, datum)
	}

	return res, nil
}
