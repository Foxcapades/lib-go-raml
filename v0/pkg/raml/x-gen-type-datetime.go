package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// DatetimeType generated @ 2020-07-06T13:52:18.671712454-04:00
type DatetimeType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) DatetimeType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.String

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(string) DatetimeType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() DatetimeType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() DatetimeExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex DatetimeExample) DatetimeType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() DatetimeType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() DatetimeExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples DatetimeExampleMap) DatetimeType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() DatetimeType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) DatetimeType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() DatetimeType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) DatetimeType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() DatetimeType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) DatetimeType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() DatetimeType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) DatetimeType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() DatetimeType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) DatetimeType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() DatetimeType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []string

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]string) DatetimeType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() DatetimeType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) DatetimeType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() DatetimeType

	SetRequired(bool) DatetimeType
	// Format returns the value of the "format" facet on the
	// current datetime type definition or nil if the facet is
	// not set.
	Format() DateFormat

	// SetFormat sets the value of the "format" facet on the
	// current datetime type definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFormat.
	//
	// The default value is Rfc3339.  If this facet is set to
	// its default value when rendered, it will not appear in
	// the output RAML.
	SetFormat(format DateFormat) DatetimeType

	// UnsetFormat removes the "format" facet from the current
	// datetime type definition.
	UnsetFormat() DatetimeType
}
