package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
)

func NewAnnotation() raml.Annotation {
	return raml.Annotation(NewUntypedMap())
}
