package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// IntegerType generated @ 2020-05-20T20:54:26.833516016-04:00
type IntegerType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) IntegerType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Int64

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(int64) IntegerType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() IntegerType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() IntegerExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex IntegerExample) IntegerType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() IntegerType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() IntegerExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples IntegerExampleMap) IntegerType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() IntegerType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) IntegerType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() IntegerType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) IntegerType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() IntegerType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) IntegerType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() IntegerType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) IntegerType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() IntegerType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) IntegerType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() IntegerType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []int64

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]int64) IntegerType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() IntegerType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) IntegerType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() IntegerType

	SetRequired(bool) IntegerType
	// Minimum returns an option which will contain the value
	// of the "minimum" facet for the current integer type
	// definition if it is set.
	Minimum() option.Int64

	// SetMinimum sets the value of the "minimum" facet for
	// current integer type definition.
	SetMinimum(min int64) IntegerType

	// UnsetMinimum removes the "minimum" facet from the
	// current integer type definition.
	UnsetMinimum() IntegerType

	// Maximum returns an option which will contain the value
	// of the "maximum" facet for the current integer type
	// definition if it is set.
	Maximum() option.Int64

	// SetMaximum sets the value of the "maximum" facet for
	// current integer type definition.
	SetMaximum(max int64) IntegerType

	// UnsetMinimum removes the "maximum" facet from the
	// current integer type definition.
	UnsetMaximum() IntegerType

	// Format returns the value of the current integer type
	// definition's "format" facet, or nil if no format is
	// set.
	Format() IntegerFormat

	// SetFormat sets the value of the current integer type
	// definition's "format" facet.
	SetFormat(format IntegerFormat) IntegerType

	// UnsetFormat removes the "format" facet from the current
	// integer type definition.
	UnsetFormat() IntegerType

	// MultipleOf returns an option which will contain the
	// value of the current integer type definition's
	// "multipleOf" facet if it is set.
	MultipleOf() option.Float64

	// SetMultipleOf sets the value of the "multipleOf" facet
	// for the current integer type definition.
	SetMultipleOf(of float64) IntegerType

	// UnsetMultipleOf removes the "multipleOf" facet from the
	// current integer type definition.
	UnsetMultipleOf() IntegerType
}
