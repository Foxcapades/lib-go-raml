package assign

import (
	"reflect"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/sirupsen/logrus"
)

func AsInt64Ptr(v interface{}, ptr **int64, log *logrus.Entry) error {
	log = log.WithField("func", "AsInt64Ptr")

	if val, ok := v.(int64); ok {
		*ptr = &val
		return nil
	}

	if val, ok := v.(int); ok {
		tv := int64(val)
		*ptr = &tv
		return nil
	}

	return xlog.Errorf(log, errReqType, "int64", reflect.TypeOf(v))
}

func AsInt64PtrExtra(v interface{}, ptr **int64, log *logrus.Entry) (bool, error) {
	log = log.WithField("func", "AsInt64PtrExtra")

	if val, ok := v.(int64); ok {
		*ptr = &val
		return true, nil
	}

	if val, ok := v.(int); ok {
		tv := int64(val)
		*ptr = &tv
		return true, nil
	}

	return false, xlog.Errorf(log, errReqType, "int64", reflect.TypeOf(v))
}
