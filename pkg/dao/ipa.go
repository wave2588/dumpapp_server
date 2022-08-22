//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// IpaDAO ...
type IpaDAO interface {
	Insert(ctx context.Context, data *models.Ipa) error
	Update(ctx context.Context, data *models.Ipa) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.Ipa, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.Ipa, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	GetByLikeName(ctx context.Context, name string) ([]int64, error)
	CountByLikeName(ctx context.Context, name string) (int64, error)
}
