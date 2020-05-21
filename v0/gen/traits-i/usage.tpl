package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

type hasUsage interface {
	{{define "hasUsage"}}
	// Usage returns an option of the value of the `usage` property for the
	// current RAML element.
	Usage() option.String

	// SetUsage sets the value of the usage property on the current RAML element.
	SetUsage(usage string) {{.}}

	// UnsetUsage removes the usage property from the current RAML element.
	UnsetUsage() {{.}}
	{{- end}}
}
