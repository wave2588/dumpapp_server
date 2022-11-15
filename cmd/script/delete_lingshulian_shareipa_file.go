package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller/impl"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"fmt"
)

func main() {

	ctx := context.Background()

	var (
		marker *string
		limit  = 1000
		bucket = config.DumpConfig.AppConfig.LingshulianShareIpaBucket
	)

	ids, err := impl2.DefaultIpaVersionDAO.ListIDs(ctx, 0, limit, nil, []string{})
	util.PanicIf(err)

	ipaVersionMap, err := impl2.DefaultIpaVersionDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	tokenMap := make(map[string]*models.IpaVersion)
	for _, version := range ipaVersionMap {
		tokenMap[version.TokenPath] = version
	}

	res, err := impl.DefaultLingshulianController.List(ctx, bucket, marker, int64(limit))
	util.PanicIf(err)

	for _, content := range res.Contents {
		if content.Key == nil {
			continue
		}
		key := *content.Key
		_, ok := tokenMap[key]
		if ok {
			fmt.Println("不需要删除--->: ", key)
			continue
		}

		fmt.Println("需要删除--->: ", key)
		/// 说明在库里没有找到, 需要删除
		util.PanicIf(impl.DefaultLingshulianController.Delete(ctx, bucket, key))
	}
}
