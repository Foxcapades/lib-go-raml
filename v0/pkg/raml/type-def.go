package raml

import (
	"errors"
	"strings"
)

type TypeDef struct {
	object Type
}

func (t *TypeDef) GetType() string {
	return t.object.GetType()
}

func (t *TypeDef) SetType(k string) {
	t.object.SetType(k)
}

func (t *TypeDef) GetRawObject() interface{} {
	return t.object
}

func (t *TypeDef) ToRAML() (string, error) {
	panic("implement me")
}

type typeContainer struct {
	Type string `yaml:"type"`
}

func (t TypeDef) MarshalYAML() (interface{}, error) {
	return t.object, nil
}

func (t *TypeDef) UnmarshalYAML(fn func(interface{}) error) error {
	raw := "object"

	if err := fn(&raw); err == nil {
		tmp, _ := getType(parseType(raw), func(interface{}) error { return nil })
		t.object = tmp
		return nil
	}

	container := typeContainer{}
	err := fn(&container)
	if err != nil {
		return errors.New("Could not parse type definition as string or object, is your raml syntactically correct?")
	}

	raw = container.Type

	if raw == "" {
		raw = "object"
	}

	if tmp, err := getType(parseType(raw), fn); err != nil {
		return err
	} else {
		t.object = tmp
	}
	return nil
}

func getType(kind string, fn func(interface{}) error) (Type, error) {
	switch kind {
	case TypeObject:
		tmp := NewObject()
		return tmp, fn(tmp)
	case TypeArray:
		tmp := NewArray()
		return tmp, fn(tmp)
	case TypeString:
		tmp := NewString()
		return tmp, fn(tmp)
	case TypeNumber:
		tmp := NewNumber()
		return tmp, fn(tmp)
	case TypeInteger:
		tmp := NewInteger()
		return tmp, fn(tmp)
	case TypeBoolean:
		tmp := NewBoolean()
		return tmp, fn(tmp)
	case TypeDateOnly:
		tmp := NewDateOnly()
		return tmp, fn(tmp)
	case TypeTimeOnly:
		tmp := NewTimeOnly()
		return tmp, fn(tmp)
	case TypeDatetimeOnly:
		tmp := NewDatetimeOnly()
		return tmp, fn(tmp)
	case TypeDatetime:
		tmp := NewDatetime()
		return tmp, fn(tmp)
	case TypeFile:
		tmp := NewFile()
		return tmp, fn(tmp)
	case TypeAny:
		tmp := NewAny()
		return tmp, fn(tmp)
	case TypeNil:
		tmp := NewNil()
		return tmp, fn(tmp)
	case TypeUnion:
		tmp := NewUnion()
		tmp.SetType(kind)
		return tmp, nil
	default:
		tmp := NewCustomType()
		tmp.SetType(kind)
		return tmp, nil
	}
}

func parseType(kind string) string {
	if strings.Index(kind, "|") > 0 {
		return "union"
	} else if strings.Index(kind, "!include") > -1 {
		return "include"
	}
	return kind
}
