//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberInviteCodeDAO ...
type MemberInviteCodeDAO interface {
	Insert(ctx context.Context, data *models.MemberInviteCode) error
	Update(ctx context.Context, data *models.MemberInviteCode) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberInviteCode, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberInviteCode, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByMemberID retrieves a single record by uniq key memberID from db.
	GetByMemberID(ctx context.Context, memberID int64) (*models.MemberInviteCode, error)
	// BatchGetByMemberID retrieves multiple records by uniq key memberID from db.
	BatchGetByMemberID(ctx context.Context, memberIDs []int64) (map[int64]*models.MemberInviteCode, error)
	// GetByCode retrieves a single record by uniq key code from db.
	GetByCode(ctx context.Context, code string) (*models.MemberInviteCode, error)
	// BatchGetByCode retrieves multiple records by uniq key code from db.
	BatchGetByCode(ctx context.Context, codes []string) (map[string]*models.MemberInviteCode, error)
}
