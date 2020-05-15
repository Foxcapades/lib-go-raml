package raml

type RamlType interface {
	GetType() string
	ToRaml() (string, error)
}
