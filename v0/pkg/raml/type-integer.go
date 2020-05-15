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
	Base `yaml:",inline"`

	Minimum    *int64         `yaml:"minimum,omitempty"`
	Maximum    *int64         `yaml:"maximum,omitempty"`
	Format     *IntegerFormat `yaml:"format,omitempty"`
	MultipleOf *int64         `yaml:"multipleOf,omitempty"`
}

func (i *Integer) ToRAML() (string, error) {
	return dataTypeRaml(i)
}

type intAlias Integer

func (i *Integer) MarshalYAML() (interface{}, error) {
	if i.canSimplify() {
		return i.Type, nil
	}
	return intAlias(*i), nil
}

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
