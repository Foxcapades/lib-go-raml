package uses

type HasImports interface {
	{{define "hasImports"}}
	// Uses returns the map of defined library imports for the current RAML
	// element.
	Uses() StringMap

	// SetUses replaces the map of defined library imports for the current RAML
	// element with the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetUses.
	SetUses(uses StringMap) {{.}}

	// UnsetUses removes the map of defined library imports from the current RAML
	// element.
	UnsetUses() {{.}}
	{{- end}}
}
