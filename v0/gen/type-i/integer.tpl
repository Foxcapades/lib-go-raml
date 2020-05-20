package x

type x interface {
	{{define "integer" -}}
	// Minimum returns an option which will contain the value
	// of the "minimum" facet for the current integer type
	// definition if it is set.
	Minimum() option.Int64

	// SetMinimum sets the value of the "minimum" facet for
	// current integer type definition.
	SetMinimum(min int64) IntegerType

	// UnsetMinimum removes the "minimum" facet from the
	// current integer type definition.
	UnsetMinimum() IntegerType

	// Maximum returns an option which will contain the value
	// of the "maximum" facet for the current integer type
	// definition if it is set.
	Maximum() option.Int64

	// SetMaximum sets the value of the "maximum" facet for
	// current integer type definition.
	SetMaximum(max int64) IntegerType

	// UnsetMinimum removes the "maximum" facet from the
	// current integer type definition.
	UnsetMaximum() IntegerType

	// Format returns the value of the current integer type
	// definition's "format" facet, or nil if no format is
	// set.
	Format() IntegerFormat

	// SetFormat sets the value of the current integer type
	// definition's "format" facet.
	SetFormat(format IntegerFormat) IntegerType

	// UnsetFormat removes the "format" facet from the current
	// integer type definition.
	UnsetFormat() IntegerType

	// MultipleOf returns an option which will contain the
	// value of the current integer type definition's
	// "multipleOf" facet if it is set.
	MultipleOf() option.Float64

	// SetMultipleOf sets the value of the "multipleOf" facet
	// for the current integer type definition.
	SetMultipleOf(of float64) IntegerType

	// UnsetMultipleOf removes the "multipleOf" facet from the
	// current integer type definition.
	UnsetMultipleOf() IntegerType
	{{- end}}
}