package raml

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

type DateFormat string

const (
	DateFormatRfc3339 DateFormat = "rfc3339"
	DateFormatRfc2616 DateFormat = "rfc2616"
)

func (d DateFormat) String() string {
	return string(d)
}

func (d *DateFormat) UnmarshalYAML(y *yaml.Node) error {
	if err := xyml.RequireString(y); err != nil {
		return err
	}

	switch DateFormat(y.Value) {
	case DateFormatRfc3339:
		*d = DateFormatRfc3339
	case DateFormatRfc2616:
		*d = DateFormatRfc2616
	default:
		// TODO: validate here when enabled
		*d = DateFormat(y.Value)
	}

	return nil
}

func (d DateFormat) MarshalYAML() (interface{}, error) {
	return string(d), nil
}
