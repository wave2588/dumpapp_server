package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {

	ctx := context.Background()

	ids, err := impl.DefaultAccountDAO.ListIDs(ctx, 0, 3000, []qm.QueryMod{
		models.AccountWhere.Phone.EQ(""),
	}, nil)
	util.PanicIf(err)

	res, err := impl.DefaultAccountDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	for idx, account := range res {
		if account.Phone == "" {
			account.Phone = cast.ToString(idx)
			util.PanicIf(impl.DefaultAccountDAO.Update(ctx, account))
		}
	}
}
