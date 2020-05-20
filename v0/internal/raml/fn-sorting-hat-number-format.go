package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"reflect"
)

const badNumFormatType = "invalid integer format value. expected string, got %s"

// NumberFormatSortingHat parses the input into the appropriate implementation
// of the raml.NumberFormat interface.
func NumberFormatSortingHat(
	val interface{},
	log *logrus.Entry,
) (raml.NumberFormat, error) {
	log.Trace("IntegerFormatSortingHat")

	if val, ok := val.(string); ok {
		df := NumberFormat(val)
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

	return nil, xlog.Errorf(log, badNumFormatType, reflect.TypeOf(val))
}
