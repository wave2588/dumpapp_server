// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MemberRebateRecord is an object representing the database table.
type MemberRebateRecord struct {
	ID               int64     `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	OrderID          int64     `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	ReceiverMemberID int64     `boil:"receiver_member_id" json:"receiver_member_id" toml:"receiver_member_id" yaml:"receiver_member_id"`
	Count            int       `boil:"count" json:"count" toml:"count" yaml:"count"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *memberRebateRecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L memberRebateRecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MemberRebateRecordColumns = struct {
	ID               string
	OrderID          string
	ReceiverMemberID string
	Count            string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	OrderID:          "order_id",
	ReceiverMemberID: "receiver_member_id",
	Count:            "count",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var MemberRebateRecordWhere = struct {
	ID               whereHelperint64
	OrderID          whereHelperint64
	ReceiverMemberID whereHelperint64
	Count            whereHelperint
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
}{
	ID:               whereHelperint64{field: "`member_rebate_record`.`id`"},
	OrderID:          whereHelperint64{field: "`member_rebate_record`.`order_id`"},
	ReceiverMemberID: whereHelperint64{field: "`member_rebate_record`.`receiver_member_id`"},
	Count:            whereHelperint{field: "`member_rebate_record`.`count`"},
	CreatedAt:        whereHelpertime_Time{field: "`member_rebate_record`.`created_at`"},
	UpdatedAt:        whereHelpertime_Time{field: "`member_rebate_record`.`updated_at`"},
}

// MemberRebateRecordRels is where relationship names are stored.
var MemberRebateRecordRels = struct {
}{}

// memberRebateRecordR is where relationships are stored.
type memberRebateRecordR struct {
}

// NewStruct creates a new relationship struct
func (*memberRebateRecordR) NewStruct() *memberRebateRecordR {
	return &memberRebateRecordR{}
}

// memberRebateRecordL is where Load methods for each relationship are stored.
type memberRebateRecordL struct{}

var (
	memberRebateRecordAllColumns            = []string{"id", "order_id", "receiver_member_id", "count", "created_at", "updated_at"}
	memberRebateRecordColumnsWithoutDefault = []string{"order_id", "receiver_member_id", "count"}
	memberRebateRecordColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	memberRebateRecordPrimaryKeyColumns     = []string{"id"}
)

type (
	// MemberRebateRecordSlice is an alias for a slice of pointers to MemberRebateRecord.
	// This should almost always be used instead of []MemberRebateRecord.
	MemberRebateRecordSlice []*MemberRebateRecord
	// MemberRebateRecordHook is the signature for custom MemberRebateRecord hook methods
	MemberRebateRecordHook func(context.Context, boil.ContextExecutor, *MemberRebateRecord) error

	memberRebateRecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	memberRebateRecordType                 = reflect.TypeOf(&MemberRebateRecord{})
	memberRebateRecordMapping              = queries.MakeStructMapping(memberRebateRecordType)
	memberRebateRecordPrimaryKeyMapping, _ = queries.BindMapping(memberRebateRecordType, memberRebateRecordMapping, memberRebateRecordPrimaryKeyColumns)
	memberRebateRecordInsertCacheMut       sync.RWMutex
	memberRebateRecordInsertCache          = make(map[string]insertCache)
	memberRebateRecordUpdateCacheMut       sync.RWMutex
	memberRebateRecordUpdateCache          = make(map[string]updateCache)
	memberRebateRecordUpsertCacheMut       sync.RWMutex
	memberRebateRecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var memberRebateRecordBeforeInsertHooks []MemberRebateRecordHook
var memberRebateRecordBeforeUpdateHooks []MemberRebateRecordHook
var memberRebateRecordBeforeDeleteHooks []MemberRebateRecordHook
var memberRebateRecordBeforeUpsertHooks []MemberRebateRecordHook

var memberRebateRecordAfterInsertHooks []MemberRebateRecordHook
var memberRebateRecordAfterSelectHooks []MemberRebateRecordHook
var memberRebateRecordAfterUpdateHooks []MemberRebateRecordHook
var memberRebateRecordAfterDeleteHooks []MemberRebateRecordHook
var memberRebateRecordAfterUpsertHooks []MemberRebateRecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MemberRebateRecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MemberRebateRecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MemberRebateRecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MemberRebateRecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MemberRebateRecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MemberRebateRecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MemberRebateRecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MemberRebateRecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MemberRebateRecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberRebateRecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMemberRebateRecordHook registers your hook function for all future operations.
func AddMemberRebateRecordHook(hookPoint boil.HookPoint, memberRebateRecordHook MemberRebateRecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		memberRebateRecordBeforeInsertHooks = append(memberRebateRecordBeforeInsertHooks, memberRebateRecordHook)
	case boil.BeforeUpdateHook:
		memberRebateRecordBeforeUpdateHooks = append(memberRebateRecordBeforeUpdateHooks, memberRebateRecordHook)
	case boil.BeforeDeleteHook:
		memberRebateRecordBeforeDeleteHooks = append(memberRebateRecordBeforeDeleteHooks, memberRebateRecordHook)
	case boil.BeforeUpsertHook:
		memberRebateRecordBeforeUpsertHooks = append(memberRebateRecordBeforeUpsertHooks, memberRebateRecordHook)
	case boil.AfterInsertHook:
		memberRebateRecordAfterInsertHooks = append(memberRebateRecordAfterInsertHooks, memberRebateRecordHook)
	case boil.AfterSelectHook:
		memberRebateRecordAfterSelectHooks = append(memberRebateRecordAfterSelectHooks, memberRebateRecordHook)
	case boil.AfterUpdateHook:
		memberRebateRecordAfterUpdateHooks = append(memberRebateRecordAfterUpdateHooks, memberRebateRecordHook)
	case boil.AfterDeleteHook:
		memberRebateRecordAfterDeleteHooks = append(memberRebateRecordAfterDeleteHooks, memberRebateRecordHook)
	case boil.AfterUpsertHook:
		memberRebateRecordAfterUpsertHooks = append(memberRebateRecordAfterUpsertHooks, memberRebateRecordHook)
	}
}

// One returns a single memberRebateRecord record from the query.
func (q memberRebateRecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MemberRebateRecord, error) {
	o := &MemberRebateRecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for member_rebate_record")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MemberRebateRecord records from the query.
func (q memberRebateRecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (MemberRebateRecordSlice, error) {
	var o []*MemberRebateRecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MemberRebateRecord slice")
	}

	if len(memberRebateRecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MemberRebateRecord records in the query.
func (q memberRebateRecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count member_rebate_record rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q memberRebateRecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if member_rebate_record exists")
	}

	return count > 0, nil
}

// MemberRebateRecords retrieves all the records using an executor.
func MemberRebateRecords(mods ...qm.QueryMod) memberRebateRecordQuery {
	mods = append(mods, qm.From("`member_rebate_record`"))
	return memberRebateRecordQuery{NewQuery(mods...)}
}

// FindMemberRebateRecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMemberRebateRecord(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MemberRebateRecord, error) {
	memberRebateRecordObj := &MemberRebateRecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `member_rebate_record` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, memberRebateRecordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from member_rebate_record")
	}

	if err = memberRebateRecordObj.doAfterSelectHooks(ctx, exec); err != nil {
		return memberRebateRecordObj, err
	}

	return memberRebateRecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MemberRebateRecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_rebate_record provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(memberRebateRecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	memberRebateRecordInsertCacheMut.RLock()
	cache, cached := memberRebateRecordInsertCache[key]
	memberRebateRecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			memberRebateRecordAllColumns,
			memberRebateRecordColumnsWithDefault,
			memberRebateRecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(memberRebateRecordType, memberRebateRecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(memberRebateRecordType, memberRebateRecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `member_rebate_record` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `member_rebate_record` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `member_rebate_record` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, memberRebateRecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into member_rebate_record")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberRebateRecordMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for member_rebate_record")
	}

CacheNoHooks:
	if !cached {
		memberRebateRecordInsertCacheMut.Lock()
		memberRebateRecordInsertCache[key] = cache
		memberRebateRecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MemberRebateRecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MemberRebateRecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	memberRebateRecordUpdateCacheMut.RLock()
	cache, cached := memberRebateRecordUpdateCache[key]
	memberRebateRecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			memberRebateRecordAllColumns,
			memberRebateRecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update member_rebate_record, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `member_rebate_record` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, memberRebateRecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(memberRebateRecordType, memberRebateRecordMapping, append(wl, memberRebateRecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update member_rebate_record row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for member_rebate_record")
	}

	if !cached {
		memberRebateRecordUpdateCacheMut.Lock()
		memberRebateRecordUpdateCache[key] = cache
		memberRebateRecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q memberRebateRecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for member_rebate_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for member_rebate_record")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MemberRebateRecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberRebateRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `member_rebate_record` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberRebateRecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in memberRebateRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all memberRebateRecord")
	}
	return rowsAff, nil
}

var mySQLMemberRebateRecordUniqueColumns = []string{
	"id",
	"order_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MemberRebateRecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_rebate_record provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(memberRebateRecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMemberRebateRecordUniqueColumns, o)

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

	memberRebateRecordUpsertCacheMut.RLock()
	cache, cached := memberRebateRecordUpsertCache[key]
	memberRebateRecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			memberRebateRecordAllColumns,
			memberRebateRecordColumnsWithDefault,
			memberRebateRecordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			memberRebateRecordAllColumns,
			memberRebateRecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert member_rebate_record, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`member_rebate_record`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `member_rebate_record` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(memberRebateRecordType, memberRebateRecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(memberRebateRecordType, memberRebateRecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for member_rebate_record")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberRebateRecordMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(memberRebateRecordType, memberRebateRecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for member_rebate_record")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for member_rebate_record")
	}

CacheNoHooks:
	if !cached {
		memberRebateRecordUpsertCacheMut.Lock()
		memberRebateRecordUpsertCache[key] = cache
		memberRebateRecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MemberRebateRecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MemberRebateRecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MemberRebateRecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), memberRebateRecordPrimaryKeyMapping)
	sql := "DELETE FROM `member_rebate_record` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from member_rebate_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for member_rebate_record")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q memberRebateRecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no memberRebateRecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from member_rebate_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_rebate_record")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MemberRebateRecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(memberRebateRecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberRebateRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `member_rebate_record` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberRebateRecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from memberRebateRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_rebate_record")
	}

	if len(memberRebateRecordAfterDeleteHooks) != 0 {
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
func (o *MemberRebateRecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMemberRebateRecord(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MemberRebateRecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MemberRebateRecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberRebateRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `member_rebate_record`.* FROM `member_rebate_record` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberRebateRecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MemberRebateRecordSlice")
	}

	*o = slice

	return nil
}

// MemberRebateRecordExists checks if the MemberRebateRecord row exists.
func MemberRebateRecordExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `member_rebate_record` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if member_rebate_record exists")
	}

	return exists, nil
}
