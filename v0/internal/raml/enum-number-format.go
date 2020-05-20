package raml

type NumberFormat string

const (
	NumberFormatInt    NumberFormat = "int"
	NumberFormatInt8   NumberFormat = "int8"
	NumberFormatInt16  NumberFormat = "int16"
	NumberFormatInt32  NumberFormat = "int32"
	NumberFormatInt64  NumberFormat = "int64"
	NumberFormatLong   NumberFormat = "long"
	NumberFormatFloat  NumberFormat = "float"
	NumberFormatDouble NumberFormat = "double"
)

func (d NumberFormat) String() string {
	return string(d)
}

func (d *NumberFormat) UnmarshalYAML(fn func(interface{}) error) error {
	return fn(d)
}

func (d NumberFormat) MarshalYAML() (interface{}, error) {
	return string(d), nil
}

func (d NumberFormat) render() bool {
	return true
}
