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

type AdminDumpOrderDAO struct {
	mysqlPool *sql.DB
}

var DefaultAdminDumpOrderDAO *AdminDumpOrderDAO

func init() {
	DefaultAdminDumpOrderDAO = NewAdminDumpOrderDAO()
}

func NewAdminDumpOrderDAO() *AdminDumpOrderDAO {
	d := &AdminDumpOrderDAO{
		mysqlPool: clients.MySQLConnectionsPool,
	}
	return d
}

func (d *AdminDumpOrderDAO) Insert(ctx context.Context, data *models.AdminDumpOrder) error {
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

func (d *AdminDumpOrderDAO) Update(ctx context.Context, data *models.AdminDumpOrder) error {
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

func (d *AdminDumpOrderDAO) Delete(ctx context.Context, id int64) error {
	qs := []qm.QueryMod{
		models.AdminDumpOrderWhere.ID.EQ(id),
	}

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	_, err := models.AdminDumpOrders(qs...).DeleteAll(ctx, exec)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	return nil
}

func (d *AdminDumpOrderDAO) Get(ctx context.Context, id int64) (*models.AdminDumpOrder, error) {
	result, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	if v, ok := result[id]; !ok {
		return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=admin_dump_order, id=%d", id)
	} else {
		return v, nil
	}
}

// BatchGet retrieves multiple records by primary key from db.
func (d *AdminDumpOrderDAO) BatchGet(ctx context.Context, ids []int64) (map[int64]*models.AdminDumpOrder, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.AdminDumpOrders(models.AdminDumpOrderWhere.ID.IN(ids)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[int64]*models.AdminDumpOrder)
	for _, c := range datas {
		result[c.ID] = c
	}

	return result, nil
}

// 后台和脚本使用：倒序列出所有
func (d *AdminDumpOrderDAO) ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error) {
	if offset < 0 || limit <= 0 || limit > 10000 {
		return nil, pkgErr.Errorf("invalid offset or limit")
	}
	qs := []qm.QueryMod{qm.Select(models.AdminDumpOrderColumns.ID)}
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

	datas, err := models.AdminDumpOrders(qs...).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.Wrap(err, fmt.Sprintf("table=admin_dump_order offset=%d limit=%d filters=%v", offset, limit, filters))
	}

	result := make([]int64, 0)
	for _, c := range datas {
		result = append(result, c.ID)
	}
	return result, nil
}

func (d *AdminDumpOrderDAO) Count(ctx context.Context, filters []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{qm.Select(models.AdminDumpOrderColumns.ID)}
	qs = append(qs, filters...)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	return models.AdminDumpOrders(qs...).Count(ctx, exec)
}

// GetByIpaIDIpaVersion retrieves a single record by uniq key ipaID, ipaVersion from db.
func (d *AdminDumpOrderDAO) GetByIpaIDIpaVersion(ctx context.Context, ipaID int64, ipaVersion string) (*models.AdminDumpOrder, error) {
	adminDumpOrderObj := &models.AdminDumpOrder{}

	sel := "*"
	query := fmt.Sprintf(
		"select %s from `admin_dump_order` where `ipa_id`=? AND `ipa_version`=?", sel,
	)

	q := queries.Raw(query, ipaID, ipaVersion)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	err := q.Bind(ctx, exec, adminDumpOrderObj)
	if err != nil {
		if pkgErr.Cause(err) == sql.ErrNoRows {
			return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=admin_dump_order, query=%s, args=ipaID:%v ipaVersion :%v", query, ipaID, ipaVersion)
		}
		return nil, pkgErr.Wrap(err, "dao: unable to select from admin_dump_order")
	}

	return adminDumpOrderObj, nil
}

// GetAdminDumpOrderSliceByIpaID retrieves a slice of records by first field of uniq key [ipaID] with an executor.
func (d *AdminDumpOrderDAO) GetAdminDumpOrderSliceByIpaID(ctx context.Context, ipaID int64) ([]*models.AdminDumpOrder, error) {
	var o []*models.AdminDumpOrder

	query := "select `admin_dump_order`.* from `admin_dump_order` where `ipa_id`=?"

	q := queries.Raw(query, ipaID)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		if pkgErr.Cause(err) == sql.ErrNoRows {
			return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=admin_dump_order, query=%s, args=ipaID :%v", query, ipaID)
		}
		return nil, pkgErr.Wrap(err, "dao: unable to select from admin_dump_order")
	}
	return o, nil
}
