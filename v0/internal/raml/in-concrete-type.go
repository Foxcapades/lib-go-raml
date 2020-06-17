package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"gopkg.in/yaml.v3"
)

type concreteType interface {
	raml.Unmarshaler
	raml.Marshaler
	raml.DataType

	assign(key, val *yaml.Node) error
	marshal(out raml.AnyMap) error
}
