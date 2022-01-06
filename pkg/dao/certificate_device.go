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
	// GetByDeviceIDCertificateID retrieves a single record by uniq key deviceID, certificateID from db.
	GetByDeviceIDCertificateID(ctx context.Context, deviceID int64, certificateID int64) (*models.CertificateDevice, error)
	// GetCertificateDeviceSliceByDeviceID retrieves a slice of records by first field of uniq key [deviceID] with an executor.
	GetCertificateDeviceSliceByDeviceID(ctx context.Context, deviceID int64) ([]*models.CertificateDevice, error)
	GetByDeviceID(ctx context.Context, deviceID int64) ([]*models.CertificateDevice, error)
	BatchGetByDeviceIDs(ctx context.Context, deviceIDs []int64) (map[int64][]*models.CertificateDevice, error)
	GetByCertificateID(ctx context.Context, certificateID int64) ([]*models.CertificateDevice, error)
}
