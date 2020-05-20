package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
)

func NewAnnotation(log *logrus.Entry) raml.Annotation {
	return raml.Annotation(NewUntypedMap(log))
}
