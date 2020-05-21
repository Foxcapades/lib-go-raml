{{define "example" -}}
package raml

// {{.Name}}Example defines a single example attached to a DataType
// or Property definition.
//
// Generated @ {{.Time}}
type {{.Name}}Example interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) {{.Name}}Example

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() {{.Name}}Example

	// SetDescription sets this example's description value.
	SetDescription(string) {{.Name}}Example

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() {{.Name}}Example

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) {{.Name}}Example

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() {{.Name}}Example

	// Value returns this example's value.
	Value() {{.Type}}

	// SetValue sets this example's value.
	SetValue({{.Type}}) {{.Name}}Example

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) {{.Name}}Example
}
{{end}}