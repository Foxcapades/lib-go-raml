{{define "object"}}
	// Properties returns a mutable map of the properties
	// defined on the current object type definition.
	Properties() PropertyMap

	// SetProperties replaces the map of properties on the
	// current object type definition with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetProperties.
	SetProperties(props PropertyMap) ObjectType

	// UnsetProperties removes the properties from the current
	// object type definition.
	UnsetProperties() ObjectType

	// MinProperties returns an option which will contain the
	// value of the "minProperties" facet on the current
	// object type definition if it is set.
	MinProperties() option.Uint

	// SetMinProperties sets the value of the "minProperties"
	// facet on the current object type definition.
	SetMinProperties(min uint) ObjectType

	// UnsetMinProperties removes the "minProperties" facet
	// from the current object type definition.
	UnsetMinProperties() ObjectType

	// MaxProperties returns an option which will contain the
	// value of the "maxProperties" facet on the current
	// object type definition if it is set.
	MaxProperties() option.Uint

	// SetMaxProperties sets the value of the "maxProperties"
	// facet on the current object type definition.
	SetMaxProperties(uint) ObjectType

	// UnsetMaxProperties removes the "maxProperties" facet
	// from the current object type definition.
	UnsetMaxProperties() ObjectType

	// AdditionalProperties returns the value of the
	// "additionalProperties" facet on the current object type
	// definition.
	//
	// The default value is true.  If this facet is set to its
	// default value when rendered it will not appear in the
	// output RAML.
	AdditionalProperties() bool

	// SetAdditionalProperties sets the value of the
	// "additionalProperties" facet on the current object type
	// definition.
	//
	// The default value is true.  If this facet is set to its
	// default value when rendered it will not appear in the
	// output RAML.
	SetAdditionalProperties(val bool) ObjectType

	// Discriminator returns an option which will contain the
	// value of the "discriminator" facet on the current
	// object type definition if it is set.
	Discriminator() option.String

	// SetDiscriminator sets the value of the "discriminator"
	// facet on the current object type definition.
	SetDiscriminator(facet string) ObjectType

	// UnsetDiscriminator removes the "discriminator" facet
	// from the current object type definition.
	UnsetDiscriminator() ObjectType

	// DiscriminatorValue returns an option which will contain
	// the value of the "discriminatorValue" facet on the
	// current object type definition if it is set.
	DiscriminatorValue() option.Untyped

	// SetDiscriminatorValue sets the value of the
	// "discriminatorValue" facet on the current object type
	// definition.
	SetDiscriminatorValue(val interface{}) ObjectType

	// UnsetDiscriminatorValue removes the
	// "discriminatorValue" facet from the current object type
	// definition.
	UnsetDiscriminatorValue() ObjectType
{{end}}