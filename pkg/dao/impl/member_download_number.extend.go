package impl

import (
	"context"
	"github.com/volatiletech/null/v8"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberDownloadNumberDAO) BatchGetMemberNormalCount(ctx context.Context, memberIDs []int64) (map[int64]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.MemberDownloadNumberColumns.MemberID, "count(*) as count"),
		qm.From("member_download_number"),
		models.MemberDownloadNumberWhere.MemberID.IN(memberIDs),
		models.MemberDownloadNumberWhere.Status.EQ(enum.MemberDownloadNumberStatusNormal),
		qm.GroupBy(models.MemberDownloadNumberColumns.MemberID),
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

func (d *MemberDownloadNumberDAO) GetByMemberIDAndIpaID(ctx context.Context, memberID, ipaID int64) ([]*models.MemberDownloadNumber, error) {
	qs := []qm.QueryMod{
		models.MemberDownloadNumberWhere.MemberID.EQ(memberID),
		models.MemberDownloadNumberWhere.IpaID.EQ(null.Int64From(ipaID)),
	}

	data, err := models.MemberDownloadNumbers(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}

	res := make([]*models.MemberDownloadNumber, 0)
	for _, re := range data {
		res = append(res, re)
	}
	return res, nil
}
