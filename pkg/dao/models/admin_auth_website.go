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

	"dumpapp_server/pkg/common/datatype"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AdminAuthWebsite is an object representing the database table.
type AdminAuthWebsite struct {
	ID       int64                           `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	MemberID int64                           `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	Domain   string                          `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`
	BizExt   datatype.AdminAuthWebsiteBizExt `boil:"biz_ext" json:"biz_ext" toml:"biz_ext" yaml:"biz_ext"`
	// ????
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// ????
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *adminAuthWebsiteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminAuthWebsiteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminAuthWebsiteColumns = struct {
	ID        string
	MemberID  string
	Domain    string
	BizExt    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	MemberID:  "member_id",
	Domain:    "domain",
	BizExt:    "biz_ext",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var AdminAuthWebsiteTableColumns = struct {
	ID        string
	MemberID  string
	Domain    string
	BizExt    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "admin_auth_website.id",
	MemberID:  "admin_auth_website.member_id",
	Domain:    "admin_auth_website.domain",
	BizExt:    "admin_auth_website.biz_ext",
	CreatedAt: "admin_auth_website.created_at",
	UpdatedAt: "admin_auth_website.updated_at",
}

// Generated where

type whereHelperdatatype_AdminAuthWebsiteBizExt struct{ field string }

func (w whereHelperdatatype_AdminAuthWebsiteBizExt) EQ(x datatype.AdminAuthWebsiteBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperdatatype_AdminAuthWebsiteBizExt) NEQ(x datatype.AdminAuthWebsiteBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperdatatype_AdminAuthWebsiteBizExt) LT(x datatype.AdminAuthWebsiteBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperdatatype_AdminAuthWebsiteBizExt) LTE(x datatype.AdminAuthWebsiteBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperdatatype_AdminAuthWebsiteBizExt) GT(x datatype.AdminAuthWebsiteBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperdatatype_AdminAuthWebsiteBizExt) GTE(x datatype.AdminAuthWebsiteBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AdminAuthWebsiteWhere = struct {
	ID        whereHelperint64
	MemberID  whereHelperint64
	Domain    whereHelperstring
	BizExt    whereHelperdatatype_AdminAuthWebsiteBizExt
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`admin_auth_website`.`id`"},
	MemberID:  whereHelperint64{field: "`admin_auth_website`.`member_id`"},
	Domain:    whereHelperstring{field: "`admin_auth_website`.`domain`"},
	BizExt:    whereHelperdatatype_AdminAuthWebsiteBizExt{field: "`admin_auth_website`.`biz_ext`"},
	CreatedAt: whereHelpertime_Time{field: "`admin_auth_website`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`admin_auth_website`.`updated_at`"},
}

// AdminAuthWebsiteRels is where relationship names are stored.
var AdminAuthWebsiteRels = struct {
}{}

// adminAuthWebsiteR is where relationships are stored.
type adminAuthWebsiteR struct {
}

// NewStruct creates a new relationship struct
func (*adminAuthWebsiteR) NewStruct() *adminAuthWebsiteR {
	return &adminAuthWebsiteR{}
}

// adminAuthWebsiteL is where Load methods for each relationship are stored.
type adminAuthWebsiteL struct{}

var (
	adminAuthWebsiteAllColumns            = []string{"id", "member_id", "domain", "biz_ext", "created_at", "updated_at"}
	adminAuthWebsiteColumnsWithoutDefault = []string{"member_id", "domain", "biz_ext"}
	adminAuthWebsiteColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	adminAuthWebsitePrimaryKeyColumns     = []string{"id"}
	adminAuthWebsiteGeneratedColumns      = []string{}
)

type (
	// AdminAuthWebsiteSlice is an alias for a slice of pointers to AdminAuthWebsite.
	// This should almost always be used instead of []AdminAuthWebsite.
	AdminAuthWebsiteSlice []*AdminAuthWebsite
	// AdminAuthWebsiteHook is the signature for custom AdminAuthWebsite hook methods
	AdminAuthWebsiteHook func(context.Context, boil.ContextExecutor, *AdminAuthWebsite) error

	adminAuthWebsiteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminAuthWebsiteType                 = reflect.TypeOf(&AdminAuthWebsite{})
	adminAuthWebsiteMapping              = queries.MakeStructMapping(adminAuthWebsiteType)
	adminAuthWebsitePrimaryKeyMapping, _ = queries.BindMapping(adminAuthWebsiteType, adminAuthWebsiteMapping, adminAuthWebsitePrimaryKeyColumns)
	adminAuthWebsiteInsertCacheMut       sync.RWMutex
	adminAuthWebsiteInsertCache          = make(map[string]insertCache)
	adminAuthWebsiteUpdateCacheMut       sync.RWMutex
	adminAuthWebsiteUpdateCache          = make(map[string]updateCache)
	adminAuthWebsiteUpsertCacheMut       sync.RWMutex
	adminAuthWebsiteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var adminAuthWebsiteAfterSelectHooks []AdminAuthWebsiteHook

var adminAuthWebsiteBeforeInsertHooks []AdminAuthWebsiteHook
var adminAuthWebsiteAfterInsertHooks []AdminAuthWebsiteHook

var adminAuthWebsiteBeforeUpdateHooks []AdminAuthWebsiteHook
var adminAuthWebsiteAfterUpdateHooks []AdminAuthWebsiteHook

var adminAuthWebsiteBeforeDeleteHooks []AdminAuthWebsiteHook
var adminAuthWebsiteAfterDeleteHooks []AdminAuthWebsiteHook

var adminAuthWebsiteBeforeUpsertHooks []AdminAuthWebsiteHook
var adminAuthWebsiteAfterUpsertHooks []AdminAuthWebsiteHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AdminAuthWebsite) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AdminAuthWebsite) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AdminAuthWebsite) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AdminAuthWebsite) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AdminAuthWebsite) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AdminAuthWebsite) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AdminAuthWebsite) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AdminAuthWebsite) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AdminAuthWebsite) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range adminAuthWebsiteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAdminAuthWebsiteHook registers your hook function for all future operations.
func AddAdminAuthWebsiteHook(hookPoint boil.HookPoint, adminAuthWebsiteHook AdminAuthWebsiteHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		adminAuthWebsiteAfterSelectHooks = append(adminAuthWebsiteAfterSelectHooks, adminAuthWebsiteHook)
	case boil.BeforeInsertHook:
		adminAuthWebsiteBeforeInsertHooks = append(adminAuthWebsiteBeforeInsertHooks, adminAuthWebsiteHook)
	case boil.AfterInsertHook:
		adminAuthWebsiteAfterInsertHooks = append(adminAuthWebsiteAfterInsertHooks, adminAuthWebsiteHook)
	case boil.BeforeUpdateHook:
		adminAuthWebsiteBeforeUpdateHooks = append(adminAuthWebsiteBeforeUpdateHooks, adminAuthWebsiteHook)
	case boil.AfterUpdateHook:
		adminAuthWebsiteAfterUpdateHooks = append(adminAuthWebsiteAfterUpdateHooks, adminAuthWebsiteHook)
	case boil.BeforeDeleteHook:
		adminAuthWebsiteBeforeDeleteHooks = append(adminAuthWebsiteBeforeDeleteHooks, adminAuthWebsiteHook)
	case boil.AfterDeleteHook:
		adminAuthWebsiteAfterDeleteHooks = append(adminAuthWebsiteAfterDeleteHooks, adminAuthWebsiteHook)
	case boil.BeforeUpsertHook:
		adminAuthWebsiteBeforeUpsertHooks = append(adminAuthWebsiteBeforeUpsertHooks, adminAuthWebsiteHook)
	case boil.AfterUpsertHook:
		adminAuthWebsiteAfterUpsertHooks = append(adminAuthWebsiteAfterUpsertHooks, adminAuthWebsiteHook)
	}
}

