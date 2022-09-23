package impl

import (
	"context"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/dao/models"
)

func (d *AdminConfigInfoDAO) GetConfig(ctx context.Context) (*models.AdminConfigInfo, error) {
	id := int64(1)
	data, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	res, ok := data[id]
	if !ok {
		return nil, errors2.ErrNotFound
	}
	return res, nil
}
