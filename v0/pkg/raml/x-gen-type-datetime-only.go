package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// DatetimeOnlyType generated @ 2020-05-20T18:40:13.095690448-04:00
type DatetimeOnlyType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) DatetimeOnlyType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.String

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(string) DatetimeOnlyType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() DatetimeOnlyType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() DatetimeOnlyExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex DatetimeOnlyExample) DatetimeOnlyType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() DatetimeOnlyType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() DatetimeOnlyExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples DatetimeOnlyExampleMap) DatetimeOnlyType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() DatetimeOnlyType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) DatetimeOnlyType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() DatetimeOnlyType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) DatetimeOnlyType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() DatetimeOnlyType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) DatetimeOnlyType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() DatetimeOnlyType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) DatetimeOnlyType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() DatetimeOnlyType

	// SetXml sets the current DataType definition's xml
	// serialization settings object to the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetXml.
	SetXml(Xml) DatetimeOnlyType

	// Removes the xml serialization settings object from the
	// current DataType definition.
	UnsetXml() DatetimeOnlyType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []string

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]string) DatetimeOnlyType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() DatetimeOnlyType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) DatetimeOnlyType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() DatetimeOnlyType

	SetRequired(bool) DatetimeOnlyType
}
