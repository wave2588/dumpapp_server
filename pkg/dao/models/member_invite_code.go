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

// MemberInviteCode is an object representing the database table.
type MemberInviteCode struct {
	ID       int64  `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	MemberID int64  `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	Code     string `boil:"code" json:"code" toml:"code" yaml:"code"`
	// ????
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// ????
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *memberInviteCodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L memberInviteCodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MemberInviteCodeColumns = struct {
	ID        string
	MemberID  string
	Code      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	MemberID:  "member_id",
	Code:      "code",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var MemberInviteCodeTableColumns = struct {
	ID        string
	MemberID  string
	Code      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "member_invite_code.id",
	MemberID:  "member_invite_code.member_id",
	Code:      "member_invite_code.code",
	CreatedAt: "member_invite_code.created_at",
	UpdatedAt: "member_invite_code.updated_at",
}

// Generated where

var MemberInviteCodeWhere = struct {
	ID        whereHelperint64
	MemberID  whereHelperint64
	Code      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`member_invite_code`.`id`"},
	MemberID:  whereHelperint64{field: "`member_invite_code`.`member_id`"},
	Code:      whereHelperstring{field: "`member_invite_code`.`code`"},
	CreatedAt: whereHelpertime_Time{field: "`member_invite_code`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`member_invite_code`.`updated_at`"},
}

// MemberInviteCodeRels is where relationship names are stored.
var MemberInviteCodeRels = struct {
}{}

// memberInviteCodeR is where relationships are stored.
type memberInviteCodeR struct {
}

// NewStruct creates a new relationship struct
func (*memberInviteCodeR) NewStruct() *memberInviteCodeR {
	return &memberInviteCodeR{}
}

// memberInviteCodeL is where Load methods for each relationship are stored.
type memberInviteCodeL struct{}

var (
	memberInviteCodeAllColumns            = []string{"id", "member_id", "code", "created_at", "updated_at"}
	memberInviteCodeColumnsWithoutDefault = []string{"member_id", "code"}
	memberInviteCodeColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	memberInviteCodePrimaryKeyColumns     = []string{"id"}
	memberInviteCodeGeneratedColumns      = []string{}
)

type (
	// MemberInviteCodeSlice is an alias for a slice of pointers to MemberInviteCode.
	// This should almost always be used instead of []MemberInviteCode.
	MemberInviteCodeSlice []*MemberInviteCode
	// MemberInviteCodeHook is the signature for custom MemberInviteCode hook methods
	MemberInviteCodeHook func(context.Context, boil.ContextExecutor, *MemberInviteCode) error

	memberInviteCodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	memberInviteCodeType                 = reflect.TypeOf(&MemberInviteCode{})
	memberInviteCodeMapping              = queries.MakeStructMapping(memberInviteCodeType)
	memberInviteCodePrimaryKeyMapping, _ = queries.BindMapping(memberInviteCodeType, memberInviteCodeMapping, memberInviteCodePrimaryKeyColumns)
	memberInviteCodeInsertCacheMut       sync.RWMutex
	memberInviteCodeInsertCache          = make(map[string]insertCache)
	memberInviteCodeUpdateCacheMut       sync.RWMutex
	memberInviteCodeUpdateCache          = make(map[string]updateCache)
	memberInviteCodeUpsertCacheMut       sync.RWMutex
	memberInviteCodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var memberInviteCodeAfterSelectHooks []MemberInviteCodeHook

var memberInviteCodeBeforeInsertHooks []MemberInviteCodeHook
var memberInviteCodeAfterInsertHooks []MemberInviteCodeHook

var memberInviteCodeBeforeUpdateHooks []MemberInviteCodeHook
var memberInviteCodeAfterUpdateHooks []MemberInviteCodeHook

var memberInviteCodeBeforeDeleteHooks []MemberInviteCodeHook
var memberInviteCodeAfterDeleteHooks []MemberInviteCodeHook

var memberInviteCodeBeforeUpsertHooks []MemberInviteCodeHook
var memberInviteCodeAfterUpsertHooks []MemberInviteCodeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MemberInviteCode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MemberInviteCode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MemberInviteCode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MemberInviteCode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MemberInviteCode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MemberInviteCode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MemberInviteCode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MemberInviteCode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MemberInviteCode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberInviteCodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMemberInviteCodeHook registers your hook function for all future operations.
func AddMemberInviteCodeHook(hookPoint boil.HookPoint, memberInviteCodeHook MemberInviteCodeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		memberInviteCodeAfterSelectHooks = append(memberInviteCodeAfterSelectHooks, memberInviteCodeHook)
	case boil.BeforeInsertHook:
		memberInviteCodeBeforeInsertHooks = append(memberInviteCodeBeforeInsertHooks, memberInviteCodeHook)
	case boil.AfterInsertHook:
		memberInviteCodeAfterInsertHooks = append(memberInviteCodeAfterInsertHooks, memberInviteCodeHook)
	case boil.BeforeUpdateHook:
		memberInviteCodeBeforeUpdateHooks = append(memberInviteCodeBeforeUpdateHooks, memberInviteCodeHook)
	case boil.AfterUpdateHook:
		memberInviteCodeAfterUpdateHooks = append(memberInviteCodeAfterUpdateHooks, memberInviteCodeHook)
	case boil.BeforeDeleteHook:
		memberInviteCodeBeforeDeleteHooks = append(memberInviteCodeBeforeDeleteHooks, memberInviteCodeHook)
	case boil.AfterDeleteHook:
		memberInviteCodeAfterDeleteHooks = append(memberInviteCodeAfterDeleteHooks, memberInviteCodeHook)
	case boil.BeforeUpsertHook:
		memberInviteCodeBeforeUpsertHooks = append(memberInviteCodeBeforeUpsertHooks, memberInviteCodeHook)
	case boil.AfterUpsertHook:
		memberInviteCodeAfterUpsertHooks = append(memberInviteCodeAfterUpsertHooks, memberInviteCodeHook)
	}
}

// One returns a single memberInviteCode record from the query.
func (q memberInviteCodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MemberInviteCode, error) {
	o := &MemberInviteCode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for member_invite_code")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MemberInviteCode records from the query.
func (q memberInviteCodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (MemberInviteCodeSlice, error) {
	var o []*MemberInviteCode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MemberInviteCode slice")
	}

	if len(memberInviteCodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MemberInviteCode records in the query.
func (q memberInviteCodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count member_invite_code rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q memberInviteCodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if member_invite_code exists")
	}

	return count > 0, nil
}

// MemberInviteCodes retrieves all the records using an executor.
func MemberInviteCodes(mods ...qm.QueryMod) memberInviteCodeQuery {
	mods = append(mods, qm.From("`member_invite_code`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`member_invite_code`.*"})
	}

	return memberInviteCodeQuery{q}
}

// FindMemberInviteCode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMemberInviteCode(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MemberInviteCode, error) {
	memberInviteCodeObj := &MemberInviteCode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `member_invite_code` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, memberInviteCodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from member_invite_code")
	}

	if err = memberInviteCodeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return memberInviteCodeObj, err
	}

	return memberInviteCodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MemberInviteCode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_invite_code provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(memberInviteCodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	memberInviteCodeInsertCacheMut.RLock()
	cache, cached := memberInviteCodeInsertCache[key]
	memberInviteCodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			memberInviteCodeAllColumns,
			memberInviteCodeColumnsWithDefault,
			memberInviteCodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(memberInviteCodeType, memberInviteCodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(memberInviteCodeType, memberInviteCodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `member_invite_code` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `member_invite_code` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `member_invite_code` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, memberInviteCodePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into member_invite_code")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberInviteCodeMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for member_invite_code")
	}

CacheNoHooks:
	if !cached {
		memberInviteCodeInsertCacheMut.Lock()
		memberInviteCodeInsertCache[key] = cache
		memberInviteCodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MemberInviteCode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MemberInviteCode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	memberInviteCodeUpdateCacheMut.RLock()
	cache, cached := memberInviteCodeUpdateCache[key]
	memberInviteCodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			memberInviteCodeAllColumns,
			memberInviteCodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update member_invite_code, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `member_invite_code` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, memberInviteCodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(memberInviteCodeType, memberInviteCodeMapping, append(wl, memberInviteCodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update member_invite_code row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for member_invite_code")
	}

	if !cached {
		memberInviteCodeUpdateCacheMut.Lock()
		memberInviteCodeUpdateCache[key] = cache
		memberInviteCodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q memberInviteCodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for member_invite_code")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for member_invite_code")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MemberInviteCodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberInviteCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `member_invite_code` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberInviteCodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in memberInviteCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all memberInviteCode")
	}
	return rowsAff, nil
}

var mySQLMemberInviteCodeUniqueColumns = []string{
	"id",
	"member_id",
	"code",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MemberInviteCode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no member_invite_code provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(memberInviteCodeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMemberInviteCodeUniqueColumns, o)

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

	memberInviteCodeUpsertCacheMut.RLock()
	cache, cached := memberInviteCodeUpsertCache[key]
	memberInviteCodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			memberInviteCodeAllColumns,
			memberInviteCodeColumnsWithDefault,
			memberInviteCodeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			memberInviteCodeAllColumns,
			memberInviteCodePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert member_invite_code, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`member_invite_code`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `member_invite_code` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(memberInviteCodeType, memberInviteCodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(memberInviteCodeType, memberInviteCodeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for member_invite_code")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberInviteCodeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(memberInviteCodeType, memberInviteCodeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for member_invite_code")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for member_invite_code")
	}

CacheNoHooks:
	if !cached {
		memberInviteCodeUpsertCacheMut.Lock()
		memberInviteCodeUpsertCache[key] = cache
		memberInviteCodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MemberInviteCode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MemberInviteCode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MemberInviteCode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), memberInviteCodePrimaryKeyMapping)
	sql := "DELETE FROM `member_invite_code` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from member_invite_code")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for member_invite_code")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q memberInviteCodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no memberInviteCodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from member_invite_code")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_invite_code")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MemberInviteCodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(memberInviteCodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberInviteCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `member_invite_code` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberInviteCodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from memberInviteCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for member_invite_code")
	}

	if len(memberInviteCodeAfterDeleteHooks) != 0 {
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
func (o *MemberInviteCode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMemberInviteCode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MemberInviteCodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MemberInviteCodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberInviteCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `member_invite_code`.* FROM `member_invite_code` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberInviteCodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MemberInviteCodeSlice")
	}

	*o = slice

	return nil
}

// MemberInviteCodeExists checks if the MemberInviteCode row exists.
func MemberInviteCodeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `member_invite_code` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if member_invite_code exists")
	}

	return exists, nil
}

// Exists checks if the MemberInviteCode row exists.
func (o *MemberInviteCode) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MemberInviteCodeExists(ctx, exec, o.ID)
}
