package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// DatetimeOnlyExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-21T22:19:27.027107376-04:00
type DatetimeOnlyExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) DatetimeOnlyExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() DatetimeOnlyExample

	// SetDescription sets this example's description value.
	SetDescription(string) DatetimeOnlyExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() DatetimeOnlyExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) DatetimeOnlyExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() DatetimeOnlyExample

	// Value returns this example's value.
	Value() option.String

	// SetValue sets this example's value.
	SetValue(v string) DatetimeOnlyExample

	// UnsetValue removes this example's value.
	UnsetValue() DatetimeOnlyExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) DatetimeOnlyExample
}