// One returns a single adminAuthWebsite record from the query.
func (q adminAuthWebsiteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AdminAuthWebsite, error) {
	o := &AdminAuthWebsite{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for admin_auth_website")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AdminAuthWebsite records from the query.
func (q adminAuthWebsiteQuery) All(ctx context.Context, exec boil.ContextExecutor) (AdminAuthWebsiteSlice, error) {
	var o []*AdminAuthWebsite

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AdminAuthWebsite slice")
	}

	if len(adminAuthWebsiteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AdminAuthWebsite records in the query.
func (q adminAuthWebsiteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count admin_auth_website rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q adminAuthWebsiteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if admin_auth_website exists")
	}

	return count > 0, nil
}

// AdminAuthWebsites retrieves all the records using an executor.
func AdminAuthWebsites(mods ...qm.QueryMod) adminAuthWebsiteQuery {
	mods = append(mods, qm.From("`admin_auth_website`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`admin_auth_website`.*"})
	}

	return adminAuthWebsiteQuery{q}
}

// FindAdminAuthWebsite retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdminAuthWebsite(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AdminAuthWebsite, error) {
	adminAuthWebsiteObj := &AdminAuthWebsite{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `admin_auth_website` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, adminAuthWebsiteObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from admin_auth_website")
	}

	if err = adminAuthWebsiteObj.doAfterSelectHooks(ctx, exec); err != nil {
		return adminAuthWebsiteObj, err
	}

	return adminAuthWebsiteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AdminAuthWebsite) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_auth_website provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(adminAuthWebsiteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	adminAuthWebsiteInsertCacheMut.RLock()
	cache, cached := adminAuthWebsiteInsertCache[key]
	adminAuthWebsiteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			adminAuthWebsiteAllColumns,
			adminAuthWebsiteColumnsWithDefault,
			adminAuthWebsiteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(adminAuthWebsiteType, adminAuthWebsiteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminAuthWebsiteType, adminAuthWebsiteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `admin_auth_website` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `admin_auth_website` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `admin_auth_website` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, adminAuthWebsitePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into admin_auth_website")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == adminAuthWebsiteMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for admin_auth_website")
	}

