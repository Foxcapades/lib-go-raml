package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// ObjectType generated @ 2020-10-19T13:48:24.9771134-04:00
type ObjectType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) ObjectType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Untyped

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(interface{}) ObjectType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() ObjectType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() ObjectExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex ObjectExample) ObjectType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() ObjectType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() ObjectExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples ObjectExampleMap) ObjectType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() ObjectType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) ObjectType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() ObjectType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) ObjectType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() ObjectType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) ObjectType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() ObjectType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) ObjectType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() ObjectType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) ObjectType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() ObjectType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []interface{}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]interface{}) ObjectType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() ObjectType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) ObjectType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() ObjectType

	SetRequired(bool) ObjectType
	// Properties returns a mutable map of the properties
	// defined on the current object type definition.
	Properties() PropertyMap

	// SetProperties replaces the map of properties on the
	// current object type definition with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetProperties.
	SetProperties(props PropertyMap) ObjectType

	// UnsetProperties removes the properties from the current
	// object type definition.
	UnsetProperties() ObjectType

	// MinProperties returns an option which will contain the
	// value of the "minProperties" facet on the current
	// object type definition if it is set.
	MinProperties() option.Uint

	// SetMinProperties sets the value of the "minProperties"
	// facet on the current object type definition.
	SetMinProperties(min uint) ObjectType

	// UnsetMinProperties removes the "minProperties" facet
	// from the current object type definition.
	UnsetMinProperties() ObjectType

	// MaxProperties returns an option which will contain the
	// value of the "maxProperties" facet on the current
	// object type definition if it is set.
	MaxProperties() option.Uint

	// SetMaxProperties sets the value of the "maxProperties"
	// facet on the current object type definition.
	SetMaxProperties(uint) ObjectType

	// UnsetMaxProperties removes the "maxProperties" facet
	// from the current object type definition.
	UnsetMaxProperties() ObjectType

	// AdditionalProperties returns the value of the
	// "additionalProperties" facet on the current object type
	// definition.
	//
	// The default value is true.  If this facet is set to its
	// default value when rendered it will not appear in the
	// output RAML.
	AdditionalProperties() bool

	// SetAdditionalProperties sets the value of the
	// "additionalProperties" facet on the current object type
	// definition.
	//
	// The default value is true.  If this facet is set to its
	// default value when rendered it will not appear in the
	// output RAML.
	SetAdditionalProperties(val bool) ObjectType

	// Discriminator returns an option which will contain the
	// value of the "discriminator" facet on the current
	// object type definition if it is set.
	Discriminator() option.String

	// SetDiscriminator sets the value of the "discriminator"
	// facet on the current object type definition.
	SetDiscriminator(facet string) ObjectType

	// UnsetDiscriminator removes the "discriminator" facet
	// from the current object type definition.
	UnsetDiscriminator() ObjectType

	// DiscriminatorValue returns an option which will contain
	// the value of the "discriminatorValue" facet on the
	// current object type definition if it is set.
	DiscriminatorValue() option.Untyped

	// SetDiscriminatorValue sets the value of the
	// "discriminatorValue" facet on the current object type
	// definition.
	SetDiscriminatorValue(val interface{}) ObjectType

	// UnsetDiscriminatorValue removes the
	// "discriminatorValue" facet from the current object type
	// definition.
	UnsetDiscriminatorValue() ObjectType
}
