//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AdminDumpOrderDAO ...
type AdminDumpOrderDAO interface {
	Insert(ctx context.Context, data *models.AdminDumpOrder) error
	Update(ctx context.Context, data *models.AdminDumpOrder) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.AdminDumpOrder, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.AdminDumpOrder, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByIpaIDIpaVersion retrieves a single record by uniq key ipaID, ipaVersion from db.
	GetByIpaIDIpaVersion(ctx context.Context, ipaID int64, ipaVersion string) (*models.AdminDumpOrder, error)
	// GetAdminDumpOrderSliceByIpaID retrieves a slice of records by first field of uniq key [ipaID] with an executor.
	GetAdminDumpOrderSliceByIpaID(ctx context.Context, ipaID int64) ([]*models.AdminDumpOrder, error)
}
