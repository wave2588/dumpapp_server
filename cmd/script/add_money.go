package main

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"fmt"
)

func main() {
	ctx := context.Background()

	accountMap, err := impl.DefaultAccountDAO.BatchGetByEmail(ctx, emails)
	util.PanicIf(err)

	for _, email := range emails {
		account, ok := accountMap[email]
		if !ok {
			fmt.Println("not found", email)
			continue
		}
		for i := 0; i < 9; i++ {
			id := util2.MustGenerateID(ctx)
			util.PanicIf(impl.DefaultMemberPayCountDAO.Insert(ctx, &models.MemberPayCount{
				ID:       id,
				MemberID: account.ID,
				Status:   enum.MemberPayCountStatusNormal,
				Source:   enum.MemberPayCountSourceAdminPresented,
			}))
		}
	}
}

var emails = []string{}
