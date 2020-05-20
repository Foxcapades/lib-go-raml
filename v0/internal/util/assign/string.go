package assign

import (
	"github.com/sirupsen/logrus"
	"reflect"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
)

func AsString(v interface{}, ptr *string, log *logrus.Entry) error {
	log.Trace("assign.AsString")

	if val, ok := v.(string); ok {
		*ptr = val
		return nil
	}

	return xlog.Errorf(log, errReqType, "string", reflect.TypeOf(v))
}

func AsStringPtr(v interface{}, ptr **string, log *logrus.Entry) error {
	log.Trace("assign.AsStringPtr")

	if val, ok := v.(string); ok {
		*ptr = &val
		return nil
	}

	return xlog.Errorf(log, errReqType, "string", reflect.TypeOf(v))
}
