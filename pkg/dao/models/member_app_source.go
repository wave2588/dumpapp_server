// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// MemberAppSource is an object representing the database table.
type MemberAppSource struct {
	ID          int64 `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	MemberID    int64 `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	AppSourceID int64 `boil:"app_source_id" json:"app_source_id" toml:"app_source_id" yaml:"app_source_id"`
	// ????
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// ????
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *memberAppSourceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L memberAppSourceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MemberAppSourceColumns = struct {
	ID          string
	MemberID    string
	AppSourceID string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	MemberID:    "member_id",
	AppSourceID: "app_source_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var MemberAppSourceTableColumns = struct {
	ID          string
	MemberID    string
	AppSourceID string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "member_app_source.id",
	MemberID:    "member_app_source.member_id",
	AppSourceID: "member_app_source.app_source_id",
	CreatedAt:   "member_app_source.created_at",
	UpdatedAt:   "member_app_source.updated_at",
}

// Generated where

var MemberAppSourceWhere = struct {
	ID          whereHelperint64
	MemberID    whereHelperint64
	AppSourceID whereHelperint64
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint64{field: "`member_app_source`.`id`"},
	MemberID:    whereHelperint64{field: "`member_app_source`.`member_id`"},
	AppSourceID: whereHelperint64{field: "`member_app_source`.`app_source_id`"},
	CreatedAt:   whereHelpertime_Time{field: "`member_app_source`.`created_at`"},
	UpdatedAt:   whereHelpertime_Time{field: "`member_app_source`.`updated_at`"},
}

// MemberAppSourceRels is where relationship names are stored.
var MemberAppSourceRels = struct {
}{}

// memberAppSourceR is where relationships are stored.
type memberAppSourceR struct {
}

// NewStruct creates a new relationship struct
func (*memberAppSourceR) NewStruct() *memberAppSourceR {
	return &memberAppSourceR{}
}

// memberAppSourceL is where Load methods for each relationship are stored.
type memberAppSourceL struct{}

var (
	memberAppSourceAllColumns            = []string{"id", "member_id", "app_source_id", "created_at", "updated_at"}
	memberAppSourceColumnsWithoutDefault = []string{"member_id", "app_source_id"}
	memberAppSourceColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	memberAppSourcePrimaryKeyColumns     = []string{"id"}
	memberAppSourceGeneratedColumns      = []string{}
)

type (
	// MemberAppSourceSlice is an alias for a slice of pointers to MemberAppSource.
	// This should almost always be used instead of []MemberAppSource.
	MemberAppSourceSlice []*MemberAppSource
	// MemberAppSourceHook is the signature for custom MemberAppSource hook methods
	MemberAppSourceHook func(context.Context, boil.ContextExecutor, *MemberAppSource) error

	memberAppSourceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	memberAppSourceType                 = reflect.TypeOf(&MemberAppSource{})
	memberAppSourceMapping              = queries.MakeStructMapping(memberAppSourceType)
	memberAppSourcePrimaryKeyMapping, _ = queries.BindMapping(memberAppSourceType, memberAppSourceMapping, memberAppSourcePrimaryKeyColumns)
	memberAppSourceInsertCacheMut       sync.RWMutex
	memberAppSourceInsertCache          = make(map[string]insertCache)
	memberAppSourceUpdateCacheMut       sync.RWMutex
	memberAppSourceUpdateCache          = make(map[string]updateCache)
	memberAppSourceUpsertCacheMut       sync.RWMutex
	memberAppSourceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var memberAppSourceAfterSelectHooks []MemberAppSourceHook

var memberAppSourceBeforeInsertHooks []MemberAppSourceHook
var memberAppSourceAfterInsertHooks []MemberAppSourceHook

var memberAppSourceBeforeUpdateHooks []MemberAppSourceHook
var memberAppSourceAfterUpdateHooks []MemberAppSourceHook

var memberAppSourceBeforeDeleteHooks []MemberAppSourceHook
var memberAppSourceAfterDeleteHooks []MemberAppSourceHook

var memberAppSourceBeforeUpsertHooks []MemberAppSourceHook
var memberAppSourceAfterUpsertHooks []MemberAppSourceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MemberAppSource) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MemberAppSource) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MemberAppSource) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MemberAppSource) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MemberAppSource) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MemberAppSource) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MemberAppSource) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MemberAppSource) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MemberAppSource) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAppSourceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMemberAppSourceHook registers your hook function for all future operations.
func AddMemberAppSourceHook(hookPoint boil.HookPoint, memberAppSourceHook MemberAppSourceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		memberAppSourceAfterSelectHooks = append(memberAppSourceAfterSelectHooks, memberAppSourceHook)
	case boil.BeforeInsertHook:
		memberAppSourceBeforeInsertHooks = append(memberAppSourceBeforeInsertHooks, memberAppSourceHook)
	case boil.AfterInsertHook:
		memberAppSourceAfterInsertHooks = append(memberAppSourceAfterInsertHooks, memberAppSourceHook)
	case boil.BeforeUpdateHook:
		memberAppSourceBeforeUpdateHooks = append(memberAppSourceBeforeUpdateHooks, memberAppSourceHook)
	case boil.AfterUpdateHook:
		memberAppSourceAfterUpdateHooks = append(memberAppSourceAfterUpdateHooks, memberAppSourceHook)
	case boil.BeforeDeleteHook:
		memberAppSourceBeforeDeleteHooks = append(memberAppSourceBeforeDeleteHooks, memberAppSourceHook)
	case boil.AfterDeleteHook:
		memberAppSourceAfterDeleteHooks = append(memberAppSourceAfterDeleteHooks, memberAppSourceHook)
	case boil.BeforeUpsertHook:
		memberAppSourceBeforeUpsertHooks = append(memberAppSourceBeforeUpsertHooks, memberAppSourceHook)
	case boil.AfterUpsertHook:
		memberAppSourceAfterUpsertHooks = append(memberAppSourceAfterUpsertHooks, memberAppSourceHook)
	}
}

