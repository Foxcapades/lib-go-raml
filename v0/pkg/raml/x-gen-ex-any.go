package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// AnyExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-07-06T13:52:18.264181542-04:00
type AnyExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) AnyExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() AnyExample

	// SetDescription sets this example's description value.
	SetDescription(string) AnyExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() AnyExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) AnyExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() AnyExample

	// Value returns this example's value.
	Value() option.Untyped

	// SetValue sets this example's value.
	SetValue(v interface{}) AnyExample

	// UnsetValue removes this example's value.
	UnsetValue() AnyExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) AnyExample
}
