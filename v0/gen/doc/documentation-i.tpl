{{define "documentation-interface" -}}
package raml

type Documentation interface{
	Marshaler
	Unmarshaler
	{{template "hasTitle"  "Documentation"}}
	{{template "hasFacets" "Documentation"}}
	// Content returns the value of the content property defined on the current
	// RAML element.
	Content() string

	// SetContent sets the value of the content property on the current RAML
	// element.
	SetContent(content string) Documentation
}
{{- end}}