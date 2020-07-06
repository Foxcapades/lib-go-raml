package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// AnyType generated @ 2020-07-06T12:49:37.941034901-04:00
type AnyType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) AnyType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Untyped

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(interface{}) AnyType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() AnyType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() AnyExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex AnyExample) AnyType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() AnyType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() AnyExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples AnyExampleMap) AnyType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() AnyType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) AnyType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() AnyType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) AnyType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() AnyType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) AnyType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() AnyType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) AnyType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() AnyType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) AnyType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() AnyType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []interface{}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]interface{}) AnyType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() AnyType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) AnyType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() AnyType

	SetRequired(bool) AnyType
}
