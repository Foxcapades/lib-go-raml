package raml

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

type IntegerFormat string

const (
	IntegerFormatInt   IntegerFormat = "int"
	IntegerFormatInt8  IntegerFormat = "int8"
	IntegerFormatInt16 IntegerFormat = "int16"
	IntegerFormatInt32 IntegerFormat = "int32"
	IntegerFormatInt64 IntegerFormat = "int64"
	IntegerFormatLong  IntegerFormat = "long"
)

var validIntFormats = []IntegerFormat{
	IntegerFormatInt,
	IntegerFormatInt8,
	IntegerFormatInt16,
	IntegerFormatInt32,
	IntegerFormatInt64,
	IntegerFormatLong,
}

func (d IntegerFormat) String() string {
	return string(d)
}

func (d *IntegerFormat) UnmarshalYAML(y *yaml.Node) error {
	if err := xyml.RequireString(y); err != nil {
		return err
	}

	conv := IntegerFormat(y.Value)
	for _, v := range validIntFormats {
		if v == conv {
			*d = v
			return nil
		}
	}

	// TODO: apply validation here
	*d = conv
	return nil
}

func (d IntegerFormat) MarshalYAML() (interface{}, error) {
	return d.ToYAML()
}

func (d IntegerFormat) ToYAML() (*yaml.Node, error) {
	return xyml.NewStringNode(string(d)), nil
}
