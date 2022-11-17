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

// InstallAppCertificate is an object representing the database table.
type InstallAppCertificate struct {
	ID                         int64                      `boil:"id" json:"id,string" toml:"id" yaml:"id"`
	Udid                       string                     `boil:"udid" json:"udid" toml:"udid" yaml:"udid"`
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

	R *installAppCertificateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L installAppCertificateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InstallAppCertificateColumns = struct {
	ID                         string
	Udid                       string
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
	Udid:                       "udid",
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

var InstallAppCertificateTableColumns = struct {
	ID                         string
	Udid                       string
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
	ID:                         "install_app_certificate.id",
	Udid:                       "install_app_certificate.udid",
	P12FileData:                "install_app_certificate.p12_file_data",
	P12FileDataMD5:             "install_app_certificate.p12_file_data_md5",
	ModifiedP12FileDate:        "install_app_certificate.modified_p12_file_date",
	MobileProvisionFileData:    "install_app_certificate.mobile_provision_file_data",
	MobileProvisionFileDataMD5: "install_app_certificate.mobile_provision_file_data_md5",
	Source:                     "install_app_certificate.source",
	BizExt:                     "install_app_certificate.biz_ext",
	CreatedAt:                  "install_app_certificate.created_at",
	UpdatedAt:                  "install_app_certificate.updated_at",
}

// Generated where

var InstallAppCertificateWhere = struct {
	ID                         whereHelperint64
	Udid                       whereHelperstring
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
	ID:                         whereHelperint64{field: "`install_app_certificate`.`id`"},
	Udid:                       whereHelperstring{field: "`install_app_certificate`.`udid`"},
	P12FileData:                whereHelperstring{field: "`install_app_certificate`.`p12_file_data`"},
	P12FileDataMD5:             whereHelperstring{field: "`install_app_certificate`.`p12_file_data_md5`"},
	ModifiedP12FileDate:        whereHelperstring{field: "`install_app_certificate`.`modified_p12_file_date`"},
	MobileProvisionFileData:    whereHelperstring{field: "`install_app_certificate`.`mobile_provision_file_data`"},
	MobileProvisionFileDataMD5: whereHelperstring{field: "`install_app_certificate`.`mobile_provision_file_data_md5`"},
	Source:                     whereHelperenum_CertificateSource{field: "`install_app_certificate`.`source`"},
	BizExt:                     whereHelperdatatype_CertificateBizExt{field: "`install_app_certificate`.`biz_ext`"},
	CreatedAt:                  whereHelpertime_Time{field: "`install_app_certificate`.`created_at`"},
	UpdatedAt:                  whereHelpertime_Time{field: "`install_app_certificate`.`updated_at`"},
}

// InstallAppCertificateRels is where relationship names are stored.
var InstallAppCertificateRels = struct {
}{}

// installAppCertificateR is where relationships are stored.
type installAppCertificateR struct {
}

// NewStruct creates a new relationship struct
func (*installAppCertificateR) NewStruct() *installAppCertificateR {
	return &installAppCertificateR{}
}

// installAppCertificateL is where Load methods for each relationship are stored.
type installAppCertificateL struct{}

var (
	installAppCertificateAllColumns            = []string{"id", "udid", "p12_file_data", "p12_file_data_md5", "modified_p12_file_date", "mobile_provision_file_data", "mobile_provision_file_data_md5", "source", "biz_ext", "created_at", "updated_at"}
	installAppCertificateColumnsWithoutDefault = []string{"udid", "p12_file_data", "p12_file_data_md5", "modified_p12_file_date", "mobile_provision_file_data", "mobile_provision_file_data_md5", "biz_ext"}
	installAppCertificateColumnsWithDefault    = []string{"id", "source", "created_at", "updated_at"}
	installAppCertificatePrimaryKeyColumns     = []string{"id"}
	installAppCertificateGeneratedColumns      = []string{}
)

type (
	// InstallAppCertificateSlice is an alias for a slice of pointers to InstallAppCertificate.
	// This should almost always be used instead of []InstallAppCertificate.
	InstallAppCertificateSlice []*InstallAppCertificate
	// InstallAppCertificateHook is the signature for custom InstallAppCertificate hook methods
	InstallAppCertificateHook func(context.Context, boil.ContextExecutor, *InstallAppCertificate) error

	installAppCertificateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	installAppCertificateType                 = reflect.TypeOf(&InstallAppCertificate{})
	installAppCertificateMapping              = queries.MakeStructMapping(installAppCertificateType)
	installAppCertificatePrimaryKeyMapping, _ = queries.BindMapping(installAppCertificateType, installAppCertificateMapping, installAppCertificatePrimaryKeyColumns)
	installAppCertificateInsertCacheMut       sync.RWMutex
	installAppCertificateInsertCache          = make(map[string]insertCache)
	installAppCertificateUpdateCacheMut       sync.RWMutex
	installAppCertificateUpdateCache          = make(map[string]updateCache)
	installAppCertificateUpsertCacheMut       sync.RWMutex
	installAppCertificateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var installAppCertificateAfterSelectHooks []InstallAppCertificateHook

var installAppCertificateBeforeInsertHooks []InstallAppCertificateHook
var installAppCertificateAfterInsertHooks []InstallAppCertificateHook

var installAppCertificateBeforeUpdateHooks []InstallAppCertificateHook
var installAppCertificateAfterUpdateHooks []InstallAppCertificateHook

var installAppCertificateBeforeDeleteHooks []InstallAppCertificateHook
var installAppCertificateAfterDeleteHooks []InstallAppCertificateHook

var installAppCertificateBeforeUpsertHooks []InstallAppCertificateHook
var installAppCertificateAfterUpsertHooks []InstallAppCertificateHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *InstallAppCertificate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *InstallAppCertificate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *InstallAppCertificate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *InstallAppCertificate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *InstallAppCertificate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *InstallAppCertificate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *InstallAppCertificate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *InstallAppCertificate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *InstallAppCertificate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range installAppCertificateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInstallAppCertificateHook registers your hook function for all future operations.
func AddInstallAppCertificateHook(hookPoint boil.HookPoint, installAppCertificateHook InstallAppCertificateHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		installAppCertificateAfterSelectHooks = append(installAppCertificateAfterSelectHooks, installAppCertificateHook)
	case boil.BeforeInsertHook:
		installAppCertificateBeforeInsertHooks = append(installAppCertificateBeforeInsertHooks, installAppCertificateHook)
	case boil.AfterInsertHook:
		installAppCertificateAfterInsertHooks = append(installAppCertificateAfterInsertHooks, installAppCertificateHook)
	case boil.BeforeUpdateHook:
		installAppCertificateBeforeUpdateHooks = append(installAppCertificateBeforeUpdateHooks, installAppCertificateHook)
	case boil.AfterUpdateHook:
		installAppCertificateAfterUpdateHooks = append(installAppCertificateAfterUpdateHooks, installAppCertificateHook)
	case boil.BeforeDeleteHook:
		installAppCertificateBeforeDeleteHooks = append(installAppCertificateBeforeDeleteHooks, installAppCertificateHook)
	case boil.AfterDeleteHook:
		installAppCertificateAfterDeleteHooks = append(installAppCertificateAfterDeleteHooks, installAppCertificateHook)
	case boil.BeforeUpsertHook:
		installAppCertificateBeforeUpsertHooks = append(installAppCertificateBeforeUpsertHooks, installAppCertificateHook)
	case boil.AfterUpsertHook:
		installAppCertificateAfterUpsertHooks = append(installAppCertificateAfterUpsertHooks, installAppCertificateHook)
	}
}

// One returns a single installAppCertificate record from the query.
func (q installAppCertificateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*InstallAppCertificate, error) {
	o := &InstallAppCertificate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for install_app_certificate")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all InstallAppCertificate records from the query.
func (q installAppCertificateQuery) All(ctx context.Context, exec boil.ContextExecutor) (InstallAppCertificateSlice, error) {
	var o []*InstallAppCertificate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to InstallAppCertificate slice")
	}

	if len(installAppCertificateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all InstallAppCertificate records in the query.
func (q installAppCertificateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count install_app_certificate rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q installAppCertificateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if install_app_certificate exists")
	}

	return count > 0, nil
}

// InstallAppCertificates retrieves all the records using an executor.
func InstallAppCertificates(mods ...qm.QueryMod) installAppCertificateQuery {
	mods = append(mods, qm.From("`install_app_certificate`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`install_app_certificate`.*"})
	}

	return installAppCertificateQuery{q}
}

// FindInstallAppCertificate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInstallAppCertificate(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*InstallAppCertificate, error) {
	installAppCertificateObj := &InstallAppCertificate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `install_app_certificate` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, installAppCertificateObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from install_app_certificate")
	}

	if err = installAppCertificateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return installAppCertificateObj, err
	}

	return installAppCertificateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *InstallAppCertificate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no install_app_certificate provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(installAppCertificateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	installAppCertificateInsertCacheMut.RLock()
	cache, cached := installAppCertificateInsertCache[key]
	installAppCertificateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			installAppCertificateAllColumns,
			installAppCertificateColumnsWithDefault,
			installAppCertificateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(installAppCertificateType, installAppCertificateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(installAppCertificateType, installAppCertificateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `install_app_certificate` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `install_app_certificate` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `install_app_certificate` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, installAppCertificatePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into install_app_certificate")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == installAppCertificateMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for install_app_certificate")
	}

CacheNoHooks:
	if !cached {
		installAppCertificateInsertCacheMut.Lock()
		installAppCertificateInsertCache[key] = cache
		installAppCertificateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the InstallAppCertificate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *InstallAppCertificate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	installAppCertificateUpdateCacheMut.RLock()
	cache, cached := installAppCertificateUpdateCache[key]
	installAppCertificateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			installAppCertificateAllColumns,
			installAppCertificatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update install_app_certificate, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `install_app_certificate` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, installAppCertificatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(installAppCertificateType, installAppCertificateMapping, append(wl, installAppCertificatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update install_app_certificate row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for install_app_certificate")
	}

	if !cached {
		installAppCertificateUpdateCacheMut.Lock()
		installAppCertificateUpdateCache[key] = cache
		installAppCertificateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q installAppCertificateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for install_app_certificate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for install_app_certificate")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InstallAppCertificateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), installAppCertificatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `install_app_certificate` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, installAppCertificatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in installAppCertificate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all installAppCertificate")
	}
	return rowsAff, nil
}

var mySQLInstallAppCertificateUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *InstallAppCertificate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no install_app_certificate provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(installAppCertificateColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLInstallAppCertificateUniqueColumns, o)

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

	installAppCertificateUpsertCacheMut.RLock()
	cache, cached := installAppCertificateUpsertCache[key]
	installAppCertificateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			installAppCertificateAllColumns,
			installAppCertificateColumnsWithDefault,
			installAppCertificateColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			installAppCertificateAllColumns,
			installAppCertificatePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert install_app_certificate, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`install_app_certificate`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `install_app_certificate` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(installAppCertificateType, installAppCertificateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(installAppCertificateType, installAppCertificateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for install_app_certificate")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == installAppCertificateMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(installAppCertificateType, installAppCertificateMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for install_app_certificate")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for install_app_certificate")
	}

CacheNoHooks:
	if !cached {
		installAppCertificateUpsertCacheMut.Lock()
		installAppCertificateUpsertCache[key] = cache
		installAppCertificateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single InstallAppCertificate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *InstallAppCertificate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InstallAppCertificate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), installAppCertificatePrimaryKeyMapping)
	sql := "DELETE FROM `install_app_certificate` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from install_app_certificate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for install_app_certificate")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q installAppCertificateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no installAppCertificateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from install_app_certificate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for install_app_certificate")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InstallAppCertificateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(installAppCertificateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), installAppCertificatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `install_app_certificate` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, installAppCertificatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from installAppCertificate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for install_app_certificate")
	}

	if len(installAppCertificateAfterDeleteHooks) != 0 {
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
func (o *InstallAppCertificate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInstallAppCertificate(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InstallAppCertificateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InstallAppCertificateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), installAppCertificatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `install_app_certificate`.* FROM `install_app_certificate` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, installAppCertificatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InstallAppCertificateSlice")
	}

	*o = slice

	return nil
}

// InstallAppCertificateExists checks if the InstallAppCertificate row exists.
func InstallAppCertificateExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `install_app_certificate` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if install_app_certificate exists")
	}

	return exists, nil
}

// Exists checks if the InstallAppCertificate row exists.
func (o *InstallAppCertificate) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return InstallAppCertificateExists(ctx, exec, o.ID)
}
