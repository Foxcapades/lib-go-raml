package f

type f interface {
	{{define "array" -}}
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
	{{- end}}
}