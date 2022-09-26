package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller/impl"
	impl2 "dumpapp_server/pkg/dao/impl"
	"fmt"
)

var bucket = config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket

func main() {
	ctx := context.Background()

	resp, err := impl.DefaultLingshulianController.List(ctx, bucket)
	util.PanicIf(err)

	ipaFileTokens := make([]string, 0)
	for _, content := range resp.Contents {
		if content.Key == nil {
			continue
		}
		ipaFileTokens = append(ipaFileTokens, *content.Key)
	}

	signIpaMap, err := impl2.DefaultMemberSignIpaDAO.BatchGetByIpaFileToken(ctx, ipaFileTokens)
	util.PanicIf(err)

	for _, token := range ipaFileTokens {
		_, ok := signIpaMap[token]
		if ok {
			continue
		}
		fmt.Println("delete-->: ", token)
		err := deleteIpaFile(ctx, token)
		if err != nil {
			fmt.Println("delete fail", err)
		}
	}
}

func deleteIpaFile(ctx context.Context, token string) error {
	return impl.DefaultLingshulianController.Delete(ctx, bucket, token)
}
