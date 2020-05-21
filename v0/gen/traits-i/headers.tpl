package raml

type headers interface {
	{{define "hasHeaders"}}
	// Headers returns the map of headers defined on the current RAML element.
	Headers() DataTypeMap

	// SetHeaders replaces the map of headers for the current RAML element with
	// the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetHeaders.
	SetHeaders(headers DataTypeMap) {{.}}

	// UnsetHeaders removes all headers from the current RAML element.
	UnsetHeaders() {{.}}
	{{- end}}
}
