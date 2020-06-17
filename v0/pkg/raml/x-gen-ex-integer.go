package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// IntegerExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-25T18:15:33.802521588-04:00
type IntegerExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) IntegerExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() IntegerExample

	// SetDescription sets this example's description value.
	SetDescription(string) IntegerExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() IntegerExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) IntegerExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() IntegerExample

	// Value returns this example's value.
	Value() option.Int64

	// SetValue sets this example's value.
	SetValue(v int64) IntegerExample

	// UnsetValue removes this example's value.
	UnsetValue() IntegerExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) IntegerExample
}
