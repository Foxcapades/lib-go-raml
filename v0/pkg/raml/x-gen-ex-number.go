package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// NumberExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2023-01-17T11:42:46.993195814-05:00
type NumberExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) NumberExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() NumberExample

	// SetDescription sets this example's description value.
	SetDescription(string) NumberExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() NumberExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) NumberExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() NumberExample

	// Value returns this example's value.
	Value() option.Float64

	// SetValue sets this example's value.
	SetValue(v float64) NumberExample

	// UnsetValue removes this example's value.
	UnsetValue() NumberExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) NumberExample
}
