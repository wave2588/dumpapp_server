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

type MemberAppSourceDAO struct {
	mysqlPool *sql.DB
}

var DefaultMemberAppSourceDAO *MemberAppSourceDAO

func init() {
	DefaultMemberAppSourceDAO = NewMemberAppSourceDAO()
}

func NewMemberAppSourceDAO() *MemberAppSourceDAO {
	d := &MemberAppSourceDAO{
		mysqlPool: clients.MySQLConnectionsPool,
	}
	return d
}

func (d *MemberAppSourceDAO) Insert(ctx context.Context, data *models.MemberAppSource) error {
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

func (d *MemberAppSourceDAO) Update(ctx context.Context, data *models.MemberAppSource) error {
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

func (d *MemberAppSourceDAO) Delete(ctx context.Context, id int64) error {
	qs := []qm.QueryMod{
		models.MemberAppSourceWhere.ID.EQ(id),
	}

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	_, err := models.MemberAppSources(qs...).DeleteAll(ctx, exec)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	return nil
}

func (d *MemberAppSourceDAO) Get(ctx context.Context, id int64) (*models.MemberAppSource, error) {
	result, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	if v, ok := result[id]; !ok {
		return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=member_app_source, id=%d", id)
	} else {
		return v, nil
	}
}

// BatchGet retrieves multiple records by primary key from db.
func (d *MemberAppSourceDAO) BatchGet(ctx context.Context, ids []int64) (map[int64]*models.MemberAppSource, error) {
	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}
	datas, err := models.MemberAppSources(models.MemberAppSourceWhere.ID.IN(ids)).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	result := make(map[int64]*models.MemberAppSource)
	for _, c := range datas {
		result[c.ID] = c
	}

	return result, nil
}

// 后台和脚本使用：倒序列出所有
func (d *MemberAppSourceDAO) ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error) {
	if offset < 0 || limit <= 0 || limit > 10000 {
		return nil, pkgErr.Errorf("invalid offset or limit")
	}
	qs := []qm.QueryMod{qm.Select(models.MemberAppSourceColumns.ID)}
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

	datas, err := models.MemberAppSources(qs...).All(ctx, exec)
	if err != nil {
		return nil, pkgErr.Wrap(err, fmt.Sprintf("table=member_app_source offset=%d limit=%d filters=%v", offset, limit, filters))
	}

	result := make([]int64, 0)
	for _, c := range datas {
		result = append(result, c.ID)
	}
	return result, nil
}

func (d *MemberAppSourceDAO) Count(ctx context.Context, filters []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{qm.Select(models.MemberAppSourceColumns.ID)}
	qs = append(qs, filters...)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	return models.MemberAppSources(qs...).Count(ctx, exec)
}

// GetByMemberIDAppSourceID retrieves a single record by uniq key memberID, appSourceID from db.
func (d *MemberAppSourceDAO) GetByMemberIDAppSourceID(ctx context.Context, memberID int64, appSourceID int64) (*models.MemberAppSource, error) {
	memberAppSourceObj := &models.MemberAppSource{}

	sel := "*"
	query := fmt.Sprintf(
		"select %s from `member_app_source` where `member_id`=? AND `app_source_id`=?", sel,
	)

	q := queries.Raw(query, memberID, appSourceID)

	var exec boil.ContextExecutor
	txn := ctx.Value("txn")
	if txn == nil {
		exec = d.mysqlPool
	} else {
		exec = txn.(*sql.Tx)
	}

	err := q.Bind(ctx, exec, memberAppSourceObj)
	if err != nil {
		if pkgErr.Cause(err) == sql.ErrNoRows {
			return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=member_app_source, query=%s, args=memberID:%v appSourceID :%v", query, memberID, appSourceID)
		}
		return nil, pkgErr.Wrap(err, "dao: unable to select from member_app_source")
	}

	return memberAppSourceObj, nil
}

// GetMemberAppSourceSliceByMemberID retrieves a slice of records by first field of uniq key [memberID] with an executor.
func (d *MemberAppSourceDAO) GetMemberAppSourceSliceByMemberID(ctx context.Context, memberID int64) ([]*models.MemberAppSource, error) {
	var o []*models.MemberAppSource

	query := "select `member_app_source`.* from `member_app_source` where `member_id`=?"

	q := queries.Raw(query, memberID)

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
			return nil, pkgErr.Wrapf(errors.ErrNotFound, "table=member_app_source, query=%s, args=memberID :%v", query, memberID)
		}
		return nil, pkgErr.Wrap(err, "dao: unable to select from member_app_source")
	}
	return o, nil
}
