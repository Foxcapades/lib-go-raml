package raml

type hasAnnotations interface {
	{{define "hasAnnotations"}}
	// Annotations returns a map of the annotations applied to the current
	// RAML element.
	Annotations() AnnotationMap

	// SetAnnotations replaces all annotations applied to the current RAML element
	// with those defined in the given map.
	//
	// Passing this method a nil value is equivalent to calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) {{.}}

	// UnsetAnnotations removes all annotations applied to the current RAML
	// element.
	UnsetAnnotations() {{.}}
	{{- end}}
}
