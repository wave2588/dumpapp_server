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

	"dumpapp_server/pkg/common/enum"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MemberDownloadNumber is an object representing the database table.
type MemberDownloadNumber struct {
	ID        int64                           `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	MemberID  int64                           `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	Status    enum.MemberDownloadNumberStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	IpaID     null.Int64                      `boil:"ipa_id" json:"ipa_id,omitempty" toml:"ipa_id" yaml:"ipa_id,omitempty"`
	CreatedAt time.Time                       `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time                       `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *memberDownloadNumberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L memberDownloadNumberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MemberDownloadNumberColumns = struct {
	ID        string
	MemberID  string
	Status    string
	IpaID     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	MemberID:  "member_id",
	Status:    "status",
	IpaID:     "ipa_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelperenum_MemberDownloadNumberStatus struct{ field string }

func (w whereHelperenum_MemberDownloadNumberStatus) EQ(x enum.MemberDownloadNumberStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperenum_MemberDownloadNumberStatus) NEQ(x enum.MemberDownloadNumberStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperenum_MemberDownloadNumberStatus) LT(x enum.MemberDownloadNumberStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperenum_MemberDownloadNumberStatus) LTE(x enum.MemberDownloadNumberStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperenum_MemberDownloadNumberStatus) GT(x enum.MemberDownloadNumberStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperenum_MemberDownloadNumberStatus) GTE(x enum.MemberDownloadNumberStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MemberDownloadNumberWhere = struct {
	ID        whereHelperint64
	MemberID  whereHelperint64
	Status    whereHelperenum_MemberDownloadNumberStatus
	IpaID     whereHelpernull_Int64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`member_download_number`.`id`"},
	MemberID:  whereHelperint64{field: "`member_download_number`.`member_id`"},
	Status:    whereHelperenum_MemberDownloadNumberStatus{field: "`member_download_number`.`status`"},
	IpaID:     whereHelpernull_Int64{field: "`member_download_number`.`ipa_id`"},
	CreatedAt: whereHelpertime_Time{field: "`member_download_number`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`member_download_number`.`updated_at`"},
}

// MemberDownloadNumberRels is where relationship names are stored.
var MemberDownloadNumberRels = struct {
}{}

// memberDownloadNumberR is where relationships are stored.
type memberDownloadNumberR struct {
}

// NewStruct creates a new relationship struct
func (*memberDownloadNumberR) NewStruct() *memberDownloadNumberR {
	return &memberDownloadNumberR{}
}

// memberDownloadNumberL is where Load methods for each relationship are stored.
type memberDownloadNumberL struct{}

var (
	memberDownloadNumberAllColumns            = []string{"id", "member_id", "status", "ipa_id", "created_at", "updated_at"}
	memberDownloadNumberColumnsWithoutDefault = []string{"member_id", "status", "ipa_id"}
	memberDownloadNumberColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	memberDownloadNumberPrimaryKeyColumns     = []string{"id"}
)

type (
	// MemberDownloadNumberSlice is an alias for a slice of pointers to MemberDownloadNumber.
	// This should almost always be used instead of []MemberDownloadNumber.
	MemberDownloadNumberSlice []*MemberDownloadNumber
	// MemberDownloadNumberHook is the signature for custom MemberDownloadNumber hook methods
	MemberDownloadNumberHook func(context.Context, boil.ContextExecutor, *MemberDownloadNumber) error

	memberDownloadNumberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	memberDownloadNumberType                 = reflect.TypeOf(&MemberDownloadNumber{})
	memberDownloadNumberMapping              = queries.MakeStructMapping(memberDownloadNumberType)
	memberDownloadNumberPrimaryKeyMapping, _ = queries.BindMapping(memberDownloadNumberType, memberDownloadNumberMapping, memberDownloadNumberPrimaryKeyColumns)
	memberDownloadNumberInsertCacheMut       sync.RWMutex
	memberDownloadNumberInsertCache          = make(map[string]insertCache)
	memberDownloadNumberUpdateCacheMut       sync.RWMutex
	memberDownloadNumberUpdateCache          = make(map[string]updateCache)
	memberDownloadNumberUpsertCacheMut       sync.RWMutex
	memberDownloadNumberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var memberDownloadNumberBeforeInsertHooks []MemberDownloadNumberHook
var memberDownloadNumberBeforeUpdateHooks []MemberDownloadNumberHook
var memberDownloadNumberBeforeDeleteHooks []MemberDownloadNumberHook
var memberDownloadNumberBeforeUpsertHooks []MemberDownloadNumberHook

var memberDownloadNumberAfterInsertHooks []MemberDownloadNumberHook
var memberDownloadNumberAfterSelectHooks []MemberDownloadNumberHook
var memberDownloadNumberAfterUpdateHooks []MemberDownloadNumberHook
var memberDownloadNumberAfterDeleteHooks []MemberDownloadNumberHook
var memberDownloadNumberAfterUpsertHooks []MemberDownloadNumberHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MemberDownloadNumber) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MemberDownloadNumber) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MemberDownloadNumber) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MemberDownloadNumber) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MemberDownloadNumber) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MemberDownloadNumber) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MemberDownloadNumber) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MemberDownloadNumber) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MemberDownloadNumber) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberDownloadNumberAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMemberDownloadNumberHook registers your hook function for all future operations.
func AddMemberDownloadNumberHook(hookPoint boil.HookPoint, memberDownloadNumberHook MemberDownloadNumberHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		memberDownloadNumberBeforeInsertHooks = append(memberDownloadNumberBeforeInsertHooks, memberDownloadNumberHook)
	case boil.BeforeUpdateHook:
		memberDownloadNumberBeforeUpdateHooks = append(memberDownloadNumberBeforeUpdateHooks, memberDownloadNumberHook)
	case boil.BeforeDeleteHook:
		memberDownloadNumberBeforeDeleteHooks = append(memberDownloadNumberBeforeDeleteHooks, memberDownloadNumberHook)
	case boil.BeforeUpsertHook:
		memberDownloadNumberBeforeUpsertHooks = append(memberDownloadNumberBeforeUpsertHooks, memberDownloadNumberHook)
	case boil.AfterInsertHook:
		memberDownloadNumberAfterInsertHooks = append(memberDownloadNumberAfterInsertHooks, memberDownloadNumberHook)
	case boil.AfterSelectHook:
		memberDownloadNumberAfterSelectHooks = append(memberDownloadNumberAfterSelectHooks, memberDownloadNumberHook)
	case boil.AfterUpdateHook:
		memberDownloadNumberAfterUpdateHooks = append(memberDownloadNumberAfterUpdateHooks, memberDownloadNumberHook)
	case boil.AfterDeleteHook:
		memberDownloadNumberAfterDeleteHooks = append(memberDownloadNumberAfterDeleteHooks, memberDownloadNumberHook)
	case boil.AfterUpsertHook:
		memberDownloadNumberAfterUpsertHooks = append(memberDownloadNumberAfterUpsertHooks, memberDownloadNumberHook)
	}
}

