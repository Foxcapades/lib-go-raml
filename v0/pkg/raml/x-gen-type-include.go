package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// IncludeType generated @ 2025-04-30T18:50:42.201241268-04:00
type IncludeType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) IncludeType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Untyped

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(interface{}) IncludeType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() IncludeType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) IncludeType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() IncludeType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) IncludeType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() IncludeType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) IncludeType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() IncludeType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) IncludeType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() IncludeType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) IncludeType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() IncludeType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []interface{}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]interface{}) IncludeType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() IncludeType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) IncludeType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() IncludeType

	SetRequired(bool) IncludeType
}
