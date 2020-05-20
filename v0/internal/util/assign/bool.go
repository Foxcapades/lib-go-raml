package assign

import (
	"reflect"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/sirupsen/logrus"
)


// Helper function for bool values in assign functions
func AsBool(v interface{}, ptr *bool, log *logrus.Entry) error {
	log.Trace("assign.AsBool")

	if val, ok := v.(bool); ok {
		*ptr = val
		return nil
	}

	return xlog.Errorf(log, errReqType, "bool", reflect.TypeOf(v))
}

// Helper function for bool values in assign functions
func AsBoolPtr(v interface{}, ptr **bool, log *logrus.Entry) error {
	log.Trace("assign.AsBoolPtr")

	if val, ok := v.(bool); ok {
		*ptr = &val
		return nil
	}

	return xlog.Errorf(log, errReqType, "bool", reflect.TypeOf(v))
}
