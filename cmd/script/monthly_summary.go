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

	var amount int64 = 0
	month := time.Month(1)

	startAt := time.Date(time.Now().Year(), month, 1, 0, 0, 0, 0, time.Local)
	//endAt := time.Date(time.Now().Year(), month, 30, 23, 59, 59, 0, time.Local)

	resIDs := make([]int64, 0)
	for hasNext {
		filter := []qm.QueryMod{
			models.MemberDownloadOrderWhere.CreatedAt.GTE(startAt),
			//models.MemberDownloadOrderWhere.CreatedAt.LTE(endAt),
			models.MemberDownloadOrderWhere.Status.EQ(enum.MemberDownloadOrderStatusPaid),
		}
		ids, err := impl.DefaultMemberDownloadOrderDAO.ListIDs(ctx, offset, bulkSize, filter, nil)
		util.PanicIf(err)

		resIDs = append(resIDs, ids...)

		hasNext = len(ids) >= bulkSize
		offset += len(ids)

		res, err := impl.DefaultMemberDownloadOrderDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		for _, order := range res {
			amount += order.Number * 9
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
	fmt.Println(fmt.Sprintf("%d 总收入-->: %d", month, amount))
}
