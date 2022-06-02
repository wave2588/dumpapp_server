package main

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"os"
	"sort"
	"time"
)

type MemberDCoinDetail struct {
	MemberID    int64
	MemberEmail string

	TotalAmount     float64 /// 总充值金额
	TotalDCoinCount int64   /// 总 D 币数量

	DCoinFormPayCount            int64 /// 用户充值 D 币
	DCoinFormPayForFreeCount     int64 /// 用户充值赠送
	DCoinFormInviteCount         int64 /// 邀请用户送的 D 币
	DCoinFormRebateCount         int64 /// 返利获取 D 币
	DCoinFormAdminPresentedCount int64 /// 后台赠送

	DCoinUseByIpaCount int64 /// 用来买 ipa D 币
	DCoinUseByCerCount int64 /// 用来买证书 D 币
}

func main() {

	ctx := context.Background()

	csvFile, err := os.OpenFile("5 月用户 D 币明细.csv", os.O_CREATE|os.O_RDWR, 0644)
	util.PanicIf(err)
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	util.PanicIf(writer.Write([]string{"用户邮箱", "用户 ID", "充值总金额", "总 D 币数", "--分割线--", "充值获取 D 币", "充值赠送 D 币", "邀请赠送 D 币", "返利 D 币", "后台赠送 D 币", "--分割线--", "购买 ipa D 币数", "购买证书 D 币数"}))

	startAt := time.Date(time.Now().Year(), 5, 1, 0, 0, 0, 0, time.Local)
	endAt := time.Date(time.Now().Year(), 5, 31, 23, 59, 59, 0, time.Local)

	/// 获取用户充值记录
	filter := []qm.QueryMod{
		models.MemberPayOrderWhere.CreatedAt.GTE(startAt),
		models.MemberPayOrderWhere.CreatedAt.LTE(endAt),
		models.MemberPayOrderWhere.Status.EQ(enum.MemberPayOrderStatusPaid),
	}
	orderIDs, err := impl.DefaultMemberPayOrderDAO.ListIDs(ctx, 0, 100000, filter, nil)
	util.PanicIf(err)
	orderMap, err := impl.DefaultMemberPayOrderDAO.BatchGet(ctx, orderIDs)
	util.PanicIf(err)
	memberPayOrderMap := make(map[int64]float64)
	for _, order := range orderMap {
		memberPayOrderMap[order.MemberID] = memberPayOrderMap[order.MemberID] + order.Amount
	}

	payCountFilter := []qm.QueryMod{
		models.MemberPayCountWhere.CreatedAt.GTE(startAt),
		models.MemberPayCountWhere.CreatedAt.LTE(endAt),
	}
	ids, err := impl.DefaultMemberPayCountDAO.ListIDs(ctx, 0, 100000, payCountFilter, nil)
	util.PanicIf(err)
	memberPayCountMap, err := impl.DefaultMemberPayCountDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	resultMemberPayCountMap := make(map[int64][]*models.MemberPayCount)
	for _, count := range memberPayCountMap {
		resultMemberPayCountMap[count.MemberID] = append(resultMemberPayCountMap[count.MemberID], count)
	}

	memberIDs := make([]int64, 0)
	for _, count := range memberPayCountMap {
		memberIDs = append(memberIDs, count.MemberID)
	}
	for _, order := range orderMap {
		memberIDs = append(memberIDs, order.MemberID)
	}
	accountMap, err := impl.DefaultAccountDAO.BatchGet(ctx, memberIDs)
	util.PanicIf(err)

	result := make(map[int64]*MemberDCoinDetail)

	for memberID, counts := range resultMemberPayCountMap {
		account, ok := accountMap[memberID]
		if !ok {
			fmt.Println("member not found  ", memberID)
			continue
		}
		var payCount, payForFreeCount, rebateCount, invitedCount, adminPresentedCount int64
		var useIpaCount, useCerCount int64

		for _, count := range counts {
			switch count.Source {
			case enum.MemberPayCountSourceNormal:
				payCount += 1
			case enum.MemberPayCountSourcePayForFree:
				payForFreeCount += 1
			case enum.MemberPayCountSourceAdminPresented:
				rebateCount += 1
			case enum.MemberPayCountSourceInvitedPresented:
				invitedCount += 1
			case enum.MemberPayCountSourceRebate:
				adminPresentedCount += 1
			}

			switch count.Use.String {
			case "ipa":
				useIpaCount += 1
			case "certificate":
				useCerCount += 1
			}
		}

		result[memberID] = &MemberDCoinDetail{
			MemberID:                     memberID,
			MemberEmail:                  account.Email,
			TotalAmount:                  memberPayOrderMap[memberID],
			TotalDCoinCount:              int64(len(counts)),
			DCoinFormPayCount:            payCount,
			DCoinFormPayForFreeCount:     payForFreeCount,
			DCoinFormInviteCount:         invitedCount,
			DCoinFormRebateCount:         rebateCount,
			DCoinFormAdminPresentedCount: adminPresentedCount,
			DCoinUseByIpaCount:           useIpaCount,
			DCoinUseByCerCount:           useCerCount,
		}
	}

	res := make([]*MemberDCoinDetail, 0)
	for _, detail := range result {
		res = append(res, detail)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].TotalAmount > res[j].TotalAmount
	})

	for _, re := range res {
		util.PanicIf(writer.Write([]string{
			re.MemberEmail,
			cast.ToString(re.MemberID),
			cast.ToString(re.TotalAmount),
			cast.ToString(re.TotalDCoinCount),
			cast.ToString("----"),
			cast.ToString(re.DCoinFormPayCount),
			cast.ToString(re.DCoinFormPayForFreeCount),
			cast.ToString(re.DCoinFormInviteCount),
			cast.ToString(re.DCoinFormRebateCount),
			cast.ToString(re.DCoinFormAdminPresentedCount),
			cast.ToString("----"),
			cast.ToString(re.DCoinUseByIpaCount),
			cast.ToString(re.DCoinUseByCerCount),
		}))
	}

	writer.Flush()

	fmt.Println("Done")
}
