package main

import (
	"context"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	util3 "dumpapp_server/pkg/util"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/strmangle"
)

func main() {

	ctx := context.Background()

	orderID := util3.MustGenerateID(ctx)

	count := 10
	price := 68
	outIDs, err := getOutIDs(ctx, count, 2)
	util.PanicIf(err)

	bizExt := constant.InstallAppCDKEYOrderBizExt{
		IsAgent: true,
	}
	util.PanicIf(impl.DefaultInstallAppCdkeyOrderDAO.Insert(ctx, &models.InstallAppCdkeyOrder{
		ID:     orderID,
		Status: enum.MemberPayOrderStatusPaid,
		Number: int64(count),
		Amount: cast.ToFloat64(price * count),
		BizExt: bizExt.String(),
	}))

	for _, oID := range outIDs {
		fmt.Println(oID)
		id := util3.MustGenerateID(ctx)
		util.PanicIf(impl.DefaultInstallAppCdkeyDAO.Insert(ctx, &models.InstallAppCdkey{
			ID:      id,
			OutID:   oID,
			Status:  enum.InstallAppCDKeyStatusNormal,
			OrderID: orderID,
		}))
	}
}

func getOutIDs(ctx context.Context, number, level int) ([]string, error) {
	outIDs := make([]string, 0)
	/// 生成 number * 10 的数量，以防重复
	for i := 0; i < number*10; i++ {
		oID := fmt.Sprintf("%sL%d", util3.MustGenerateAppCDKEY(), level)
		outIDs = append(outIDs, oID)
	}
	outIDs = strmangle.RemoveDuplicates(outIDs)

	cMap, err := impl.DefaultInstallAppCdkeyDAO.BatchGetByOutID(ctx, outIDs)
	if err != nil {
		return nil, err
	}

	resultOutIDs := make([]string, 0)
	for _, oID := range outIDs {
		if len(resultOutIDs) == number {
			return resultOutIDs, nil
		}
		if _, ok := cMap[oID]; !ok {
			resultOutIDs = append(resultOutIDs, oID)
		}
	}
	return nil, errors2.ErrInstallAppGenerateCDKeyFail
}
