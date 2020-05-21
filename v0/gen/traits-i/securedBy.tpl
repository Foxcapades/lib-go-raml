package raml

type securedBy interface {
	{{define "hasSecuredBy"}}
	// SecuredBy returns the value of the securedBy array defined on the current
	// RAML element.
	SecuredBy() []string

	// SetSecuredBy sets the securedBy array for the current RAML element.
	//
	// Passing this method a nil value is equivalent to calling UnsetSecuredBy.
	SetSecuredBy(securedBy []string) {{.}}

	// UnsetSecuredBy removes the securedBy property and all its values from the
	// current RAML element.
	UnsetSecuredBy() {{.}}
	{{- end}}
}