// One returns a single memberDownloadNumber record from the query.
func (q memberDownloadNumberQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MemberDownloadNumber, error) {
	o := &MemberDownloadNumber{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for member_download_number")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MemberDownloadNumber records from the query.
func (q memberDownloadNumberQuery) All(ctx context.Context, exec boil.ContextExecutor) (MemberDownloadNumberSlice, error) {
	var o []*MemberDownloadNumber

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MemberDownloadNumber slice")
	}

	if len(memberDownloadNumberAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MemberDownloadNumber records in the query.
func (q memberDownloadNumberQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count member_download_number rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q memberDownloadNumberQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if member_download_number exists")
	}

	return count > 0, nil
}

// MemberDownloadNumbers retrieves all the records using an executor.
func MemberDownloadNumbers(mods ...qm.QueryMod) memberDownloadNumberQuery {
	mods = append(mods, qm.From("`member_download_number`"))
	return memberDownloadNumberQuery{NewQuery(mods...)}
}

// FindMemberDownloadNumber retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMemberDownloadNumber(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MemberDownloadNumber, error) {
	memberDownloadNumberObj := &MemberDownloadNumber{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `member_download_number` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, memberDownloadNumberObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from member_download_number")
	}

	if err = memberDownloadNumberObj.doAfterSelectHooks(ctx, exec); err != nil {
		return memberDownloadNumberObj, err
	}

	return memberDownloadNumberObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MemberDownloadNumber) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_download_number provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(memberDownloadNumberColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	memberDownloadNumberInsertCacheMut.RLock()
	cache, cached := memberDownloadNumberInsertCache[key]
	memberDownloadNumberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			memberDownloadNumberAllColumns,
			memberDownloadNumberColumnsWithDefault,
			memberDownloadNumberColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(memberDownloadNumberType, memberDownloadNumberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(memberDownloadNumberType, memberDownloadNumberMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `member_download_number` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `member_download_number` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `member_download_number` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, memberDownloadNumberPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into member_download_number")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberDownloadNumberMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for member_download_number")
	}

CacheNoHooks:
	if !cached {
		memberDownloadNumberInsertCacheMut.Lock()
		memberDownloadNumberInsertCache[key] = cache
		memberDownloadNumberInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MemberDownloadNumber.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MemberDownloadNumber) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	memberDownloadNumberUpdateCacheMut.RLock()
	cache, cached := memberDownloadNumberUpdateCache[key]
	memberDownloadNumberUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			memberDownloadNumberAllColumns,
			memberDownloadNumberPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update member_download_number, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `member_download_number` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, memberDownloadNumberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(memberDownloadNumberType, memberDownloadNumberMapping, append(wl, memberDownloadNumberPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update member_download_number row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for member_download_number")
	}

	if !cached {
		memberDownloadNumberUpdateCacheMut.Lock()
		memberDownloadNumberUpdateCache[key] = cache
		memberDownloadNumberUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q memberDownloadNumberQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for member_download_number")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for member_download_number")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MemberDownloadNumberSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberDownloadNumberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `member_download_number` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberDownloadNumberPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in memberDownloadNumber slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all memberDownloadNumber")
	}
	return rowsAff, nil
}

var mySQLMemberDownloadNumberUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MemberDownloadNumber) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_download_number provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(memberDownloadNumberColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMemberDownloadNumberUniqueColumns, o)

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

	memberDownloadNumberUpsertCacheMut.RLock()
	cache, cached := memberDownloadNumberUpsertCache[key]
	memberDownloadNumberUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			memberDownloadNumberAllColumns,
			memberDownloadNumberColumnsWithDefault,
			memberDownloadNumberColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			memberDownloadNumberAllColumns,
			memberDownloadNumberPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert member_download_number, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`member_download_number`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `member_download_number` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(memberDownloadNumberType, memberDownloadNumberMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(memberDownloadNumberType, memberDownloadNumberMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for member_download_number")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberDownloadNumberMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(memberDownloadNumberType, memberDownloadNumberMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for member_download_number")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for member_download_number")
	}

