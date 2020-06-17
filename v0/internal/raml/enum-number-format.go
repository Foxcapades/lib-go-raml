package raml

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

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

var validNumForms = []NumberFormat{
	NumberFormatInt,
	NumberFormatInt8,
	NumberFormatInt16,
	NumberFormatInt32,
	NumberFormatInt64,
	NumberFormatLong,
	NumberFormatFloat,
	NumberFormatDouble,
}

func (d NumberFormat) String() string {
	return string(d)
}

func (d *NumberFormat) UnmarshalYAML(y *yaml.Node) error {
	if err := xyml.RequireString(y); err != nil {
		return err
	}

	f := NumberFormat(y.Value)
	for _, v := range validNumForms {
		if v == f {
			*d = v
		}
	}

	// TODO: validation point
	*d = f
	return nil
}

func (d NumberFormat) MarshalYAML() (interface{}, error) {
	return string(d), nil
}
