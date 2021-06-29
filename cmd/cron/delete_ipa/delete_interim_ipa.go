package delete_ipa

import (
	"context"
	"dumpapp_server/pkg/common/util"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func Run() {
	fmt.Println("DeleteInterimIpa")
	run()
}

func run() {
	ctx := context.Background()

	offset := 0
	bulkSize := 100
	hasNext := true

	tm := time.Now().AddDate(0, 0, -3)

	for hasNext {
		fmt.Println(fmt.Sprintf("offset: %d...", offset))

		filters := []qm.QueryMod{
			models.IpaWhere.IsInterim.EQ(1), /// 被标记为临时ipa
			models.IpaWhere.UpdatedAt.LT(tm),
		}
		ids, err := impl.DefaultIpaDAO.ListIDs(ctx, offset, bulkSize, filters, nil)
		util.PanicIf(err)

		hasNext = len(ids) >= bulkSize
		offset += len(ids)

		fmt.Println(ids)

		for _, id := range ids {
			util.PanicIf(impl.DefaultIpaDAO.Delete(ctx, id))
		}
		ipaVersions, err := impl.DefaultIpaVersionDAO.BatchGetIpaVersions(ctx, ids)
		util.PanicIf(err)

		for _, versions := range ipaVersions {
			for _, version := range versions {
				util.PanicIf(impl.DefaultIpaVersionDAO.Delete(ctx, version.ID))
				util.PanicIf(impl2.DefaultTencentController.DeleteFile(ctx, version.TokenPath))
			}
		}
	}
}
