package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
)

type renderable interface {
	render() bool
}

type concreteType interface {
	raml.Unmarshaler
	raml.Marshaler
	raml.DataType
	renderable

	assign(key, val interface{}, log *logrus.Entry) error
	marshal(out raml.AnyMap) error
}