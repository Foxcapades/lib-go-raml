package raml

type hasDescription interface {
	{{define "hasDescription"}}
	// Description returns an option of the contents of the `description` facet
	// applied to the current RAML element.
	//
	// If the element does not have a description applied, the returned option
	// will be empty.
	Description() option.String

	// SetDescription sets the value of the description property for the current
	// RAML element.
	SetDescription(desc string) {{.}}

	// UnsetDescription removes the description property from the current RAML
	// element.
	UnsetDescription() {{.}}
	{{- end}}
}
