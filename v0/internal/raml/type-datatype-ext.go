package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func NewExtendedDataType(kind rmeta.DataTypeKind, self concreteType) *ExtendedDataType {
	return &ExtendedDataType{
		DataType:       NewDataType(kind, self),
		hasAnnotations: makeAnnotations(),
		facets:         raml.NewFacetMap(0),
		required:       true,
	}
}

type ExtendedDataType struct {
	hasAnnotations
	*DataType

	displayName *string
	description *string
	facets      raml.FacetMap
	xml         raml.XML
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

func (e *ExtendedDataType) XML() raml.XML {
	return e.xml
}

func (e *ExtendedDataType) Required() bool {
	return e.required
}

func (e *ExtendedDataType) marshal(out raml.AnyMap) error {
	out.PutIfNotNil(rmeta.KeyDisplayName, e.displayName).
		PutIfNotNil(rmeta.KeyDescription, e.description)

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

	out.PutIfNotNil(rmeta.KeyXML, e.xml)

	return nil
}

func (e *ExtendedDataType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyDisplayName:
		return assign.AsStringPtr(val, &e.displayName)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(val, &e.description)
	case rmeta.KeyFacets:
		return UnmarshalFacetMapRAML(e.facets, val)
	case rmeta.KeyXML:
		xml := NewXML()
		if err := xml.UnmarshalRAML(val); err != nil {
			return err
		}
		e.xml = xml
		return nil
	}

	if xyml.IsString(key) {
		if used, err := e.hasAnnotations.in(key.Value, val); err != nil {
			return err
		} else if used {
			return nil
		}
	}

	return e.DataType.assign(key, val)
}
