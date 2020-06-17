package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
)

type DataType interface {
	hasFacets

	// Type returns the value of the type property of this RAML element.
	Type() string

	// Schema is an alias for Type
	Schema() string

	// OverrideType allows setting the raw type for the current RAML element.
	//
	// Similar to SetType on higher level implementations, this sets the type
	// property of the current element, but without the builder pattern setter.
	//
	// Additionally this is the only way to set a custom (invalid) type on basic
	// nodes such as Nil or Any.
	OverrideType(t string)

	// Kind returns the the base or computed type for the RAML element.
	Kind() rmeta.DataTypeKind

	// ExtraFacets returns a mutable map of the extra facets
	// applied to the current DataType definition.
	ExtraFacets() AnyMap

	ToRAML() (string, error)
}
