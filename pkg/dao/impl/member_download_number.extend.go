package impl

import (
	"context"
	"database/sql"
	"time"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	pkgErr "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
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

func (d *MemberDownloadNumberDAO) GetIpaDownloadCount(ctx context.Context, ipaID int64) (int64, int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.MemberDownloadNumberColumns.UpdatedAt, "count(*) as count"),
		qm.From("member_download_number"),
		models.MemberDownloadNumberWhere.IpaID.EQ(null.Int64From(ipaID)),
		models.MemberDownloadNumberWhere.Status.EQ(enum.MemberDownloadNumberStatusUsed),
		qm.Limit(1),
		qm.GroupBy(models.MemberDownloadNumberColumns.IpaID),
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
