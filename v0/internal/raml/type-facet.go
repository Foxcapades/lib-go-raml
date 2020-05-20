package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
)

func NewFacet(log *logrus.Entry) raml.Facet {
	return NewDataTypeMap(log)
}
