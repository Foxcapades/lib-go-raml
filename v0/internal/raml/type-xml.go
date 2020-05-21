package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

func NewXML() *XML {
	return &XML{
		extra: make(map[interface{}]interface{}),
	}
}

type XML struct {
	isAttr *bool
	isWrap *bool
	name   *string
	ns     *string
	pref   *string
	extra  map[interface{}]interface{}
}

func (x *XML) IsAttribute() option.Bool {
	return option.NewMaybeBool(x.isAttr)
}

func (x *XML) SetIsAttribute(b bool) raml.XML {
	x.isAttr = &b
	return x
}

func (x *XML) UnsetIsAttribute() raml.XML {
	x.isAttr = nil
	return x
}

func (x *XML) IsWrapped() option.Bool {
	return option.NewMaybeBool(x.isWrap)
}

func (x *XML) SetIsWrapped(b bool) raml.XML {
	x.isWrap = &b
	return x
}

func (x *XML) UnsetIsWrapped() raml.XML {
	x.isWrap = nil
	return x
}

func (x *XML) Name() option.String {
	return option.NewMaybeString(x.name)
}

func (x *XML) SetName(s string) raml.XML {
	x.name = &s
	return x
}

func (x *XML) UnsetName() raml.XML {
	x.name = nil
	return x
}

func (x *XML) Namespace() option.String {
	return option.NewMaybeString(x.ns)
}

func (x *XML) SetNamespace(s string) raml.XML {
	x.ns = &s
	return x
}

func (x *XML) UnsetNamespace() raml.XML {
	x.ns = nil
	return x
}

func (x *XML) Prefix() option.String {
	return option.NewMaybeString(x.pref)
}

func (x *XML) SetPrefix(s string) raml.XML {
	x.pref = &s
	return x
}

func (x *XML) UnsetPrefix() raml.XML {
	x.pref = nil
	return x
}

func (x *XML) UnmarshalRAML(v *yaml.Node) error {
	return xyml.ForEachMap(v, x.assign)
}

func (x XML) MarshalRAML(out raml.AnyMap) (bool, error) {
	out.PutNonNil(rmeta.KeyAttribute, x.isAttr).
		PutNonNil(rmeta.KeyWrapped, x.isWrap).
		PutNonNil(rmeta.KeyName, x.name).
		PutNonNil(rmeta.KeyNamespace, x.ns).
		PutNonNil(rmeta.KeyPrefix, x.pref)

	return false, nil
}

func (x *XML) assign(key, val *yaml.Node) error {
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
