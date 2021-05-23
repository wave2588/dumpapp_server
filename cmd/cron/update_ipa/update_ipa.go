package update_ipa

import (
	"context"
	"dumpapp_server/pkg/common/util"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao/impl"
	"fmt"
)

type UpdateIpa struct {
}

func (r *UpdateIpa) Run() {
	fmt.Println("UpdateIpa")

}

func (r *UpdateIpa) sss() {

	ctx := context.Background()

	hasNext := true
	offset := 0
	limit := 100

	for hasNext {
		ids, err := impl.DefaultIpaDAO.ListIDs(ctx, offset, limit, nil, nil)
		util.PanicIf(err)

		ipaMap, err := impl.DefaultIpaDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		appInfoMap, err := impl2.DefaultAppleController.BatchGetAppInfoByAppIDs(ctx, ids)
		util.PanicIf(err)

		for _, ipa := range ipaMap {
			appInfo := appInfoMap[ipa.ID]
			if appInfo == nil {
				continue
			}
			ipa.Name = appInfo.Name
			ipa.BundleID = appInfo.BundleID
			util.PanicIf(impl.DefaultIpaDAO.Update(ctx, ipa))
		}

		offset += limit
		hasNext = len(ids) < limit
	}
}
