{{ define "root" -}}
package raml
{{ if .DefIsOpt }}
import "github.com/Foxcapades/goop/v1/pkg/option"
{{ end -}}

// {{.Name}}Type generated @ {{.Time}}
type {{.Name}}Type interface {
	Unmarshaler
	Marshaler
	{{if .Base -}}
		DataType
	{{- else -}}
		ExtendedDataType
{{ template "all" $ }}
{{- if eq .Type "array" -}}
	{{ template "array" $ }}
{{- else if eq .Type "datetime" -}}
	{{ template "datetime" $ }}
{{- else if eq .Type "file" -}}
	{{ template "file" $ }}
{{- else if eq .Type "integer" -}}
	{{ template "integer" $ }}
{{- else if eq .Type "number" -}}
	{{ template "number" $ }}
{{- else if eq .Type "object" -}}
	{{ template "object" $ }}
{{- else if eq .Type "string" -}}
	{{ template "string" $ }}
{{- end -}}
{{end}}
}
{{end}}