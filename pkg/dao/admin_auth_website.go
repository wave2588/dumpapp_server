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
	// GetByMemberIDDomain retrieves a single record by uniq key memberID, domain from db.
	GetByMemberIDDomain(ctx context.Context, memberID int64, domain string) (*models.AdminAuthWebsite, error)
	// GetAdminAuthWebsiteSliceByMemberID retrieves a slice of records by first field of uniq key [memberID] with an executor.
	GetAdminAuthWebsiteSliceByMemberID(ctx context.Context, memberID int64) ([]*models.AdminAuthWebsite, error)
	IsExistDomain(ctx context.Context, domain string) (bool, error)
}
