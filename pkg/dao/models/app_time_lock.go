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

	"dumpapp_server/pkg/common/datatype"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AppTimeLock is an object representing the database table.
type AppTimeLock struct {
	ID        int64                      `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	MemberID  int64                      `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	IsDelete  bool                       `boil:"is_delete" json:"is_delete" toml:"is_delete" yaml:"is_delete"`
	IsStop    bool                       `boil:"is_stop" json:"is_stop" toml:"is_stop" yaml:"is_stop"`
	StartAt   time.Time                  `boil:"start_at" json:"start_at" toml:"start_at" yaml:"start_at"`
	EndAt     time.Time                  `boil:"end_at" json:"end_at" toml:"end_at" yaml:"end_at"`
	BizExt    datatype.AppTimeLockBizExt `boil:"biz_ext" json:"biz_ext" toml:"biz_ext" yaml:"biz_ext"`
	CreatedAt time.Time                  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time                  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *appTimeLockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L appTimeLockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AppTimeLockColumns = struct {
	ID        string
	MemberID  string
	IsDelete  string
	IsStop    string
	StartAt   string
	EndAt     string
	BizExt    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	MemberID:  "member_id",
	IsDelete:  "is_delete",
	IsStop:    "is_stop",
	StartAt:   "start_at",
	EndAt:     "end_at",
	BizExt:    "biz_ext",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperdatatype_AppTimeLockBizExt struct{ field string }

func (w whereHelperdatatype_AppTimeLockBizExt) EQ(x datatype.AppTimeLockBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperdatatype_AppTimeLockBizExt) NEQ(x datatype.AppTimeLockBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperdatatype_AppTimeLockBizExt) LT(x datatype.AppTimeLockBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperdatatype_AppTimeLockBizExt) LTE(x datatype.AppTimeLockBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperdatatype_AppTimeLockBizExt) GT(x datatype.AppTimeLockBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperdatatype_AppTimeLockBizExt) GTE(x datatype.AppTimeLockBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AppTimeLockWhere = struct {
	ID        whereHelperint64
	MemberID  whereHelperint64
	IsDelete  whereHelperbool
	IsStop    whereHelperbool
	StartAt   whereHelpertime_Time
	EndAt     whereHelpertime_Time
	BizExt    whereHelperdatatype_AppTimeLockBizExt
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`app_time_lock`.`id`"},
	MemberID:  whereHelperint64{field: "`app_time_lock`.`member_id`"},
	IsDelete:  whereHelperbool{field: "`app_time_lock`.`is_delete`"},
	IsStop:    whereHelperbool{field: "`app_time_lock`.`is_stop`"},
	StartAt:   whereHelpertime_Time{field: "`app_time_lock`.`start_at`"},
	EndAt:     whereHelpertime_Time{field: "`app_time_lock`.`end_at`"},
	BizExt:    whereHelperdatatype_AppTimeLockBizExt{field: "`app_time_lock`.`biz_ext`"},
	CreatedAt: whereHelpertime_Time{field: "`app_time_lock`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`app_time_lock`.`updated_at`"},
}

// AppTimeLockRels is where relationship names are stored.
var AppTimeLockRels = struct {
}{}

// appTimeLockR is where relationships are stored.
type appTimeLockR struct {
}

// NewStruct creates a new relationship struct
func (*appTimeLockR) NewStruct() *appTimeLockR {
	return &appTimeLockR{}
}

// appTimeLockL is where Load methods for each relationship are stored.
type appTimeLockL struct{}

var (
	appTimeLockAllColumns            = []string{"id", "member_id", "is_delete", "is_stop", "start_at", "end_at", "biz_ext", "created_at", "updated_at"}
	appTimeLockColumnsWithoutDefault = []string{"member_id", "biz_ext"}
	appTimeLockColumnsWithDefault    = []string{"id", "is_delete", "is_stop", "start_at", "end_at", "created_at", "updated_at"}
	appTimeLockPrimaryKeyColumns     = []string{"id"}
)

type (
	// AppTimeLockSlice is an alias for a slice of pointers to AppTimeLock.
	// This should almost always be used instead of []AppTimeLock.
	AppTimeLockSlice []*AppTimeLock
	// AppTimeLockHook is the signature for custom AppTimeLock hook methods
	AppTimeLockHook func(context.Context, boil.ContextExecutor, *AppTimeLock) error

	appTimeLockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	appTimeLockType                 = reflect.TypeOf(&AppTimeLock{})
	appTimeLockMapping              = queries.MakeStructMapping(appTimeLockType)
	appTimeLockPrimaryKeyMapping, _ = queries.BindMapping(appTimeLockType, appTimeLockMapping, appTimeLockPrimaryKeyColumns)
	appTimeLockInsertCacheMut       sync.RWMutex
	appTimeLockInsertCache          = make(map[string]insertCache)
	appTimeLockUpdateCacheMut       sync.RWMutex
	appTimeLockUpdateCache          = make(map[string]updateCache)
	appTimeLockUpsertCacheMut       sync.RWMutex
	appTimeLockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var appTimeLockBeforeInsertHooks []AppTimeLockHook
var appTimeLockBeforeUpdateHooks []AppTimeLockHook
var appTimeLockBeforeDeleteHooks []AppTimeLockHook
var appTimeLockBeforeUpsertHooks []AppTimeLockHook

var appTimeLockAfterInsertHooks []AppTimeLockHook
var appTimeLockAfterSelectHooks []AppTimeLockHook
var appTimeLockAfterUpdateHooks []AppTimeLockHook
var appTimeLockAfterDeleteHooks []AppTimeLockHook
var appTimeLockAfterUpsertHooks []AppTimeLockHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AppTimeLock) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AppTimeLock) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AppTimeLock) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AppTimeLock) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AppTimeLock) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AppTimeLock) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AppTimeLock) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AppTimeLock) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AppTimeLock) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range appTimeLockAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAppTimeLockHook registers your hook function for all future operations.
func AddAppTimeLockHook(hookPoint boil.HookPoint, appTimeLockHook AppTimeLockHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		appTimeLockBeforeInsertHooks = append(appTimeLockBeforeInsertHooks, appTimeLockHook)
	case boil.BeforeUpdateHook:
		appTimeLockBeforeUpdateHooks = append(appTimeLockBeforeUpdateHooks, appTimeLockHook)
	case boil.BeforeDeleteHook:
		appTimeLockBeforeDeleteHooks = append(appTimeLockBeforeDeleteHooks, appTimeLockHook)
	case boil.BeforeUpsertHook:
		appTimeLockBeforeUpsertHooks = append(appTimeLockBeforeUpsertHooks, appTimeLockHook)
	case boil.AfterInsertHook:
		appTimeLockAfterInsertHooks = append(appTimeLockAfterInsertHooks, appTimeLockHook)
	case boil.AfterSelectHook:
		appTimeLockAfterSelectHooks = append(appTimeLockAfterSelectHooks, appTimeLockHook)
	case boil.AfterUpdateHook:
		appTimeLockAfterUpdateHooks = append(appTimeLockAfterUpdateHooks, appTimeLockHook)
	case boil.AfterDeleteHook:
		appTimeLockAfterDeleteHooks = append(appTimeLockAfterDeleteHooks, appTimeLockHook)
	case boil.AfterUpsertHook:
		appTimeLockAfterUpsertHooks = append(appTimeLockAfterUpsertHooks, appTimeLockHook)
	}
}

