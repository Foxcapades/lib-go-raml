package raml

type Response interface {
	hasAnnotations
	hasDescription
	headered

	// TODO: Body
}
