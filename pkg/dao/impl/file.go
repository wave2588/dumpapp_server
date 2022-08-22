// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

type FileDAO struct {
	mysqlPool *sql.DB
}

var DefaultFileDAO *FileDAO

func init() {
	DefaultFileDAO = NewFileDAO()
}

func NewFileDAO() *FileDAO {
	d := &FileDAO{
		mysqlPool: clients.MySQLConnectionsPool,
	}
	return d
}

func (d *FileDAO) Insert(ctx context.Context, data *models.File) error {
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

func (d *FileDAO) Update(ctx context.Context, data *models.File) error {
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

func (d *FileDAO) Delete(ctx context.Context, id int64) error {
	qs := []qm.QueryMod{
		models.FileWhere.ID.EQ(id),
	}

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	_, err := models.Files(qs...).DeleteAll(ctx, exec)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	return nil
}

func (d *FileDAO) Get(ctx context.Context, id int64) (*models.File, error) {
	result, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	if v, ok := result[id]; !ok {
		return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=file, id=%d", id)
	} else {
		return v, nil
	}
}

// BatchGet retrieves multiple records by primary key from db.
func (d *FileDAO) BatchGet(ctx context.Context, ids []int64) (map[int64]*models.File, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.Files(models.FileWhere.ID.IN(ids)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[int64]*models.File)
	for _, c := range datas {
		result[c.ID] = c
	}

	return result, nil
}

// 后台和脚本使用：倒序列出所有
func (d *FileDAO) ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error) {
	if offset < 0 || limit <= 0 || limit > 10000 {
		return nil, pkgErr.Errorf("invalid offset or limit")
	}
	qs := []qm.QueryMod{qm.Select(models.FileColumns.ID)}
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

	datas, err := models.Files(qs...).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.Wrap(err, fmt.Sprintf("table=file offset=%d limit=%d filters=%v", offset, limit, filters))
	}

	result := make([]int64, 0)
	for _, c := range datas {
		result = append(result, c.ID)
	}
	return result, nil
}

func (d *FileDAO) Count(ctx context.Context, filters []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{qm.Select(models.FileColumns.ID)}
	qs = append(qs, filters...)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	return models.Files(qs...).Count(ctx, exec)
}
