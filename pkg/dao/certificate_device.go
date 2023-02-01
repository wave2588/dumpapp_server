//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CertificateDeviceDAO ...
type CertificateDeviceDAO interface {
	Insert(ctx context.Context, data *models.CertificateDevice) error
	Update(ctx context.Context, data *models.CertificateDevice) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.CertificateDevice, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.CertificateDevice, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByCertificateID retrieves a single record by uniq key certificateID from db.
	GetByCertificateID(ctx context.Context, certificateID int64) (*models.CertificateDevice, error)
	// BatchGetByCertificateID retrieves multiple records by uniq key certificateID from db.
	BatchGetByCertificateID(ctx context.Context, certificateIDs []int64) (map[int64]*models.CertificateDevice, error)
	GetLastByDeviceID(ctx context.Context, deviceID int64) (*models.CertificateDevice, error)
}
