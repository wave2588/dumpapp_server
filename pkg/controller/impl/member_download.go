package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemberDownloadController struct {
	memberPayCountDAO dao.MemberPayCountDAO
}

var DefaultMemberDownloadController *MemberDownloadController

func init() {
	DefaultMemberDownloadController = NewMemberDownloadController()
}

func NewMemberDownloadController() *MemberDownloadController {
	return &MemberDownloadController{
		memberPayCountDAO: impl.DefaultMemberPayCountDAO,
	}
}

func (c *MemberDownloadController) CheckPayCount(ctx context.Context, loginID, limit int64) error {
	filter := []qm.QueryMod{
		models.MemberPayCountWhere.MemberID.EQ(loginID),
		models.MemberPayCountWhere.Status.EQ(enum.MemberPayCountStatusNormal),
	}
	ids, err := c.memberPayCountDAO.ListIDs(ctx, 0, int(limit), filter, nil)
	if err != nil {
		return err
	}
	if len(ids) < int(limit) {
		return errors2.ErrNotPayCount
	}
	return nil
}

func (c *MemberDownloadController) DeductPayCount(ctx context.Context, loginID, limit int64, use enum.MemberPayCountUse) error {
	filter := []qm.QueryMod{
		models.MemberPayCountWhere.MemberID.EQ(loginID),
		models.MemberPayCountWhere.Status.EQ(enum.MemberPayCountStatusNormal),
	}
	ids, err := c.memberPayCountDAO.ListIDs(ctx, 0, int(limit), filter, nil)
	if err != nil {
		return err
	}
	if len(ids) < int(limit) {
		return errors2.ErrNotPayCount
	}
	res, err := c.memberPayCountDAO.BatchGet(ctx, ids)
	if err != nil {
		return err
	}
	for _, count := range res {
		count.Status = enum.MemberPayCountStatusUsed
		count.Use = null.StringFrom(use.String())
		err = c.memberPayCountDAO.Update(ctx, count)
		if err != nil {
			return err
		}
	}
	return nil
}
