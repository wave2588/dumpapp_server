package impl

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberPayCountController struct {
	memberPayCountDAO       dao.MemberPayCountDAO
	memberPayCountRecordDAO dao.MemberPayCountRecordDAO
}

var DefaultMemberPayCountController *MemberPayCountController

func init() {
	DefaultMemberPayCountController = NewMemberPayCountController()
}

func NewMemberPayCountController() *MemberPayCountController {
	return &MemberPayCountController{
		memberPayCountDAO:       impl.DefaultMemberPayCountDAO,
		memberPayCountRecordDAO: impl.DefaultMemberPayCountRecordDAO,
	}
}

func (c *MemberPayCountController) AddCount(ctx context.Context, memberID, count int64, source enum.MemberPayCountSource, recordBizExt datatype.MemberPayCountRecordBizExt) error {
	if count == 0 {
		return nil
	}

	for i := 0; i < int(count); i++ {
		err := c.memberPayCountDAO.Insert(ctx, &models.MemberPayCount{
			MemberID: memberID,
			Status:   enum.MemberPayCountStatusNormal,
			Source:   source,
		})
		if err != nil {
			return err
		}
	}

	return c.memberPayCountRecordDAO.Insert(ctx, &models.MemberPayCountRecord{
		MemberID: memberID,
		Type:     enum.ConvertMemberPayCountSourceToRecordType(source),
		Count:    count,
		BizExt:   recordBizExt,
	})
}

func (c *MemberPayCountController) CheckPayCount(ctx context.Context, loginID, limit int64) error {
	filter := []qm.QueryMod{
		models.MemberPayCountWhere.MemberID.EQ(loginID),
		models.MemberPayCountWhere.Status.EQ(enum.MemberPayCountStatusNormal),
	}
	ids, err := c.memberPayCountDAO.ListIDs(ctx, 0, int(limit), filter, nil)
	if err != nil {
		return err
	}
	if len(ids) < int(limit) {
		return errors2.ErrNotPayCountFunc(fmt.Sprintf("D 币余额不足 %d 个，请充值 D 币。", limit))
	}
	return nil
}

func (c *MemberPayCountController) DeductPayCount(ctx context.Context, loginID, limit int64, status enum.MemberPayCountStatus, use enum.MemberPayCountUse, recordBizExt datatype.MemberPayCountRecordBizExt) error {
	if limit == 0 {
		return nil
	}

	filter := []qm.QueryMod{
		models.MemberPayCountWhere.MemberID.EQ(loginID),
		models.MemberPayCountWhere.Status.EQ(enum.MemberPayCountStatusNormal),
	}
	ids, err := c.memberPayCountDAO.ListIDs(ctx, 0, int(limit), filter, nil)
	if err != nil {
		return err
	}
	if len(ids) < int(limit) {
		return errors2.ErrNotPayCountFunc(fmt.Sprintf("D 币余额不足 %d 个，请充值 D 币。", limit))
	}
	res, err := c.memberPayCountDAO.BatchGet(ctx, ids)
	if err != nil {
		return err
	}
	for _, count := range res {
		count.Status = status
		count.Use = null.StringFrom(use.String())
		err = c.memberPayCountDAO.Update(ctx, count)
		if err != nil {
			return err
		}
	}

	return c.memberPayCountRecordDAO.Insert(ctx, &models.MemberPayCountRecord{
		MemberID: loginID,
		Type:     enum.ConvertMemberPayCountUseToRecordType(use),
		Count:    limit,
		BizExt:   recordBizExt,
	})
}
