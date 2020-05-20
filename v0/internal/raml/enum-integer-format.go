package raml

type IntegerFormat string

const (
	IntegerFormatInt   IntegerFormat = "int"
	IntegerFormatInt8  IntegerFormat = "int8"
	IntegerFormatInt16 IntegerFormat = "int16"
	IntegerFormatInt32 IntegerFormat = "int32"
	IntegerFormatInt64 IntegerFormat = "int64"
	IntegerFormatLong  IntegerFormat = "long"
)

func (d IntegerFormat) String() string {
	return string(d)
}

func (d *IntegerFormat) UnmarshalYAML(fn func(interface{}) error) error {
	return fn(d)
}

func (d IntegerFormat) MarshalYAML() (interface{}, error) {
	return string(d), nil
}

func (d IntegerFormat) render() bool {
	return true
}
