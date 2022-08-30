package delete_plist

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/impl"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
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

	err := deleteFile(ctx)
	util.PanicIf(err)

	/// 获取过去三天的 file
}

func deleteFile(ctx context.Context) error {
	localFileNameMap, err := getLocalPlistFile(ctx)
	if err != nil {
		return err
	}

	/// 删除文件
	files, err := getNeedDeleteFiles(ctx)
	if err != nil {
		return err
	}
	for _, file := range files {
		if filePath, ok := localFileNameMap[file.Token]; ok {
			_ = impl.DefaultFileController.DeleteFile(ctx, filePath)
		}
		file.IsDelete = true
		err = impl2.DefaultFileDAO.Update(ctx, file)
		if err != nil {
			return err
		}
	}
	fmt.Println(len(files))

	/// 删除签名文件
	signIpas, err := getNeedDeleteSignIpa(ctx)
	if err != nil {
		return err
	}
	for _, si := range signIpas {
		if filePath, ok := localFileNameMap[si.IpaPlistFileToken]; ok {
			_ = impl.DefaultFileController.DeleteFile(ctx, filePath)
		}
		si.IsDelete = true
		err = impl2.DefaultMemberSignIpaDAO.Update(ctx, si)
		if err != nil {
			return err
		}
	}

	return nil
}

func getLocalPlistFile(ctx context.Context) (map[string]string, error) {

	filterFileMap := map[string]struct{}{
		"ipa1.plist":  {},
		"ipa2.plist":  {},
		"ipa3.plist":  {},
		"ipa4.plist":  {},
		"ipa5.plist":  {},
		"ipa6.plist":  {},
		"ipa7.plist":  {},
		"ipa8.plist":  {},
		"ipa9.plist":  {},
		"ipa10.plist": {},
		"logo.png":    {},
	}

	plistFolderPath := impl.DefaultFileController.GetPlistFolderPath(ctx)
	fileNames, err := impl.DefaultFileController.ListFolder(ctx, plistFolderPath)
	if err != nil {
		return nil, err
	}

	resultFileNameMap := make(map[string]string, 0)
	for _, name := range fileNames {
		if _, ok := filterFileMap[name]; ok {
			continue
		}
		resultFileNameMap[name] = fmt.Sprintf("%s/%s", plistFolderPath, name)
	}
	return resultFileNameMap, nil
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
			signIpas = append(signIpas, si)
		}
	}

	return signIpas, nil
}
