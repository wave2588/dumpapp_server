package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *CertificateV2DAO) ListIDsByDeviceIDs(ctx context.Context, deviceIDs []int64) (map[int64][]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.CertificateV2Columns.ID, models.CertificateV2Columns.DeviceID),
		qm.From("certificate_v2"),
		models.CertificateV2Where.DeviceID.IN(deviceIDs),
	}

	var data []struct {
		ID       int64 `boil:"id"`
		DeviceID int64 `boil:"device_id"`
	}
	err := models.NewQuery(qs...).Bind(ctx, d.mysqlPool, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[int64][]int64)
	for _, datum := range data {
		result[datum.DeviceID] = append(result[datum.DeviceID], datum.ID)
	}
	return result, nil
}
