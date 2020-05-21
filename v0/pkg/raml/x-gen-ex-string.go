package raml

// StringExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-20T21:46:00.638880955-04:00
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
	Value() string

	// SetValue sets this example's value.
	SetValue(string) StringExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) StringExample
}
