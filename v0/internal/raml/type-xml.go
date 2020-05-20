package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewXml() *Xml {
	return &Xml{
		extra: make(map[interface{}]interface{}),
	}
}

type Xml struct {
	isAttr *bool
	isWrap *bool
	name   *string
	ns     *string
	pref   *string
	extra  map[interface{}]interface{}
}

func (x *Xml) IsAttribute() option.Bool {
	return option.NewMaybeBool(x.isAttr)
}

func (x *Xml) SetIsAttribute(b bool) raml.Xml {
	x.isAttr = &b
	return x
}

func (x *Xml) UnsetIsAttribute() raml.Xml {
	x.isAttr = nil
	return x
}

func (x *Xml) IsWrapped() option.Bool {
	return option.NewMaybeBool(x.isWrap)
}

func (x *Xml) SetIsWrapped(b bool) raml.Xml {
	x.isWrap = &b
	return x
}

func (x *Xml) UnsetIsWrapped() raml.Xml {
	x.isWrap = nil
	return x
}

func (x *Xml) Name() option.String {
	return option.NewMaybeString(x.name)
}

func (x *Xml) SetName(s string) raml.Xml {
	x.name = &s
	return x
}

func (x *Xml) UnsetName() raml.Xml {
	x.name = nil
	return x
}

func (x *Xml) Namespace() option.String {
	return option.NewMaybeString(x.ns)
}

func (x *Xml) SetNamespace(s string) raml.Xml {
	x.ns = &s
	return x
}

func (x *Xml) UnsetNamespace() raml.Xml {
	x.ns = nil
	return x
}

func (x *Xml) Prefix() option.String {
	return option.NewMaybeString(x.pref)
}

func (x *Xml) SetPrefix(s string) raml.Xml {
	x.pref = &s
	return x
}

func (x *Xml) UnsetPrefix() raml.Xml {
	x.pref = nil
	return x
}

func (x *Xml) UnmarshalRAML(v *yaml.Node) error {
	logrus.Trace("internal.Xml.UnmarshalRaml")
	return xyml.ForEachMap(v, x.assign)
}

func (x Xml) MarshalRAML(out raml.AnyMap) (bool, error) {
	out.PutNonNil(rmeta.KeyAttribute, x.isAttr).
		PutNonNil(rmeta.KeyWrapped, x.isWrap).
		PutNonNil(rmeta.KeyName, x.name).
		PutNonNil(rmeta.KeyNamespace, x.ns).
		PutNonNil(rmeta.KeyPrefix, x.pref)
	return false, nil
}

func (x *Xml) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyAttribute:
		return assign.AsBoolPtr(val, &x.isAttr)
	case rmeta.KeyWrapped:
		return assign.AsBoolPtr(val, &x.isWrap)
	case rmeta.KeyName:
		return assign.AsStringPtr(val, &x.name)
	case rmeta.KeyNamespace:
		return assign.AsStringPtr(val, &x.ns)
	case rmeta.KeyPrefix:
		return assign.AsStringPtr(val, &x.pref)
	default:
		if k, err := xyml.CastYmlTypeToScalar(key); err != nil {
			return err
		} else {
			x.extra[k] = val
		}
	}

	return nil
}
