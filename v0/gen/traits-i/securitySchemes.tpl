package raml

type hasSecuritySchemes interface {
	{{define "hasSecuritySchemes"}}
	SecuritySchemes() UntypedMap
	{{- end}}
}
