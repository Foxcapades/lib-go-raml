package raml

type Type interface {
	GetType() string
	ToRAML() (string, error)
}
