//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CertificateDAO ...
type CertificateDAO interface {
	Insert(ctx context.Context, data *models.Certificate) error
	Update(ctx context.Context, data *models.Certificate) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.Certificate, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.Certificate, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByP12FileDateMD5MobileProvisionFileDataMD5 retrieves a single record by uniq key p12FileDateMD5, mobileProvisionFileDataMD5 from db.
	GetByP12FileDateMD5MobileProvisionFileDataMD5(ctx context.Context, p12FileDateMD5 string, mobileProvisionFileDataMD5 string) (*models.Certificate, error)
	// GetCertificateSliceByP12FileDateMD5 retrieves a slice of records by first field of uniq key [p12FileDateMD5] with an executor.
	GetCertificateSliceByP12FileDateMD5(ctx context.Context, p12FileDateMD5 string) ([]*models.Certificate, error)
}
