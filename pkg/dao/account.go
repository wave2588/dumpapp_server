//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AccountDAO ...
type AccountDAO interface {
	Insert(ctx context.Context, data *models.Account) error
	Update(ctx context.Context, data *models.Account) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.Account, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.Account, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByEmail retrieves a single record by uniq key email from db.
	GetByEmail(ctx context.Context, email string) (*models.Account, error)
	// BatchGetByEmail retrieves multiple records by uniq key email from db.
	BatchGetByEmail(ctx context.Context, emails []string) (map[string]*models.Account, error)
	// GetByPhone retrieves a single record by uniq key phone from db.
	GetByPhone(ctx context.Context, phone string) (*models.Account, error)
	// BatchGetByPhone retrieves multiple records by uniq key phone from db.
	BatchGetByPhone(ctx context.Context, phones []string) (map[string]*models.Account, error)
	BatchGetByPhones(ctx context.Context, phones []string) (map[string]*models.Account, error)
}
