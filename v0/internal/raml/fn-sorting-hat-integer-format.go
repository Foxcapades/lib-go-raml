package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// IntegerFormatSortingHat parses the input into the appropriate implementation
// of the raml.NumberFormat interface.
func IntegerFormatSortingHat(val *yaml.Node) (raml.IntegerFormat, error) {
	logrus.Trace("IntegerFormatSortingHat")

	if err := xyml.RequireString(val); err != nil {
		return nil, err
	}

	df := IntegerFormat(val.Value)

	switch df {
	case IntegerFormatInt:
		return &df, nil
	case IntegerFormatInt8:
		return &df, nil
	case IntegerFormatInt16:
		return &df, nil
	case IntegerFormatInt32:
		return &df, nil
	case IntegerFormatInt64:
		return &df, nil
	case IntegerFormatLong:
		return &df, nil
	}

	// TODO: validation point, if validation is enabled, this should be an error
	return &df, nil
}
