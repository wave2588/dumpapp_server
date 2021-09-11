package impl

import (
	"context"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/models"
	"github.com/spf13/cast"
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

func (d *SearchRecordV2DAO) GetOrderBySearchCount(ctx context.Context, offset, limit int, filter []qm.QueryMod) ([]*dao.SearchCount, error) {
	query := models.SearchRecordV2S(
		qm.Select("ipa_id, name, count(id) as count"),
		qm.GroupBy("ipa_id"),
		qm.OrderBy("count desc"),
		qm.Offset(offset),
		qm.Limit(limit),
	)
	res, err := query.QueryContext(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}

	result := make([]*dao.SearchCount, 0)
	for res.Next() {
		r := &dao.SearchCount{}
		err = res.Scan(&r.IpaID, &r.Name, &r.Count)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (d *SearchRecordV2DAO) CountOrderBySearchCount(ctx context.Context, filter []qm.QueryMod) (int64, error) {
	qm := []qm.QueryMod{
		qm.Select("ipa_id, name, count(id) as count"),
		qm.GroupBy("ipa_id"),
		qm.OrderBy("count desc"),
	}
	qm = append(qm, filter...)
	query := models.SearchRecordV2S(qm...)

	res, err := query.All(ctx, d.mysqlPool)
	if err != nil {
		return 0, err
	}
	return cast.ToInt64(len(res)), nil
}