CacheNoHooks:
	if !cached {
		adminAuthWebsiteInsertCacheMut.Lock()
		adminAuthWebsiteInsertCache[key] = cache
		adminAuthWebsiteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AdminAuthWebsite.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AdminAuthWebsite) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	adminAuthWebsiteUpdateCacheMut.RLock()
	cache, cached := adminAuthWebsiteUpdateCache[key]
	adminAuthWebsiteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			adminAuthWebsiteAllColumns,
			adminAuthWebsitePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update admin_auth_website, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `admin_auth_website` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, adminAuthWebsitePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminAuthWebsiteType, adminAuthWebsiteMapping, append(wl, adminAuthWebsitePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update admin_auth_website row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for admin_auth_website")
	}

	if !cached {
		adminAuthWebsiteUpdateCacheMut.Lock()
		adminAuthWebsiteUpdateCache[key] = cache
		adminAuthWebsiteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q adminAuthWebsiteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for admin_auth_website")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for admin_auth_website")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminAuthWebsiteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminAuthWebsitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `admin_auth_website` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, adminAuthWebsitePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in adminAuthWebsite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all adminAuthWebsite")
	}
	return rowsAff, nil
}

var mySQLAdminAuthWebsiteUniqueColumns = []string{
	"id",
	"domain",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AdminAuthWebsite) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_auth_website provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(adminAuthWebsiteColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAdminAuthWebsiteUniqueColumns, o)

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

	adminAuthWebsiteUpsertCacheMut.RLock()
	cache, cached := adminAuthWebsiteUpsertCache[key]
	adminAuthWebsiteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			adminAuthWebsiteAllColumns,
			adminAuthWebsiteColumnsWithDefault,
			adminAuthWebsiteColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			adminAuthWebsiteAllColumns,
			adminAuthWebsitePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert admin_auth_website, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`admin_auth_website`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `admin_auth_website` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(adminAuthWebsiteType, adminAuthWebsiteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminAuthWebsiteType, adminAuthWebsiteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for admin_auth_website")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == adminAuthWebsiteMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(adminAuthWebsiteType, adminAuthWebsiteMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for admin_auth_website")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for admin_auth_website")
	}

CacheNoHooks:
	if !cached {
		adminAuthWebsiteUpsertCacheMut.Lock()
		adminAuthWebsiteUpsertCache[key] = cache
		adminAuthWebsiteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AdminAuthWebsite record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AdminAuthWebsite) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AdminAuthWebsite provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminAuthWebsitePrimaryKeyMapping)
	sql := "DELETE FROM `admin_auth_website` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from admin_auth_website")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for admin_auth_website")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q adminAuthWebsiteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no adminAuthWebsiteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from admin_auth_website")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_auth_website")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminAuthWebsiteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(adminAuthWebsiteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminAuthWebsitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `admin_auth_website` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, adminAuthWebsitePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from adminAuthWebsite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_auth_website")
	}

	if len(adminAuthWebsiteAfterDeleteHooks) != 0 {
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
func (o *AdminAuthWebsite) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAdminAuthWebsite(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminAuthWebsiteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AdminAuthWebsiteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminAuthWebsitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `admin_auth_website`.* FROM `admin_auth_website` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, adminAuthWebsitePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AdminAuthWebsiteSlice")
	}

	*o = slice

	return nil
}

// AdminAuthWebsiteExists checks if the AdminAuthWebsite row exists.
func AdminAuthWebsiteExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `admin_auth_website` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if admin_auth_website exists")
	}

	return exists, nil
}

// Exists checks if the AdminAuthWebsite row exists.
func (o *AdminAuthWebsite) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AdminAuthWebsiteExists(ctx, exec, o.ID)
}