// One returns a single appTimeLock record from the query.
func (q appTimeLockQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AppTimeLock, error) {
	o := &AppTimeLock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for app_time_lock")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AppTimeLock records from the query.
func (q appTimeLockQuery) All(ctx context.Context, exec boil.ContextExecutor) (AppTimeLockSlice, error) {
	var o []*AppTimeLock

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AppTimeLock slice")
	}

	if len(appTimeLockAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AppTimeLock records in the query.
func (q appTimeLockQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count app_time_lock rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q appTimeLockQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if app_time_lock exists")
	}

	return count > 0, nil
}

// AppTimeLocks retrieves all the records using an executor.
func AppTimeLocks(mods ...qm.QueryMod) appTimeLockQuery {
	mods = append(mods, qm.From("`app_time_lock`"))
	return appTimeLockQuery{NewQuery(mods...)}
}

// FindAppTimeLock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAppTimeLock(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AppTimeLock, error) {
	appTimeLockObj := &AppTimeLock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `app_time_lock` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, appTimeLockObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from app_time_lock")
	}

	if err = appTimeLockObj.doAfterSelectHooks(ctx, exec); err != nil {
		return appTimeLockObj, err
	}

	return appTimeLockObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AppTimeLock) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no app_time_lock provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(appTimeLockColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	appTimeLockInsertCacheMut.RLock()
	cache, cached := appTimeLockInsertCache[key]
	appTimeLockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			appTimeLockAllColumns,
			appTimeLockColumnsWithDefault,
			appTimeLockColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(appTimeLockType, appTimeLockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(appTimeLockType, appTimeLockMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `app_time_lock` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `app_time_lock` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `app_time_lock` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, appTimeLockPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into app_time_lock")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == appTimeLockMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for app_time_lock")
	}

CacheNoHooks:
	if !cached {
		appTimeLockInsertCacheMut.Lock()
		appTimeLockInsertCache[key] = cache
		appTimeLockInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AppTimeLock.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AppTimeLock) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	appTimeLockUpdateCacheMut.RLock()
	cache, cached := appTimeLockUpdateCache[key]
	appTimeLockUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			appTimeLockAllColumns,
			appTimeLockPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update app_time_lock, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `app_time_lock` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, appTimeLockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(appTimeLockType, appTimeLockMapping, append(wl, appTimeLockPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update app_time_lock row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for app_time_lock")
	}

	if !cached {
		appTimeLockUpdateCacheMut.Lock()
		appTimeLockUpdateCache[key] = cache
		appTimeLockUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q appTimeLockQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for app_time_lock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for app_time_lock")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AppTimeLockSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appTimeLockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `app_time_lock` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, appTimeLockPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in appTimeLock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all appTimeLock")
	}
	return rowsAff, nil
}

var mySQLAppTimeLockUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AppTimeLock) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no app_time_lock provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(appTimeLockColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAppTimeLockUniqueColumns, o)

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

	appTimeLockUpsertCacheMut.RLock()
	cache, cached := appTimeLockUpsertCache[key]
	appTimeLockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			appTimeLockAllColumns,
			appTimeLockColumnsWithDefault,
			appTimeLockColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			appTimeLockAllColumns,
			appTimeLockPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert app_time_lock, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`app_time_lock`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `app_time_lock` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(appTimeLockType, appTimeLockMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(appTimeLockType, appTimeLockMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for app_time_lock")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == appTimeLockMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(appTimeLockType, appTimeLockMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for app_time_lock")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for app_time_lock")
	}

CacheNoHooks:
	if !cached {
		appTimeLockUpsertCacheMut.Lock()
		appTimeLockUpsertCache[key] = cache
		appTimeLockUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AppTimeLock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AppTimeLock) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AppTimeLock provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), appTimeLockPrimaryKeyMapping)
	sql := "DELETE FROM `app_time_lock` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from app_time_lock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for app_time_lock")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q appTimeLockQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no appTimeLockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from app_time_lock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for app_time_lock")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AppTimeLockSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(appTimeLockBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appTimeLockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `app_time_lock` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, appTimeLockPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from appTimeLock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for app_time_lock")
	}

	if len(appTimeLockAfterDeleteHooks) != 0 {
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
func (o *AppTimeLock) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAppTimeLock(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AppTimeLockSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AppTimeLockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appTimeLockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `app_time_lock`.* FROM `app_time_lock` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, appTimeLockPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AppTimeLockSlice")
	}

	*o = slice

	return nil
}

// AppTimeLockExists checks if the AppTimeLock row exists.
func AppTimeLockExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `app_time_lock` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if app_time_lock exists")
	}

	return exists, nil
}
