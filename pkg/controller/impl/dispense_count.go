package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DispenseCountController struct {
	dispenseCountDAO       dao.DispenseCountDAO
	dispenseCountRecordDAO dao.DispenseCountRecordDAO
}

var DefaultDispenseCountController *DispenseCountController

func init() {
	DefaultDispenseCountController = NewDispenseCountController()
}

func NewDispenseCountController() *DispenseCountController {
	return &DispenseCountController{
		dispenseCountDAO:       impl.DefaultDispenseCountDAO,
		dispenseCountRecordDAO: impl.DefaultDispenseCountRecordDAO,
	}
}

func (c *DispenseCountController) AddCount(ctx context.Context, memberID, count int64) error {
	if count <= 0 {
		return nil
	}

	for i := 0; i < int(count); i++ {
		err := c.dispenseCountDAO.Insert(ctx, &models.DispenseCount{
			MemberID: memberID,
			Status:   enum.DispenseCountStatusNormal,
		})
		if err != nil {
			return err
		}
	}

	/// 添加记录
	return c.dispenseCountRecordDAO.Insert(ctx, &models.DispenseCountRecord{
		MemberID: memberID,
		Type:     enum.DispenseCountRecordTypePay,
		Count:    count,
	})
}

func (c *DispenseCountController) Check(ctx context.Context, memberID, count int64) error {
	filter := []qm.QueryMod{
		models.DispenseCountWhere.MemberID.EQ(memberID),
		models.DispenseCountWhere.Status.EQ(enum.DispenseCountStatusNormal),
	}
	ids, err := c.dispenseCountDAO.ListIDs(ctx, 0, int(count), filter, nil)
	if err != nil {
		return err
	}
	if len(ids) < int(count) {
		return errors2.ErrDispenseCountFunc("下载次数不足, 请联系签名作者充值")
	}
	return nil
}

func (c *DispenseCountController) DeductCount(ctx context.Context, memberID, count int64, recordType enum.DispenseCountRecordType) error {
	if count <= 0 {
		return nil
	}

	filter := []qm.QueryMod{
		models.DispenseCountWhere.MemberID.EQ(memberID),
		models.DispenseCountWhere.Status.EQ(enum.DispenseCountStatusNormal),
	}
	ids, err := c.dispenseCountDAO.ListIDs(ctx, 0, int(count), filter, nil)
	if err != nil {
		return err
	}
	if len(ids) < int(count) {
		return errors2.ErrDispenseCountFunc("下载次数不足, 请联系签名作者充值")
	}

	dispenseCountMap, err := c.dispenseCountDAO.BatchGet(ctx, ids)
	if err != nil {
		return err
	}

	for _, dispenseCount := range dispenseCountMap {
		dispenseCount.Status = enum.DispenseCountStatusUsed
		err = c.dispenseCountDAO.Update(ctx, dispenseCount)
		if err != nil {
			return err
		}
	}

	/// 添加记录
	return c.dispenseCountRecordDAO.Insert(ctx, &models.DispenseCountRecord{
		MemberID: memberID,
		Type:     recordType,
		Count:    count,
	})
}
