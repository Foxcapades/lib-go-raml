package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// ExampleSortingHat returns a blank instance of the correct `raml.Example`
// sub-type implementation for the given datatype kind.
func ExampleSortingHat(
	kind rmeta.DataTypeKind,
	log *logrus.Entry,
) (raml.Example, error) {
	log.Trace("internal.ExampleSortingHat")

	switch kind {
	case rmeta.TypeString:
		return NewStringExample(log), nil
	case rmeta.TypeObject:
		return NewObjectExample(log), nil
	case rmeta.TypeArray:
		return NewArrayExample(log), nil
	case rmeta.TypeNumber:
		return NewNumberExample(log), nil
	case rmeta.TypeInteger:
		return NewIntegerExample(log), nil
	case rmeta.TypeDateOnly:
		return NewDateOnlyExample(log), nil
	case rmeta.TypeDatetimeOnly:
		return NewDatetimeOnlyExample(log), nil
	case rmeta.TypeFile:
		return NewFileExample(log), nil
	case rmeta.TypeBool:
		return NewBoolExample(log), nil
	case rmeta.TypeCustom:
		return NewCustomExample(log), nil
	case rmeta.TypeDatetime:
		return NewDatetimeExample(log), nil
	case rmeta.TypeUnion:
		return NewUnionExample(log), nil
	}

	return nil, xlog.Errorf(log, "No example implementation available for type %s", kind)
}
