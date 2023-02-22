package main

import (
	"context"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	errors2 "dumpapp_server/pkg/errors"
	util3 "dumpapp_server/pkg/util"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/strmangle"
	"os"
)

func main() {

	csvFile, err := os.OpenFile("l2.csv", os.O_CREATE|os.O_RDWR, 0644)
	util.PanicIf(err)
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	util.PanicIf(writer.Write([]string{"OutID"}))

	ctx := context.Background()

	orderID := util3.MustGenerateID(ctx)

	count := 50
	price := 0
	level := 2
	outIDs, err := getOutIDs(ctx, count, level)
	util.PanicIf(err)

	util.PanicIf(impl.DefaultInstallAppCdkeyOrderDAO.Insert(ctx, &models.InstallAppCdkeyOrder{
		ID:     orderID,
		Status: enum.MemberPayOrderStatusPaid,
		Number: int64(count),
		Amount: cast.ToFloat64(price * count),
		BizExt: datatype.InstallAppCdkeyOrderBizExt{
			IsAgent: true,
		},
	}))

	for _, oID := range outIDs {
		util.PanicIf(writer.Write([]string{oID}))
		id := util3.MustGenerateID(ctx)
		util.PanicIf(impl.DefaultInstallAppCdkeyDAO.Insert(ctx, &models.InstallAppCdkey{
			ID:      id,
			OutID:   oID,
			Status:  enum.InstallAppCDKeyStatusNormal,
			OrderID: orderID,
		}))
	}

	writer.Flush()
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
