package assign

import (
	"reflect"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/sirupsen/logrus"
)

func AsAnyList(v interface{}, log *logrus.Entry) ([]interface{}, error) {
	log.Trace("-> assign.AsAnyList")
	defer log.Trace("<- assign.AsAnyList")

	if val, ok := v.([]interface{}); ok {
		return val, nil
	}

	return nil, xlog.Errorf(log, errReqType, "array", reflect.TypeOf(v))
}
