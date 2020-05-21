package raml

type protocols interface {
	{{define "hasProtocols"}}
	// Protocols returns the values of the protocol array set on the current RAML
	// element.
	Protocols() []string

	// SetProtocols sets the protocol array for the current RAML element to the
	// given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetProtocols.
	SetProtocols(protocols []string) {{.}}

	// UnsetProtocols removes the protocols property and all its values from the
	// current RAML element.
	UnsetProtocols() {{.}}
	{{- end}}
}