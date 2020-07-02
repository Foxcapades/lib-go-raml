package raml

import (
	"fmt"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
	"strings"
)

const (
	multiTypeNotAllowed = "multi-type definitions are not currently supported " +
		"%d:%d"
	fullTypeKeyBadValue = "the type key in an expanded definition must be an " +
		"array, a string, or an import. got %s at %d:%d"
	badTypeDefType = "a type definition should be empty, a string, an array, an" +
		"include, or a map.  instead got %s at %d:%d"
)

// TypeSortingHat takes the given YAML node and attempts to parse it into a
// RAML data type object.
//
// If the YAML node cannot be parsed into a RAML element, returns an error.
func TypeSortingHat(val *yaml.Node) (out raml.DataType, err error) {
	if xyml.IsString(val) {
		return typeToKind(val.Value), nil
	}

	if xyml.IsSequence(val) {
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
	tmp := rmeta.DataTypeKind(val)

	switch tmp {
	case rmeta.TypeAny:
		return NewAnyType()
	case rmeta.TypeArray:
		return NewArrayType()
	case rmeta.TypeBool:
		return NewBoolType()
	case rmeta.TypeDateOnly:
		return NewDateOnlyType()
	case rmeta.TypeDatetime:
		return NewDatetimeType()
	case rmeta.TypeDatetimeOnly:
		return NewDatetimeOnlyType()
	case rmeta.TypeFile:
		return NewFileType()
	case rmeta.TypeInteger:
		return NewIntegerType()
	case rmeta.TypeNil:
		return NewNilType()
	case rmeta.TypeNumber:
		return NewNumberType()
	case rmeta.TypeObject:
		return NewObjectType()
	case rmeta.TypeString:
		return NewStringType()
	case rmeta.TypeTimeOnly:
		return NewTimeOnlyType()
	}

	if strings.HasPrefix(val, "!include ") {
		out := NewIncludeType()
		out.schema = val
		return out
	}

	if -1 < strings.IndexByte(val, '|') {
		tmp := NewUnionType()
		tmp.schema = val
		return tmp
	}

	out := NewCustomType()
	out.schema = val
	return out
}

// siftType attempts to parse the given node for either a `type` facet or the
// characteristics of either an object or an array.  If the input cannot be
// matched to one of those, falls back to a string type node as per the RAML
// specification.
func siftType(val *yaml.Node) (concreteType, error) {
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
		if tmp.LongTag() == xyml.TagString {
			return typeToKind(tmp.Value), nil
		} else if tmp.Kind == yaml.SequenceNode {
			return nil, fmt.Errorf(multiTypeNotAllowed, tmp.Line, tmp.Column)
		} else if tmp.Tag == "!include" {
			return NewIncludeType(), nil
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
