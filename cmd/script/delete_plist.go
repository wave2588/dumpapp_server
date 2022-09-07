package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller/impl"
	impl2 "dumpapp_server/pkg/dao/impl"
	"fmt"
	"time"
)

func main() {
	/// 文件只存 3 天
	interval := 3

	ctx := context.Background()

	now := time.Now()

	localFileNameMap, err := impl.DefaultFileController.GetLocalPlistFiles(ctx)
	util.PanicIf(err)

	fmt.Println("localFileNameMap-->: ", len(localFileNameMap))

	fileKeys := make([]string, 0)
	for key := range localFileNameMap {
		fileKeys = append(fileKeys, key)
	}

	/// 删除 file
	fileMap, err := impl2.DefaultFileDAO.BatchGetByTokens(ctx, fileKeys)
	util.PanicIf(err)
	for fileKey, filePath := range localFileNameMap {
		file, ok := fileMap[fileKey]
		if !ok {
			continue
		}
		if timeSub(now, file.CreatedAt) < interval {
			continue
		}
		delete(localFileNameMap, fileKey)
		util.PanicIf(impl.DefaultFileController.DeleteFile(ctx, filePath))
		file.IsDelete = true
		util.PanicIf(impl2.DefaultFileDAO.Update(ctx, file))
	}

	/// 删除 member_sign_ipa
	signIpaMap, err := impl2.DefaultMemberSignIpaDAO.BatchGetByIpaPlistFileToken(ctx, fileKeys)
	util.PanicIf(err)
	for fileKey, filePath := range localFileNameMap {
		signIpa, ok := signIpaMap[fileKey]
		if !ok {
			continue
		}
		if timeSub(now, signIpa.CreatedAt) < interval {
			continue
		}
		delete(localFileNameMap, fileKey)
		util.PanicIf(impl.DefaultFileController.DeleteFile(ctx, filePath))
		util.PanicIf(impl.DefaultLingshulianController.Delete(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, signIpa.IpaFileToken))
		signIpa.IsDelete = true
		util.PanicIf(impl2.DefaultMemberSignIpaDAO.Update(ctx, signIpa))
	}

	fmt.Println("delete localFileNameMap-->: ", len(localFileNameMap))

	/// 剩下的文件说明库里没有, 则全删掉
	for _, filePath := range localFileNameMap {
		util.PanicIf(impl.DefaultFileController.DeleteFile(ctx, filePath))
	}

}

func timeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}
