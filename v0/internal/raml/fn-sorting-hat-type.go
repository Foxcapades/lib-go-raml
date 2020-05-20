package raml

import (
	"fmt"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"strings"
)

const (
	multiTypeNotAllowed = "multi-type definitions are not currently supported " +
		"%d:%d"
	fullTypeKeyBadValue = "the type key in an expanded definition must be an " +
		"array or a string. got %s at %d:%d"
	badTypeDefType = "a type definition should be empty, a string, an array, or" +
		" a map.  instead got %s at %d:%d"
)

func TypeSortingHat(val *yaml.Node) (out raml.DataType, err error) {
	logrus.Trace("internal.TypeSortingHat")

	if xyml.IsString(val) {
		return typeToKind(val.Value), nil
	}

	if xyml.IsList(val) {
		return nil, fmt.Errorf(multiTypeNotAllowed, val.Line, val.Column)
	}

	if xyml.IsMap(val) {
		if kind, err := siftType(val); err != nil {
			return nil, err
		} else {
			if err = kind.UnmarshalRAML(val); err != nil {
				return nil, err
			}
			return kind, nil
		}
	}

	if val.Tag == "!include" {
		out := NewIncludeType()
		out.schema = val.Value
		return out, nil
	}

	if util.IsNil(val) {
		return NewStringType(), nil
	}

	return nil, fmt.Errorf(badTypeDefType, val.Tag, val.Line, val.Column)
}

func typeToKind(val string) concreteType {
	logrus.Trace("internal.typeToKind")

	tmp := rmeta.DataTypeKind(val)
	switch tmp {
	case rmeta.TypeAny:
		logrus.Debug("you go to house ", rmeta.TypeAny)
		return NewAnyType()
	case rmeta.TypeArray:
		logrus.Debug("you go to house ", rmeta.TypeArray)
		return NewArrayType()
	case rmeta.TypeBool:
		logrus.Debug("you go to house ", rmeta.TypeBool)
		return NewBoolType()
	case rmeta.TypeDateOnly:
		logrus.Debug("you go to house ", rmeta.TypeDateOnly)
		return NewDateOnlyType()
	case rmeta.TypeDatetime:
		logrus.Debug("you go to house ", rmeta.TypeDatetime)
		return NewDatetimeType()
	case rmeta.TypeDatetimeOnly:
		logrus.Debug("you go to house ", rmeta.TypeDatetimeOnly)
		return NewDatetimeOnlyType()
	case rmeta.TypeFile:
		logrus.Debug("you go to house ", rmeta.TypeFile)
		return NewFileType()
	case rmeta.TypeInteger:
		logrus.Debug("you go to house ", rmeta.TypeInteger)
		return NewIntegerType()
	case rmeta.TypeNil:
		logrus.Debug("you go to house ", rmeta.TypeNil)
		return NewNilType()
	case rmeta.TypeNumber:
		logrus.Debug("you go to house ", rmeta.TypeNumber)
		return NewNumberType()
	case rmeta.TypeObject:
		logrus.Debug("you go to house ", rmeta.TypeObject)
		return NewObjectType()
	case rmeta.TypeString:
		logrus.Debug("you go to house ", rmeta.TypeString)
		return NewStringType()
	case rmeta.TypeTimeOnly:
		logrus.Debug("you go to house ", rmeta.TypeTimeOnly)
		return NewTimeOnlyType()
	}

	if strings.HasPrefix(val, "!include ") {
		logrus.Debug("you go to house ", rmeta.TypeInclude)
		out := NewIncludeType()
		out.schema = val
		return out
	}

	if -1 < strings.IndexByte(val, '|') {
		logrus.Debug("you go to house ", rmeta.TypeUnion)
		tmp := NewUnionType()
		tmp.schema = val
		return tmp
	}

	logrus.Debug("you go to house ", rmeta.TypeCustom)
	out := NewCustomType()
	out.schema = val
	return out
}

func siftType(val *yaml.Node) (concreteType, error) {
	logrus.Trace("internal.siftType")

	schema := -1
	props := false
	items := false

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i].Value

		if key == rmeta.KeyType || key == rmeta.KeySchema {
			schema = i + 1
			break
		}

		if key == rmeta.KeyProperties {
			props = true
		}

		if key == rmeta.KeyItems {
			items = true
		}
	}

	if schema > -1 {
		tmp := val.Content[schema]
		if tmp.Tag == xyml.String {
			return typeToKind(tmp.Value), nil
		} else if tmp.Kind == yaml.SequenceNode {
			return nil, fmt.Errorf(multiTypeNotAllowed, tmp.Line, tmp.Column)
		} else {
			return nil, fmt.Errorf(fullTypeKeyBadValue, tmp.Tag, tmp.Line, tmp.Column)
		}
	}

	if props {
		return NewObjectType(), nil
	}

	if items {
		return NewArrayType(), nil
	}

	return NewStringType(), nil
}
