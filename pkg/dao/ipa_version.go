//

package dao

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// IpaVersionDAO ...
type IpaVersionDAO interface {
	Insert(ctx context.Context, data *models.IpaVersion) error
	Update(ctx context.Context, data *models.IpaVersion) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.IpaVersion, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.IpaVersion, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByTokenPath retrieves a single record by uniq key tokenPath from db.
	GetByTokenPath(ctx context.Context, tokenPath string) (*models.IpaVersion, error)
	// BatchGetByTokenPath retrieves multiple records by uniq key tokenPath from db.
	BatchGetByTokenPath(ctx context.Context, tokenPaths []string) (map[string]*models.IpaVersion, error)
	BatchGetIpaVersions(ctx context.Context, ipaIDs []int64) (map[int64][]*models.IpaVersion, error)
	GetByIpaIDAndIpaType(ctx context.Context, ipaID int64, ipaType enum.IpaType) ([]*models.IpaVersion, error)
	GetByIpaIDAndIpaTypeAndVersion(ctx context.Context, ipaID int64, ipaType enum.IpaType, version string) ([]*models.IpaVersion, error)
	GetByIpaID(ctx context.Context, ipaID int64) ([]*models.IpaVersion, error)
	GetByIpaType(ctx context.Context, ipaType enum.IpaType) ([]*models.IpaVersion, error)
}
