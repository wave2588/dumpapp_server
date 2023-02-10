package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"fmt"
)

func main() {

	ctx := context.Background()

	//createdAt := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	//endAt := createdAt.AddDate(0, 0, 1)
	//for true {
	//
	//	count, err := impl.DefaultMemberDownloadIpaRecordDAO.Count(ctx, []qm.QueryMod{
	//		models.MemberDownloadIpaRecordWhere.CreatedAt.GTE(createdAt),
	//		models.MemberDownloadIpaRecordWhere.CreatedAt.LTE(endAt),
	//	})
	//	util.PanicIf(err)
	//
	//	fmt.Println(createdAt.Format("2006-01-02"), count)
	//
	//	createdAt = createdAt.AddDate(0, 0, -1)
	//	endAt = endAt.AddDate(0, 0, -1)
	//}

	//resp, err := impl2.DefaultCertificateDeviceController.IsReplenish(ctx, 1597582789929078784, "00008020-00115DDA26F3002E")
	//util.PanicIf(err)
	//
	//fmt.Println(resp)

	ac, err := impl.DefaultAccountDAO.GetByEmail(ctx, "15711367321@163.com")
	util.PanicIf(err)
	fmt.Println(ac)

	res, err := impl.DefaultAdminConfigDAO.GetAdminBusy(ctx)
	util.PanicIf(err)
	fmt.Println(res)

	util.PanicIf(impl.DefaultIpaRankingDAO.RemoveIpaRankingData(ctx))

}
