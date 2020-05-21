package raml

type hasResourceTypes interface {
	{{define "hasResourceTypes"}}
	// ResourceTypes returns the resourceTypes map defined on the current RAML
	// element.
	ResourceTypes() UntypedMap

	// SetResourceTypes replaces the resourceTypes map defined on the current RAML
	// element with the given value.
	//
	// Passing this method a nil value is equivalent to calling
	// UnsetResourceTypes.
	SetResourceTypes(resTypes UntypedMap) {{.}}

	// UnsetResourceTypes removes the resourceTypes map from the current RAML
	// element.
	UnsetResourceTypes() {{.}}
	{{- end}}
}
