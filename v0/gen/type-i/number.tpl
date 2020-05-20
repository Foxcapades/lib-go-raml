{{define "number"}}
	// Minimum returns an option which will contain the value
	// of the "minimum" facet for the current number type
	// definition if it is set.
	Minimum() option.Float64

	// SetMinimum sets the value of the "minimum" facet for
	// current number type definition.
	SetMinimum(min float64) NumberType

	// UnsetMinimum removes the "minimum" facet from the
	// current number type definition.
	UnsetMinimum() NumberType

	// Maximum returns an option which will contain the value
	// of the "maximum" facet for the current number type
	// definition if it is set.
	Maximum() option.Float64

	// SetMaximum sets the value of the "maximum" facet for
	// current number type definition.
	SetMaximum(max float64) NumberType

	// UnsetMinimum removes the "maximum" facet from the
	// current number type definition.
	UnsetMaximum() NumberType

	// Format returns the value of the current number type
	// definition's "format" facet, or nil if no format is
	// set.
	Format() NumberFormat

	// SetFormat sets the value of the current number type
	// definition's "format" facet.
	SetFormat(format NumberFormat) NumberType

	// UnsetFormat removes the "format" facet from the current
	// number type definition.
	UnsetFormat() NumberType

	// MultipleOf returns an option which will contain the
	// value of the current number type definition's
	// "multipleOf" facet if it is set.
	MultipleOf() option.Float64

	// SetMultipleOf sets the value of the "multipleOf" facet
	// for the current number type definition.
	SetMultipleOf(of float64) NumberType

	// UnsetMultipleOf removes the "multipleOf" facet from the
	// current number type definition.
	UnsetMultipleOf() NumberType
{{end}}