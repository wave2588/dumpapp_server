//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AppSourceDAO ...
type AppSourceDAO interface {
	Insert(ctx context.Context, data *models.AppSource) error
	Update(ctx context.Context, data *models.AppSource) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.AppSource, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.AppSource, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByURL retrieves a single record by uniq key uRL from db.
	GetByURL(ctx context.Context, uRL string) (*models.AppSource, error)
	// BatchGetByURL retrieves multiple records by uniq key uRL from db.
	BatchGetByURL(ctx context.Context, uRLs []string) (map[string]*models.AppSource, error)
}
