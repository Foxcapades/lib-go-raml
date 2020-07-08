package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// StringExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-07-08T13:31:34.46078727-04:00
type StringExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) StringExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() StringExample

	// SetDescription sets this example's description value.
	SetDescription(string) StringExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() StringExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) StringExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() StringExample

	// Value returns this example's value.
	Value() option.String

	// SetValue sets this example's value.
	SetValue(v string) StringExample

	// UnsetValue removes this example's value.
	UnsetValue() StringExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) StringExample
}
