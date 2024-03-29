package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// TimeOnlyExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2023-01-17T11:42:46.993195814-05:00
type TimeOnlyExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) TimeOnlyExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() TimeOnlyExample

	// SetDescription sets this example's description value.
	SetDescription(string) TimeOnlyExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() TimeOnlyExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) TimeOnlyExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() TimeOnlyExample

	// Value returns this example's value.
	Value() option.String

	// SetValue sets this example's value.
	SetValue(v string) TimeOnlyExample

	// UnsetValue removes this example's value.
	UnsetValue() TimeOnlyExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) TimeOnlyExample
}
