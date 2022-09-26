package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberSignIpaDAO) BatchGetByIpaPlistFileToken(ctx context.Context, plistFileTokens []string) (map[string]*models.MemberSignIpa, error) {
	qs := []qm.QueryMod{
		models.MemberSignIpaWhere.IpaPlistFileToken.IN(plistFileTokens),
	}
	data, err := models.MemberSignIpas(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[string]*models.MemberSignIpa)
	for _, datum := range data {
		res[datum.IpaPlistFileToken] = datum
	}
	return res, nil
}

func (d *MemberSignIpaDAO) BatchGetByIpaFileToken(ctx context.Context, ipaFileTokens []string) (map[string]*models.MemberSignIpa, error) {
	qs := []qm.QueryMod{
		models.MemberSignIpaWhere.IpaFileToken.IN(ipaFileTokens),
	}
	data, err := models.MemberSignIpas(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	res := make(map[string]*models.MemberSignIpa)
	for _, datum := range data {
		res[datum.IpaFileToken] = datum
	}
	return res, nil
}
