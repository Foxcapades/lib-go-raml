package raml

type Property interface {
	DataType

	Required() bool
}
