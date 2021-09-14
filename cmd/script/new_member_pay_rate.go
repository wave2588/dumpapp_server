package main

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func main() {

	ctx := context.Background()

	for i := 0; i <= 100; i++ {
		offset := 0
		bulkSize := 100
		hasNext := true

		n := time.Now()
		now := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
		startAt := now.AddDate(0, 0, -i)
		endAt := now.AddDate(0, 0, -i+1)

		memberIDs := make([]int64, 0)
		for hasNext {
			filters := []qm.QueryMod{
				models.AccountWhere.CreatedAt.GT(startAt),
				models.AccountWhere.CreatedAt.LT(endAt),
			}
			ids, err := impl.DefaultAccountDAO.ListIDs(ctx, offset, bulkSize, filters, nil)
			util.PanicIf(err)

			hasNext = len(ids) >= bulkSize
			offset += len(ids)
			memberIDs = append(memberIDs, ids...)
		}
		memberOrderMap, err := impl.DefaultMemberDownloadOrderDAO.BatchGetByMemberIDs(ctx, memberIDs)
		util.PanicIf(err)
		paidMember := make(map[int64]struct{})
		for _, orders := range memberOrderMap {
			for _, order := range orders {
				if order.Status == enum.MemberDownloadOrderStatusPaid {
					paidMember[order.MemberID] = struct{}{}
				}
			}
		}

		fmt.Println(fmt.Sprintf("%v 新用户付费率:  %.2f%%", startAt, cast.ToFloat64(len(paidMember))/cast.ToFloat64(len(memberIDs))*100))
	}
}
