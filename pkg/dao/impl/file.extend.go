package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *FileDAO) BatchGetByTokens(ctx context.Context, tokens []string) (map[string]*models.File, error) {
	qs := []qm.QueryMod{
		models.FileWhere.Token.IN(tokens),
	}
	data, err := models.Files(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[string]*models.File)
	for _, datum := range data {
		res[datum.Token] = datum
	}
	return res, nil
}
