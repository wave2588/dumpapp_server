//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// InstallAppCdkeyOrderDAO ...
type InstallAppCdkeyOrderDAO interface {
	Insert(ctx context.Context, data *models.InstallAppCdkeyOrder) error
	Update(ctx context.Context, data *models.InstallAppCdkeyOrder) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.InstallAppCdkeyOrder, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.InstallAppCdkeyOrder, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	BatchGetByContact(ctx context.Context, contacts []string) (map[string][]*models.InstallAppCdkeyOrder, error)
}
