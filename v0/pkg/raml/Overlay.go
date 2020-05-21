package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

// Overlay represents the contents of a RAML API overlay.
type Overlay interface {
	yaml.Unmarshaler
	yaml.Marshaler

	// Title returns the value of the title property set on the current RAML
	// element.
	Title() string

	// SetTitle sets the value of the title property on the current RAML element.
	SetTitle(title string) APISpec

	// Annotations returns a map of the annotations applied to the current
	// RAML element.
	Annotations() AnnotationMap

	// SetAnnotations replaces all annotations applied to the current RAML element
	// with those defined in the given map.
	//
	// Passing this method a nil value is equivalent to calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) APISpec

	// UnsetAnnotations removes all annotations applied to the current RAML
	// element.
	UnsetAnnotations() APISpec

	AnnotationTypes() UntypedMap

	// Description returns an option of the contents of the `description` facet
	// applied to the current RAML element.
	//
	// If the element does not have a description applied, the returned option
	// will be empty.
	Description() option.String

	// SetDescription sets the value of the description property for the current
	// RAML element.
	SetDescription(desc string) APISpec

	// UnsetDescription removes the description property from the current RAML
	// element.
	UnsetDescription() APISpec

	// ExtraFacets returns a mutable map of extra/custom facets applied to the
	// current RAML element.
	ExtraFacets() AnyMap

	// ResourceTypes returns the resourceTypes map defined on the current RAML
	// element.
	ResourceTypes() UntypedMap

	// SetResourceTypes replaces the resourceTypes map defined on the current RAML
	// element with the given value.
	//
	// Passing this method a nil value is equivalent to calling
	// UnsetResourceTypes.
	SetResourceTypes(resTypes UntypedMap) APISpec

	// UnsetResourceTypes removes the resourceTypes map from the current RAML
	// element.
	UnsetResourceTypes() APISpec

	SecuritySchemes() UntypedMap

	// Traits returns the map of trait definitions for the current RAML element.
	Traits() UntypedMap

	// SetTraits replaces the map of trait definitions on the current RAML element
	// with the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetTraits.
	SetTraits(traits UntypedMap) APISpec

	// UnsetTraits removes all trait definitions from the current RAML element.
	UnsetTraits() APISpec

	// Types returns the value of the `types` or `schemas` property defined on
	// the current document.
	Types() DataTypeMap

	// SetTypes replaces the map of types defined on the current RAML element with
	// the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetTypes.
	SetTypes(types DataTypeMap) APISpec

	// UnsetTypes removes all types from the current RAML element.
	UnsetTypes() APISpec

	// Schemas is an alias for the Types method.
	Schemas() DataTypeMap

	// SetSchemas is an alias for the SetTypes method.
	SetSchemas(types DataTypeMap) APISpec

	// UnsetSchemas is an alias for the UnsetTypes method.
	UnsetSchemas() APISpec

	// Documentation returns the array of documentation elements defined on the
	// current RAML element.
	Documentation() []Documentation

	// SetDocumentation sets the documentation array for the current RAML element.
	//
	// Passing this method a nil value is equivalent to calling
	// UnsetDocumentation.
	SetDocumentation(docs []Documentation) APISpec

	// UnsetDocumentation removes the documentation property and all its values
	// from the current RAML element.
	UnsetDocumentation() APISpec

	BaseURIParameters() UntypedMap

	Resources() UntypedMap
}
