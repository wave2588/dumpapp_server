package delete_ipa

import (
	"context"
	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func Run() {
	fmt.Println("DeleteInterimIpa")
	run()
}

type deleteIpa struct {
	ID      int64
	Name    string
	Version string
}

func run() {
	ctx := context.Background()

	offset := 0
	bulkSize := 100
	hasNext := true

	tm := time.Now().AddDate(0, 0, -3)

	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	deleteIpaMap := make(map[int64][]*deleteIpa, 0)

	for hasNext {
		fmt.Println(fmt.Sprintf("offset: %d...", offset))

		filters := []qm.QueryMod{
			models.IpaWhere.IsInterim.EQ(1), /// 被标记为临时ipa
			models.IpaWhere.UpdatedAt.LT(tm),
		}
		ids, err := impl.DefaultIpaDAO.ListIDs(ctx, offset, bulkSize, filters, nil)
		util.PanicIf(err)

		ipaMap, err := impl.DefaultIpaDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		hasNext = len(ids) >= bulkSize
		offset += len(ids)

		for _, ipa := range ipaMap {
			util.PanicIf(impl.DefaultIpaDAO.Delete(ctx, ipa.ID))
		}
		ipaVersions, err := impl.DefaultIpaVersionDAO.BatchGetIpaVersions(ctx, ids)
		util.PanicIf(err)

		for ipaID, versions := range ipaVersions {
			ipa, ok := ipaMap[ipaID]
			if !ok {
				continue
			}
			for _, version := range versions {
				util.PanicIf(impl.DefaultIpaVersionDAO.Delete(ctx, version.ID))
				util.PanicIf(impl2.DefaultTencentController.DeleteFile(ctx, version.TokenPath))
				deleteIpaMap[ipaID] = append(deleteIpaMap[ipaID], &deleteIpa{
					ID:      ipaID,
					Name:    ipa.Name,
					Version: version.Version,
				})
			}
		}
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	nameStr := ""
	for _, is := range deleteIpaMap {
		for _, ipa := range is {
			nameStr += fmt.Sprintf("<font color=\"info\">名称：</font>%s %s\n>", ipa.Name, ipa.Version)
		}
	}

	contentStr := fmt.Sprintf("<font color=\"warning\">已删除临时 ipa：</font>\n>")
	deleteIpaStr := fmt.Sprintf("<font color=\"comment\">%s</font>", nameStr)
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": contentStr +
				deleteIpaStr,
		},
	}
	util2.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{"@all"})

}
