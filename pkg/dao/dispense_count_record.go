//

package dao

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// DispenseCountRecordDAO ...
type DispenseCountRecordDAO interface {
	Insert(ctx context.Context, data *models.DispenseCountRecord) error
	Update(ctx context.Context, data *models.DispenseCountRecord) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.DispenseCountRecord, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.DispenseCountRecord, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	BatchGetByObjectIDsAndRecordType(ctx context.Context, objectIDs []int64, recordType enum.DispenseCountRecordType) (map[int64][]*models.DispenseCountRecord, error)
}
