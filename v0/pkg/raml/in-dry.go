package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

type hasAnnotations interface {
	// Annotations returns a map of the annotations applied to the current
	// RAML element.
	Annotations() AnnotationMap
}

type hasDescription interface {
	// Description returns an option of the contents of the `description` facet
	// applied to the current RAML element.
	//
	// If the element does not have a description applied, the returned option
	// will be empty.
	Description() option.String
}

type displayed interface {
	// DisplayName returns an option of the contents of the `displayName` facet
	// applied to the current RAML element.
	//
	// If the element does not have a display name applied, the returned option
	// will be empty.
	DisplayName() option.String
}

type hasFacets interface {
	// ExtraFacets returns a mutable map of extra/custom facets applied to the
	// current RAML element.
	ExtraFacets() AnyMap
}

type headered interface {
	Headers() DataTypeMap
}

type hasUsage interface {
	// Usage returns an option of the value of the `usage` property for the
	// current RAML element.
	Usage() option.String
}