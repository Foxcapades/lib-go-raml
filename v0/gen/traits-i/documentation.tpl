package raml

type hasDocumentation interface {
	{{define "hasDocumentation"}}
	// Documentation returns the array of documentation elements defined on the
	// current RAML element.
	Documentation() []Documentation

	// SetDocumentation sets the documentation array for the current RAML element.
	//
	// Passing this method a nil value is equivalent to calling
	// UnsetDocumentation.
	SetDocumentation(docs []Documentation) {{.}}

	// UnsetDocumentation removes the documentation property and all its values
	// from the current RAML element.
	UnsetDocumentation() {{.}}
	{{- end}}
}
