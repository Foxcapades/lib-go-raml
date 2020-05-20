package raml

import (
	"fmt"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// ExampleSortingHat returns a blank instance of the correct `raml.Example`
// sub-type implementation for the given datatype kind.
func ExampleSortingHat(kind rmeta.DataTypeKind) (raml.Example, error) {
	logrus.Trace("internal.ExampleSortingHat")

	switch kind {
	case rmeta.TypeString:
		return NewStringExample(), nil
	case rmeta.TypeObject:
		return NewObjectExample(), nil
	case rmeta.TypeArray:
		return NewArrayExample(), nil
	case rmeta.TypeNumber:
		return NewNumberExample(), nil
	case rmeta.TypeInteger:
		return NewIntegerExample(), nil
	case rmeta.TypeDateOnly:
		return NewDateOnlyExample(), nil
	case rmeta.TypeDatetimeOnly:
		return NewDatetimeOnlyExample(), nil
	case rmeta.TypeFile:
		return NewFileExample(), nil
	case rmeta.TypeBool:
		return NewBoolExample(), nil
	case rmeta.TypeCustom:
		return NewCustomExample(), nil
	case rmeta.TypeDatetime:
		return NewDatetimeExample(), nil
	case rmeta.TypeUnion:
		return NewUnionExample(), nil
	}

	return nil, fmt.Errorf("no example implementation available for type %s", kind)
}
