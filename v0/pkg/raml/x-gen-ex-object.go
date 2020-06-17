package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// ObjectExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-25T18:15:33.802521588-04:00
type ObjectExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) ObjectExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() ObjectExample

	// SetDescription sets this example's description value.
	SetDescription(string) ObjectExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() ObjectExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) ObjectExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() ObjectExample

	// Value returns this example's value.
	Value() option.Untyped

	// SetValue sets this example's value.
	SetValue(v interface{}) ObjectExample

	// UnsetValue removes this example's value.
	UnsetValue() ObjectExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) ObjectExample
}
