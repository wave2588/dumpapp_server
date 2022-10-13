// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type InstallAppCdkeyDAO struct {
	mysqlPool *sql.DB
}

var DefaultInstallAppCdkeyDAO *InstallAppCdkeyDAO

func init() {
	DefaultInstallAppCdkeyDAO = NewInstallAppCdkeyDAO()
}

func NewInstallAppCdkeyDAO() *InstallAppCdkeyDAO {
	d := &InstallAppCdkeyDAO{
		mysqlPool: clients.MySQLConnectionsPool,
	}
	return d
}

func (d *InstallAppCdkeyDAO) Insert(ctx context.Context, data *models.InstallAppCdkey) error {
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

func (d *InstallAppCdkeyDAO) Update(ctx context.Context, data *models.InstallAppCdkey) error {
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

func (d *InstallAppCdkeyDAO) Delete(ctx context.Context, id int64) error {
	qs := []qm.QueryMod{
		models.InstallAppCdkeyWhere.ID.EQ(id),
	}

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	_, err := models.InstallAppCdkeys(qs...).DeleteAll(ctx, exec)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	return nil
}

func (d *InstallAppCdkeyDAO) Get(ctx context.Context, id int64) (*models.InstallAppCdkey, error) {
	result, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	if v, ok := result[id]; !ok {
		return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=install_app_cdkey, id=%d", id)
	} else {
		return v, nil
	}
}

// BatchGet retrieves multiple records by primary key from db.
func (d *InstallAppCdkeyDAO) BatchGet(ctx context.Context, ids []int64) (map[int64]*models.InstallAppCdkey, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.InstallAppCdkeys(models.InstallAppCdkeyWhere.ID.IN(ids)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[int64]*models.InstallAppCdkey)
	for _, c := range datas {
		result[c.ID] = c
	}

	return result, nil
}

// 后台和脚本使用：倒序列出所有
func (d *InstallAppCdkeyDAO) ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error) {
	if offset < 0 || limit <= 0 || limit > 10000 {
		return nil, pkgErr.Errorf("invalid offset or limit")
	}
	qs := []qm.QueryMod{qm.Select(models.InstallAppCdkeyColumns.ID)}
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

	datas, err := models.InstallAppCdkeys(qs...).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.Wrap(err, fmt.Sprintf("table=install_app_cdkey offset=%d limit=%d filters=%v", offset, limit, filters))
	}

	result := make([]int64, 0)
	for _, c := range datas {
		result = append(result, c.ID)
	}
	return result, nil
}

func (d *InstallAppCdkeyDAO) Count(ctx context.Context, filters []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{qm.Select(models.InstallAppCdkeyColumns.ID)}
	qs = append(qs, filters...)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	return models.InstallAppCdkeys(qs...).Count(ctx, exec)
}

// GetByOutID retrieves a single record by uniq key outID from db.
func (d *InstallAppCdkeyDAO) GetByOutID(ctx context.Context, outID string) (*models.InstallAppCdkey, error) {
	installAppCdkeyObj := &models.InstallAppCdkey{}

	sel := "*"
	query := fmt.Sprintf(
		"select %s from `install_app_cdkey` where `out_id`=?", sel,
	)

	q := queries.Raw(query, outID)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	err := q.Bind(ctx, exec, installAppCdkeyObj)
	if err != nil {
		if pkgErr.Cause(err) == sql.ErrNoRows {
			return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=install_app_cdkey, query=%s, args=outID :%v", query, outID)
		}
		return nil, pkgErr.Wrap(err, "dao: unable to select from install_app_cdkey")
	}

	return installAppCdkeyObj, nil
}

// BatchGetByOutID retrieves multiple records by uniq key outID from db.
func (d *InstallAppCdkeyDAO) BatchGetByOutID(ctx context.Context, outIDs []string) (map[string]*models.InstallAppCdkey, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.InstallAppCdkeys(models.InstallAppCdkeyWhere.OutID.IN(outIDs)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[string]*models.InstallAppCdkey)
	for _, c := range datas {
		result[c.OutID] = c
	}

	return result, nil
}
