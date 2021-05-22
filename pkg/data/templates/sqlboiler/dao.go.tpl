{{- $alias := .Aliases.Table .Table.Name -}}
{{- $colDefs := sqlColDefinitions .Table.Columns .Table.PKey.Columns -}}
{{- $pkNames := $colDefs.Names | stringMap (aliasCols $alias) | stringMap .StringFuncs.camelCase | stringMap .StringFuncs.replaceReserved -}}
{{- $pkArgs := joinSlices " " $pkNames $colDefs.Types | join ", " -}}
{{- $schemaTable := .Table.Name | .SchemaTable -}}
{{- $canSoftDelete := .Table.CanSoftDelete }}
{{- $uniqKeys := .Table.UniqKeys -}}

type {{$alias.UpSingular}}DAO interface {
    Insert(ctx context.Context, data *models.{{$alias.UpSingular}}) error
    Update(ctx context.Context, data *models.{{$alias.UpSingular}}) error
    Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*models.{{$alias.UpSingular}}, error)
	BatchGet(ctx context.Context, ids []int64) (map[int64]*models.{{$alias.UpSingular}}, error)
	BatchGetMySQL(ctx context.Context, ids []int64) (map[int64]*models.{{$alias.UpSingular}}, error)
	ListIDs(ctx context.Context, offset, limit int, filters []qm.QueryMod, orderBys []string) ([]int64, error)
	Count(ctx context.Context, filters []qm.QueryMod) (int64, error)

{{ range $uniqKey := .Table.UniqKeys -}}
{{- $ukColDefs := sqlColDefinitions $.Table.Columns $uniqKey.Columns -}}
{{- $ukColNames := $ukColDefs.Names | stringMap (aliasCols $alias) | stringMap $.StringFuncs.camelCase | stringMap $.StringFuncs.replaceReserved -}}
{{- $ukArgs := joinSlices " " $ukColNames $ukColDefs.Types | join ", " -}}
{{- $fnName := $ukColNames | stringMap $.StringFuncs.titleCase | join "" -}}

    Get{{$alias.UpSingular}}By{{$fnName}}(ctx context.Context, {{$ukArgs}}) (*models.{{$alias.UpSingular}}, error)
{{ if gt (len $uniqKey.Columns) 1 -}}
    {{- $columns := slice $uniqKey.Columns 0 1 -}}
    {{- $ukColDefs := sqlColDefinitions $.Table.Columns $columns -}}
    {{- $ukColNames := $ukColDefs.Names | stringMap (aliasCols $alias) | stringMap $.StringFuncs.camelCase | stringMap $.StringFuncs.replaceReserved -}}
    {{- $ukArgs := joinSlices " " $ukColNames $ukColDefs.Types | join ", " -}}
    {{- $fnName := $ukColNames | stringMap $.StringFuncs.titleCase | join "" -}}
    Get{{$alias.UpSingular}}SliceBy{{$fnName}} (ctx context.Context, {{$ukArgs}}) ([]*models.{{$alias.UpSingular}}, error)
{{ end }}
{{ end }}
}