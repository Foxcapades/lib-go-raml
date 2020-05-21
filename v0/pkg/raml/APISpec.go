package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
	"io"
)

// APISpec represents a full root RAML document.
type APISpec interface {
	yaml.Marshaler
	yaml.Unmarshaler

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

	// Uses returns the map of defined library imports for the current RAML
	// element.
	Uses() StringMap

	// SetUses replaces the map of defined library imports for the current RAML
	// element with the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetUses.
	SetUses(uses StringMap) APISpec

	// UnsetUses removes the map of defined library imports from the current RAML
	// element.
	UnsetUses() APISpec

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

	// Protocols returns the values of the protocol array set on the current RAML
	// element.
	Protocols() []string

	// SetProtocols sets the protocol array for the current RAML element to the
	// given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetProtocols.
	SetProtocols(protocols []string) APISpec

	// UnsetProtocols removes the protocols property and all its values from the
	// current RAML element.
	UnsetProtocols() APISpec

	// SecuredBy returns the value of the securedBy array defined on the current
	// RAML element.
	SecuredBy() []string

	// SetSecuredBy sets the securedBy array for the current RAML element.
	//
	// Passing this method a nil value is equivalent to calling UnsetSecuredBy.
	SetSecuredBy(securedBy []string) APISpec

	// UnsetSecuredBy removes the securedBy property and all its values from the
	// current RAML element.
	UnsetSecuredBy() APISpec
	// Version returns an option of the declared version for this API spec.
	//
	// If no version has been declared, the returned option will be empty.
	Version() option.String

	// SetVersion sets the version value for this API spec.
	SetVersion(version string) APISpec

	// UnsetVersion removes the version declaration from this API spec.
	UnsetVersion() APISpec

	// BaseURI returns an option of the baseUri value declared on the current
	// RAML element.
	//
	// If no baseUri value has been set, the returned option will be empty.
	BaseURI() option.String

	// SetBaseURI sets the baseUri value for the current RAML element.
	SetBaseURI(baseUri string) APISpec

	// UnsetBaseURI removes the baseUri value from the current RAML element.
	UnsetBaseURI() APISpec

	// MediaTypes returns the values of the mediaTypes array set on the current
	// RAML element.
	//
	// When parsed, if the RAML document has a single string value set for the
	// mediaTypes property, this slice will be populated with that single value.
	//
	// When rendered, if this slice contains a single value, it will be rendered
	// as a single string rather than an array.
	MediaTypes() []string

	// SetMediaTypes sets the mediaTypes array for the current RAML element to the
	// given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetMediaTypes.
	SetMediaTypes(mediaTypes []string) APISpec

	// UnsetMediaTypes removes the mediaTypes property and all its values from the
	// current RAML element.
	UnsetMediaTypes() APISpec

	BaseURIParameters() UntypedMap

	Resources() UntypedMap

	// WriteRAML writes out the current RAML element as a standalone document to
	// the given writer.
	WriteRAML(w io.Writer) error
}
