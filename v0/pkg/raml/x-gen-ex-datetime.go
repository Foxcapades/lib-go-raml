package raml

// DatetimeExample defines a single example attached to a DataType
// or Property definition.
// Generated @ 2020-05-20T18:40:12.806494881-04:00
type DatetimeExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) DatetimeExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() DatetimeExample

	// SetDescription sets this example's description value.
	SetDescription(string) DatetimeExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() DatetimeExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) DatetimeExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() DatetimeExample

	// Value returns this example's value.
	Value() string

	// SetValue sets this example's value.
	SetValue(string) DatetimeExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) DatetimeExample
}
