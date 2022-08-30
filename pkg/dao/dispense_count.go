//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// DispenseCountDAO ...
type DispenseCountDAO interface {
	Insert(ctx context.Context, data *models.DispenseCount) error
	Update(ctx context.Context, data *models.DispenseCount) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.DispenseCount, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.DispenseCount, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	BatchGetMemberNormalCount(ctx context.Context, memberIDs []int64) (map[int64]int64, error)
}
