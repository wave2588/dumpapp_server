//

package dao

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MemberDownloadIpaRecordDAO ...
type MemberDownloadIpaRecordDAO interface {
	Insert(ctx context.Context, data *models.MemberDownloadIpaRecord) error
	Update(ctx context.Context, data *models.MemberDownloadIpaRecord) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.MemberDownloadIpaRecord, error)
	// BatchGet retrieves multiple records by primary key from db.
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberDownloadIpaRecord, error)
	// 后台和脚本使用：倒序列出所有
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)
	// GetByMemberIDIpaIDIpaTypeVersion retrieves a single record by uniq key memberID, ipaID, ipaType, version from db.
	GetByMemberIDIpaIDIpaTypeVersion(ctx context.Context, memberID int64, ipaID null.Int64, ipaType null.String, version null.String) (*models.MemberDownloadIpaRecord, error)
	// GetMemberDownloadIpaRecordSliceByMemberID retrieves a slice of records by first field of uniq key [memberID] with an executor.
	GetMemberDownloadIpaRecordSliceByMemberID(ctx context.Context, memberID int64) ([]*models.MemberDownloadIpaRecord, error)
	BatchGetMemberNormalCount(ctx context.Context, memberIDs []int64) (map[int64]int64, error)
	GetIpaDownloadCount(ctx context.Context, ipaID int64) (int64, int64, error)
	BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberDownloadIpaRecord, error)
	GetByMemberIDAndIpaID(ctx context.Context, memberID, ipaID int64) ([]*models.MemberDownloadIpaRecord, error)
}
