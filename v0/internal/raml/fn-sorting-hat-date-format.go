package raml

import (
	"fmt"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const badDateFormatType = "invalid date format value. expected string, got %s at %d:%d"

// DateFormatSortingHat parses the input into the appropriate implementation of
// the raml.DateFormat interface.
func DateFormatSortingHat(val *yaml.Node) (raml.DateFormat, error) {
	logrus.Trace("DateFormatSortingHat")

	if xyml.IsString(val) {
		df := DateFormat(val.Value)
		switch df {
		case DateFormatRfc3339:
			return &df, nil
		case DateFormatRfc2616:
			return &df, nil
		}
		// TODO: validation point, if validation is enabled, this should be an error
		return &df, nil
	}

	return nil, fmt.Errorf(badDateFormatType, val.Tag, val.Line, val.Column)
}
