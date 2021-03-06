//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberRebateRecordDAO ...
type MemberRebateRecordDAO interface {
	Insert(ctx context.Context, data *models.MemberRebateRecord) error
	Update(ctx context.Context, data *models.MemberRebateRecord) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberRebateRecord, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberRebateRecord, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByOrderID retrieves a single record by uniq key orderID from db.
	GetByOrderID(ctx context.Context, orderID int64) (*models.MemberRebateRecord, error)
	// BatchGetByOrderID retrieves multiple records by uniq key orderID from db.
	BatchGetByOrderID(ctx context.Context, orderIDs []int64) (map[int64]*models.MemberRebateRecord, error)
}
