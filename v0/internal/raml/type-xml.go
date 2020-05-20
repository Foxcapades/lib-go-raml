package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

func NewXml(log *logrus.Entry) *Xml {
	return &Xml{
		log:   log.WithField(xlog.KeyType, "internal.Xml"),
		extra: make(map[interface{}]interface{}),
	}
}

type Xml struct {
	log *logrus.Entry

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

func (x *Xml) UnmarshalRAML(value interface{}, log *logrus.Entry) error {
	x.log.Trace("internal.Xml.UnmarshalRaml")
	ms, err := assign.AsMapSlice(value)

	if err != nil {
		return xlog.Error(x.log, err)
	}

	for i := range ms {
		l2 := xlog.AddPath(x.log, ms[i].Key)

		if err = x.assign(ms[i].Key, ms[i].Value); err != nil {
			return xlog.Error(l2, err)
		}
	}

	return nil
}

func (x Xml) MarshalRAML(out raml.AnyMap) (bool, error) {
	out.PutNonNil(rmeta.KeyAttribute, x.isAttr).
		PutNonNil(rmeta.KeyWrapped, x.isWrap).
		PutNonNil(rmeta.KeyName, x.name).
		PutNonNil(rmeta.KeyNamespace, x.ns).
		PutNonNil(rmeta.KeyPrefix, x.pref)
	return false, nil
}

func (x *Xml) assign(key, val interface{}) error {
	str, ok := key.(string)
	if !ok {
		x.extra[key] = val
		return nil
	}
	var err error

	switch str {
	case rmeta.KeyAttribute:
		err = assign.AsBoolPtr(val, &x.isAttr, x.log)
	case rmeta.KeyWrapped:
		err = assign.AsBoolPtr(val, &x.isWrap, x.log)
	case rmeta.KeyName:
		err = assign.AsStringPtr(val, &x.name, x.log)
	case rmeta.KeyNamespace:
		err = assign.AsStringPtr(val, &x.ns, x.log)
	case rmeta.KeyPrefix:
		err = assign.AsStringPtr(val, &x.pref, x.log)
	default:
		x.extra[str] = val
	}

	return err
}
