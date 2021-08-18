//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// IpaVersionDAO ...
type IpaVersionDAO interface {
	Insert(ctx context.Context, data *models.IpaVersion) error
	Update(ctx context.Context, data *models.IpaVersion) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.IpaVersion, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.IpaVersion, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByIpaIDVersion retrieves a single record by uniq key ipaID, version from db.
	GetByIpaIDVersion(ctx context.Context, ipaID int64, version string) (*models.IpaVersion, error)
	// GetIpaVersionSliceByIpaID retrieves a slice of records by first field of uniq key [ipaID] with an executor.
	GetIpaVersionSliceByIpaID(ctx context.Context, ipaID int64) ([]*models.IpaVersion, error)
	BatchGetIpaVersions(ctx context.Context, ipaIDs []int64) (map[int64][]*models.IpaVersion, error)
}
