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

type MemberDownloadController struct {
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO
}

var DefaultMemberDownloadController *MemberDownloadController

func init() {
	DefaultMemberDownloadController = NewMemberDownloadController()
}

func NewMemberDownloadController() *MemberDownloadController {
	return &MemberDownloadController{
		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,
	}
}

func (c *MemberDownloadController) GetDownloadNumber(ctx context.Context, loginID int64) (*models.MemberDownloadNumber, error) {
	filter := []qm.QueryMod{
		models.MemberDownloadNumberWhere.MemberID.EQ(loginID),
		models.MemberDownloadNumberWhere.Status.EQ(enum.MemberDownloadNumberStatusNormal),
	}
	ids, err := c.memberDownloadNumberDAO.ListIDs(ctx, 0, 1, filter, nil)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors2.ErrNotDownloadNumber
	}
	return c.memberDownloadNumberDAO.Get(ctx, ids[0])
}
