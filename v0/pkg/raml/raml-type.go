package raml

type RamlType interface {
	ToRaml() (string, error)
}
