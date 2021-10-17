package main

import (
	"context"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"fmt"
	pkgErr "github.com/pkg/errors"
)

func main() {

	ctx := context.Background()

	hasNext := true
	offset := 0
	size := 100

	for hasNext {
		fmt.Println("offset-->: ", offset)

		ids, err := impl2.DefaultAccountDAO.ListIDs(ctx, offset, size, nil, nil)
		util.PanicIf(err)
		offset += len(ids)
		hasNext = len(ids) == size

		accountMap, err := impl2.DefaultAccountDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		for _, id := range ids {
			_, ok := accountMap[id]
			if !ok {
				fmt.Println(fmt.Sprintf("account not found. id=%d", id))
				break
			}
			e, err := impl2.DefaultMemberIDEncryptionDAO.GetByMemberID(ctx, id)
			if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
				util.PanicIf(err)
			}
			if e == nil {
				code := util2.MustGenerateCode(ctx, 10)
				util.PanicIf(impl2.DefaultMemberIDEncryptionDAO.Insert(ctx, &models.MemberIDEncryption{
					MemberID: id,
					Code:     code,
				}))
			}
		}
	}

	fmt.Println("done")
}
