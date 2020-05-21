package raml

// ArrayType generated @ 2020-05-20T20:54:26.833516016-04:00
type ArrayType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) ArrayType

	// Default returns the current default value or nil of no
	// default value is set.
	Default() []interface{}

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault([]interface{}) ArrayType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() ArrayType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() ArrayExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex ArrayExample) ArrayType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() ArrayType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() ArrayExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples ArrayExampleMap) ArrayType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() ArrayType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) ArrayType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() ArrayType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) ArrayType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() ArrayType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) ArrayType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() ArrayType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) ArrayType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() ArrayType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) ArrayType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() ArrayType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []interface{}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]interface{}) ArrayType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() ArrayType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) ArrayType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() ArrayType

	SetRequired(bool) ArrayType
	// UniqueItems returns whether or not the current array
	// type definition requires items in the defined array to
	// be unique.
	//
	// Default value is false.  If this facet is set to false
	// when rendered, it will not appear in the output RAML.
	UniqueItems() bool

	// SetUniqueItems sets whether the current array type
	// definition requires items in the defined array to be
	// unique.
	//
	// Default value is false.  If this facet is set to false
	// when rendered, it will not appear in the output RAML.
	SetUniqueItems(uniq bool) ArrayType

	// MinItems returns the minimum number of items that are
	// required to appear in the array defined by the current
	// type definition.
	//
	// Default value is 0.  If this facet is set to 0 when
	// rendered, it will not appear in the output RAML.
	MinItems() uint

	// SetMinItems sets the minimum number of items that are
	// required to appear in the array defined by the current
	// type definition.
	//
	// Default value is 0.  If this facet is set to 0 when
	// rendered, it will not appear in the output RAML.
	SetMinItems(min uint) ArrayType

	// MaxItems returns the maximum number of items that may
	// appear in the array defined by the current type
	// definition.
	//
	// Default value is 2_147_483_647.  If this facet is set
	// to 2_147_483_647 when rendered, it will not appear in
	// the output RAML.
	MaxItems() uint

	// SetMaxItems sets the maximum number of items that may
	// appear in the array defined by the current type
	// definition.
	//
	// Default value is 2_147_483_647.  If this facet is set
	// to 2_147_483_647 when rendered, it will not appear in
	// the output RAML.
	SetMaxItems(max uint) ArrayType

	// Items returns the type definition for the items that
	// are allowed to appear in the array defined by the
	// current type definition.
	Items() DataType

	// SetItems sets the type definition for the items that
	// are allowed to appear in the array defined by the
	// current type definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetItems.
	SetItems(items DataType) ArrayType

	// UnsetItems removes the type definition for the items
	// that are allowed to appear in the array defined by the
	// current type definition.
	UnsetItems() ArrayType
}
