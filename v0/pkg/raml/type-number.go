package raml

const (
	keyMinimum    = "minimum"
	keyMaximum    = "maximum"
	keyFormat     = "format"
	keyMultipleOf = "multipleOf"
)

type NumberFormat string

const (
	NumFormatInt    = "int"
	NumFormatInt8   = "int8"
	NumFormatInt16  = "int16"
	NumFormatInt32  = "int32"
	NumFormatInt64  = "int64"
	NumFormatLong   = "long"
	NumFormatFloat  = "float"
	NumFormatDouble = "double"

	TypeNumber = "number"
)

func NewNumber() *Number {
	return &Number{Base: Base{Type: TypeNumber}}
}

type Number struct {
	Base

	Minimum    *float64      `yaml:"minimum,omitempty"`
	Maximum    *float64      `yaml:"maximum,omitempty"`
	Format     *NumberFormat `yaml:"format,omitempty"`
	MultipleOf *float64      `yaml:"multipleOf,omitempty"`
}

func (n *Number) ToRaml() (string, error) {
	return dataTypeRaml(n)
}

type numAlias Number
func (n *Number) MarshalYAML() (interface{}, error) {
	if n.canSimplify() {
		return n.Type, nil
	}
	return numAlias(*n), nil
}

//func (o *Number) UnmarshalYAML(value *yaml.Node) error {
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
//func (o *Number) assign(key string, val *yaml.Node) error {
//	switch key {
//	case keyMinimum:
//		tmp, err := strconv.ParseFloat(val.Value, 64)
//		if err != nil {
//			return err
//		}
//		o.Minimum = &tmp
//	case keyMaximum:
//		tmp, err := strconv.ParseFloat(val.Value, 64)
//		if err != nil {
//			return err
//		}
//		o.Minimum = &tmp
//	case keyFormat:
//		tmp := NumberFormat(val.Value)
//		o.Format = &tmp
//	case keyMultipleOf:
//		tmp, err := strconv.ParseFloat(val.Value, 64)
//		if err != nil {
//			return err
//		}
//		o.MultipleOf = &tmp
//	}
//	return nil
//}

func (n *Number) canSimplify() bool {
	if !n.Base.canSimplify() {
		return false
	}
	if n.Maximum != nil {
		return false
	}
	if n.Minimum != nil {
		return false
	}
	if n.Format != nil {
		return false
	}
	if n.MultipleOf != nil {
		return false
	}

	return true
}
