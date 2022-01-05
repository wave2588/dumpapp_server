package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *IpaSignDAO) GetByStatus(ctx context.Context, status enum.IpaSignStatus) ([]*models.IpaSign, error) {
	qs := []qm.QueryMod{
		models.IpaSignWhere.Status.EQ(status),
	}
	return models.IpaSigns(qs...).All(ctx, d.mysqlPool)
}
