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

	localFileNameMap, err := impl.DefaultFileController.GetLocalPlistFiles(ctx)
	util.PanicIf(err)

	fmt.Println("localFileNameMap-->: ", len(localFileNameMap))

	localFileNames := make([]string, 0)
	for key := range localFileNameMap {
		localFileNames = append(localFileNames, key)
	}

	fileMap, err := impl2.DefaultFileDAO.BatchGetByTokens(ctx, localFileNames)
	util.PanicIf(err)

	memberSignIpaMap, err := impl2.DefaultMemberSignIpaDAO.BatchGetByIpaPlistFileToken(ctx, localFileNames)
	util.PanicIf(err)

	for key, path := range localFileNameMap {
		isDelete := false

		file, fOk := fileMap[key]
		if fOk && file.IsDelete {
			isDelete = true
		}

		mSignIpa, sOK := memberSignIpaMap[key]
		if sOK && mSignIpa.IsDelete {
			isDelete = true
		}

		if !fOk && !sOK {
			/// 都没找到的话, 就需要删除
			isDelete = true
		}

		if isDelete {
			fmt.Println("需要删除-->: ", path, key)
			util.PanicIf(impl.DefaultFileController.DeleteFile(ctx, path))
		}
	}
}
