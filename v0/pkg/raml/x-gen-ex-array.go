package raml

// ArrayExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-20T21:46:00.638880955-04:00
type ArrayExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) ArrayExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() ArrayExample

	// SetDescription sets this example's description value.
	SetDescription(string) ArrayExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() ArrayExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) ArrayExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() ArrayExample

	// Value returns this example's value.
	Value() []interface{}

	// SetValue sets this example's value.
	SetValue([]interface{}) ArrayExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) ArrayExample
}
