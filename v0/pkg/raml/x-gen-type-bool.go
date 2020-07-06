package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// BoolType generated @ 2020-07-06T12:49:37.941034901-04:00
type BoolType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) BoolType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Bool

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(bool) BoolType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() BoolType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() BoolExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex BoolExample) BoolType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() BoolType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() BoolExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples BoolExampleMap) BoolType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() BoolType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) BoolType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() BoolType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) BoolType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() BoolType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) BoolType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() BoolType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) BoolType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() BoolType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) BoolType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() BoolType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []bool

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]bool) BoolType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() BoolType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) BoolType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() BoolType

	SetRequired(bool) BoolType
}
