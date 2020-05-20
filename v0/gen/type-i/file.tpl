{{define "file"}}
	// FileTypes returns a slice of the "fileTypes" facet set
	// on the current file type definition.
	FileTypes() []string

	// SetFileTypes replaces the values of the "fileTypes"
	// facet on the current file type definition with the
	// given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFileTypes.
	SetFileTypes([]string) FileType

	// UnsetFileTypes clears the "fileTypes" facet for the
	// current file type definition.
	UnsetFileTypes() FileType

	// MinLength returns the value for the "minLength" facet
	// on the current file type definition.
	//
	// The default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	MinLength() uint

	// SetMinLength sets the value of the "minLength" facet on
	// the current file type definition.
	//
	// The default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	SetMinLength(min uint) FileType

	// MaxLength returns the value for the "maxLength" facet
	// on the current file type definition.
	//
	// The default value is 2_147_483_647.  If this facet is
	// set to its default value when rendered, it will not
	// appear in the output RAML.
	MaxLength() uint

	// SetMaxLength sets the value of the "maxLength" facet on
	// the current file type definition.
	//
	// The default value is 2_147_483_647.  If this facet is
	// set to its default value when rendered, it will not
	// appear in the output RAML.
	SetMaxLength(max uint) FileType
{{end}}