package main

import (
	"context"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"fmt"
)

func main() {
	run()
}

func run() {
	ctx := context.Background()

	offset := 0
	bulkSize := 100
	hasNext := true

	for hasNext {
		fmt.Println(fmt.Sprintf("offset: %d...", offset))

		memberIDs, err := impl.DefaultMemberVipDAO.ListIDs(ctx, offset, bulkSize, nil, nil)
		util.PanicIf(err)

		memberVipMap, err := impl.DefaultMemberVipDAO.BatchGet(ctx, memberIDs)
		util.PanicIf(err)

		hasNext = len(memberIDs) >= bulkSize
		offset += len(memberIDs)

		fmt.Println(memberIDs)

		memberOrderMap, err := impl.DefaultMemberVipOrderDAO.BatchGetOrdersByMemberIDs(ctx, memberIDs)
		util.PanicIf(err)

		for _, memberID := range memberIDs {
			memberVip, ok := memberVipMap[memberID]
			if !ok {
				fmt.Println("memberVip 没找到", memberID)
				continue
			}
			orders, ok := memberOrderMap[memberID]
			if !ok {
				fmt.Println("memberID 没支付", memberID)
				continue
			}
			for _, order := range orders {
				if order.Status != enum.MemberVipOrderStatusPaid {
					continue
				}
				durationType, ok := constant.DurationToMemberVipDurationType[order.Duration.String]
				if !ok {
					continue
				}
				endAt := constant.GetMemberVipDays(durationType)
				/// 再多给 五天
				endAt = endAt.AddDate(0, 0, 5)
				memberVip.EndAt = endAt
				util.PanicIf(impl.DefaultMemberVipDAO.Update(ctx, memberVip))
			}
		}
	}
}
