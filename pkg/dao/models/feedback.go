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

// Feedback is an object representing the database table.
type Feedback struct {
	ID       int64  `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	MemberID int64  `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	Content  string `boil:"content" json:"content" toml:"content" yaml:"content"`
	// ????
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// ????
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *feedbackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L feedbackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FeedbackColumns = struct {
	ID        string
	MemberID  string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	MemberID:  "member_id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var FeedbackTableColumns = struct {
	ID        string
	MemberID  string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "feedback.id",
	MemberID:  "feedback.member_id",
	Content:   "feedback.content",
	CreatedAt: "feedback.created_at",
	UpdatedAt: "feedback.updated_at",
}

// Generated where

var FeedbackWhere = struct {
	ID        whereHelperint64
	MemberID  whereHelperint64
	Content   whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`feedback`.`id`"},
	MemberID:  whereHelperint64{field: "`feedback`.`member_id`"},
	Content:   whereHelperstring{field: "`feedback`.`content`"},
	CreatedAt: whereHelpertime_Time{field: "`feedback`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`feedback`.`updated_at`"},
}

// FeedbackRels is where relationship names are stored.
var FeedbackRels = struct {
}{}

// feedbackR is where relationships are stored.
type feedbackR struct {
}

// NewStruct creates a new relationship struct
func (*feedbackR) NewStruct() *feedbackR {
	return &feedbackR{}
}

// feedbackL is where Load methods for each relationship are stored.
type feedbackL struct{}

var (
	feedbackAllColumns            = []string{"id", "member_id", "content", "created_at", "updated_at"}
	feedbackColumnsWithoutDefault = []string{"member_id", "content"}
	feedbackColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	feedbackPrimaryKeyColumns     = []string{"id"}
	feedbackGeneratedColumns      = []string{}
)

type (
	// FeedbackSlice is an alias for a slice of pointers to Feedback.
	// This should almost always be used instead of []Feedback.
	FeedbackSlice []*Feedback
	// FeedbackHook is the signature for custom Feedback hook methods
	FeedbackHook func(context.Context, boil.ContextExecutor, *Feedback) error

	feedbackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	feedbackType                 = reflect.TypeOf(&Feedback{})
	feedbackMapping              = queries.MakeStructMapping(feedbackType)
	feedbackPrimaryKeyMapping, _ = queries.BindMapping(feedbackType, feedbackMapping, feedbackPrimaryKeyColumns)
	feedbackInsertCacheMut       sync.RWMutex
	feedbackInsertCache          = make(map[string]insertCache)
	feedbackUpdateCacheMut       sync.RWMutex
	feedbackUpdateCache          = make(map[string]updateCache)
	feedbackUpsertCacheMut       sync.RWMutex
	feedbackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var feedbackAfterSelectHooks []FeedbackHook

var feedbackBeforeInsertHooks []FeedbackHook
var feedbackAfterInsertHooks []FeedbackHook

var feedbackBeforeUpdateHooks []FeedbackHook
var feedbackAfterUpdateHooks []FeedbackHook

var feedbackBeforeDeleteHooks []FeedbackHook
var feedbackAfterDeleteHooks []FeedbackHook

var feedbackBeforeUpsertHooks []FeedbackHook
var feedbackAfterUpsertHooks []FeedbackHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Feedback) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Feedback) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Feedback) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Feedback) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Feedback) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Feedback) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Feedback) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Feedback) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Feedback) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range feedbackAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeedbackHook registers your hook function for all future operations.
func AddFeedbackHook(hookPoint boil.HookPoint, feedbackHook FeedbackHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		feedbackAfterSelectHooks = append(feedbackAfterSelectHooks, feedbackHook)
	case boil.BeforeInsertHook:
		feedbackBeforeInsertHooks = append(feedbackBeforeInsertHooks, feedbackHook)
	case boil.AfterInsertHook:
		feedbackAfterInsertHooks = append(feedbackAfterInsertHooks, feedbackHook)
	case boil.BeforeUpdateHook:
		feedbackBeforeUpdateHooks = append(feedbackBeforeUpdateHooks, feedbackHook)
	case boil.AfterUpdateHook:
		feedbackAfterUpdateHooks = append(feedbackAfterUpdateHooks, feedbackHook)
	case boil.BeforeDeleteHook:
		feedbackBeforeDeleteHooks = append(feedbackBeforeDeleteHooks, feedbackHook)
	case boil.AfterDeleteHook:
		feedbackAfterDeleteHooks = append(feedbackAfterDeleteHooks, feedbackHook)
	case boil.BeforeUpsertHook:
		feedbackBeforeUpsertHooks = append(feedbackBeforeUpsertHooks, feedbackHook)
	case boil.AfterUpsertHook:
		feedbackAfterUpsertHooks = append(feedbackAfterUpsertHooks, feedbackHook)
	}
}

