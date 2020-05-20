package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"reflect"
)

const badDateFormatType = "invalid date format value. expected string, got %s"

// DateFormatSortingHat parses the input into the appropriate implementation of
// the raml.DateFormat interface.
func DateFormatSortingHat(
	val interface{},
	log *logrus.Entry,
) (raml.DateFormat, error) {
	log.Trace("DateFormatSortingHat")

	if val, ok := val.(string); ok {
		df := DateFormat(val)
		switch df {
		case DateFormatRfc3339:
			return &df, nil
		case DateFormatRfc2616:
			return &df, nil
		}
		// TODO: validation point, if validation is enabled, this should be an error
		return &df, nil
	}

	return nil, xlog.Errorf(log, badDateFormatType, reflect.TypeOf(val))
}
