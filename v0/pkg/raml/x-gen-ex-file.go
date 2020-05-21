package raml

// FileExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-20T20:54:25.600656868-04:00
type FileExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) FileExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() FileExample

	// SetDescription sets this example's description value.
	SetDescription(string) FileExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() FileExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) FileExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() FileExample

	// Value returns this example's value.
	Value() interface{}

	// SetValue sets this example's value.
	SetValue(interface{}) FileExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) FileExample
}
