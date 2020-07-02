package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// BoolExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-07-02T14:19:59.953300998-04:00
type BoolExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) BoolExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() BoolExample

	// SetDescription sets this example's description value.
	SetDescription(string) BoolExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() BoolExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) BoolExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() BoolExample

	// Value returns this example's value.
	Value() option.Bool

	// SetValue sets this example's value.
	SetValue(v bool) BoolExample

	// UnsetValue removes this example's value.
	UnsetValue() BoolExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) BoolExample
}
