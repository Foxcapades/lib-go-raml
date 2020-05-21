package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

type displayed interface {
	{{define "hasDisplayName"}}
	// DisplayName returns an option of the contents of the `displayName` facet
	// applied to the current RAML element.
	//
	// If the element does not have a display name applied, the returned option
	// will be empty.
	DisplayName() option.String

	// SetDisplayName sets the value of the displayName property for the current
	// RAML element.
	SetDisplayName(name string) {{.}}

	// UnsetDisplayName removes the displayName property from the current RAML
	// element.
	UnsetDisplayName() {{.}}
	{{- end}}
}
