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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// IpaVersion is an object representing the database table.
type IpaVersion struct {
	ID          int64        `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	IpaID       int64        `boil:"ipa_id" json:"ipa_id" toml:"ipa_id" yaml:"ipa_id"`
	Version     string       `boil:"version" json:"version" toml:"version" yaml:"version"`
	IpaType     enum.IpaType `boil:"ipa_type" json:"ipa_type" toml:"ipa_type" yaml:"ipa_type"`
	TokenPath   string       `boil:"token_path" json:"token_path" toml:"token_path" yaml:"token_path"`
	BizExt      string       `boil:"biz_ext" json:"biz_ext" toml:"biz_ext" yaml:"biz_ext"`
	IsTemporary int64        `boil:"is_temporary" json:"is_temporary" toml:"is_temporary" yaml:"is_temporary"`
	CreatedAt   time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *ipaVersionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ipaVersionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IpaVersionColumns = struct {
	ID          string
	IpaID       string
	Version     string
	IpaType     string
	TokenPath   string
	BizExt      string
	IsTemporary string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	IpaID:       "ipa_id",
	Version:     "version",
	IpaType:     "ipa_type",
	TokenPath:   "token_path",
	BizExt:      "biz_ext",
	IsTemporary: "is_temporary",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// Generated where

type whereHelperenum_IpaType struct{ field string }

func (w whereHelperenum_IpaType) EQ(x enum.IpaType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperenum_IpaType) NEQ(x enum.IpaType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperenum_IpaType) LT(x enum.IpaType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperenum_IpaType) LTE(x enum.IpaType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperenum_IpaType) GT(x enum.IpaType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperenum_IpaType) GTE(x enum.IpaType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var IpaVersionWhere = struct {
	ID          whereHelperint64
	IpaID       whereHelperint64
	Version     whereHelperstring
	IpaType     whereHelperenum_IpaType
	TokenPath   whereHelperstring
	BizExt      whereHelperstring
	IsTemporary whereHelperint64
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint64{field: "`ipa_version`.`id`"},
	IpaID:       whereHelperint64{field: "`ipa_version`.`ipa_id`"},
	Version:     whereHelperstring{field: "`ipa_version`.`version`"},
	IpaType:     whereHelperenum_IpaType{field: "`ipa_version`.`ipa_type`"},
	TokenPath:   whereHelperstring{field: "`ipa_version`.`token_path`"},
	BizExt:      whereHelperstring{field: "`ipa_version`.`biz_ext`"},
	IsTemporary: whereHelperint64{field: "`ipa_version`.`is_temporary`"},
	CreatedAt:   whereHelpertime_Time{field: "`ipa_version`.`created_at`"},
	UpdatedAt:   whereHelpertime_Time{field: "`ipa_version`.`updated_at`"},
}

// IpaVersionRels is where relationship names are stored.
var IpaVersionRels = struct {
}{}

// ipaVersionR is where relationships are stored.
type ipaVersionR struct {
}

// NewStruct creates a new relationship struct
func (*ipaVersionR) NewStruct() *ipaVersionR {
	return &ipaVersionR{}
}

// ipaVersionL is where Load methods for each relationship are stored.
type ipaVersionL struct{}

var (
	ipaVersionAllColumns            = []string{"id", "ipa_id", "version", "ipa_type", "token_path", "biz_ext", "is_temporary", "created_at", "updated_at"}
	ipaVersionColumnsWithoutDefault = []string{"ipa_id", "version", "ipa_type", "token_path", "biz_ext", "is_temporary"}
	ipaVersionColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	ipaVersionPrimaryKeyColumns     = []string{"id"}
)

type (
	// IpaVersionSlice is an alias for a slice of pointers to IpaVersion.
	// This should almost always be used instead of []IpaVersion.
	IpaVersionSlice []*IpaVersion
	// IpaVersionHook is the signature for custom IpaVersion hook methods
	IpaVersionHook func(context.Context, boil.ContextExecutor, *IpaVersion) error

	ipaVersionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ipaVersionType                 = reflect.TypeOf(&IpaVersion{})
	ipaVersionMapping              = queries.MakeStructMapping(ipaVersionType)
	ipaVersionPrimaryKeyMapping, _ = queries.BindMapping(ipaVersionType, ipaVersionMapping, ipaVersionPrimaryKeyColumns)
	ipaVersionInsertCacheMut       sync.RWMutex
	ipaVersionInsertCache          = make(map[string]insertCache)
	ipaVersionUpdateCacheMut       sync.RWMutex
	ipaVersionUpdateCache          = make(map[string]updateCache)
	ipaVersionUpsertCacheMut       sync.RWMutex
	ipaVersionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ipaVersionBeforeInsertHooks []IpaVersionHook
var ipaVersionBeforeUpdateHooks []IpaVersionHook
var ipaVersionBeforeDeleteHooks []IpaVersionHook
var ipaVersionBeforeUpsertHooks []IpaVersionHook

var ipaVersionAfterInsertHooks []IpaVersionHook
var ipaVersionAfterSelectHooks []IpaVersionHook
var ipaVersionAfterUpdateHooks []IpaVersionHook
var ipaVersionAfterDeleteHooks []IpaVersionHook
var ipaVersionAfterUpsertHooks []IpaVersionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *IpaVersion) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *IpaVersion) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *IpaVersion) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *IpaVersion) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *IpaVersion) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *IpaVersion) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *IpaVersion) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *IpaVersion) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *IpaVersion) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ipaVersionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIpaVersionHook registers your hook function for all future operations.
func AddIpaVersionHook(hookPoint boil.HookPoint, ipaVersionHook IpaVersionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ipaVersionBeforeInsertHooks = append(ipaVersionBeforeInsertHooks, ipaVersionHook)
	case boil.BeforeUpdateHook:
		ipaVersionBeforeUpdateHooks = append(ipaVersionBeforeUpdateHooks, ipaVersionHook)
	case boil.BeforeDeleteHook:
		ipaVersionBeforeDeleteHooks = append(ipaVersionBeforeDeleteHooks, ipaVersionHook)
	case boil.BeforeUpsertHook:
		ipaVersionBeforeUpsertHooks = append(ipaVersionBeforeUpsertHooks, ipaVersionHook)
	case boil.AfterInsertHook:
		ipaVersionAfterInsertHooks = append(ipaVersionAfterInsertHooks, ipaVersionHook)
	case boil.AfterSelectHook:
		ipaVersionAfterSelectHooks = append(ipaVersionAfterSelectHooks, ipaVersionHook)
	case boil.AfterUpdateHook:
		ipaVersionAfterUpdateHooks = append(ipaVersionAfterUpdateHooks, ipaVersionHook)
	case boil.AfterDeleteHook:
		ipaVersionAfterDeleteHooks = append(ipaVersionAfterDeleteHooks, ipaVersionHook)
	case boil.AfterUpsertHook:
		ipaVersionAfterUpsertHooks = append(ipaVersionAfterUpsertHooks, ipaVersionHook)
	}
}

// One returns a single ipaVersion record from the query.
func (q ipaVersionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*IpaVersion, error) {
	o := &IpaVersion{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ipa_version")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all IpaVersion records from the query.
func (q ipaVersionQuery) All(ctx context.Context, exec boil.ContextExecutor) (IpaVersionSlice, error) {
	var o []*IpaVersion

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to IpaVersion slice")
	}

	if len(ipaVersionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all IpaVersion records in the query.
func (q ipaVersionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ipa_version rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ipaVersionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ipa_version exists")
	}

	return count > 0, nil
}

// IpaVersions retrieves all the records using an executor.
func IpaVersions(mods ...qm.QueryMod) ipaVersionQuery {
	mods = append(mods, qm.From("`ipa_version`"))
	return ipaVersionQuery{NewQuery(mods...)}
}

// FindIpaVersion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIpaVersion(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*IpaVersion, error) {
	ipaVersionObj := &IpaVersion{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `ipa_version` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ipaVersionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ipa_version")
	}

	if err = ipaVersionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ipaVersionObj, err
	}

	return ipaVersionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *IpaVersion) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ipa_version provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(ipaVersionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ipaVersionInsertCacheMut.RLock()
	cache, cached := ipaVersionInsertCache[key]
	ipaVersionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ipaVersionAllColumns,
			ipaVersionColumnsWithDefault,
			ipaVersionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ipaVersionType, ipaVersionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ipaVersionType, ipaVersionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `ipa_version` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `ipa_version` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `ipa_version` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, ipaVersionPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into ipa_version")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == ipaVersionMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for ipa_version")
	}

CacheNoHooks:
	if !cached {
		ipaVersionInsertCacheMut.Lock()
		ipaVersionInsertCache[key] = cache
		ipaVersionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the IpaVersion.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *IpaVersion) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ipaVersionUpdateCacheMut.RLock()
	cache, cached := ipaVersionUpdateCache[key]
	ipaVersionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ipaVersionAllColumns,
			ipaVersionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ipa_version, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `ipa_version` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, ipaVersionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ipaVersionType, ipaVersionMapping, append(wl, ipaVersionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ipa_version row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ipa_version")
	}

	if !cached {
		ipaVersionUpdateCacheMut.Lock()
		ipaVersionUpdateCache[key] = cache
		ipaVersionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ipaVersionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ipa_version")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ipa_version")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IpaVersionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipaVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `ipa_version` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ipaVersionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ipaVersion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ipaVersion")
	}
	return rowsAff, nil
}

var mySQLIpaVersionUniqueColumns = []string{
	"id",
	"token_path",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *IpaVersion) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ipa_version provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(ipaVersionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLIpaVersionUniqueColumns, o)

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

	ipaVersionUpsertCacheMut.RLock()
	cache, cached := ipaVersionUpsertCache[key]
	ipaVersionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ipaVersionAllColumns,
			ipaVersionColumnsWithDefault,
			ipaVersionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ipaVersionAllColumns,
			ipaVersionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert ipa_version, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`ipa_version`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `ipa_version` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(ipaVersionType, ipaVersionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ipaVersionType, ipaVersionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for ipa_version")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == ipaVersionMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(ipaVersionType, ipaVersionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for ipa_version")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for ipa_version")
	}

CacheNoHooks:
	if !cached {
		ipaVersionUpsertCacheMut.Lock()
		ipaVersionUpsertCache[key] = cache
		ipaVersionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single IpaVersion record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *IpaVersion) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no IpaVersion provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ipaVersionPrimaryKeyMapping)
	sql := "DELETE FROM `ipa_version` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ipa_version")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ipa_version")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ipaVersionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ipaVersionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ipa_version")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ipa_version")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IpaVersionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ipaVersionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipaVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `ipa_version` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ipaVersionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ipaVersion slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ipa_version")
	}

	if len(ipaVersionAfterDeleteHooks) != 0 {
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
func (o *IpaVersion) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIpaVersion(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IpaVersionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IpaVersionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipaVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `ipa_version`.* FROM `ipa_version` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ipaVersionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IpaVersionSlice")
	}

	*o = slice

	return nil
}

// IpaVersionExists checks if the IpaVersion row exists.
func IpaVersionExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `ipa_version` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ipa_version exists")
	}

	return exists, nil
}
