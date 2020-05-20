package raml

// BoolExample defines a single example attached to a DataType
// or Property definition.
// Generated @ 2020-05-19T21:52:50.990676962-04:00
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
	Value() bool

	// SetValue sets this example's value.
	SetValue(bool) BoolExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) BoolExample
}
