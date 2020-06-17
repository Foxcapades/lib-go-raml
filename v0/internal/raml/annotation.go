package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
)

func NewAnnotation() raml.Annotation {
	return raml.Annotation(raml.NewUntypedMap(1))
}
