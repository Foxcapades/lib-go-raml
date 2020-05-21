package raml

// DateOnlyExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-20T20:54:25.600656868-04:00
type DateOnlyExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) DateOnlyExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() DateOnlyExample

	// SetDescription sets this example's description value.
	SetDescription(string) DateOnlyExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() DateOnlyExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) DateOnlyExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() DateOnlyExample

	// Value returns this example's value.
	Value() string

	// SetValue sets this example's value.
	SetValue(string) DateOnlyExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) DateOnlyExample
}
