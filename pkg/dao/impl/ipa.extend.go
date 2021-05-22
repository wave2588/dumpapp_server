package impl

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *IpaDAO) GetByLikeName(ctx context.Context, name string) ([]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.IpaColumns.ID),
		qm.From("ipa"),
		qm.Where("name like ?", fmt.Sprintf(`%s%%`, name)),
	}

	var data []struct {
		ID int64 `boil:"id"`
	}

	err := models.NewQuery(qs...).Bind(ctx, d.mysqlPool, &data)
	if err != nil {
		return nil, err
	}

	result := make([]int64, 0)
	for _, r := range data {
		result = append(result, r.ID)
	}
	return result, nil
}

func (d *IpaDAO) CountByLikeName(ctx context.Context, name string) (int64, error) {
	qs := []qm.QueryMod{
		qm.Select("count(*) as count"),
		qm.From("ipa"),
		qm.Where("name like ?", fmt.Sprintf(`%%%s%%`, name)),
	}

	var data struct {
		Count int64 `boil:"count"`
	}
	err := models.NewQuery(qs...).Bind(ctx, d.mysqlPool, &data)
	if err != nil {
		return 0, err
	}
	return data.Count, nil
}
