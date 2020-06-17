package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// NumberFormatSortingHat parses the input into the appropriate implementation
// of the raml.NumberFormat interface.
func NumberFormatSortingHat(val *yaml.Node) (raml.NumberFormat, error) {
	logrus.Trace("IntegerFormatSortingHat")

	if err := xyml.RequireString(val); err != nil {
		return nil, err
	}

	df := NumberFormat(val.Value)

	switch df {
	case NumberFormatInt:
		return &df, nil
	case NumberFormatInt8:
		return &df, nil
	case NumberFormatInt16:
		return &df, nil
	case NumberFormatInt32:
		return &df, nil
	case NumberFormatInt64:
		return &df, nil
	case NumberFormatLong:
		return &df, nil
	case NumberFormatFloat:
		return &df, nil
	case NumberFormatDouble:
		return &df, nil
	}

	// TODO: validation point, if validation is enabled, this should be an error
	return &df, nil
}
