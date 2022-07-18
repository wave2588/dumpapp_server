//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberDeviceDAO ...
type MemberDeviceDAO interface {
	Insert(ctx context.Context, data *models.MemberDevice) error
	Update(ctx context.Context, data *models.MemberDevice) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberDevice, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberDevice, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByMemberIDUdid retrieves a single record by uniq key memberID, udid from db.
	GetByMemberIDUdid(ctx context.Context, memberID int64, udid string) (*models.MemberDevice, error)
	// GetMemberDeviceSliceByMemberID retrieves a slice of records by first field of uniq key [memberID] with an executor.
	GetMemberDeviceSliceByMemberID(ctx context.Context, memberID int64) ([]*models.MemberDevice, error)
	BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberDevice, error)
	GetByMemberIDUdidSafe(ctx context.Context, memberID int64, udid string) (*models.MemberDevice, error)
}
