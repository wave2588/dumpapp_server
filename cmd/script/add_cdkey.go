package main

import (
	"context"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

func main() {

	ctx := context.Background()

	for i := 0; i < 3; i++ {
		orderID := util2.MustGenerateID(ctx)
		bizExt := datatype.InstallAppCdkeyOrderBizExt{
			ContactWay: "15711367321",
			IsTest:     true,
		}
		util.PanicIf(impl.DefaultInstallAppCdkeyOrderDAO.Insert(ctx, &models.InstallAppCdkeyOrder{
			ID:     orderID,
			Status: enum.MemberPayOrderStatusPaid,
			Number: 1,
			Amount: 10,
			BizExt: bizExt,
		}))

		cdkeyID := util2.MustGenerateID(ctx)
		outID := util2.MustGenerateAppCDKEY()
		util.PanicIf(impl.DefaultInstallAppCdkeyDAO.Insert(ctx, &models.InstallAppCdkey{
			ID:      cdkeyID,
			OutID:   outID,
			Status:  enum.InstallAppCDKeyStatusNormal,
			OrderID: orderID,
		}))
	}
}
