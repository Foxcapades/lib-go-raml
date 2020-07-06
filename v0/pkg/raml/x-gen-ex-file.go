package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// FileExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-07-06T12:49:37.48714807-04:00
type FileExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) FileExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() FileExample

	// SetDescription sets this example's description value.
	SetDescription(string) FileExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() FileExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) FileExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() FileExample

	// Value returns this example's value.
	Value() option.Untyped

	// SetValue sets this example's value.
	SetValue(v interface{}) FileExample

	// UnsetValue removes this example's value.
	UnsetValue() FileExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) FileExample
}
