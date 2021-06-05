// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"dumpapp_server/pkg/common/enum"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MemberVipOrder is an object representing the database table.
type MemberVipOrder struct {
	ID        int64                     `boil:"id" json:"id" toml:"id" yaml:"id"`
	MemberID  int64                     `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	Status    enum.MemberVipOrderStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	Duration  null.String               `boil:"duration" json:"duration,omitempty" toml:"duration" yaml:"duration,omitempty"`
	CreatedAt time.Time                 `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time                 `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *memberVipOrderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L memberVipOrderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MemberVipOrderColumns = struct {
	ID        string
	MemberID  string
	Status    string
	Duration  string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	MemberID:  "member_id",
	Status:    "status",
	Duration:  "duration",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelperenum_MemberVipOrderStatus struct{ field string }

func (w whereHelperenum_MemberVipOrderStatus) EQ(x enum.MemberVipOrderStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperenum_MemberVipOrderStatus) NEQ(x enum.MemberVipOrderStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperenum_MemberVipOrderStatus) LT(x enum.MemberVipOrderStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperenum_MemberVipOrderStatus) LTE(x enum.MemberVipOrderStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperenum_MemberVipOrderStatus) GT(x enum.MemberVipOrderStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperenum_MemberVipOrderStatus) GTE(x enum.MemberVipOrderStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MemberVipOrderWhere = struct {
	ID        whereHelperint64
	MemberID  whereHelperint64
	Status    whereHelperenum_MemberVipOrderStatus
	Duration  whereHelpernull_String
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`member_vip_order`.`id`"},
	MemberID:  whereHelperint64{field: "`member_vip_order`.`member_id`"},
	Status:    whereHelperenum_MemberVipOrderStatus{field: "`member_vip_order`.`status`"},
	Duration:  whereHelpernull_String{field: "`member_vip_order`.`duration`"},
	CreatedAt: whereHelpertime_Time{field: "`member_vip_order`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`member_vip_order`.`updated_at`"},
}

// MemberVipOrderRels is where relationship names are stored.
var MemberVipOrderRels = struct {
}{}

// memberVipOrderR is where relationships are stored.
type memberVipOrderR struct {
}

// NewStruct creates a new relationship struct
func (*memberVipOrderR) NewStruct() *memberVipOrderR {
	return &memberVipOrderR{}
}

// memberVipOrderL is where Load methods for each relationship are stored.
type memberVipOrderL struct{}

var (
	memberVipOrderAllColumns            = []string{"id", "member_id", "status", "duration", "created_at", "updated_at"}
	memberVipOrderColumnsWithoutDefault = []string{"member_id", "status", "duration"}
	memberVipOrderColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	memberVipOrderPrimaryKeyColumns     = []string{"id"}
)

type (
	// MemberVipOrderSlice is an alias for a slice of pointers to MemberVipOrder.
	// This should generally be used opposed to []MemberVipOrder.
	MemberVipOrderSlice []*MemberVipOrder
	// MemberVipOrderHook is the signature for custom MemberVipOrder hook methods
	MemberVipOrderHook func(context.Context, boil.ContextExecutor, *MemberVipOrder) error

	memberVipOrderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	memberVipOrderType                 = reflect.TypeOf(&MemberVipOrder{})
	memberVipOrderMapping              = queries.MakeStructMapping(memberVipOrderType)
	memberVipOrderPrimaryKeyMapping, _ = queries.BindMapping(memberVipOrderType, memberVipOrderMapping, memberVipOrderPrimaryKeyColumns)
	memberVipOrderInsertCacheMut       sync.RWMutex
	memberVipOrderInsertCache          = make(map[string]insertCache)
	memberVipOrderUpdateCacheMut       sync.RWMutex
	memberVipOrderUpdateCache          = make(map[string]updateCache)
	memberVipOrderUpsertCacheMut       sync.RWMutex
	memberVipOrderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var memberVipOrderBeforeInsertHooks []MemberVipOrderHook
var memberVipOrderBeforeUpdateHooks []MemberVipOrderHook
var memberVipOrderBeforeDeleteHooks []MemberVipOrderHook
var memberVipOrderBeforeUpsertHooks []MemberVipOrderHook

var memberVipOrderAfterInsertHooks []MemberVipOrderHook
var memberVipOrderAfterSelectHooks []MemberVipOrderHook
var memberVipOrderAfterUpdateHooks []MemberVipOrderHook
var memberVipOrderAfterDeleteHooks []MemberVipOrderHook
var memberVipOrderAfterUpsertHooks []MemberVipOrderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MemberVipOrder) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MemberVipOrder) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MemberVipOrder) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MemberVipOrder) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MemberVipOrder) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MemberVipOrder) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MemberVipOrder) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MemberVipOrder) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MemberVipOrder) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberVipOrderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMemberVipOrderHook registers your hook function for all future operations.
func AddMemberVipOrderHook(hookPoint boil.HookPoint, memberVipOrderHook MemberVipOrderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		memberVipOrderBeforeInsertHooks = append(memberVipOrderBeforeInsertHooks, memberVipOrderHook)
	case boil.BeforeUpdateHook:
		memberVipOrderBeforeUpdateHooks = append(memberVipOrderBeforeUpdateHooks, memberVipOrderHook)
	case boil.BeforeDeleteHook:
		memberVipOrderBeforeDeleteHooks = append(memberVipOrderBeforeDeleteHooks, memberVipOrderHook)
	case boil.BeforeUpsertHook:
		memberVipOrderBeforeUpsertHooks = append(memberVipOrderBeforeUpsertHooks, memberVipOrderHook)
	case boil.AfterInsertHook:
		memberVipOrderAfterInsertHooks = append(memberVipOrderAfterInsertHooks, memberVipOrderHook)
	case boil.AfterSelectHook:
		memberVipOrderAfterSelectHooks = append(memberVipOrderAfterSelectHooks, memberVipOrderHook)
	case boil.AfterUpdateHook:
		memberVipOrderAfterUpdateHooks = append(memberVipOrderAfterUpdateHooks, memberVipOrderHook)
	case boil.AfterDeleteHook:
		memberVipOrderAfterDeleteHooks = append(memberVipOrderAfterDeleteHooks, memberVipOrderHook)
	case boil.AfterUpsertHook:
		memberVipOrderAfterUpsertHooks = append(memberVipOrderAfterUpsertHooks, memberVipOrderHook)
	}
}

// One returns a single memberVipOrder record from the query.
func (q memberVipOrderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MemberVipOrder, error) {
	o := &MemberVipOrder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for member_vip_order")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MemberVipOrder records from the query.
func (q memberVipOrderQuery) All(ctx context.Context, exec boil.ContextExecutor) (MemberVipOrderSlice, error) {
	var o []*MemberVipOrder

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MemberVipOrder slice")
	}

	if len(memberVipOrderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MemberVipOrder records in the query.
func (q memberVipOrderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count member_vip_order rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q memberVipOrderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if member_vip_order exists")
	}

	return count > 0, nil
}

// MemberVipOrders retrieves all the records using an executor.
func MemberVipOrders(mods ...qm.QueryMod) memberVipOrderQuery {
	mods = append(mods, qm.From("`member_vip_order`"))
	return memberVipOrderQuery{NewQuery(mods...)}
}

// FindMemberVipOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMemberVipOrder(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MemberVipOrder, error) {
	memberVipOrderObj := &MemberVipOrder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `member_vip_order` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, memberVipOrderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from member_vip_order")
	}

	return memberVipOrderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MemberVipOrder) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_vip_order provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(memberVipOrderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	memberVipOrderInsertCacheMut.RLock()
	cache, cached := memberVipOrderInsertCache[key]
	memberVipOrderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			memberVipOrderAllColumns,
			memberVipOrderColumnsWithDefault,
			memberVipOrderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(memberVipOrderType, memberVipOrderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(memberVipOrderType, memberVipOrderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `member_vip_order` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `member_vip_order` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `member_vip_order` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, memberVipOrderPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into member_vip_order")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberVipOrderMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for member_vip_order")
	}

CacheNoHooks:
	if !cached {
		memberVipOrderInsertCacheMut.Lock()
		memberVipOrderInsertCache[key] = cache
		memberVipOrderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MemberVipOrder.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MemberVipOrder) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	memberVipOrderUpdateCacheMut.RLock()
	cache, cached := memberVipOrderUpdateCache[key]
	memberVipOrderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			memberVipOrderAllColumns,
			memberVipOrderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update member_vip_order, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `member_vip_order` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, memberVipOrderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(memberVipOrderType, memberVipOrderMapping, append(wl, memberVipOrderPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update member_vip_order row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for member_vip_order")
	}

	if !cached {
		memberVipOrderUpdateCacheMut.Lock()
		memberVipOrderUpdateCache[key] = cache
		memberVipOrderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q memberVipOrderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for member_vip_order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for member_vip_order")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MemberVipOrderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberVipOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `member_vip_order` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberVipOrderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in memberVipOrder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all memberVipOrder")
	}
	return rowsAff, nil
}

var mySQLMemberVipOrderUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MemberVipOrder) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_vip_order provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(memberVipOrderColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMemberVipOrderUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	memberVipOrderUpsertCacheMut.RLock()
	cache, cached := memberVipOrderUpsertCache[key]
	memberVipOrderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			memberVipOrderAllColumns,
			memberVipOrderColumnsWithDefault,
			memberVipOrderColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			memberVipOrderAllColumns,
			memberVipOrderPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert member_vip_order, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`member_vip_order`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `member_vip_order` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(memberVipOrderType, memberVipOrderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(memberVipOrderType, memberVipOrderMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for member_vip_order")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberVipOrderMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(memberVipOrderType, memberVipOrderMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for member_vip_order")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for member_vip_order")
	}

CacheNoHooks:
	if !cached {
		memberVipOrderUpsertCacheMut.Lock()
		memberVipOrderUpsertCache[key] = cache
		memberVipOrderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MemberVipOrder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MemberVipOrder) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MemberVipOrder provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), memberVipOrderPrimaryKeyMapping)
	sql := "DELETE FROM `member_vip_order` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from member_vip_order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for member_vip_order")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q memberVipOrderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no memberVipOrderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from member_vip_order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_vip_order")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MemberVipOrderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(memberVipOrderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberVipOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `member_vip_order` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberVipOrderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from memberVipOrder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_vip_order")
	}

	if len(memberVipOrderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MemberVipOrder) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMemberVipOrder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MemberVipOrderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MemberVipOrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberVipOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `member_vip_order`.* FROM `member_vip_order` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberVipOrderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MemberVipOrderSlice")
	}

	*o = slice

	return nil
}

// MemberVipOrderExists checks if the MemberVipOrder row exists.
func MemberVipOrderExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `member_vip_order` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if member_vip_order exists")
	}

	return exists, nil
}
