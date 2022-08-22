package impl

import (
	"context"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberSignIpaDownloadCountController struct {
	memberSignIpaDownloadCountDAO       dao.MemberSignIpaDownloadCountDAO
	memberSignIpaDownloadCountRecordDAO dao.MemberSignIpaDownloadCountRecordDAO
}

var DefaultMemberSignIpaDownloadCountController *MemberSignIpaDownloadCountController

func init() {
	DefaultMemberSignIpaDownloadCountController = NewMemberSignIpaDownloadCountController()
}

func NewMemberSignIpaDownloadCountController() *MemberSignIpaDownloadCountController {
	return &MemberSignIpaDownloadCountController{
		memberSignIpaDownloadCountDAO:       impl.DefaultMemberSignIpaDownloadCountDAO,
		memberSignIpaDownloadCountRecordDAO: impl.DefaultMemberSignIpaDownloadCountRecordDAO,
	}
}

func (c *MemberSignIpaDownloadCountController) CheckCount(ctx context.Context, memberID, limit int64) (bool, error) {
	filter := []qm.QueryMod{
		models.MemberSignIpaDownloadCountWhere.MemberID.EQ(memberID),
		models.MemberSignIpaDownloadCountWhere.Status.EQ(enum.MemberSignIpaDownloadCountStatusNormal),
	}
	count, err := c.memberSignIpaDownloadCountDAO.Count(ctx, filter)
	if err != nil {
		return false, err
	}
	if count < limit {
		return false, nil
	}
	return true, nil
}

func (c *MemberSignIpaDownloadCountController) AddCount(ctx context.Context, memberID, count int64, recordType enum.MemberSignIpaDownloadCountRecordType) error {
	if count == 0 {
		return nil
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	/// 写入次数
	for i := 0; i < int(count); i++ {
		err := c.memberSignIpaDownloadCountDAO.Insert(ctx, &models.MemberSignIpaDownloadCount{
			MemberID: memberID,
			Status:   enum.MemberSignIpaDownloadCountStatusNormal,
		})
		if err != nil {
			return err
		}
	}

	/// 写入记录
	err := c.memberSignIpaDownloadCountRecordDAO.Insert(ctx, &models.MemberSignIpaDownloadCountRecord{
		MemberID: memberID,
		Type:     recordType,
		Count:    count,
	})
	if err != nil {
		return err
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return nil
}

func (c *MemberSignIpaDownloadCountController) DeductPayCount(ctx context.Context, memberID, count int64, recordType enum.MemberSignIpaDownloadCountRecordType) error {
	if count == 0 {
		return nil
	}

	filter := []qm.QueryMod{
		models.MemberSignIpaDownloadCountWhere.MemberID.EQ(memberID),
		models.MemberSignIpaDownloadCountWhere.Status.EQ(enum.MemberSignIpaDownloadCountStatusNormal),
	}
	ids, err := c.memberSignIpaDownloadCountDAO.ListIDs(ctx, 0, int(count), filter, nil)
	if err != nil {
		return err
	}
	if len(ids) < int(count) {
		return errors2.ErrNotMemberSignIpaDownloadCount
	}

	counts, err := c.memberSignIpaDownloadCountDAO.BatchGet(ctx, ids)
	if err != nil {
		return err
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	/// 修改次数
	for _, downloadCount := range counts {
		downloadCount.Status = enum.MemberSignIpaDownloadCountStatusUsed
		err = c.memberSignIpaDownloadCountDAO.Update(ctx, downloadCount)
		if err != nil {
			return err
		}
	}

	/// 写入记录
	err = c.memberSignIpaDownloadCountRecordDAO.Insert(ctx, &models.MemberSignIpaDownloadCountRecord{
		MemberID: memberID,
		Type:     recordType,
		Count:    count,
	})
	if err != nil {
		return err
	}

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	return nil
}
