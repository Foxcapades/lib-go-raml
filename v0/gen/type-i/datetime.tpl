package x

type x interface {
	{{define "datetime" -}}
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
	{{- end}}
}