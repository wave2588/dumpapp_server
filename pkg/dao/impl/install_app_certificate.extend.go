package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *InstallAppCertificateDAO) BatchGetByUDIDs(ctx context.Context, UDIDs []string) (map[string][]*models.InstallAppCertificate, error) {
	qs := []qm.QueryMod{
		models.InstallAppCertificateWhere.Udid.IN(UDIDs),
	}
	data, err := models.InstallAppCertificates(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]*models.InstallAppCertificate)
	for _, datum := range data {
		result[datum.Udid] = append(result[datum.Udid], datum)
	}
	return result, nil
}
