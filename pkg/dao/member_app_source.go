//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberAppSourceDAO ...
type MemberAppSourceDAO interface {
	Insert(ctx context.Context, data *models.MemberAppSource) error
	Update(ctx context.Context, data *models.MemberAppSource) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberAppSource, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberAppSource, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByMemberIDAppSourceID retrieves a single record by uniq key memberID, appSourceID from db.
	GetByMemberIDAppSourceID(ctx context.Context, memberID int64, appSourceID int64) (*models.MemberAppSource, error)
	// GetMemberAppSourceSliceByMemberID retrieves a slice of records by first field of uniq key [memberID] with an executor.
	GetMemberAppSourceSliceByMemberID(ctx context.Context, memberID int64) ([]*models.MemberAppSource, error)
}
