package raml

type hasTraits interface {
	{{define "hasTraits"}}
	// Traits returns the map of trait definitions for the current RAML element.
	Traits() UntypedMap

	// SetTraits replaces the map of trait definitions on the current RAML element
	// with the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetTraits.
	SetTraits(traits UntypedMap) {{.}}

	// UnsetTraits removes all trait definitions from the current RAML element.
	UnsetTraits() {{.}}
	{{- end}}
}