// One returns a single feedback record from the query.
func (q feedbackQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Feedback, error) {
	o := &Feedback{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feedback")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Feedback records from the query.
func (q feedbackQuery) All(ctx context.Context, exec boil.ContextExecutor) (FeedbackSlice, error) {
	var o []*Feedback

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Feedback slice")
	}

	if len(feedbackAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Feedback records in the query.
func (q feedbackQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feedback rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q feedbackQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feedback exists")
	}

	return count > 0, nil
}

// Feedbacks retrieves all the records using an executor.
func Feedbacks(mods ...qm.QueryMod) feedbackQuery {
	mods = append(mods, qm.From("`feedback`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`feedback`.*"})
	}

	return feedbackQuery{q}
}

// FindFeedback retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeedback(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Feedback, error) {
	feedbackObj := &Feedback{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `feedback` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, feedbackObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feedback")
	}

	if err = feedbackObj.doAfterSelectHooks(ctx, exec); err != nil {
		return feedbackObj, err
	}

	return feedbackObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Feedback) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no feedback provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(feedbackColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	feedbackInsertCacheMut.RLock()
	cache, cached := feedbackInsertCache[key]
	feedbackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			feedbackAllColumns,
			feedbackColumnsWithDefault,
			feedbackColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(feedbackType, feedbackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(feedbackType, feedbackMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `feedback` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `feedback` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `feedback` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, feedbackPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into feedback")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == feedbackMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for feedback")
	}

CacheNoHooks:
	if !cached {
		feedbackInsertCacheMut.Lock()
		feedbackInsertCache[key] = cache
		feedbackInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Feedback.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Feedback) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	feedbackUpdateCacheMut.RLock()
	cache, cached := feedbackUpdateCache[key]
	feedbackUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			feedbackAllColumns,
			feedbackPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update feedback, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `feedback` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, feedbackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(feedbackType, feedbackMapping, append(wl, feedbackPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update feedback row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for feedback")
	}

	if !cached {
		feedbackUpdateCacheMut.Lock()
		feedbackUpdateCache[key] = cache
		feedbackUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q feedbackQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for feedback")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for feedback")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeedbackSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), feedbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `feedback` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, feedbackPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in feedback slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all feedback")
	}
	return rowsAff, nil
}

var mySQLFeedbackUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Feedback) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no feedback provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(feedbackColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFeedbackUniqueColumns, o)

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

	feedbackUpsertCacheMut.RLock()
	cache, cached := feedbackUpsertCache[key]
	feedbackUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			feedbackAllColumns,
			feedbackColumnsWithDefault,
			feedbackColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			feedbackAllColumns,
			feedbackPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert feedback, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`feedback`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `feedback` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(feedbackType, feedbackMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(feedbackType, feedbackMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feedback")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == feedbackMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(feedbackType, feedbackMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for feedback")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for feedback")
	}

CacheNoHooks:
	if !cached {
		feedbackUpsertCacheMut.Lock()
		feedbackUpsertCache[key] = cache
		feedbackUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Feedback record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Feedback) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Feedback provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), feedbackPrimaryKeyMapping)
	sql := "DELETE FROM `feedback` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from feedback")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for feedback")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q feedbackQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no feedbackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from feedback")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for feedback")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeedbackSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(feedbackBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), feedbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `feedback` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, feedbackPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from feedback slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for feedback")
	}

	if len(feedbackAfterDeleteHooks) != 0 {
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
func (o *Feedback) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFeedback(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeedbackSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FeedbackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), feedbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `feedback`.* FROM `feedback` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, feedbackPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeedbackSlice")
	}

	*o = slice

	return nil
}

// FeedbackExists checks if the Feedback row exists.
func FeedbackExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `feedback` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feedback exists")
	}

	return exists, nil
}

// Exists checks if the Feedback row exists.
func (o *Feedback) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FeedbackExists(ctx, exec, o.ID)
}
