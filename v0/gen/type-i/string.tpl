package x

type x interface {
	{{define "string" -}}
	// Pattern returns an option which will contain the regex
	// pattern for the current string type definition if it is
	// set.
	Pattern() option.String

	// SetPattern sets the regex pattern for the current
	// string type definition.
	SetPattern(string) StringType

	// UnsetPattern removes the regex pattern facet from the
	// current string type definition.
	UnsetPattern() StringType

	// MinLength returns the minLength facet value set on the
	// current string type definition.
	//
	// Default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	MinLength() uint

	// SetMinLength sets the value for the minLength facet on
	// the current string type definition.
	//
	// Default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	SetMinLength(min uint) StringType

	// MaxLength returns the maxLength facet value set on the
	// current string type definition.
	MaxLength() option.Uint

	// SetMaxLength sets the value for the maxLength facet on
	// the current string type definition.
	SetMaxLength(ln uint) StringType

	UnsetMaxLength() StringType
	{{- end}}
}