package raml

import (
	"errors"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewDataType(kind rmeta.DataTypeKind, self concreteType) *DataType {
	return &DataType{
		self:     self,
		schema:   string(kind),
		kind:     kind,
		hasExtra: makeExtra(),
	}
}

type DataType struct {
	hasExtra

	self concreteType

	schema string
	kind   rmeta.DataTypeKind
}

func (d *DataType) Schema() string {
	return d.schema
}

func (d *DataType) Type() string {
	return d.schema
}

func (d *DataType) OverrideType(t string) {
	d.schema = t
}

func (d *DataType) Kind() rmeta.DataTypeKind {
	return d.kind
}

func (d *DataType) ToRAML() (string, error) {
	panic("implement me")
}

func (d DataType) ToYAML() (*yaml.Node, error) {
	out := raml.NewAnyMap(0).SerializeOrdered(false)
	if short, err := d.MarshalRAML(out); err != nil {
		return nil, err
	} else if short {
		schema, _ := out.Get(rmeta.KeyType)
		logrus.Debug("Printing RAML type short form ", schema)
		return xyml.ToYamlNode(schema)
	}
	logrus.Debug("Printing RAML type long form")
	return out.ToYAML()
}

func (d DataType) MarshalYAML() (interface{}, error) {
	return d.ToYAML()
}

func (d *DataType) MarshalRAML(out raml.AnyMap) (bool, error) {
	if err := d.self.marshal(out); err != nil {
		return false, err
	}
	if out.Len() == 1 && out.Has(rmeta.KeyType) {
		return true, nil
	}
	return false, nil
}

func (d *DataType) UnmarshalRAML(val *yaml.Node) error {
	if xyml.IsString(val) {
		d.schema = val.Value
		return nil
	}

	if xyml.IsSequence(val) {
		return errors.New("multi-type declarations are not currently supported")
	}

	if xyml.IsMap(val) {
		return xyml.MapForEach(val, d.self.assign)
	}

	return errors.New("type definitions must be an array, a string, or a map")
}

func (d *DataType) marshal(out raml.AnyMap) error {
	out.Put(rmeta.KeyType, d.schema)
	d.hasExtra.out(out)
	return nil
}

func (d *DataType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyType, rmeta.KeySchema:
		if err := xyml.RequireString(val); err != nil {
			return err
		} else {
			d.schema = val.Value
		}
		return nil
	}

	d.hasExtra.in(key, val)
	return nil
}
