package raml

type Type interface {
	GetType() string
	SetType(t string)
	ToRAML() (string, error)
}
