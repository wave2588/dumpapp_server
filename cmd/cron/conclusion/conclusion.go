package conclusion

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func Run() {
	run()
}

func run() {
	ctx := context.Background()

	offset := 0
	bulkSize := 100
	hasNext := true

	memberIDs := make([]int64, 0)

	now := time.Now()
	startAt := now.AddDate(0, 0, -1)
	for hasNext {
		filters := []qm.QueryMod{
			models.AccountWhere.CreatedAt.GT(startAt),
		}
		ids, err := impl.DefaultAccountDAO.ListIDs(ctx, offset, bulkSize, filters, nil)
		util.PanicIf(err)

		hasNext = len(ids) >= bulkSize
		offset += len(ids)

		memberIDs = append(memberIDs, ids...)
	}

	/// 计算支付过的新用户
	memberOrderMap, err := impl.DefaultMemberDownloadOrderDAO.BatchGetByMemberIDs(ctx, memberIDs)
	util.PanicIf(err)
	paidMemberIDs := make([]int64, 0)
	for memberID, orders := range memberOrderMap {
		for _, order := range orders {
			if order.Status == enum.MemberDownloadOrderStatusPaid {
				paidMemberIDs = append(paidMemberIDs, memberID)
			}
		}
	}

	/// 计算总费用
	filters := []qm.QueryMod{
		models.MemberDownloadOrderWhere.CreatedAt.GT(startAt),
	}
	totalMemberOrders, err := impl.DefaultMemberDownloadOrderDAO.GetByFilters(ctx, filters, nil)
	util.PanicIf(err)
	amount := 0.0
	orderCount := 0
	for _, order := range totalMemberOrders {
		if order.Status == enum.MemberDownloadOrderStatusPaid {
			amount += order.Amount.Float64
			orderCount += 1
		}
	}

	/// 计算使用人数
	numbers := getMemberDownloadMap(ctx, startAt)
	downloadedCount := 0
	downloadedMember := make(map[int64]struct{}, 0)
	for _, number := range numbers {
		if number.Status == enum.MemberDownloadNumberStatusUsed {
			downloadedCount++
			downloadedMember[number.MemberID] = struct{}{}
		}
	}

	contentStr := fmt.Sprintf("<font color=\"info\">每日总结\n截止昨日此时数据统计如下：</font>\n>")
	newMemberStr := fmt.Sprintf("新注册用户：<font color=\"comment\">%d</font> 人\n", len(memberIDs))
	paidMemberStr := fmt.Sprintf("新用户付费率：：<font color=\"comment\">%.2f%%</font>\n", cast.ToFloat64(len(paidMemberIDs))/cast.ToFloat64(len(memberIDs))*100)
	orderCountStr := fmt.Sprintf("总订单：<font color=\"comment\">%d</font>\n", orderCount)
	amountStr := fmt.Sprintf("总收入：<font color=\"comment\">%v</font>\n", amount)
	downloadedStr := fmt.Sprintf("使用次数：<font color=\"comment\">%d</font>\n", downloadedCount)
	downloadedMemberStr := fmt.Sprintf("下载人数：<font color=\"comment\">%d</font>\n", len(downloadedMember))
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", now.Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": contentStr +
				newMemberStr + paidMemberStr + orderCountStr + amountStr + downloadedStr + downloadedMemberStr + timeStr,
		},
	}
	util2.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{"@all"})
}

func getMemberDownloadMap(ctx context.Context, startAt time.Time) []*models.MemberDownloadNumber {
	offset := 0
	bulkSize := 100
	hasNext := true

	result := make([]*models.MemberDownloadNumber, 0)
	for hasNext {
		fmt.Println(fmt.Sprintf("offset: %d...", offset))

		filter := []qm.QueryMod{
			models.MemberDownloadNumberWhere.CreatedAt.GT(startAt),
		}
		ids, err := impl.DefaultMemberDownloadNumberDAO.ListIDs(ctx, offset, bulkSize, filter, nil)
		util.PanicIf(err)

		data, err := impl.DefaultMemberDownloadNumberDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		hasNext = len(ids) >= bulkSize
		offset += len(ids)

		for _, number := range data {
			result = append(result, number)
		}
	}

	return result
}