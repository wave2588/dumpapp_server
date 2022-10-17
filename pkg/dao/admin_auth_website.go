//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AdminAuthWebsiteDAO ...
type AdminAuthWebsiteDAO interface {
	Insert(ctx context.Context, data *models.AdminAuthWebsite) error
	Update(ctx context.Context, data *models.AdminAuthWebsite) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.AdminAuthWebsite, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.AdminAuthWebsite, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByDomain retrieves a single record by uniq key domain from db.
	GetByDomain(ctx context.Context, domain string) (*models.AdminAuthWebsite, error)
	GetByDomainSafe(ctx context.Context, domain string) (*models.AdminAuthWebsite, error)
	// BatchGetByDomain retrieves multiple records by uniq key domain from db.
	BatchGetByDomain(ctx context.Context, domains []string) (map[string]*models.AdminAuthWebsite, error)
}
