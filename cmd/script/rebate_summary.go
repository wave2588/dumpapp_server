package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	util2 "dumpapp_server/pkg/util"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cast"
	"os"
)

func main() {
	ctx := context.Background()

	csvFile, err := os.OpenFile("返利明细.csv", os.O_CREATE|os.O_RDWR, 0644)
	util.PanicIf(err)
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	util.PanicIf(writer.Write([]string{"邀请人", "被邀请人", "被邀请人充值金额", "返利金额", "返利时间"}))

	offset := 0
	bulkSize := 1000
	hasNext := true

	//startAt := time.Date(time.Now().Year(), 5, 1, 0, 0, 0, 0, time.Local)
	//endAt := time.Date(time.Now().Year(), 5, 31, 23, 59, 59, 0, time.Local)

	//filter := []qm.QueryMod{
	//	models.MemberRebateRecordWhere.CreatedAt.GTE(startAt),
	//	models.MemberRebateRecordWhere.CreatedAt.LTE(endAt),
	//}

	for hasNext {
		ids, err := impl.DefaultMemberRebateRecordDAO.ListIDs(ctx, offset, bulkSize, nil, nil)
		util.PanicIf(err)

		offset += offset
		hasNext = len(ids) == bulkSize

		rebateRecordMap, err := impl.DefaultMemberRebateRecordDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		memberIDs := make([]int64, 0)

		orderIDs := make([]int64, 0)
		for _, record := range rebateRecordMap {
			orderIDs = append(orderIDs, record.OrderID)
			memberIDs = append(memberIDs, record.ReceiverMemberID)
		}

		orderMap, err := impl.DefaultMemberPayOrderDAO.BatchGet(ctx, orderIDs)
		util.PanicIf(err)
		for _, order := range orderMap {
			memberIDs = append(memberIDs, order.MemberID)
		}
		memberIDs = util2.RemoveDuplicates(memberIDs)

		accountMap, err := impl.DefaultAccountDAO.BatchGet(ctx, memberIDs)
		util.PanicIf(err)

		for _, rebateRecordID := range ids {
			rebate, ok := rebateRecordMap[rebateRecordID]
			if !ok {
				fmt.Println("rebate not found", rebateRecordID)
				continue
			}

			receiverMember, ok := accountMap[rebate.ReceiverMemberID]
			if !ok {
				fmt.Println("receiverMember not found", rebate.ReceiverMemberID)
				continue
			}
			order, ok := orderMap[rebate.OrderID]
			if !ok {
				fmt.Println("order not found", rebate.OrderID)
				continue
			}
			orderMember, ok := accountMap[order.MemberID]
			if !ok {
				fmt.Println("order_member not found", order.MemberID)
				continue
			}
			util.PanicIf(writer.Write([]string{receiverMember.Email, orderMember.Email, cast.ToString(order.Amount), cast.ToString(rebate.Count), cast.ToString(rebate.CreatedAt.Format("2006-01-02 15:04:05"))}))
		}
	}

	writer.Flush()

	fmt.Println("Done")
}
