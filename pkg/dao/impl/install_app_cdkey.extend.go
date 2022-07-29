package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *InstallAppCdkeyDAO) BatchGetByCertificateIDs(ctx context.Context, certificateIDs []int64) (map[int64][]*models.InstallAppCdkey, error) {
	qs := []qm.QueryMod{
		models.InstallAppCdkeyWhere.CertificateID.IN(certificateIDs),
	}

	data, err := models.InstallAppCdkeys(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}

	res := make(map[int64][]*models.InstallAppCdkey)
	for _, datum := range data {
		res[datum.CertificateID] = append(res[datum.CertificateID], datum)
	}
	return res, nil
}
