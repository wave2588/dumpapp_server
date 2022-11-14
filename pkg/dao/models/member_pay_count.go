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

	"dumpapp_server/pkg/common/enum"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MemberPayCount is an object representing the database table.
type MemberPayCount struct {
	ID       int64                     `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	MemberID int64                     `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	Status   enum.MemberPayCountStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	Source   enum.MemberPayCountSource `boil:"source" json:"source" toml:"source" yaml:"source"`
	Use      null.String               `boil:"use" json:"use,omitempty" toml:"use" yaml:"use,omitempty"`
	// ????
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// ????
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *memberPayCountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L memberPayCountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MemberPayCountColumns = struct {
	ID        string
	MemberID  string
	Status    string
	Source    string
	Use       string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	MemberID:  "member_id",
	Status:    "status",
	Source:    "source",
	Use:       "use",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var MemberPayCountTableColumns = struct {
	ID        string
	MemberID  string
	Status    string
	Source    string
	Use       string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "member_pay_count.id",
	MemberID:  "member_pay_count.member_id",
	Status:    "member_pay_count.status",
	Source:    "member_pay_count.source",
	Use:       "member_pay_count.use",
	CreatedAt: "member_pay_count.created_at",
	UpdatedAt: "member_pay_count.updated_at",
}

// Generated where

type whereHelperenum_MemberPayCountStatus struct{ field string }

func (w whereHelperenum_MemberPayCountStatus) EQ(x enum.MemberPayCountStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperenum_MemberPayCountStatus) NEQ(x enum.MemberPayCountStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperenum_MemberPayCountStatus) LT(x enum.MemberPayCountStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperenum_MemberPayCountStatus) LTE(x enum.MemberPayCountStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperenum_MemberPayCountStatus) GT(x enum.MemberPayCountStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperenum_MemberPayCountStatus) GTE(x enum.MemberPayCountStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperenum_MemberPayCountSource struct{ field string }

func (w whereHelperenum_MemberPayCountSource) EQ(x enum.MemberPayCountSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperenum_MemberPayCountSource) NEQ(x enum.MemberPayCountSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperenum_MemberPayCountSource) LT(x enum.MemberPayCountSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperenum_MemberPayCountSource) LTE(x enum.MemberPayCountSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperenum_MemberPayCountSource) GT(x enum.MemberPayCountSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperenum_MemberPayCountSource) GTE(x enum.MemberPayCountSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MemberPayCountWhere = struct {
	ID        whereHelperint64
	MemberID  whereHelperint64
	Status    whereHelperenum_MemberPayCountStatus
	Source    whereHelperenum_MemberPayCountSource
	Use       whereHelpernull_String
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`member_pay_count`.`id`"},
	MemberID:  whereHelperint64{field: "`member_pay_count`.`member_id`"},
	Status:    whereHelperenum_MemberPayCountStatus{field: "`member_pay_count`.`status`"},
	Source:    whereHelperenum_MemberPayCountSource{field: "`member_pay_count`.`source`"},
	Use:       whereHelpernull_String{field: "`member_pay_count`.`use`"},
	CreatedAt: whereHelpertime_Time{field: "`member_pay_count`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`member_pay_count`.`updated_at`"},
}

// MemberPayCountRels is where relationship names are stored.
var MemberPayCountRels = struct {
}{}

// memberPayCountR is where relationships are stored.
type memberPayCountR struct {
}

// NewStruct creates a new relationship struct
func (*memberPayCountR) NewStruct() *memberPayCountR {
	return &memberPayCountR{}
}

// memberPayCountL is where Load methods for each relationship are stored.
type memberPayCountL struct{}

var (
	memberPayCountAllColumns            = []string{"id", "member_id", "status", "source", "use", "created_at", "updated_at"}
	memberPayCountColumnsWithoutDefault = []string{"member_id", "status", "source", "use"}
	memberPayCountColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	memberPayCountPrimaryKeyColumns     = []string{"id"}
	memberPayCountGeneratedColumns      = []string{}
)

type (
	// MemberPayCountSlice is an alias for a slice of pointers to MemberPayCount.
	// This should almost always be used instead of []MemberPayCount.
	MemberPayCountSlice []*MemberPayCount
	// MemberPayCountHook is the signature for custom MemberPayCount hook methods
	MemberPayCountHook func(context.Context, boil.ContextExecutor, *MemberPayCount) error

	memberPayCountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	memberPayCountType                 = reflect.TypeOf(&MemberPayCount{})
	memberPayCountMapping              = queries.MakeStructMapping(memberPayCountType)
	memberPayCountPrimaryKeyMapping, _ = queries.BindMapping(memberPayCountType, memberPayCountMapping, memberPayCountPrimaryKeyColumns)
	memberPayCountInsertCacheMut       sync.RWMutex
	memberPayCountInsertCache          = make(map[string]insertCache)
	memberPayCountUpdateCacheMut       sync.RWMutex
	memberPayCountUpdateCache          = make(map[string]updateCache)
	memberPayCountUpsertCacheMut       sync.RWMutex
	memberPayCountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var memberPayCountAfterSelectHooks []MemberPayCountHook

var memberPayCountBeforeInsertHooks []MemberPayCountHook
var memberPayCountAfterInsertHooks []MemberPayCountHook

var memberPayCountBeforeUpdateHooks []MemberPayCountHook
var memberPayCountAfterUpdateHooks []MemberPayCountHook

var memberPayCountBeforeDeleteHooks []MemberPayCountHook
var memberPayCountAfterDeleteHooks []MemberPayCountHook

var memberPayCountBeforeUpsertHooks []MemberPayCountHook
var memberPayCountAfterUpsertHooks []MemberPayCountHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MemberPayCount) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MemberPayCount) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MemberPayCount) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MemberPayCount) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MemberPayCount) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MemberPayCount) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MemberPayCount) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MemberPayCount) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MemberPayCount) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberPayCountAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMemberPayCountHook registers your hook function for all future operations.
func AddMemberPayCountHook(hookPoint boil.HookPoint, memberPayCountHook MemberPayCountHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		memberPayCountAfterSelectHooks = append(memberPayCountAfterSelectHooks, memberPayCountHook)
	case boil.BeforeInsertHook:
		memberPayCountBeforeInsertHooks = append(memberPayCountBeforeInsertHooks, memberPayCountHook)
	case boil.AfterInsertHook:
		memberPayCountAfterInsertHooks = append(memberPayCountAfterInsertHooks, memberPayCountHook)
	case boil.BeforeUpdateHook:
		memberPayCountBeforeUpdateHooks = append(memberPayCountBeforeUpdateHooks, memberPayCountHook)
	case boil.AfterUpdateHook:
		memberPayCountAfterUpdateHooks = append(memberPayCountAfterUpdateHooks, memberPayCountHook)
	case boil.BeforeDeleteHook:
		memberPayCountBeforeDeleteHooks = append(memberPayCountBeforeDeleteHooks, memberPayCountHook)
	case boil.AfterDeleteHook:
		memberPayCountAfterDeleteHooks = append(memberPayCountAfterDeleteHooks, memberPayCountHook)
	case boil.BeforeUpsertHook:
		memberPayCountBeforeUpsertHooks = append(memberPayCountBeforeUpsertHooks, memberPayCountHook)
	case boil.AfterUpsertHook:
		memberPayCountAfterUpsertHooks = append(memberPayCountAfterUpsertHooks, memberPayCountHook)
	}
}

