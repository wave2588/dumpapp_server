package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberDownloadOrderDAO) BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberDownloadOrder, error) {
	qs := []qm.QueryMod{
		models.MemberDownloadOrderWhere.MemberID.IN(memberIDs),
	}
	data, err := models.MemberDownloadOrders(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[int64][]*models.MemberDownloadOrder)
	for _, datum := range data {
		res[datum.MemberID] = append(res[datum.MemberID], datum)
	}
	return res, nil
}

func (d *MemberDownloadOrderDAO) GetMemberIDsOrderByPaidCount(ctx context.Context, offset, limit int, filter []qm.QueryMod) ([]int64, error) {
	qs := []qm.QueryMod{
		qm.Select("member_id, count(id) as count"),
		qm.GroupBy("member_id"),
		qm.OrderBy("count desc"),
		qm.Offset(offset),
		qm.Limit(limit),
	}
	qs = append(qs, filter...)
	query := models.MemberDownloadOrders(qs...)
	rows, err := query.QueryContext(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	result := make([]int64, 0)
	for rows.Next() {
		var MemberID, Count int64
		err = rows.Scan(&MemberID, &Count)
		result = append(result, MemberID)
	}
	return result, nil
}

func (d *MemberDownloadOrderDAO) CountMemberIDsOrderByPaidCount(ctx context.Context, filter []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{
		qm.Select("member_id, count(id) as count"),
		qm.GroupBy("member_id"),
		qm.OrderBy("count desc"),
	}
	qs = append(qs, filter...)
	query := models.MemberDownloadOrders(qs...)
	data, err := query.All(ctx, d.mysqlPool)
	if err != nil {
		return 0, err
	}
	return cast.ToInt64(len(data)), nil
}

func (d *MemberDownloadOrderDAO) GetByFilters(ctx context.Context, filters []qm.QueryMod, orderBys []string) ([]*models.MemberDownloadOrder, error) {
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
	return models.MemberDownloadOrders(qs...).All(ctx, d.mysqlPool)
}
