//

package dao

import (
	"context"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// SearchRecordV2DAO ...
type SearchRecordV2DAO interface {
	Insert(ctx context.Context, data *models.SearchRecordV2) error
	Update(ctx context.Context, data *models.SearchRecordV2) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.SearchRecordV2, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.SearchRecordV2, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	BatchGetByIpaIDs(ctx context.Context, filters []qm.QueryMod) ([]*models.SearchRecordV2, error)
	GetOrderBySearchCount(ctx context.Context, offset, limit int, filter []qm.QueryMod) ([]*constant.SearchCount, error)
	CountOrderBySearchCount(ctx context.Context, filter []qm.QueryMod) (int64, error)
}
