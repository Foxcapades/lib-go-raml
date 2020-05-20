{{- /* gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/type.extTypeProps */ -}}
{{define "root" -}}
package raml
{{if .Base -}}
	{{template "base" $}}
{{- else -}}
	{{template "extended" $}}
{{- end}}
{{end}}