
{{- $alias := .Aliases.Table .Table.Name -}}
{{- $colDefs := sqlColDefinitions .Table.Columns .Table.PKey.Columns -}}
{{- $pkNames := $colDefs.Names | stringMap (aliasCols $alias) | stringMap .StringFuncs.camelCase | stringMap .StringFuncs.replaceReserved -}}
{{- $pkArgs := joinSlices " " $pkNames $colDefs.Types | join ", " -}}
{{- $schemaTable := .Table.Name | .SchemaTable -}}
{{- $uniqKeys := .Table.UniqKeys -}}

type {{$alias.UpSingular}}DAO struct {
	mysqlPool        *sql.DB
}

var Default{{$alias.UpSingular}}DAO *{{$alias.UpSingular}}DAO

func init() {
	Default{{$alias.UpSingular}}DAO = New{{$alias.UpSingular}}DAO()
}

func New{{$alias.UpSingular}}DAO() *{{$alias.UpSingular}}DAO {
	d := &{{$alias.UpSingular}}DAO{
        mysqlPool:        clients.MySQLConnectionsPool,
	}
	return d
}

func (d *{{$alias.UpSingular}}DAO) Insert(ctx context.Context, data *models.{{$alias.UpSingular}}) error {
	var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }

    err := data.Insert(ctx, exec, boil.Infer())
	if err != nil {
		if mysqlError, ok := pkgErr.Cause(err).(*mysqlDriver.MySQLError); !(ok && mysqlError.Number == 1062) {
			return pkgErr.WithStack(err)
		}
	}
	return nil
}

func (d *{{$alias.UpSingular}}DAO) Update(ctx context.Context, data *models.{{$alias.UpSingular}}) error {
	var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }
    _, err := data.Update(ctx, exec, boil.Infer())
	return pkgErr.WithStack(err)
}

func (d *{{$alias.UpSingular}}DAO) Delete(ctx context.Context, id int64) error {
	qs := []qm.QueryMod{
		models.{{$alias.UpSingular}}Where.ID.EQ(id),
	}

	var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }
    _, err := models.{{$alias.UpPlural}}(qs...).DeleteAll(ctx, exec)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	return nil
}

func (d *{{$alias.UpSingular}}DAO) Get(ctx context.Context, id int64) (*models.{{$alias.UpSingular}}, error) {
   result, err := d.BatchGet(ctx, []int64{id})
	if err != nil {
		return nil, err
	}
	if v, ok := result[id]; !ok {
		return nil, pkgErr.Wrapf(errors.ErrNotFound, "table={{$.Table.Name}}, id=%d", id)
	} else {
		return v, nil
	}
}

// BatchGet retrieves multiple records by primary key from db.
func (d *{{$alias.UpSingular}}DAO) BatchGet(ctx context.Context, ids []int64) (map[int64]*models.{{$alias.UpSingular}}, error) {
    var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }
    datas, err := models.{{$alias.UpPlural}}(models.{{$alias.UpSingular}}Where.ID.IN(ids)).All(ctx, exec)
    if err != nil {
    	return nil, pkgErr.WithStack(err)
    }

    result := make(map[int64]*models.{{$alias.UpSingular}})
    for _, c := range datas {
    		result[c.ID] = c
    }

    return result, nil
}

// 后台和脚本使用：倒序列出所有
func (d *{{$alias.UpSingular}}DAO) ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error) {
    if offset < 0 || limit <= 0 || limit > 10000 {
        return nil, pkgErr.Errorf("invalid offset or limit")
    }
	qs := []qm.QueryMod{qm.Select(models.{{$alias.UpSingular}}Columns.ID)}
	qs = append(qs, filters...)

    if len(orderBys) > 0 {
		orderBys = append(orderBys, "id desc")
		for _, orderBy := range orderBys {
			qs = append(qs, qm.OrderBy(orderBy))
		}
	} else {
		qs = append(qs, qm.OrderBy("created_at DESC, id DESC"))
	}

	if offset >= 0 && limit >= 0 {
		qs = append(qs, qm.Offset(offset), qm.Limit(limit))
	}

    var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }

	datas, err := models.{{$alias.UpPlural}}(qs...).All(ctx, exec)
    if err != nil {
   		return nil, pkgErr.Wrap(err, fmt.Sprintf("table={{$.Table.Name}} offset=%d limit=%d filters=%v", offset, limit, filters))
   	}

   	result := make([]int64, 0)
   	for _, c := range datas {
        result = append(result, c.ID)
    }
    return result, nil
}

func (d *{{$alias.UpSingular}}DAO) Count(ctx context.Context, filters []qm.QueryMod) (int64, error) {
	qs := []qm.QueryMod{qm.Select(models.{{$alias.UpSingular}}Columns.ID)}
	qs = append(qs, filters...)

    var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }

	return models.{{$alias.UpPlural}}(qs...).Count(ctx, exec)
}

{{ range $uniqKey := .Table.UniqKeys -}}
	{{- $ukColDefs := sqlColDefinitions $.Table.Columns $uniqKey.Columns -}}
	{{- $ukColNames := $ukColDefs.Names | stringMap (aliasCols $alias) | stringMap $.StringFuncs.camelCase | stringMap $.StringFuncs.replaceReserved -}}
	{{- $ukFieldNames := $ukColDefs.Names | stringMap (aliasCols $alias) | stringMap $.StringFuncs.titleCase | stringMap $.StringFuncs.replaceReserved -}}
	{{- $ukArgs := joinSlices " " $ukColNames $ukColDefs.Types | join ", " -}}
	{{- $ukNameArgs := $ukColNames | join ", " -}}
	{{- $fnName := $ukColNames | stringMap $.StringFuncs.titleCase | join "" -}}


