package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

/// 拉黑刷邀请注册的用户
func main() {

	ctx := context.Background()

	offset := 0
	limit := 100
	hasNext := true

	var id int64 = 1520621090479542272

	acc, err := impl.DefaultAccountDAO.Get(ctx, id)
	util.PanicIf(err)

	acc.Status = 2
	util.PanicIf(impl.DefaultAccountDAO.Update(ctx, acc))

	for hasNext {

		ids, err := impl.DefaultMemberInviteDAO.ListIDs(ctx, offset, limit, []qm.QueryMod{models.MemberInviteWhere.InviterID.EQ(acc.ID)}, nil)
		util.PanicIf(err)

		offset += len(ids)
		hasNext = len(ids) == limit

		invites, err := impl.DefaultMemberInviteDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		memberIDs := make([]int64, 0)
		for _, invite := range invites {
			memberIDs = append(memberIDs, invite.InviteeID)
		}

		accounts, err := impl.DefaultAccountDAO.BatchGet(ctx, memberIDs)
		util.PanicIf(err)

		for _, account := range accounts {
			account.Status = 2
			util.PanicIf(impl.DefaultAccountDAO.Update(ctx, account))
		}
	}
}
