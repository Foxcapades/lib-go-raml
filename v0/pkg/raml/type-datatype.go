package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
)

type DataType interface {
	hasFacets

	Schema() string
	Type() string
	Kind() rmeta.DataTypeKind

	// ExtraFacets returns a mutable map of the extra facets
	// applied to the current DataType definition.
	ExtraFacets() AnyMap

	ToRAML() (string, error)
}
