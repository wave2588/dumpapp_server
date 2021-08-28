package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/impl"
	impl2 "dumpapp_server/pkg/dao/impl"
	"fmt"
)

func main() {

	ctx := context.Background()

	hasNext := true
	var nextMarker *string = nil

	for hasNext {
		res, err := impl.DefaultTencentController.ListFile(ctx, nextMarker, 100)
		util.PanicIf(err)

		hasNext = len(res.Contents) != 0
		nextMarker = util.StringPtr(res.NextMarker)

		tokens := make([]string, 0)
		for _, content := range res.Contents {
			tokens = append(tokens, content.Key)
		}

		ipaVersionMap, err := impl2.DefaultIpaVersionDAO.BatchGetByTokenPath(ctx, tokens)
		util.PanicIf(err)

		for _, content := range res.Contents {
			_, ok := ipaVersionMap[content.Key]
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
