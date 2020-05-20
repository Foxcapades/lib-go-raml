package raml

import "github.com/Foxcapades/goop/v1/pkg/option"
// UnionType generated @ 2020-05-20T00:33:46.349824232-04:00
type UnionType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) UnionType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Untyped

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(interface{}) UnionType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() UnionType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() UnionExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex UnionExample) UnionType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() UnionType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() UnionExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples UnionExampleMap) UnionType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() UnionType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) UnionType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() UnionType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) UnionType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() UnionType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) UnionType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() UnionType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) UnionType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() UnionType

	// SetXml sets the current DataType definition's xml
	// serialization settings object to the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetXml.
	SetXml(Xml) UnionType

	// Removes the xml serialization settings object from the
	// current DataType definition.
	UnsetXml() UnionType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []interface{}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]interface{}) UnionType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() UnionType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) UnionType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() UnionType

	SetRequired(bool) UnionType

}
