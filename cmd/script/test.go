package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func main() {

	ctx := context.Background()

	util.PanicIf(impl.DefaultIpaRankingDAO.RemoveIpaRankingData(ctx))

	now := time.Now()
	data, err := getIpaRankingData(ctx, now.AddDate(0, 0, -1).Unix(), now.Unix())
	util.PanicIf(err)

	util.PanicIf(impl.DefaultIpaRankingDAO.SetIpaRankingData(ctx, &dao.IpaRanking{
		Data: data,
	}))

	res, err := impl.DefaultIpaRankingDAO.GetIpaRankingData(ctx)
	util.PanicIf(err)

	fmt.Println(len(res.Data))
}

func getIpaRankingData(ctx context.Context, startAt, endAt int64) ([]interface{}, error) {

	filter := make([]qm.QueryMod, 0)
	filter = append(filter, models.SearchRecordV2Where.CreatedAt.GTE(cast.ToTime(startAt)))
	filter = append(filter, models.SearchRecordV2Where.CreatedAt.LTE(cast.ToTime(endAt)))

	data, err := impl.DefaultSearchRecordV2DAO.GetOrderBySearchCount(ctx, 0, 20, filter)
	if err != nil {
		return nil, err
	}

	ipaIDs := make([]int64, 0)
	for _, datum := range data {
		ipaIDs = append(ipaIDs, datum.IpaID)
	}
	appleDataMap, err := impl2.DefaultAppleController.BatchGetAppInfoByAppIDs(ctx, ipaIDs)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0)
	for _, ipaID := range ipaIDs {
		appleData, ok := appleDataMap[ipaID]
		if !ok {
			continue
		}
		result = append(result, appleData)
	}
	return result, nil
}
