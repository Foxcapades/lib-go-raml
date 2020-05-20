package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// NumberType generated @ 2020-05-20T18:40:13.095690448-04:00
type NumberType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) NumberType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Float64

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(float64) NumberType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() NumberType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() NumberExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex NumberExample) NumberType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() NumberType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() NumberExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples NumberExampleMap) NumberType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() NumberType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) NumberType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() NumberType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) NumberType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() NumberType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) NumberType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() NumberType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) NumberType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() NumberType

	// SetXml sets the current DataType definition's xml
	// serialization settings object to the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetXml.
	SetXml(Xml) NumberType

	// Removes the xml serialization settings object from the
	// current DataType definition.
	UnsetXml() NumberType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []float64

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]float64) NumberType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() NumberType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) NumberType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() NumberType

	SetRequired(bool) NumberType
	// Minimum returns an option which will contain the value
	// of the "minimum" facet for the current number type
	// definition if it is set.
	Minimum() option.Float64

	// SetMinimum sets the value of the "minimum" facet for
	// current number type definition.
	SetMinimum(min float64) NumberType

	// UnsetMinimum removes the "minimum" facet from the
	// current number type definition.
	UnsetMinimum() NumberType

	// Maximum returns an option which will contain the value
	// of the "maximum" facet for the current number type
	// definition if it is set.
	Maximum() option.Float64

	// SetMaximum sets the value of the "maximum" facet for
	// current number type definition.
	SetMaximum(max float64) NumberType

	// UnsetMinimum removes the "maximum" facet from the
	// current number type definition.
	UnsetMaximum() NumberType

	// Format returns the value of the current number type
	// definition's "format" facet, or nil if no format is
	// set.
	Format() NumberFormat

	// SetFormat sets the value of the current number type
	// definition's "format" facet.
	SetFormat(format NumberFormat) NumberType

	// UnsetFormat removes the "format" facet from the current
	// number type definition.
	UnsetFormat() NumberType

	// MultipleOf returns an option which will contain the
	// value of the current number type definition's
	// "multipleOf" facet if it is set.
	MultipleOf() option.Float64

	// SetMultipleOf sets the value of the "multipleOf" facet
	// for the current number type definition.
	SetMultipleOf(of float64) NumberType

	// UnsetMultipleOf removes the "multipleOf" facet from the
	// current number type definition.
	UnsetMultipleOf() NumberType
}
