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
	qm := []qm.QueryMod{
		qm.Select("member_id, count(id) as count"),
		qm.GroupBy("member_id"),
		qm.OrderBy("count desc"),
		qm.Offset(offset),
		qm.Limit(limit),
	}
	qm = append(qm, filter...)
	query := models.MemberDownloadOrders(qm...)
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
	qm := []qm.QueryMod{
		qm.Select("member_id, count(id) as count"),
		qm.GroupBy("member_id"),
		qm.OrderBy("count desc"),
	}
	qm = append(qm, filter...)
	query := models.MemberDownloadOrders(qm...)
	data, err := query.All(ctx, d.mysqlPool)
	if err != nil {
		return 0, err
	}
	return cast.ToInt64(len(data)), nil
}
