package raml

type hasTitle interface {
	{{define "hasTitle"}}
	// Title returns the value of the title property set on the current RAML
	// element.
	Title() string

	// SetTitle sets the value of the title property on the current RAML element.
	SetTitle(title string) {{.}}
	{{- end}}
}
