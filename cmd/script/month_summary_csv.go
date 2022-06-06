package main

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"os"
	"sort"
	"time"
)

type MemberAmount struct {
	MemberID        int64
	MemberEmail     string
	TotalAmount     float64
	TotalDCoinCount int64
}

func main() {
	ctx := context.Background()

	csvFile, err := os.OpenFile("order.csv", os.O_CREATE|os.O_RDWR, 0644)
	util.PanicIf(err)
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	util.PanicIf(writer.Write([]string{"用户邮箱", "用户 ID", "5月充值总金额", "5月总 D 币数"}))

	offset := 0
	bulkSize := 1000
	hasNext := true

	startAt := time.Date(time.Now().Year(), 5, 1, 0, 0, 0, 0, time.Local)
	endAt := time.Date(time.Now().Year(), 5, 31, 23, 59, 59, 0, time.Local)

	memberAmountMap := make(map[int64]*MemberAmount)

	for hasNext {
		filter := []qm.QueryMod{
			models.MemberPayOrderWhere.CreatedAt.GTE(startAt),
			models.MemberPayOrderWhere.CreatedAt.LTE(endAt),
			models.MemberPayOrderWhere.Status.EQ(enum.MemberPayOrderStatusPaid),
		}
		orderIDs, err := impl.DefaultMemberPayOrderDAO.ListIDs(ctx, offset, bulkSize, filter, nil)
		util.PanicIf(err)

		offset += len(orderIDs)
		hasNext = bulkSize == len(orderIDs)

		orderMap, err := impl.DefaultMemberPayOrderDAO.BatchGet(ctx, orderIDs)
		util.PanicIf(err)

		memberIDs := make([]int64, 0)
		for _, order := range orderMap {
			memberIDs = append(memberIDs, order.MemberID)
		}
		memberIDs = util2.RemoveDuplicates(memberIDs)

		accountMap, err := impl.DefaultAccountDAO.BatchGet(ctx, memberIDs)
		util.PanicIf(err)

		memberPayCountMap, err := impl.DefaultMemberPayCountDAO.BatchGetByMemberIDs(ctx, memberIDs)

		for _, orderID := range orderIDs {
			order, ok := orderMap[orderID]
			if !ok {
				fmt.Println("order not found-->; ", orderID)
				continue
			}
			account, ok := accountMap[order.MemberID]
			if !ok {
				if !ok {
					fmt.Println("account not found-->; ", orderID)
					continue
				}
			}
			mPayCounts := memberPayCountMap[order.MemberID]

			if ma, ok := memberAmountMap[order.MemberID]; ok {
				memberAmountMap[order.MemberID].TotalAmount = ma.TotalAmount + order.Amount
			} else {
				memberAmountMap[order.MemberID] = &MemberAmount{
					MemberID:        order.MemberID,
					MemberEmail:     account.Email,
					TotalAmount:     order.Amount,
					TotalDCoinCount: cast.ToInt64(len(mPayCounts)),
				}
			}
		}
	}

	memberAmounts := make([]*MemberAmount, 0)
	for _, amount := range memberAmountMap {
		memberAmounts = append(memberAmounts, amount)
	}

	sort.Slice(memberAmounts, func(i, j int) bool {
		return memberAmounts[i].TotalAmount > memberAmounts[j].TotalAmount
	})

	for _, amount := range memberAmounts {
		util.PanicIf(writer.Write([]string{amount.MemberEmail, cast.ToString(amount.MemberID), cast.ToString(amount.TotalAmount)}))
	}

	writer.Flush()

	fmt.Println("Done")
}
