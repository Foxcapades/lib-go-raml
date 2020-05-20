package raml

type Method interface {
	hasAnnotations
	displayed
	hasDescription
	hasFacets
	headered

	MethodName() string

	// TODO: QueryString() ObjectType | ScalarType
}
