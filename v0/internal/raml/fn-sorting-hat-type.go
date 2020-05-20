package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"reflect"
	"strings"
)

const (
	multiTypeNotAllowed = "multi-type definitions are not currently supported"
	fullTypeKeyBadValue = "the type key in an expanded definition must be an " +
		"array or a string. got %s"
	badTypeDefType = "a type definition should be empty, a string, an array, or" +
		" a map.  instead got %s"
)

func TypeSortingHat(val interface{}, log *logrus.Entry) (out raml.DataType, err error) {
	log.Trace("internal.TypeSortingHat")

	if str, ok := val.(string); ok {
		return typeToKind(str, log), nil
	}

	if _, ok := val.([]interface{}); ok {
		return nil, xlog.Error(log, multiTypeNotAllowed)
	}

	if tmp, ok := val.(yaml.MapSlice); ok {
		if kind, err := siftType(tmp, log); err != nil {
			return nil, err
		} else {
			if err = kind.UnmarshalRAML(val, log); err != nil {
				return nil, err
			}
			return kind, nil
		}
	}

	if util.IsNil(val) {
		return NewStringType(log), nil
	}

	return nil, xlog.Errorf(log, badTypeDefType, reflect.TypeOf(val))
}

func typeToKind(val string, log *logrus.Entry) concreteType {
	log.Trace("internal.typeToKind")

	tmp := rmeta.DataTypeKind(val)
	switch tmp {
	case rmeta.TypeAny:
		return NewAnyType(log)
	case rmeta.TypeArray:
		return NewArrayType(log)
	case rmeta.TypeBool:
		return NewBoolType(log)
	case rmeta.TypeDateOnly:
		return NewDateOnlyType(log)
	case rmeta.TypeDatetime:
		return NewDatetimeType(log)
	case rmeta.TypeDatetimeOnly:
		return NewDatetimeOnlyType(log)
	case rmeta.TypeFile:
		return NewFileType(log)
	case rmeta.TypeInteger:
		return NewIntegerType(log)
	case rmeta.TypeNil:
		return NewNilType(log)
	case rmeta.TypeNumber:
		return NewNumberType(log)
	case rmeta.TypeObject:
		return NewObjectType(log)
	case rmeta.TypeString:
		return NewStringType(log)
	case rmeta.TypeTimeOnly:
		return NewTimeOnlyType(log)
	}

	trm := strings.TrimSpace(val)
	if strings.HasPrefix(trm, "!include ") {
		out := NewIncludeType(log)
		out.schema = val
		return out
	}

	if -1 < strings.IndexByte(trm, '|') {
		tmp := NewUnionType(log)
		tmp.schema = val
		return tmp
	}

	out := NewCustomType(log)
	out.schema = val
	return out
}

func siftType(val yaml.MapSlice, log *logrus.Entry) (concreteType, error) {
	log.Trace("internal.siftType")

	schema := -1
	props := false
	items := false

	for i := range val {
		if val[i].Key == rmeta.KeyType || val[i].Key == rmeta.KeySchema {
			schema = i
			break
		}
		if val[i].Key == rmeta.KeyProperties {
			props = true
		}
		if val[i].Key == rmeta.KeyItems {
			items = true
		}
	}

	if schema > -1 {
		if str, ok := val[schema].Value.(string); ok {
			return typeToKind(str, log), nil
		} else if _, ok := val[schema].Value.([]interface{}); ok {
			return nil, xlog.Error(log, multiTypeNotAllowed)
		} else {
			return nil, xlog.Errorf(log, fullTypeKeyBadValue, reflect.TypeOf(val[schema].Value))
		}
	}

	if props {
		return NewObjectType(log), nil
	}

	if items {
		return NewArrayType(log), nil
	}

	return NewStringType(log), nil
}
