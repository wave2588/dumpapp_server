// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package impl

import (
	"context"
	"database/sql"
	"fmt"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/dao/models"
	mysqlDriver "github.com/go-sql-driver/mysql"
	pkgErr "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SearchRecordV2DAO struct {
	mysqlPool *sql.DB
}

var DefaultSearchRecordV2DAO *SearchRecordV2DAO

func init() {
	DefaultSearchRecordV2DAO = NewSearchRecordV2DAO()
}

func NewSearchRecordV2DAO() *SearchRecordV2DAO {
	d := &SearchRecordV2DAO{
		mysqlPool: clients.MySQLConnectionsPool,
	}
	return d
}

func (d *SearchRecordV2DAO) Insert(ctx context.Context, data *models.SearchRecordV2) error {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	err := data.Insert(ctx, exec, boil.Infer())
	if err != nil {
		if mysqlError, ok := pkgErr.Cause(err).(*mysqlDriver.MySQLError); !(ok && mysqlError.Number == 1062) {
			return pkgErr.WithStack(err)
		}
	}
	return nil
}

func (d *SearchRecordV2DAO) Update(ctx context.Context, data *models.SearchRecordV2) error {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	_, err := data.Update(ctx, exec, boil.Infer())
	return pkgErr.WithStack(err)
}

func (d *SearchRecordV2DAO) Delete(ctx context.Context, id int64) error {
	qs := []qm.QueryMod{
		models.SearchRecordV2Where.ID.EQ(id),
	}

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	_, err := models.SearchRecordV2S(qs...).DeleteAll(ctx, exec)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	return nil
}

func (d *SearchRecordV2DAO) Get(ctx context.Context, id int64) (*models.SearchRecordV2, error) {
	result, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	if v, ok := result[id]; !ok {
		return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=search_record_v2, id=%d", id)
	} else {
		return v, nil
	}
}

// BatchGet retrieves multiple records by primary key from db.
func (d *SearchRecordV2DAO) BatchGet(ctx context.Context, ids []int64) (map[int64]*models.SearchRecordV2, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.SearchRecordV2S(models.SearchRecordV2Where.ID.IN(ids)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[int64]*models.SearchRecordV2)
	for _, c := range datas {
		result[c.ID] = c
	}

	return result, nil
}

// 后台和脚本使用：倒序列出所有
func (d *SearchRecordV2DAO) ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error) {
	if offset < 0 || limit <= 0 || limit > 10000 {
		return nil, pkgErr.Errorf("invalid offset or limit")
	}
	qs := []qm.QueryMod{qm.Select(models.SearchRecordV2Columns.ID)}
	qs = append(qs, filters...)

	if len(orderBys) > 0 {
		orderBys = append(orderBys, "id desc")
		for _, orderBy := range orderBys {
			qs = append(qs, qm.OrderBy(orderBy))
		}
	} else {
		qs = append(qs, qm.OrderBy("created_at DESC, id DESC"))
	}

	if offset >= 0 && limit >= 0 {
		qs = append(qs, qm.Offset(offset), qm.Limit(limit))
	}

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	datas, err := models.SearchRecordV2S(qs...).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.Wrap(err, fmt.Sprintf("table=search_record_v2 offset=%d limit=%d filters=%v", offset, limit, filters))
	}

	result := make([]int64, 0)
	for _, c := range datas {
		result = append(result, c.ID)
	}
	return result, nil
}

func (d *SearchRecordV2DAO) Count(ctx context.Context, filters []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{qm.Select(models.SearchRecordV2Columns.ID)}
	qs = append(qs, filters...)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	return models.SearchRecordV2S(qs...).Count(ctx, exec)
}
