package impl

import (
	"context"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/dao/models"
	pkgErr "github.com/pkg/errors"
)

func (d *AdminAuthWebsiteDAO) GetByDomainSafe(ctx context.Context, domain string) (*models.AdminAuthWebsite, error) {
	res, err := d.GetByDomain(ctx, domain)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		return nil, err
	}
	return res, nil
}
