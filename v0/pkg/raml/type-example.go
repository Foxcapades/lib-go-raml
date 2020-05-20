package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
)

type Example interface {
	Unmarshaler
	Marshaler

	// DisplayName returns an option that will contain this
	// example's display name if one is set.
	DisplayName() option.String

	// Description returns an option that will contain this
	// example's description if one is set.
	Description() option.String

	// Annotations returns a mutable map of the annotations
	// currently applied to this example.
	Annotations() AnnotationMap

	// Strict returns whether or not this example will be
	// validated against its parent type definition.
	//
	// This value defaults to true.  If set to true when
	// rendered, it will not appear in the produced RAML.
	Strict() bool

	ExtraFacets() AnyMap
}
