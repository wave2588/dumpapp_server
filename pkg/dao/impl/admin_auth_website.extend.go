package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *AdminAuthWebsiteDAO) IsExistDomain(ctx context.Context, domain string) (bool, error) {
	qs := []qm.QueryMod{
		models.AdminAuthWebsiteWhere.Domain.EQ(domain),
	}
	return models.AdminAuthWebsites(qs...).Exists(ctx, d.mysqlPool)
}