// One returns a single memberAppSource record from the query.
func (q memberAppSourceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MemberAppSource, error) {
	o := &MemberAppSource{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for member_app_source")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MemberAppSource records from the query.
func (q memberAppSourceQuery) All(ctx context.Context, exec boil.ContextExecutor) (MemberAppSourceSlice, error) {
	var o []*MemberAppSource

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MemberAppSource slice")
	}

	if len(memberAppSourceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MemberAppSource records in the query.
func (q memberAppSourceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count member_app_source rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q memberAppSourceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if member_app_source exists")
	}

	return count > 0, nil
}

// MemberAppSources retrieves all the records using an executor.
func MemberAppSources(mods ...qm.QueryMod) memberAppSourceQuery {
	mods = append(mods, qm.From("`member_app_source`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`member_app_source`.*"})
	}

	return memberAppSourceQuery{q}
}

// FindMemberAppSource retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMemberAppSource(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MemberAppSource, error) {
	memberAppSourceObj := &MemberAppSource{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `member_app_source` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, memberAppSourceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from member_app_source")
	}

	if err = memberAppSourceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return memberAppSourceObj, err
	}

	return memberAppSourceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MemberAppSource) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_app_source provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(memberAppSourceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	memberAppSourceInsertCacheMut.RLock()
	cache, cached := memberAppSourceInsertCache[key]
	memberAppSourceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			memberAppSourceAllColumns,
			memberAppSourceColumnsWithDefault,
			memberAppSourceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(memberAppSourceType, memberAppSourceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(memberAppSourceType, memberAppSourceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `member_app_source` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `member_app_source` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `member_app_source` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, memberAppSourcePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into member_app_source")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberAppSourceMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for member_app_source")
	}

CacheNoHooks:
	if !cached {
		memberAppSourceInsertCacheMut.Lock()
		memberAppSourceInsertCache[key] = cache
		memberAppSourceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MemberAppSource.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MemberAppSource) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	memberAppSourceUpdateCacheMut.RLock()
	cache, cached := memberAppSourceUpdateCache[key]
	memberAppSourceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			memberAppSourceAllColumns,
			memberAppSourcePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update member_app_source, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `member_app_source` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, memberAppSourcePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(memberAppSourceType, memberAppSourceMapping, append(wl, memberAppSourcePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update member_app_source row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for member_app_source")
	}

	if !cached {
		memberAppSourceUpdateCacheMut.Lock()
		memberAppSourceUpdateCache[key] = cache
		memberAppSourceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q memberAppSourceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for member_app_source")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for member_app_source")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MemberAppSourceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberAppSourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `member_app_source` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberAppSourcePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in memberAppSource slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all memberAppSource")
	}
	return rowsAff, nil
}

var mySQLMemberAppSourceUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MemberAppSource) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_app_source provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(memberAppSourceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMemberAppSourceUniqueColumns, o)

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

	memberAppSourceUpsertCacheMut.RLock()
	cache, cached := memberAppSourceUpsertCache[key]
	memberAppSourceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			memberAppSourceAllColumns,
			memberAppSourceColumnsWithDefault,
			memberAppSourceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			memberAppSourceAllColumns,
			memberAppSourcePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert member_app_source, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`member_app_source`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `member_app_source` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(memberAppSourceType, memberAppSourceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(memberAppSourceType, memberAppSourceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for member_app_source")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberAppSourceMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(memberAppSourceType, memberAppSourceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for member_app_source")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for member_app_source")
	}

CacheNoHooks:
	if !cached {
		memberAppSourceUpsertCacheMut.Lock()
		memberAppSourceUpsertCache[key] = cache
		memberAppSourceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MemberAppSource record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MemberAppSource) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MemberAppSource provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), memberAppSourcePrimaryKeyMapping)
	sql := "DELETE FROM `member_app_source` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from member_app_source")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for member_app_source")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q memberAppSourceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no memberAppSourceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from member_app_source")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_app_source")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MemberAppSourceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(memberAppSourceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberAppSourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `member_app_source` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberAppSourcePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from memberAppSource slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_app_source")
	}

	if len(memberAppSourceAfterDeleteHooks) != 0 {
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
func (o *MemberAppSource) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMemberAppSource(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MemberAppSourceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MemberAppSourceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberAppSourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `member_app_source`.* FROM `member_app_source` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberAppSourcePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MemberAppSourceSlice")
	}

	*o = slice

	return nil
}

// MemberAppSourceExists checks if the MemberAppSource row exists.
func MemberAppSourceExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `member_app_source` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if member_app_source exists")
	}

	return exists, nil
}

// Exists checks if the MemberAppSource row exists.
func (o *MemberAppSource) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MemberAppSourceExists(ctx, exec, o.ID)
}
