package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

type hasFacets interface {
{{define "hasFacets"}}
	// ExtraFacets returns a mutable map of extra/custom facets applied to the
	// current RAML element.
	ExtraFacets() AnyMap
{{- end}}
}
