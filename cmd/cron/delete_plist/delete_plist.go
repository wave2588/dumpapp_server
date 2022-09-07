package delete_plist

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller/impl"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func Run() {
	fmt.Println("delete plist")
	run()
}

func run() {
	ctx := context.Background()

	deleteFileCount, deleteSignIpaCount, err := deleteFile(ctx)
	util.PanicIf(err)

	contentStr := fmt.Sprintf("<font color=\"warning\">定时删除 plist：</font>\n>")
	deleteFileStr := fmt.Sprintf("<font color=\"comment\">删除了本地签名 plist 文件 %d 个</font>\n", deleteFileCount)
	deleteSignIpaStr := fmt.Sprintf("<font color=\"comment\">删除了签名分发 plist 文件 %d 个</font>", deleteSignIpaCount)
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": contentStr + deleteFileStr + deleteSignIpaStr,
		},
	}
	util2.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{"@all"})
}

func deleteFile(ctx context.Context) (int64, int64, error) {
	localFileNameMap, err := impl.DefaultFileController.GetLocalPlistFiles(ctx)
	if err != nil {
		return 0, 0, err
	}

	/// 删除 file 库
	files, err := getNeedDeleteFiles(ctx)
	if err != nil {
		return 0, 0, err
	}
	for _, file := range files {
		if filePath, ok := localFileNameMap[file.Token]; ok {
			_ = impl.DefaultFileController.DeleteFile(ctx, filePath)
		}
		file.IsDelete = true
		err = impl2.DefaultFileDAO.Update(ctx, file)
		if err != nil {
			return 0, 0, err
		}
	}

	/// 删除 member_sign_ipa 库
	signIpas, err := getNeedDeleteSignIpa(ctx)
	if err != nil {
		return 0, 0, err
	}
	for _, si := range signIpas {
		if filePath, ok := localFileNameMap[si.IpaPlistFileToken]; ok {
			_ = impl.DefaultFileController.DeleteFile(ctx, filePath)
		}
		_ = impl.DefaultLingshulianController.Delete(ctx, config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket, si.IpaFileToken)
		si.IsDelete = true
		err = impl2.DefaultMemberSignIpaDAO.Update(ctx, si)
		if err != nil {
			return 0, 0, err
		}
	}

	return int64(len(files)), int64(len(signIpas)), nil
}

func getNeedDeleteFiles(ctx context.Context) ([]*models.File, error) {
	offset := 0
	limit := 100
	hasNext := true

	files := make([]*models.File, 0)

	tm := time.Now().AddDate(0, 0, -3)
	filter := []qm.QueryMod{
		models.FileWhere.CreatedAt.LT(tm),
	}
	for hasNext {
		ids, err := impl2.DefaultFileDAO.ListIDs(ctx, offset, limit, filter, nil)
		if err != nil {
			return nil, err
		}

		offset += len(ids)
		hasNext = limit == len(ids)

		fileMap, err := impl2.DefaultFileDAO.BatchGet(ctx, ids)
		if err != nil {
			return nil, err
		}

		for _, file := range fileMap {
			if file.IsDelete {
				continue
			}
			files = append(files, file)
		}
	}

	return files, nil
}

func getNeedDeleteSignIpa(ctx context.Context) ([]*models.MemberSignIpa, error) {
	offset := 0
	limit := 100
	hasNext := true

	signIpas := make([]*models.MemberSignIpa, 0)

	tm := time.Now().AddDate(0, 0, -3)
	filter := []qm.QueryMod{
		models.MemberSignIpaWhere.CreatedAt.LT(tm),
	}
	for hasNext {
		ids, err := impl2.DefaultMemberSignIpaDAO.ListIDs(ctx, offset, limit, filter, nil)
		if err != nil {
			return nil, err
		}

		offset += len(ids)
		hasNext = limit == len(ids)

		signIpaMap, err := impl2.DefaultMemberSignIpaDAO.BatchGet(ctx, ids)
		if err != nil {
			return nil, err
		}

		for _, si := range signIpaMap {
			if si.IsDelete {
				continue
			}
			signIpas = append(signIpas, si)
		}
	}

	return signIpas, nil
}
