package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"reflect"
)

const badIntFormatType = "invalid integer format value. expected string, got %s"

// IntegerFormatSortingHat parses the input into the appropriate implementation
// of the raml.NumberFormat interface.
func IntegerFormatSortingHat(
	val interface{},
	log *logrus.Entry,
) (raml.IntegerFormat, error) {
	log.Trace("IntegerFormatSortingHat")

	if val, ok := val.(string); ok {
		df := IntegerFormat(val)
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

	return nil, xlog.Errorf(log, badIntFormatType, reflect.TypeOf(val))
}
