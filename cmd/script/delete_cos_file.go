package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/impl"
	impl2 "dumpapp_server/pkg/dao/impl"
	"fmt"
	"strings"
)

func main() {

	ctx := context.Background()

	hasNext := true
	var nextMarker *string = nil

	ids, err := impl2.DefaultIpaVersionDAO.ListIDs(ctx, 0, 1000, nil, []string{})
	util.PanicIf(err)

	ipaVersionMap, err := impl2.DefaultIpaVersionDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	tokenMap := make(map[string]bool)
	for _, version := range ipaVersionMap {
		tokenMap[version.TokenPath] = true
	}

	for hasNext {
		res, err := impl.DefaultTencentController.ListFile(ctx, nextMarker, 1000)
		util.PanicIf(err)

		hasNext = len(res.Contents) != 0
		nextMarker = util.StringPtr(res.NextMarker)

		for _, content := range res.Contents {
			if strings.Contains(content.Key, "cos-access-log/2022") {
				continue
			}
			_, ok := tokenMap[content.Key]
			if ok {
				fmt.Println("不需要删除", content.Key)
				continue
			}
			fmt.Println("要删除", content.Key)
			util.PanicIf(impl.DefaultTencentController.DeleteFile(ctx, content.Key))
		}
	}

	fmt.Println("done")
}
