package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// DateOnlyType generated @ 2020-05-20T20:54:26.833516016-04:00
type DateOnlyType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) DateOnlyType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.String

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(string) DateOnlyType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() DateOnlyType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() DateOnlyExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex DateOnlyExample) DateOnlyType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() DateOnlyType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() DateOnlyExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples DateOnlyExampleMap) DateOnlyType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() DateOnlyType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) DateOnlyType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() DateOnlyType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) DateOnlyType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() DateOnlyType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) DateOnlyType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() DateOnlyType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) DateOnlyType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() DateOnlyType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) DateOnlyType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() DateOnlyType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []string

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]string) DateOnlyType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() DateOnlyType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) DateOnlyType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() DateOnlyType

	SetRequired(bool) DateOnlyType
}