// One returns a single memberPayCount record from the query.
func (q memberPayCountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MemberPayCount, error) {
	o := &MemberPayCount{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for member_pay_count")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MemberPayCount records from the query.
func (q memberPayCountQuery) All(ctx context.Context, exec boil.ContextExecutor) (MemberPayCountSlice, error) {
	var o []*MemberPayCount

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MemberPayCount slice")
	}

	if len(memberPayCountAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MemberPayCount records in the query.
func (q memberPayCountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count member_pay_count rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q memberPayCountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if member_pay_count exists")
	}

	return count > 0, nil
}

// MemberPayCounts retrieves all the records using an executor.
func MemberPayCounts(mods ...qm.QueryMod) memberPayCountQuery {
	mods = append(mods, qm.From("`member_pay_count`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`member_pay_count`.*"})
	}

	return memberPayCountQuery{q}
}

// FindMemberPayCount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMemberPayCount(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MemberPayCount, error) {
	memberPayCountObj := &MemberPayCount{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `member_pay_count` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, memberPayCountObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from member_pay_count")
	}

	if err = memberPayCountObj.doAfterSelectHooks(ctx, exec); err != nil {
		return memberPayCountObj, err
	}

	return memberPayCountObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MemberPayCount) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_pay_count provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(memberPayCountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	memberPayCountInsertCacheMut.RLock()
	cache, cached := memberPayCountInsertCache[key]
	memberPayCountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			memberPayCountAllColumns,
			memberPayCountColumnsWithDefault,
			memberPayCountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(memberPayCountType, memberPayCountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(memberPayCountType, memberPayCountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `member_pay_count` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `member_pay_count` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `member_pay_count` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, memberPayCountPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into member_pay_count")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberPayCountMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for member_pay_count")
	}

CacheNoHooks:
	if !cached {
		memberPayCountInsertCacheMut.Lock()
		memberPayCountInsertCache[key] = cache
		memberPayCountInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MemberPayCount.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MemberPayCount) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	memberPayCountUpdateCacheMut.RLock()
	cache, cached := memberPayCountUpdateCache[key]
	memberPayCountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			memberPayCountAllColumns,
			memberPayCountPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update member_pay_count, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `member_pay_count` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, memberPayCountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(memberPayCountType, memberPayCountMapping, append(wl, memberPayCountPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update member_pay_count row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for member_pay_count")
	}

	if !cached {
		memberPayCountUpdateCacheMut.Lock()
		memberPayCountUpdateCache[key] = cache
		memberPayCountUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q memberPayCountQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for member_pay_count")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for member_pay_count")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MemberPayCountSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberPayCountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `member_pay_count` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberPayCountPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in memberPayCount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all memberPayCount")
	}
	return rowsAff, nil
}

var mySQLMemberPayCountUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MemberPayCount) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_pay_count provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(memberPayCountColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMemberPayCountUniqueColumns, o)

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

	memberPayCountUpsertCacheMut.RLock()
	cache, cached := memberPayCountUpsertCache[key]
	memberPayCountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			memberPayCountAllColumns,
			memberPayCountColumnsWithDefault,
			memberPayCountColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			memberPayCountAllColumns,
			memberPayCountPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert member_pay_count, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`member_pay_count`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `member_pay_count` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(memberPayCountType, memberPayCountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(memberPayCountType, memberPayCountMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for member_pay_count")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberPayCountMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(memberPayCountType, memberPayCountMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for member_pay_count")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for member_pay_count")
	}

CacheNoHooks:
	if !cached {
		memberPayCountUpsertCacheMut.Lock()
		memberPayCountUpsertCache[key] = cache
		memberPayCountUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MemberPayCount record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MemberPayCount) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MemberPayCount provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), memberPayCountPrimaryKeyMapping)
	sql := "DELETE FROM `member_pay_count` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from member_pay_count")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for member_pay_count")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q memberPayCountQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no memberPayCountQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from member_pay_count")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_pay_count")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MemberPayCountSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(memberPayCountBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberPayCountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `member_pay_count` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberPayCountPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from memberPayCount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_pay_count")
	}

	if len(memberPayCountAfterDeleteHooks) != 0 {
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
func (o *MemberPayCount) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMemberPayCount(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MemberPayCountSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MemberPayCountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberPayCountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `member_pay_count`.* FROM `member_pay_count` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberPayCountPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MemberPayCountSlice")
	}

	*o = slice

	return nil
}

// MemberPayCountExists checks if the MemberPayCount row exists.
func MemberPayCountExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `member_pay_count` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if member_pay_count exists")
	}

	return exists, nil
}

// Exists checks if the MemberPayCount row exists.
func (o *MemberPayCount) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MemberPayCountExists(ctx, exec, o.ID)
}
