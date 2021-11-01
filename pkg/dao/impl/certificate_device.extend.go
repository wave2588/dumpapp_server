package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *CertificateDeviceDAO) GetByDeviceID(ctx context.Context, deviceID int64) ([]*models.CertificateDevice, error) {
	result, err := d.BatchGetByDeviceIDs(ctx, []int64{deviceID})
	if err != nil {
		return nil, err
	}
	return result[deviceID], nil
}

func (d *CertificateDeviceDAO) BatchGetByDeviceIDs(ctx context.Context, deviceIDs []int64) (map[int64][]*models.CertificateDevice, error) {
	qs := []qm.QueryMod{
		models.CertificateDeviceWhere.DeviceID.IN(deviceIDs),
	}
	data, err := models.CertificateDevices(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	result := make(map[int64][]*models.CertificateDevice)
	for _, datum := range data {
		result[datum.DeviceID] = append(result[datum.DeviceID], datum)
	}
	return result, nil
}
