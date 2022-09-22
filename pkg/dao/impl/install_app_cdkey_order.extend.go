package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *InstallAppCdkeyOrderDAO) BatchGetByContact(ctx context.Context, contacts []string) (map[string][]*models.InstallAppCdkeyOrder, error) {
	qs := []qm.QueryMod{
		models.InstallAppCdkeyOrderWhere.Contact.IN(contacts),
	}

	data, err := models.InstallAppCdkeyOrders(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}

	res := make(map[string][]*models.InstallAppCdkeyOrder)
	for _, datum := range data {
		res[datum.Contact] = append(res[datum.Contact], datum)
	}
	return res, nil
}
