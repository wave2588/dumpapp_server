//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberVipOrderDAO ...
type MemberVipOrderDAO interface {
	Insert(ctx context.Context, data *models.MemberVipOrder) error
	Update(ctx context.Context, data *models.MemberVipOrder) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberVipOrder, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberVipOrder, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	BatchGetOrdersByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberVipOrder, error)
}
