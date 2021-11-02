//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberDownloadOrderDAO ...
type MemberDownloadOrderDAO interface {
	Insert(ctx context.Context, data *models.MemberDownloadOrder) error
	Update(ctx context.Context, data *models.MemberDownloadOrder) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberDownloadOrder, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberDownloadOrder, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberDownloadOrder, error)
	GetMemberIDsOrderByPaidCount(ctx context.Context, offset, limit int, filter []qm.QueryMod) ([]int64, error)
	CountMemberIDsOrderByPaidCount(ctx context.Context, filter []qm.QueryMod) (int64, error)
	GetByFilters(ctx context.Context, filters []qm.QueryMod, orderBys []string) ([]*models.MemberDownloadOrder, error)
}
