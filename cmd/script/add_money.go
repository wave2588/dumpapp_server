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

var emails = []string{
	"likkk@163.com",
	"wangguang3@126.com",
	"495180175@qq.com",
	"1511141687@qq.com",
	"919062031@qq.com",
	"codemakertyq@163.com",
	"1330903352@qq.com",
	"2667371594@qq.com",
	"2111374699@qq.com",
	"273880058@qq.com",
	"1798814463@qq.com",
	"112013513@@qq.com",
	"807984999@qq.com",
	"2511017073@qq.com",
	"83273738@qq.com",
	"6269469@gmail.com",
	"948779023@qq.com",
	"9456258@qq.com",
	"3324554289@qq.com",
	"276880308@qq.com",
	"745690361@qq.com",
	"273880058@qq.com",
	"jinyan19931120@163.com",
	"83273738@qq.com",
	"754075907@qq.com",
	"1511141687@qq.com",
}