CacheNoHooks:
	if !cached {
		memberDownloadNumberUpsertCacheMut.Lock()
		memberDownloadNumberUpsertCache[key] = cache
		memberDownloadNumberUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MemberDownloadNumber record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MemberDownloadNumber) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MemberDownloadNumber provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), memberDownloadNumberPrimaryKeyMapping)
	sql := "DELETE FROM `member_download_number` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from member_download_number")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for member_download_number")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q memberDownloadNumberQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no memberDownloadNumberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from member_download_number")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_download_number")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MemberDownloadNumberSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(memberDownloadNumberBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberDownloadNumberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `member_download_number` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberDownloadNumberPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from memberDownloadNumber slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_download_number")
	}

	if len(memberDownloadNumberAfterDeleteHooks) != 0 {
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
func (o *MemberDownloadNumber) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMemberDownloadNumber(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MemberDownloadNumberSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MemberDownloadNumberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberDownloadNumberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `member_download_number`.* FROM `member_download_number` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberDownloadNumberPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MemberDownloadNumberSlice")
	}

	*o = slice

	return nil
}

// MemberDownloadNumberExists checks if the MemberDownloadNumber row exists.
func MemberDownloadNumberExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `member_download_number` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if member_download_number exists")
	}

	return exists, nil
}
