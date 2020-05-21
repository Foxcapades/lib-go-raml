package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

// Library represents the contents of a Library type RAML fragment.
type Library interface {
	yaml.Unmarshaler
	yaml.Marshaler

	// Annotations returns a map of the annotations applied to the current
	// RAML element.
	Annotations() AnnotationMap

	// SetAnnotations replaces all annotations applied to the current RAML element
	// with those defined in the given map.
	//
	// Passing this method a nil value is equivalent to calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) Library

	// UnsetAnnotations removes all annotations applied to the current RAML
	// element.
	UnsetAnnotations() Library

	AnnotationTypes() UntypedMap

	// ExtraFacets returns a mutable map of extra/custom facets applied to the
	// current RAML element.
	ExtraFacets() AnyMap

	// Uses returns the map of defined library imports for the current RAML
	// element.
	Uses() StringMap

	// SetUses replaces the map of defined library imports for the current RAML
	// element with the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetUses.
	SetUses(uses StringMap) Library

	// UnsetUses removes the map of defined library imports from the current RAML
	// element.
	UnsetUses() Library

	// ResourceTypes returns the resourceTypes map defined on the current RAML
	// element.
	ResourceTypes() UntypedMap

	// SetResourceTypes replaces the resourceTypes map defined on the current RAML
	// element with the given value.
	//
	// Passing this method a nil value is equivalent to calling
	// UnsetResourceTypes.
	SetResourceTypes(resTypes UntypedMap) Library

	// UnsetResourceTypes removes the resourceTypes map from the current RAML
	// element.
	UnsetResourceTypes() Library

	SecuritySchemes() UntypedMap

	// Traits returns the map of trait definitions for the current RAML element.
	Traits() UntypedMap

	// SetTraits replaces the map of trait definitions on the current RAML element
	// with the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetTraits.
	SetTraits(traits UntypedMap) Library

	// UnsetTraits removes all trait definitions from the current RAML element.
	UnsetTraits() Library

	// Types returns the value of the `types` or `schemas` property defined on
	// the current document.
	Types() DataTypeMap

	// SetTypes replaces the map of types defined on the current RAML element with
	// the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetTypes.
	SetTypes(types DataTypeMap) Library

	// UnsetTypes removes all types from the current RAML element.
	UnsetTypes() Library

	// Schemas is an alias for the Types method.
	Schemas() DataTypeMap

	// SetSchemas is an alias for the SetTypes method.
	SetSchemas(types DataTypeMap) Library

	// UnsetSchemas is an alias for the UnsetTypes method.
	UnsetSchemas() Library

	// Usage returns an option of the value of the `usage` property for the
	// current RAML element.
	Usage() option.String

	// SetUsage sets the value of the usage property on the current RAML element.
	SetUsage(usage string) Library

	// UnsetUsage removes the usage property from the current RAML element.
	UnsetUsage() Library
}
