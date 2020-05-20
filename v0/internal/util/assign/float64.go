package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/sirupsen/logrus"
	"reflect"
)

func AsFloat64Ptr(v interface{}, ptr **float64, log *logrus.Entry) error {
	log = log.WithField("func", "AsFloat64Ptr")

	if val, ok := v.(float64); ok {
		*ptr = &val
		return nil
	}

	return xlog.Errorf(log, errReqType, "float64", reflect.TypeOf(v))
}

func AsFloat64PtrExtra(v interface{}, ptr **float64, log *logrus.Entry) (bool, error) {
	log = log.WithField("func", "AsFloat64PtrExtra")

	if err := AsFloat64Ptr(v, ptr, log); err != nil {
		return false, xlog.Error(log, err)
	}

	return true, nil
}
