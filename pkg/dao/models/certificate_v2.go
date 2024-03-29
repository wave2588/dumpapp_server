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
	"dumpapp_server/pkg/common/enum"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CertificateV2 is an object representing the database table.
type CertificateV2 struct {
	ID                         int64                      `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	DeviceID                   int64                      `boil:"device_id" json:"device_id" toml:"device_id" yaml:"device_id"`
	P12FileData                string                     `boil:"p12_file_data" json:"p12_file_data" toml:"p12_file_data" yaml:"p12_file_data"`
	P12FileDataMD5             string                     `boil:"p12_file_data_md5" json:"p12_file_data_md5" toml:"p12_file_data_md5" yaml:"p12_file_data_md5"`
	ModifiedP12FileDate        string                     `boil:"modified_p12_file_date" json:"modified_p12_file_date" toml:"modified_p12_file_date" yaml:"modified_p12_file_date"`
	MobileProvisionFileData    string                     `boil:"mobile_provision_file_data" json:"mobile_provision_file_data" toml:"mobile_provision_file_data" yaml:"mobile_provision_file_data"`
	MobileProvisionFileDataMD5 string                     `boil:"mobile_provision_file_data_md5" json:"mobile_provision_file_data_md5" toml:"mobile_provision_file_data_md5" yaml:"mobile_provision_file_data_md5"`
	Source                     enum.CertificateSource     `boil:"source" json:"source" toml:"source" yaml:"source"`
	BizExt                     datatype.CertificateBizExt `boil:"biz_ext" json:"biz_ext" toml:"biz_ext" yaml:"biz_ext"`
	// ????
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// ????
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *certificateV2R `boil:"-" json:"-" toml:"-" yaml:"-"`
	L certificateV2L  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CertificateV2Columns = struct {
	ID                         string
	DeviceID                   string
	P12FileData                string
	P12FileDataMD5             string
	ModifiedP12FileDate        string
	MobileProvisionFileData    string
	MobileProvisionFileDataMD5 string
	Source                     string
	BizExt                     string
	CreatedAt                  string
	UpdatedAt                  string
}{
	ID:                         "id",
	DeviceID:                   "device_id",
	P12FileData:                "p12_file_data",
	P12FileDataMD5:             "p12_file_data_md5",
	ModifiedP12FileDate:        "modified_p12_file_date",
	MobileProvisionFileData:    "mobile_provision_file_data",
	MobileProvisionFileDataMD5: "mobile_provision_file_data_md5",
	Source:                     "source",
	BizExt:                     "biz_ext",
	CreatedAt:                  "created_at",
	UpdatedAt:                  "updated_at",
}

var CertificateV2TableColumns = struct {
	ID                         string
	DeviceID                   string
	P12FileData                string
	P12FileDataMD5             string
	ModifiedP12FileDate        string
	MobileProvisionFileData    string
	MobileProvisionFileDataMD5 string
	Source                     string
	BizExt                     string
	CreatedAt                  string
	UpdatedAt                  string
}{
	ID:                         "certificate_v2.id",
	DeviceID:                   "certificate_v2.device_id",
	P12FileData:                "certificate_v2.p12_file_data",
	P12FileDataMD5:             "certificate_v2.p12_file_data_md5",
	ModifiedP12FileDate:        "certificate_v2.modified_p12_file_date",
	MobileProvisionFileData:    "certificate_v2.mobile_provision_file_data",
	MobileProvisionFileDataMD5: "certificate_v2.mobile_provision_file_data_md5",
	Source:                     "certificate_v2.source",
	BizExt:                     "certificate_v2.biz_ext",
	CreatedAt:                  "certificate_v2.created_at",
	UpdatedAt:                  "certificate_v2.updated_at",
}

// Generated where

type whereHelperenum_CertificateSource struct{ field string }

func (w whereHelperenum_CertificateSource) EQ(x enum.CertificateSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperenum_CertificateSource) NEQ(x enum.CertificateSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperenum_CertificateSource) LT(x enum.CertificateSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperenum_CertificateSource) LTE(x enum.CertificateSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperenum_CertificateSource) GT(x enum.CertificateSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperenum_CertificateSource) GTE(x enum.CertificateSource) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperdatatype_CertificateBizExt struct{ field string }

func (w whereHelperdatatype_CertificateBizExt) EQ(x datatype.CertificateBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperdatatype_CertificateBizExt) NEQ(x datatype.CertificateBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperdatatype_CertificateBizExt) LT(x datatype.CertificateBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperdatatype_CertificateBizExt) LTE(x datatype.CertificateBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperdatatype_CertificateBizExt) GT(x datatype.CertificateBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperdatatype_CertificateBizExt) GTE(x datatype.CertificateBizExt) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CertificateV2Where = struct {
	ID                         whereHelperint64
	DeviceID                   whereHelperint64
	P12FileData                whereHelperstring
	P12FileDataMD5             whereHelperstring
	ModifiedP12FileDate        whereHelperstring
	MobileProvisionFileData    whereHelperstring
	MobileProvisionFileDataMD5 whereHelperstring
	Source                     whereHelperenum_CertificateSource
	BizExt                     whereHelperdatatype_CertificateBizExt
	CreatedAt                  whereHelpertime_Time
	UpdatedAt                  whereHelpertime_Time
}{
	ID:                         whereHelperint64{field: "`certificate_v2`.`id`"},
	DeviceID:                   whereHelperint64{field: "`certificate_v2`.`device_id`"},
	P12FileData:                whereHelperstring{field: "`certificate_v2`.`p12_file_data`"},
	P12FileDataMD5:             whereHelperstring{field: "`certificate_v2`.`p12_file_data_md5`"},
	ModifiedP12FileDate:        whereHelperstring{field: "`certificate_v2`.`modified_p12_file_date`"},
	MobileProvisionFileData:    whereHelperstring{field: "`certificate_v2`.`mobile_provision_file_data`"},
	MobileProvisionFileDataMD5: whereHelperstring{field: "`certificate_v2`.`mobile_provision_file_data_md5`"},
	Source:                     whereHelperenum_CertificateSource{field: "`certificate_v2`.`source`"},
	BizExt:                     whereHelperdatatype_CertificateBizExt{field: "`certificate_v2`.`biz_ext`"},
	CreatedAt:                  whereHelpertime_Time{field: "`certificate_v2`.`created_at`"},
	UpdatedAt:                  whereHelpertime_Time{field: "`certificate_v2`.`updated_at`"},
}

// CertificateV2Rels is where relationship names are stored.
var CertificateV2Rels = struct {
}{}

// certificateV2R is where relationships are stored.
type certificateV2R struct {
}

// NewStruct creates a new relationship struct
func (*certificateV2R) NewStruct() *certificateV2R {
	return &certificateV2R{}
}

// certificateV2L is where Load methods for each relationship are stored.
type certificateV2L struct{}

var (
	certificateV2AllColumns            = []string{"id", "device_id", "p12_file_data", "p12_file_data_md5", "modified_p12_file_date", "mobile_provision_file_data", "mobile_provision_file_data_md5", "source", "biz_ext", "created_at", "updated_at"}
	certificateV2ColumnsWithoutDefault = []string{"device_id", "p12_file_data", "p12_file_data_md5", "modified_p12_file_date", "mobile_provision_file_data", "mobile_provision_file_data_md5", "biz_ext"}
	certificateV2ColumnsWithDefault    = []string{"id", "source", "created_at", "updated_at"}
	certificateV2PrimaryKeyColumns     = []string{"id"}
	certificateV2GeneratedColumns      = []string{}
)

type (
	// CertificateV2Slice is an alias for a slice of pointers to CertificateV2.
	// This should almost always be used instead of []CertificateV2.
	CertificateV2Slice []*CertificateV2
	// CertificateV2Hook is the signature for custom CertificateV2 hook methods
	CertificateV2Hook func(context.Context, boil.ContextExecutor, *CertificateV2) error

	certificateV2Query struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	certificateV2Type                 = reflect.TypeOf(&CertificateV2{})
	certificateV2Mapping              = queries.MakeStructMapping(certificateV2Type)
	certificateV2PrimaryKeyMapping, _ = queries.BindMapping(certificateV2Type, certificateV2Mapping, certificateV2PrimaryKeyColumns)
	certificateV2InsertCacheMut       sync.RWMutex
	certificateV2InsertCache          = make(map[string]insertCache)
	certificateV2UpdateCacheMut       sync.RWMutex
	certificateV2UpdateCache          = make(map[string]updateCache)
	certificateV2UpsertCacheMut       sync.RWMutex
	certificateV2UpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var certificateV2AfterSelectHooks []CertificateV2Hook

var certificateV2BeforeInsertHooks []CertificateV2Hook
var certificateV2AfterInsertHooks []CertificateV2Hook

var certificateV2BeforeUpdateHooks []CertificateV2Hook
var certificateV2AfterUpdateHooks []CertificateV2Hook

var certificateV2BeforeDeleteHooks []CertificateV2Hook
var certificateV2AfterDeleteHooks []CertificateV2Hook

var certificateV2BeforeUpsertHooks []CertificateV2Hook
var certificateV2AfterUpsertHooks []CertificateV2Hook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CertificateV2) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2AfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CertificateV2) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2BeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CertificateV2) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2AfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CertificateV2) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2BeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CertificateV2) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2AfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CertificateV2) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2BeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CertificateV2) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2AfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CertificateV2) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2BeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CertificateV2) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range certificateV2AfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCertificateV2Hook registers your hook function for all future operations.
func AddCertificateV2Hook(hookPoint boil.HookPoint, certificateV2Hook CertificateV2Hook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		certificateV2AfterSelectHooks = append(certificateV2AfterSelectHooks, certificateV2Hook)
	case boil.BeforeInsertHook:
		certificateV2BeforeInsertHooks = append(certificateV2BeforeInsertHooks, certificateV2Hook)
	case boil.AfterInsertHook:
		certificateV2AfterInsertHooks = append(certificateV2AfterInsertHooks, certificateV2Hook)
	case boil.BeforeUpdateHook:
		certificateV2BeforeUpdateHooks = append(certificateV2BeforeUpdateHooks, certificateV2Hook)
	case boil.AfterUpdateHook:
		certificateV2AfterUpdateHooks = append(certificateV2AfterUpdateHooks, certificateV2Hook)
	case boil.BeforeDeleteHook:
		certificateV2BeforeDeleteHooks = append(certificateV2BeforeDeleteHooks, certificateV2Hook)
	case boil.AfterDeleteHook:
		certificateV2AfterDeleteHooks = append(certificateV2AfterDeleteHooks, certificateV2Hook)
	case boil.BeforeUpsertHook:
		certificateV2BeforeUpsertHooks = append(certificateV2BeforeUpsertHooks, certificateV2Hook)
	case boil.AfterUpsertHook:
		certificateV2AfterUpsertHooks = append(certificateV2AfterUpsertHooks, certificateV2Hook)
	}
}

// One returns a single certificateV2 record from the query.
func (q certificateV2Query) One(ctx context.Context, exec boil.ContextExecutor) (*CertificateV2, error) {
	o := &CertificateV2{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for certificate_v2")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CertificateV2 records from the query.
func (q certificateV2Query) All(ctx context.Context, exec boil.ContextExecutor) (CertificateV2Slice, error) {
	var o []*CertificateV2

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CertificateV2 slice")
	}

	if len(certificateV2AfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CertificateV2 records in the query.
func (q certificateV2Query) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count certificate_v2 rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q certificateV2Query) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if certificate_v2 exists")
	}

	return count > 0, nil
}

// CertificateV2S retrieves all the records using an executor.
func CertificateV2S(mods ...qm.QueryMod) certificateV2Query {
	mods = append(mods, qm.From("`certificate_v2`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`certificate_v2`.*"})
	}

	return certificateV2Query{q}
}

// FindCertificateV2 retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCertificateV2(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CertificateV2, error) {
	certificateV2Obj := &CertificateV2{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `certificate_v2` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, certificateV2Obj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from certificate_v2")
	}

	if err = certificateV2Obj.doAfterSelectHooks(ctx, exec); err != nil {
		return certificateV2Obj, err
	}

	return certificateV2Obj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CertificateV2) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no certificate_v2 provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(certificateV2ColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	certificateV2InsertCacheMut.RLock()
	cache, cached := certificateV2InsertCache[key]
	certificateV2InsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			certificateV2AllColumns,
			certificateV2ColumnsWithDefault,
			certificateV2ColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(certificateV2Type, certificateV2Mapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(certificateV2Type, certificateV2Mapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `certificate_v2` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `certificate_v2` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `certificate_v2` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, certificateV2PrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into certificate_v2")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == certificateV2Mapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for certificate_v2")
	}

CacheNoHooks:
	if !cached {
		certificateV2InsertCacheMut.Lock()
		certificateV2InsertCache[key] = cache
		certificateV2InsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CertificateV2.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CertificateV2) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	certificateV2UpdateCacheMut.RLock()
	cache, cached := certificateV2UpdateCache[key]
	certificateV2UpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			certificateV2AllColumns,
			certificateV2PrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update certificate_v2, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `certificate_v2` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, certificateV2PrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(certificateV2Type, certificateV2Mapping, append(wl, certificateV2PrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update certificate_v2 row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for certificate_v2")
	}

	if !cached {
		certificateV2UpdateCacheMut.Lock()
		certificateV2UpdateCache[key] = cache
		certificateV2UpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q certificateV2Query) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for certificate_v2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for certificate_v2")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CertificateV2Slice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), certificateV2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `certificate_v2` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, certificateV2PrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in certificateV2 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all certificateV2")
	}
	return rowsAff, nil
}

var mySQLCertificateV2UniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CertificateV2) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no certificate_v2 provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(certificateV2ColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCertificateV2UniqueColumns, o)

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

	certificateV2UpsertCacheMut.RLock()
	cache, cached := certificateV2UpsertCache[key]
	certificateV2UpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			certificateV2AllColumns,
			certificateV2ColumnsWithDefault,
			certificateV2ColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			certificateV2AllColumns,
			certificateV2PrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert certificate_v2, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`certificate_v2`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `certificate_v2` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(certificateV2Type, certificateV2Mapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(certificateV2Type, certificateV2Mapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for certificate_v2")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == certificateV2Mapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(certificateV2Type, certificateV2Mapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for certificate_v2")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for certificate_v2")
	}

CacheNoHooks:
	if !cached {
		certificateV2UpsertCacheMut.Lock()
		certificateV2UpsertCache[key] = cache
		certificateV2UpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CertificateV2 record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CertificateV2) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CertificateV2 provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), certificateV2PrimaryKeyMapping)
	sql := "DELETE FROM `certificate_v2` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from certificate_v2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for certificate_v2")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q certificateV2Query) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no certificateV2Query provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from certificate_v2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for certificate_v2")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CertificateV2Slice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(certificateV2BeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), certificateV2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `certificate_v2` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, certificateV2PrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from certificateV2 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for certificate_v2")
	}

	if len(certificateV2AfterDeleteHooks) != 0 {
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
func (o *CertificateV2) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCertificateV2(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CertificateV2Slice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CertificateV2Slice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), certificateV2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `certificate_v2`.* FROM `certificate_v2` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, certificateV2PrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CertificateV2Slice")
	}

	*o = slice

	return nil
}

// CertificateV2Exists checks if the CertificateV2 row exists.
func CertificateV2Exists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `certificate_v2` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if certificate_v2 exists")
	}

	return exists, nil
}

// Exists checks if the CertificateV2 row exists.
func (o *CertificateV2) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CertificateV2Exists(ctx, exec, o.ID)
}
