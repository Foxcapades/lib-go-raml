package raml

type Resource interface {
	hasAnnotations
	hasDescription
	displayed
	hasFacets

	FullPath() string
	PathSegment() string
}
