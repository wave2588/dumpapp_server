package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *AccountDAO) BatchGetByPhones(ctx context.Context, phones []string) (map[string]*models.Account, error) {
	qs := []qm.QueryMod{
		models.AccountWhere.Phone.IN(phones),
	}
	data, err := models.Accounts(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[string]*models.Account)
	for _, datum := range data {
		res[datum.Phone] = datum
	}
	return res, nil
}
