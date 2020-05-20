package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

func NewExtendedDataType(
	kind rmeta.DataTypeKind,
	log *logrus.Entry,
	self concreteType,
) *ExtendedDataType {
	return &ExtendedDataType{
		DataType:       NewDataType(kind, log, self),
		hasAnnotations: makeAnnotations(log),
		facets:         NewFacetMap(log),
		required:       true,
	}
}

type ExtendedDataType struct {
	hasAnnotations
	*DataType

	displayName *string
	description *string
	facets      raml.FacetMap
	xml         raml.Xml
	required    bool
}

func (e *ExtendedDataType) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *ExtendedDataType) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *ExtendedDataType) FacetDefinitions() raml.FacetMap {
	return e.facets
}

func (e *ExtendedDataType) Xml() raml.Xml {
	return e.xml
}

func (e *ExtendedDataType) Required() bool {
	return e.required
}

func (e *ExtendedDataType) marshal(out raml.AnyMap) error {
	out.PutNonNil(rmeta.KeyDisplayName, e.displayName).
		PutNonNil(rmeta.KeyDescription, e.description)

	if err := e.DataType.marshal(out); err != nil {
		return err
	}

	if !e.required {
		out.Put(rmeta.KeyRequired, e.required)
	}

	e.hasAnnotations.out(out)

	if e.facets.Len() > 0 {
		out.Put(rmeta.KeyFacets, e.facets)
	}

	out.PutNonNil(rmeta.KeyXml, e.xml)

	return nil
}

func (e *ExtendedDataType) assign(key, val interface{}, log *logrus.Entry) error {
	switch key {
	case rmeta.KeyDisplayName:
		return xlog.OptError(log, assign.AsStringPtr(val, &e.displayName, log))
	case rmeta.KeyDescription:
		return xlog.OptError(log, assign.AsStringPtr(val, &e.description, log))
	case rmeta.KeyFacets:
		return xlog.OptError(log, e.facets.UnmarshalRAML(val, log))
	case rmeta.KeyXml:
		xml := NewXml(log)
		if err := xml.UnmarshalRAML(val, log); err != nil {
			return xlog.Error(log, err)
		}
		e.xml = xml
		return nil
	}

	if str, ok := key.(string); ok {
		if used, err := e.hasAnnotations.in(str, val); err != nil {
			return err
		} else if used {
			return nil
		}
	}

	return e.DataType.assign(key, val, log)
}

func (e *ExtendedDataType) render() bool {
	return true
}
