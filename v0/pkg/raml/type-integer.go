package raml

type IntegerFormat string

const (
	IntFormatInt   = "int"
	IntFormatInt8  = "int8"
	IntFormatInt16 = "int16"
	IntFormatInt32 = "int32"
	IntFormatInt64 = "int64"
	IntFormatLong  = "long"

	TypeInteger = "integer"
)

func NewInteger() *Integer {
	return &Integer{Base: Base{Type: TypeInteger}}
}

type Integer struct {
	Base

	Minimum    *int64         `yaml:"minimum,omitempty"`
	Maximum    *int64         `yaml:"maximum,omitempty"`
	Format     *IntegerFormat `yaml:"format,omitempty"`
	MultipleOf *int64         `yaml:"multipleOf,omitempty"`
}

func (i *Integer) GetType() string {
	return TypeInteger
}

func (i *Integer) ToRaml() (string, error) {
	return dataTypeRaml(i)
}

type intAlias Integer

func (i *Integer) MarshalYAML() (interface{}, error) {
	if i.canSimplify() {
		return i.Type, nil
	}
	return intAlias(*i), nil
}

//
//func (o *Integer) UnmarshalYAML(value *yaml.Node) error {
//	if err := o.Base.UnmarshalYAML(value); err != nil {
//		return err
//	}
//
//	for i := 0; i < len(value.Content); i++ {
//		key := value.Content[i].Value
//		i++
//		if err := o.assign(key, value.Content[i]); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (o *Integer) assign(key string, val *yaml.Node) error {
//	switch key {
//	case keyMinimum:
//		tmp, err := strconv.ParseInt(val.Value, 10, 64)
//		if err != nil {
//			return err
//		}
//		o.Minimum = &tmp
//	case keyMaximum:
//		tmp, err := strconv.ParseInt(val.Value, 10, 64)
//		if err != nil {
//			return err
//		}
//		o.Minimum = &tmp
//	case keyFormat:
//		tmp := IntegerFormat(val.Value)
//		o.Format = &tmp
//	case keyMultipleOf:
//		tmp, err := strconv.ParseInt(val.Value, 10, 64)
//		if err != nil {
//			return err
//		}
//		o.MultipleOf = &tmp
//	}
//	return nil
//}

func (i *Integer) canSimplify() bool {
	if !i.Base.canSimplify() {
		return false
	}
	if i.Maximum != nil {
		return false
	}
	if i.Minimum != nil {
		return false
	}
	if i.Format != nil {
		return false
	}
	if i.MultipleOf != nil {
		return false
	}

	return true
}
