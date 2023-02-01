package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *CertificateDeviceDAO) GetLastByDeviceID(ctx context.Context, deviceID int64) (*models.CertificateDevice, error) {
	ids, err := d.ListIDs(ctx, 0, 1, []qm.QueryMod{models.CertificateDeviceWhere.DeviceID.EQ(deviceID)}, nil)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, nil
	}
	id := ids[0]

	data, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	return data[id], nil
}
