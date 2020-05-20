package raml

// DatetimeOnlyExample defines a single example attached to a DataType
// or Property definition.
// Generated @ 2020-05-19T21:52:50.990676962-04:00
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
	Value() string

	// SetValue sets this example's value.
	SetValue(string) DatetimeOnlyExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) DatetimeOnlyExample
}
