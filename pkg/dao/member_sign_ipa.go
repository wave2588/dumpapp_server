//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberSignIpaDAO ...
type MemberSignIpaDAO interface {
	Insert(ctx context.Context, data *models.MemberSignIpa) error
	Update(ctx context.Context, data *models.MemberSignIpa) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberSignIpa, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberSignIpa, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByExpenseID retrieves a single record by uniq key expenseID from db.
	GetByExpenseID(ctx context.Context, expenseID string) (*models.MemberSignIpa, error)
	// BatchGetByExpenseID retrieves multiple records by uniq key expenseID from db.
	BatchGetByExpenseID(ctx context.Context, expenseIDs []string) (map[string]*models.MemberSignIpa, error)
	BatchGetByIpaPlistFileToken(ctx context.Context, plistFileTokens []string) (map[string]*models.MemberSignIpa, error)
	BatchGetByIpaFileToken(ctx context.Context, ipaFileTokens []string) (map[string]*models.MemberSignIpa, error)
}
