//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// InstallAppCdkeyDAO ...
type InstallAppCdkeyDAO interface {
	Insert(ctx context.Context, data *models.InstallAppCdkey) error
	Update(ctx context.Context, data *models.InstallAppCdkey) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.InstallAppCdkey, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.InstallAppCdkey, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByOutID retrieves a single record by uniq key outID from db.
	GetByOutID(ctx context.Context, outID string) (*models.InstallAppCdkey, error)
	// BatchGetByOutID retrieves multiple records by uniq key outID from db.
	BatchGetByOutID(ctx context.Context, outIDs []string) (map[string]*models.InstallAppCdkey, error)
	BatchGetByCertificateIDs(ctx context.Context, certificateIDs []int64) (map[int64][]*models.InstallAppCdkey, error)
}
