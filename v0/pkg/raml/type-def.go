package raml

import (
	"fmt"
	"strings"
)

type TypeDef struct {
	object RamlType
}

type typeContainer struct {
	Type string `yaml:"type"`
}

func (t TypeDef) MarshalYAML() (interface{}, error) {
	return t.object, nil
}

func (t *TypeDef) UnmarshalYAML(fn func(interface{}) error) error {
	kind := "object"
	raw := "object"
	simple := true

	if err := fn(&raw); err != nil {
		simple = false
		tmp := typeContainer{}
		err2 := fn(&tmp)
		if err2 != nil {
			return fmt.Errorf("Could not parse type definition as string or object, is your raml syntactically correct?:\n  %s\n  %s", err, err2)
		}
		raw = tmp.Type
	} else {
		t.object = NewCustomType()
		return fn(t.object)
	}
	if raw == "" {
		raw = "object"
	}

	if simple {
		kind = parseType(raw)
	}

	var tmp RamlType
	switch kind {
	case TypeObject:
		tmp = NewObject()
	case TypeArray:
		tmp = NewArray()
	case TypeString:
		tmp = NewString()
	case TypeNumber:
		tmp = NewNumber()
	case TypeInteger:
		tmp = NewInteger()
	case TypeBoolean:
		tmp = NewBoolean()
	case TypeDateOnly:
		tmp = NewDateOnly()
	case TypeTimeOnly:
		tmp = NewTimeOnly()
	case TypeDatetimeOnly:
		tmp = NewDatetimeOnly()
	case TypeDatetime:
		tmp = NewDatetime()
	case TypeFile:
		tmp = NewFile()
	case TypeAny:
		tmp = NewAny()
	case TypeNil:
		tmp = NewNil()
	case TypeUnion:
		tmp = NewUnion()
	default:
		tmp = NewCustomType()
	}
	t.object = tmp
	return fn(tmp)
}

//func getType(val *yaml.Node) (kind, raw string) {
//	kind = TypeObject
//	raw = TypeObject
//	for i := 0; i < len(val.Content); i++ {
//		if val.Content[i].Value == keyType {
//			raw = val.Content[i+1].Value
//			kind = parseType(raw)
//			break
//		}
//	}
//	return
//}

func parseType(kind string) string {
	if strings.Index(kind, "|") > 0 {
		return "union"
	} else if strings.Index(kind, "!include") > -1 {
		return "include"
	}
	return kind
}
