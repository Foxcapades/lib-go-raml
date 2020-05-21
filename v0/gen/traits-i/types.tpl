package raml

type hasTypes interface {
	{{define "hasTypes"}}
	// Types returns the value of the `types` or `schemas` property defined on
	// the current document.
	Types() DataTypeMap

	// SetTypes replaces the map of types defined on the current RAML element with
	// the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetTypes.
	SetTypes(types DataTypeMap) {{.}}

	// UnsetTypes removes all types from the current RAML element.
	UnsetTypes() {{.}}

	// Schemas is an alias for the Types method.
	Schemas() DataTypeMap

	// SetSchemas is an alias for the SetTypes method.
	SetSchemas(types DataTypeMap) {{.}}

	// UnsetSchemas is an alias for the UnsetTypes method.
	UnsetSchemas() {{.}}
	{{- end}}
}
