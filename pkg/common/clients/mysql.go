package clients

import (
	"context"
	"database/sql"
	"fmt"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnectionsPool(master string, slaves []string, maxIdle, maxOpen int) *sql.DB {
	fmt.Println("start db")

	db, err := sql.Open("mysql", master)
	util.PanicIf(err)

	return db
}

var MySQLConnectionsPool = NewMySQLConnectionsPool(
	config.DumpConfig.AppConfig.MySQL.Master,
	config.DumpConfig.AppConfig.MySQL.Slaves,
	256, 256)

// var MySQLOfflinePool = NewMySQLConnectionsPool(
//	config.DumpConfig.AppConfig.ChaoMySQL.Master,
//	config.DumpConfig.AppConfig.ChaoMySQL.Offlines,
//	256, 256)

/// 事物
func GetMySQLTransaction(ctx context.Context, db *sql.DB, master bool) *sql.Tx {
	/// todo: master 暂时没用到
	tx, err := db.Begin()
	util.PanicIf(err)
	return tx
}

func MustClearMySQLTransaction(ctx context.Context, txn *sql.Tx) {
	err := txn.Rollback()
	if err != sql.ErrTxDone && err != nil {
		panic(err)
	}
}

func MustCommit(ctx context.Context, txn *sql.Tx) {
	if err := txn.Commit(); err != nil {
		panic(err)
	}
}

const (
	UpdateMode string = "FOR UPDATE"
	ShareMode  string = "LOCK IN SHARE MODE"
)
