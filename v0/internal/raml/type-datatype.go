package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/cast"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

func NewDataType(
	kind rmeta.DataTypeKind,
	log *logrus.Entry,
	self concreteType,
) *DataType {
	return &DataType{
		self:     self,
		log:      log,
		schema:   string(kind),
		kind:     kind,
		hasExtra: makeExtra(log),
	}
}

type DataType struct {
	hasExtra

	log  *logrus.Entry
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

func (d *DataType) Kind() rmeta.DataTypeKind {
	return d.kind
}

func (d *DataType) ToRAML() (string, error) {
	panic("implement me")
}

func (d DataType) MarshalYAML() (interface{}, error) {
	out := NewAnyMap(d.log)
	if short, err := d.MarshalRAML(out); err != nil {
		return nil, err
	} else if short {
		schema := out.Get(rmeta.KeyType).Get()
		d.log.Debug("Printing RAML type short form ", schema)
		return schema, nil
	}
	d.log.Debug("Printing RAML type long form")
	return out, nil
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

func (d *DataType) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
	if str, ok := val.(string); ok {
		d.schema = str
		return nil
	}

	if _, ok := val.([]interface{}); ok {
		return xlog.Error(log, "multi-type declarations are not currently supported")
	}

	if slice, err := assign.AsMapSlice(val); err == nil {
		for i := range slice {
			l2 := xlog.AddPath(log, slice[i].Key)

			if err := d.self.assign(slice[i].Key, slice[i].Value, l2); err != nil {
				return xlog.Error(l2, err)
			}
		}
		return nil
	}

	return xlog.Error(log, "type definitions must be an array, a string, or a map")
}

func (d *DataType) marshal(out raml.AnyMap) error {
	d.log.Trace("internal.DataType.marshal")

	out.Put(rmeta.KeyType, d.schema)
	d.hasExtra.out(out)
	return nil
}

func (d *DataType) assign(key, val interface{}, log *logrus.Entry) error {
	switch key {
	case rmeta.KeyType, rmeta.KeySchema:
		if str, err := cast.AsString(val); err != nil {
			return err
		} else {
			d.schema = str
		}
		return nil
	}

	d.hasExtra.in(key, val)
	return nil
}

func (d *DataType) render() bool {
	return true
}
