package main

import (
	"context"
	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
)

func main() {

	ctx := context.Background()

	ids, err := impl.DefaultIpaDAO.ListIDs(ctx, 0, 10000, nil, nil)
	util.PanicIf(err)

	ipas, err := impl.DefaultIpaDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	for _, ipa := range ipas {
		app, err := impl.DefaultAppDAO.GetByName(ctx, ipa.Name)
		util.PanicIf(err)
		version := &models.IpaVersion{
			IpaID:     ipa.ID,
			Version:   "",
			TokenPath: app.TokenPath,
		}
		util.PanicIf(impl.DefaultIpaVersionDAO.Insert(ctx, version))
	}

	clients.MustCommit(ctx, txn)
	util.ResetCtxKey(ctx, constant.TransactionKeyTxn)
}
