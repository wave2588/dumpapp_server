package pending_order

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func Run() {
	fmt.Println("pending order")
	run()
}

//https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=a5bc9a15-3206-4686-a828-fadfda6379b3

func run() {
	ctx := context.Background()

	endAt := time.Now()
	startAt := endAt.AddDate(0, 0, -1)

	orderIDs := make([]int64, 0)

	offset := 0
	limit := 100
	hasNext := true
	for hasNext {
		ids, err := impl.DefaultMemberPayOrderDAO.ListIDs(ctx, offset, limit, []qm.QueryMod{
			models.MemberPayOrderWhere.CreatedAt.GTE(startAt),
			models.MemberPayOrderWhere.CreatedAt.LTE(endAt),
			models.MemberPayOrderWhere.Status.EQ(enum.MemberPayOrderStatusPending),
			models.MemberPayOrderWhere.Amount.GTE(100),
		}, nil)
		util.PanicIf(err)

		offset += len(ids)
		hasNext = limit == len(ids)

		orderIDs = append(orderIDs, ids...)
	}

	orderMap, err := impl.DefaultMemberPayOrderDAO.BatchGet(ctx, orderIDs)
	util.PanicIf(err)

	memberIDs := make([]int64, 0)
	for _, order := range orderMap {
		memberIDs = append(memberIDs, order.MemberID)
	}

	memberMap, err := impl.DefaultAccountDAO.BatchGet(ctx, memberIDs)
	util.PanicIf(err)

	title := "<font color=\"warning\">过去 24h 没有未充值成功的订单</font>\n"
	message := ""
	if len(orderIDs) != 0 {
		title = "<font color=\"warning\">过去 24h 未充值成功的订单</font>\n"
		for _, oID := range orderIDs {
			order, ok := orderMap[oID]
			if !ok {
				continue
			}
			member, ok := memberMap[order.MemberID]
			if !ok {
				continue
			}
			msg := fmt.Sprintf("\n \n邮箱: <font color=\"comment\">%s</font>  \n手机号: <font color=\"comment\">%s</font>  \n金额: <font color=\"comment\">%.2f</font> \n订单时间: %s", member.Email, member.Phone, order.Amount, order.CreatedAt.Format("2006-01-02 15:04:05"))
			message += msg
		}
	}

	impl2.DefaultAlterWebController.SendCustomMsg(ctx, "a5bc9a15-3206-4686-a828-fadfda6379b3", title+message)
}
