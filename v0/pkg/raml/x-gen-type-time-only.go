package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// TimeOnlyType generated @ 2023-01-17T10:02:54.294844187-05:00
type TimeOnlyType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) TimeOnlyType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.String

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(string) TimeOnlyType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() TimeOnlyType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() TimeOnlyExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex TimeOnlyExample) TimeOnlyType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() TimeOnlyType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() TimeOnlyExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples TimeOnlyExampleMap) TimeOnlyType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() TimeOnlyType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) TimeOnlyType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() TimeOnlyType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) TimeOnlyType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() TimeOnlyType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) TimeOnlyType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() TimeOnlyType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) TimeOnlyType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() TimeOnlyType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) TimeOnlyType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() TimeOnlyType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []string

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]string) TimeOnlyType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() TimeOnlyType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) TimeOnlyType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() TimeOnlyType

	SetRequired(bool) TimeOnlyType
}
