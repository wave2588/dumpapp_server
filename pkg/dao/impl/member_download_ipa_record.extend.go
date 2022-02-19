package impl

import (
	"context"
	"database/sql"
	"time"

	"dumpapp_server/pkg/dao/models"
	pkgErr "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberDownloadIpaRecordDAO) BatchGetMemberNormalCount(ctx context.Context, memberIDs []int64) (map[int64]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.MemberDownloadIpaRecordColumns.MemberID, "count(*) as count"),
		qm.From("member_download_ipa_record"),
		models.MemberDownloadIpaRecordWhere.MemberID.IN(memberIDs),
		models.MemberDownloadIpaRecordWhere.Status.EQ("normal"),
		qm.GroupBy(models.MemberDownloadIpaRecordColumns.MemberID),
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

func (d *MemberDownloadIpaRecordDAO) GetIpaDownloadCount(ctx context.Context, ipaID int64) (int64, int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.MemberDownloadIpaRecordColumns.UpdatedAt, "count(*) as count"),
		qm.From("member_download_ipa_record"),
		models.MemberDownloadIpaRecordWhere.IpaID.EQ(null.Int64From(ipaID)),
		models.MemberDownloadIpaRecordWhere.Status.EQ("used"),
		qm.Limit(1),
		qm.GroupBy(models.MemberDownloadIpaRecordColumns.IpaID),
		qm.OrderBy("updated_at desc"),
	}

	var data struct {
		UpdatedAt time.Time `boil:"updated_at"`
		Count     int64     `boil:"count"`
	}
	err := models.NewQuery(qs...).Bind(ctx, d.mysqlPool, &data)
	if err != nil && pkgErr.Cause(err) != sql.ErrNoRows {
		return 0, 0, err
	}

	return data.Count, data.UpdatedAt.Unix(), nil
}

func (d *MemberDownloadIpaRecordDAO) BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberDownloadIpaRecord, error) {
	qs := []qm.QueryMod{
		models.MemberDownloadIpaRecordWhere.MemberID.IN(memberIDs),
	}
	data, err := models.MemberDownloadIpaRecords(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	result := make(map[int64][]*models.MemberDownloadIpaRecord)
	for _, datum := range data {
		result[datum.MemberID] = append(result[datum.MemberID], datum)
	}
	return result, nil
}
