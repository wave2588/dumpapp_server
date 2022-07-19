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
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AccountDAO struct {
	mysqlPool *sql.DB
}

var DefaultAccountDAO *AccountDAO

func init() {
	DefaultAccountDAO = NewAccountDAO()
}

func NewAccountDAO() *AccountDAO {
	d := &AccountDAO{
		mysqlPool: clients.MySQLConnectionsPool,
	}
	return d
}

func (d *AccountDAO) Insert(ctx context.Context, data *models.Account) error {
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

func (d *AccountDAO) Update(ctx context.Context, data *models.Account) error {
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

func (d *AccountDAO) Delete(ctx context.Context, id int64) error {
	qs := []qm.QueryMod{
		models.AccountWhere.ID.EQ(id),
	}

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	_, err := models.Accounts(qs...).DeleteAll(ctx, exec)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	return nil
}

func (d *AccountDAO) Get(ctx context.Context, id int64) (*models.Account, error) {
	result, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	if v, ok := result[id]; !ok {
		return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=account, id=%d", id)
	} else {
		return v, nil
	}
}

// BatchGet retrieves multiple records by primary key from db.
func (d *AccountDAO) BatchGet(ctx context.Context, ids []int64) (map[int64]*models.Account, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.Accounts(models.AccountWhere.ID.IN(ids)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[int64]*models.Account)
	for _, c := range datas {
		result[c.ID] = c
	}

	return result, nil
}

// 后台和脚本使用：倒序列出所有
func (d *AccountDAO) ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error) {
	if offset < 0 || limit <= 0 || limit > 10000 {
		return nil, pkgErr.Errorf("invalid offset or limit")
	}
	qs := []qm.QueryMod{qm.Select(models.AccountColumns.ID)}
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

	datas, err := models.Accounts(qs...).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.Wrap(err, fmt.Sprintf("table=account offset=%d limit=%d filters=%v", offset, limit, filters))
	}

	result := make([]int64, 0)
	for _, c := range datas {
		result = append(result, c.ID)
	}
	return result, nil
}

func (d *AccountDAO) Count(ctx context.Context, filters []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{qm.Select(models.AccountColumns.ID)}
	qs = append(qs, filters...)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	return models.Accounts(qs...).Count(ctx, exec)
}

// GetByPhone retrieves a single record by uniq key phone from db.
func (d *AccountDAO) GetByPhone(ctx context.Context, phone string) (*models.Account, error) {
	accountObj := &models.Account{}

	sel := "*"
	query := fmt.Sprintf(
		"select %s from `account` where `phone`=?", sel,
	)

	q := queries.Raw(query, phone)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	err := q.Bind(ctx, exec, accountObj)
	if err != nil {
		if pkgErr.Cause(err) == sql.ErrNoRows {
			return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=account, query=%s, args=phone :%v", query, phone)
		}
		return nil, pkgErr.Wrap(err, "dao: unable to select from account")
	}

	return accountObj, nil
}

// BatchGetByPhone retrieves multiple records by uniq key phone from db.
func (d *AccountDAO) BatchGetByPhone(ctx context.Context, phones []string) (map[string]*models.Account, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.Accounts(models.AccountWhere.Phone.IN(phones)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[string]*models.Account)
	for _, c := range datas {
		result[c.Phone] = c
	}

	return result, nil
}

// GetByEmail retrieves a single record by uniq key email from db.
func (d *AccountDAO) GetByEmail(ctx context.Context, email string) (*models.Account, error) {
	accountObj := &models.Account{}

	sel := "*"
	query := fmt.Sprintf(
		"select %s from `account` where `email`=?", sel,
	)

	q := queries.Raw(query, email)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	err := q.Bind(ctx, exec, accountObj)
	if err != nil {
		if pkgErr.Cause(err) == sql.ErrNoRows {
			return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=account, query=%s, args=email :%v", query, email)
		}
		return nil, pkgErr.Wrap(err, "dao: unable to select from account")
	}

	return accountObj, nil
}

// BatchGetByEmail retrieves multiple records by uniq key email from db.
func (d *AccountDAO) BatchGetByEmail(ctx context.Context, emails []string) (map[string]*models.Account, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.Accounts(models.AccountWhere.Email.IN(emails)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[string]*models.Account)
	for _, c := range datas {
		result[c.Email] = c
	}

	return result, nil
}
