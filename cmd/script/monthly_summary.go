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

	var webAmount, appAmount float64 = 0, 0
	month := time.Now().Month()
	//day := time.Now().Day()
	day := 1

	startAt := time.Date(time.Now().Year(), month, day, 0, 0, 0, 0, time.Local)
	//endAt := time.Date(time.Now().Year(), month, 31, 23, 59, 59, 0, time.Local)

	cerCount, err := impl.DefaultCertificateV2DAO.Count(ctx, []qm.QueryMod{
		models.CertificateV2Where.CreatedAt.GTE(startAt),
	})
	util.PanicIf(err)

	iCerCount, err := impl.DefaultInstallAppCertificateDAO.Count(ctx, []qm.QueryMod{
		models.InstallAppCertificateWhere.CreatedAt.GTE(startAt),
	})
	util.PanicIf(err)

	cerAmount := (cerCount + iCerCount) * 25

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
			if order.BizExt.Platform == enum.MemberPayOrderPlatformWeb {
				webAmount += order.Amount
			} else {
				appAmount += order.Amount
			}
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

	filter := []qm.QueryMod{
		models.InstallAppCdkeyOrderWhere.CreatedAt.GTE(startAt),
		models.InstallAppCdkeyOrderWhere.Status.EQ(enum.MemberPayOrderStatusPaid),
	}
	installAppIDs, err := impl.DefaultInstallAppCdkeyOrderDAO.ListIDs(ctx, 0, 9999, filter, nil)
	util.PanicIf(err)
	installAppOrderMap, err := impl.DefaultInstallAppCdkeyOrderDAO.BatchGet(ctx, installAppIDs)
	var installAppAmount float64
	for _, order := range installAppOrderMap {
		installAppAmount += order.Amount
	}

	total := webAmount + appAmount + installAppAmount
	fmt.Println(fmt.Sprintf("%d-%d 至今 新注册用户数-->: %d", month, day, len(memberIDs)))
	fmt.Println(fmt.Sprintf("%d-%d 至今 支付成功的订单-->: %d", month, day, len(resIDs)))
	fmt.Println(fmt.Sprintf("%d-%d 至今 主站收入-->: %.2f, web: %.2f  app: %.2f", month, day, webAmount+appAmount, webAmount, appAmount))
	fmt.Println(fmt.Sprintf("%d-%d 至今 app 兑换码收入-->: %.2f", month, day, installAppAmount))
	fmt.Println(fmt.Sprintf("%d-%d 至今 总收入-->: %.2f", month, day, total))
	fmt.Println(fmt.Sprintf("%d-%d 至今 支出-->: %d", month, day, cerAmount))
	fmt.Println(fmt.Sprintf("%d-%d 至今 利润-->: %.2f%%", month, day, (total-float64(cerAmount))/total*100))
	fmt.Println(fmt.Sprintf("%d-%d 至今 账户剩余-->: %.2f", month, day, total-float64(cerAmount)))
}
