package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *AdminNeedDumpIpaDAO) GetByIpaIDAndIpaVersion(ctx context.Context, ipaID int64, ipaVersion string) ([]*models.AdminNeedDumpIpa, error) {
	qs := []qm.QueryMod{
		models.AdminNeedDumpIpaWhere.IpaID.EQ(ipaID),
		models.AdminNeedDumpIpaWhere.IpaVersion.EQ(ipaVersion),
	}
	return models.AdminNeedDumpIpas(qs...).All(ctx, d.mysqlPool)
}
