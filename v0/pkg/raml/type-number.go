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
	Base `yaml:",inline"`

	Minimum    *float64      `yaml:"minimum,omitempty"`
	Maximum    *float64      `yaml:"maximum,omitempty"`
	Format     *NumberFormat `yaml:"format,omitempty"`
	MultipleOf *float64      `yaml:"multipleOf,omitempty"`
}

func (n *Number) ToRAML() (string, error) {
	return dataTypeRaml(n)
}

type numAlias Number

func (n Number) MarshalYAML() (interface{}, error) {
	if n.canSimplify() {
		return n.Type, nil
	}
	return numAlias(n), nil
}

func (n *Number) canSimplify() bool {
	return n.Base.canSimplify() &&
		n.Maximum == nil &&
		n.Minimum == nil &&
		n.Format == nil &&
		n.MultipleOf == nil
}
