package main

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func main() {

	ctx := context.Background()

	offset := 0
	bulkSize := 100
	hasNext := true

	var amount float64 = 0
	month := time.Now().Month()

	startAt := time.Date(time.Now().Year(), month, 1, 0, 0, 0, 0, time.Local)
	//endAt := time.Date(time.Now().Year(), month, 31, 23, 59, 59, 0, time.Local)

	resIDs := make([]int64, 0)
	for hasNext {
		filter := []qm.QueryMod{
			models.MemberPayOrderWhere.CreatedAt.GTE(startAt),
			//models.MemberPayOrderWhere.CreatedAt.LTE(endAt),
			models.MemberPayOrderWhere.Status.EQ(enum.MemberPayOrderStatusPaid),
		}
		ids, err := impl.DefaultMemberPayOrderDAO.ListIDs(ctx, offset, bulkSize, filter, nil)
		util.PanicIf(err)

		resIDs = append(resIDs, ids...)

		hasNext = len(ids) >= bulkSize
		offset += len(ids)

		res, err := impl.DefaultMemberPayOrderDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		for _, order := range res {
			amount += order.Amount
		}
	}

	offset = 0
	hasNext = true
	memberIDs := make([]int64, 0)
	for hasNext {
		filter := []qm.QueryMod{
			models.AccountWhere.CreatedAt.GTE(startAt),
		}
		ids, err := impl.DefaultAccountDAO.ListIDs(ctx, offset, bulkSize, filter, nil)
		util.PanicIf(err)
		hasNext = len(ids) >= bulkSize
		offset += len(ids)

		memberIDs = append(memberIDs, ids...)
	}

	fmt.Println(fmt.Sprintf("%d 月新注册用户数-->: %d", month, len(memberIDs)))
	fmt.Println(fmt.Sprintf("%d 支付成功的订单-->: %d", month, len(resIDs)))
	fmt.Println(fmt.Sprintf("%d 总收入-->: %.2f", month, amount))
}
