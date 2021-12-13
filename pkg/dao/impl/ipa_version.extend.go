package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
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

func (d *IpaVersionDAO) GetByIpaIDAndIpaType(ctx context.Context, ipaID int64, ipaType enum.IpaType) ([]*models.IpaVersion, error) {
	qs := []qm.QueryMod{
		models.IpaVersionWhere.IpaID.EQ(ipaID),
		models.IpaVersionWhere.IpaType.EQ(ipaType),
		qm.OrderBy("created_at desc"),
	}
	return models.IpaVersions(qs...).All(ctx, d.mysqlPool)
}

func (d *IpaVersionDAO) GetByIpaIDAndIpaTypeAndVersion(ctx context.Context, ipaID int64, ipaType enum.IpaType, version string) ([]*models.IpaVersion, error) {
	qs := []qm.QueryMod{
		models.IpaVersionWhere.IpaID.EQ(ipaID),
		models.IpaVersionWhere.IpaType.EQ(ipaType),
		models.IpaVersionWhere.Version.EQ(version),
		qm.OrderBy("created_at desc"),
	}
	return models.IpaVersions(qs...).All(ctx, d.mysqlPool)
}

func (d *IpaVersionDAO) GetByIpaID(ctx context.Context, ipaID int64) ([]*models.IpaVersion, error) {
	qs := []qm.QueryMod{
		models.IpaVersionWhere.IpaID.EQ(ipaID),
		qm.OrderBy("created_at desc"),
	}
	return models.IpaVersions(qs...).All(ctx, d.mysqlPool)
}
