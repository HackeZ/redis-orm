{{- define "script.mysql"}}{{- $obj := . -}}
{{- if ne $obj.DbTable ""}}
CREATE TABLE IF NOT EXISTS `{{$obj.DbTable}}` (
	{{- range $i, $field := $obj.Fields}}
	{{$field.SQLColumn "mysql"}}, 
	{{- end}}
	{{$obj.PrimaryKey.SQLColumn "mysql"}}
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

{{range $i, $unique := $obj.Uniques}}
{{- if not $unique.HasPrimaryKey}}
DROP UNIQUE INDEX `{{$unique.Name | camel2name}}` ON `{{$obj.DbTable}}`;
CREATE UNIQUE INDEX `{{$unique.Name | camel2name}}` ON `{{$obj.DbTable}}`(
	{{- range $i, $f := $unique.Fields -}}		
		{{- if eq (add $i 1) (len $unique.Fields) -}}
			`{{- $f.Name | camel2name -}}`
		{{- else -}}
			`{{- $f.Name | camel2name -}}`,
		{{- end -}}
	{{- end -}}
);
{{- end}}
{{- end}}

{{- range $i, $index := $obj.Indexes}}
{{- if not $index.HasPrimaryKey}}
DROP INDEX `{{$index.Name | camel2name}}` ON `{{$obj.DbTable}}`
CREATE INDEX `{{$index.Name | camel2name}}` ON `{{$obj.DbTable}}`(
	{{- range $i, $f := $index.Fields -}}
		{{- if eq (add $i 1) (len $index.Fields) -}}
			`{{- $f.Name | camel2name -}}`
		{{- else -}}
			`{{- $f.Name | camel2name -}}`,
		{{- end -}}
	{{- end -}}
);
{{- end}}
{{- end}}

{{- range $i, $index := $obj.Ranges}}
{{- if not $index.HasPrimaryKey}}
DROP INDEX `{{$index.Name | camel2name}}` ON `{{$obj.DbTable}}`;
CREATE INDEX `{{$index.Name | camel2name}}` ON `{{$obj.DbTable}}`(
	{{- range $i, $f := $index.Fields -}}
		{{- if eq (add $i 1) (len $index.Fields) -}}
			`{{- $f.Name | camel2name -}}`
		{{- else -}}
			`{{- $f.Name | camel2name -}}`,
		{{- end -}}
	{{- end -}}
);
{{- end}}
{{- end}}
{{- end}}

{{- if ne $obj.DbView ""}}
DROP VIEW IF EXISTS `{{$obj.DbView}}`;
CREATE VIEW `{{$obj.DbView}}` AS {{$obj.ImportSQL}};
{{- end}}

{{end}}