// GetBy{{$fnName}} retrieves a single record by uniq key {{$ukNameArgs}} from db.
func (d *{{$alias.UpSingular}}DAO) GetBy{{$fnName}}(ctx context.Context, {{$ukArgs}}) (*models.{{$alias.UpSingular}}, error) {
	{{$alias.DownSingular}}Obj := &models.{{$alias.UpSingular}}{}

	sel := "*"
	query := fmt.Sprintf(
	"select %s from {{$.Table.Name | $.SchemaTable}} where {{if $.Dialect.UseIndexPlaceholders}}{{whereClause $.LQ $.RQ 1 $uniqKey.Columns}}{{else}}{{whereClause $.LQ $.RQ 0 $uniqKey.Columns}}{{end}}{{if and $.AddSoftDeletes}} and {{"deleted_at" | $.Quotes}} is null{{end}}", sel,
	)

	q := queries.Raw(query, {{$ukColNames | join ", "}})

	var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }

	err := q.Bind(ctx, exec, {{$alias.DownSingular}}Obj)
	if err != nil {
	    if pkgErr.Cause(err) == sql.ErrNoRows {
	        return nil, pkgErr.Wrapf(errors.ErrNotFound, "table={{$.Table.Name}}, query=%s, args={{ $ukColNames | join ":%v "}} :%v", query, {{ $ukColNames | join ", "}})
	    }
	    return nil, pkgErr.Wrap(err, "{{$.PkgName}}: unable to select from {{$.Table.Name}}")
	}

	return {{$alias.DownSingular}}Obj, nil
}

	{{ if eq (len $uniqKey.Columns) 1 -}}

	{{- $ukColName := index $ukColNames 0 -}}
	{{- $ukFieldName :=  index $ukFieldNames 0 -}}
	{{- $ukColType :=  index $ukColDefs.Types 0 -}}


// BatchGetBy{{$fnName}} retrieves multiple records by uniq key {{$ukColName}} from db.
func (d *{{$alias.UpSingular}}DAO) BatchGetBy{{$fnName}}(ctx context.Context, {{$ukColName}}s []{{$ukColType}}) (map[{{$ukColType}}]*models.{{$alias.UpSingular}}, error) {
    var exec boil.ContextExecutor
    txn := ctx.Value("txn")
    if txn == nil {
        exec = d.mysqlPool
    } else {
        exec = txn.(*sql.Tx)
    }
    datas, err := models.{{$alias.UpPlural}}(models.{{$alias.UpSingular}}Where.{{$ukFieldName}}.IN({{$ukColName}}s)).All(ctx, exec)
    if err != nil {
    	return nil, pkgErr.WithStack(err)
    }

    result := make(map[{{$ukColType}}]*models.{{$alias.UpSingular}})
    for _, c := range datas {
    		result[c.{{$ukFieldName}}] = c
    }

    return result, nil
}
{{ end -}}


	{{ if gt (len $uniqKey.Columns) 1 -}}
		{{- $columns := slice $uniqKey.Columns 0 1 -}}
		{{- $ukColDefs := sqlColDefinitions $.Table.Columns $columns -}}
		{{- $ukColNames := $ukColDefs.Names | stringMap (aliasCols $alias) | stringMap $.StringFuncs.camelCase | stringMap $.StringFuncs.replaceReserved -}}
		{{- $ukArgs := joinSlices " " $ukColNames $ukColDefs.Types | join ", " -}}
		{{- $fnName := $ukColNames | stringMap $.StringFuncs.titleCase | join "" -}}

		// Get{{$alias.UpSingular}}SliceBy{{$fnName}} retrieves a slice of records by first field of uniq key {{$ukColNames}} with an executor.
		func (d *{{$alias.UpSingular}}DAO)Get{{$alias.UpSingular}}SliceBy{{$fnName}} (ctx context.Context, {{$ukArgs}}) ([]*models.{{$alias.UpSingular}}, error) {
		var o []*models.{{$alias.UpSingular}}

		query := "select {{$.Table.Name | $.SchemaTable}}.* from {{$.Table.Name | $.SchemaTable}} where {{if $.Dialect.UseIndexPlaceholders}}{{whereClause $.LQ $.RQ 1 $columns}}{{else}}{{whereClause $.LQ $.RQ 0 $columns}}{{end}}{{if and $.AddSoftDeletes}} and {{"deleted_at" | $.Quotes}} is null{{end}}"

		q := queries.Raw(query, {{$ukColNames | join ", "}})

        var exec boil.ContextExecutor
        txn := ctx.Value("txn")
        if txn == nil {
            exec = d.mysqlPool
        } else {
            exec = txn.(*sql.Tx)
        }

		err := q.Bind(ctx, exec, &o)
		if err != nil {
		    if pkgErr.Cause(err) == sql.ErrNoRows {
		        return nil, pkgErr.Wrapf(errors.ErrNotFound, "table={{$.Table.Name}}, query=%s, args={{ $ukColNames | join "%v "}} :%v", query, {{ $ukColNames | join ", "}})
		    }
		    return nil, pkgErr.Wrap(err, "{{$.PkgName}}: unable to select from {{$.Table.Name}}")
		}
		return o, nil
		}
	{{ end -}}
{{ end -}}
