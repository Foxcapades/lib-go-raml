package raml

type body interface {
	{{define "hasBody"}}
	Body() Body

	SetBody(body Body) {{.}}

	UnsetBody() {{.}}
	{{- end}}
}