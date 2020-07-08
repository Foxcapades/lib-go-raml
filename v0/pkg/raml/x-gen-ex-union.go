package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// UnionExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-07-08T13:31:34.46078727-04:00
type UnionExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) UnionExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() UnionExample

	// SetDescription sets this example's description value.
	SetDescription(string) UnionExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() UnionExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) UnionExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() UnionExample

	// Value returns this example's value.
	Value() option.Untyped

	// SetValue sets this example's value.
	SetValue(v interface{}) UnionExample

	// UnsetValue removes this example's value.
	UnsetValue() UnionExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) UnionExample
}
