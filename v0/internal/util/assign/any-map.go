package assign

import (
	"fmt"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/cast"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"reflect"

	"gopkg.in/yaml.v2"
)

func AsMapSlice(v interface{}) (yaml.MapSlice, error) {
	if val, ok := v.(yaml.MapSlice); ok {
		return val, nil
	}

	return nil, fmt.Errorf(errReqType, "map", reflect.TypeOf(v))
}

func ToStringMap(v interface{}, ref raml.StringMap, log *logrus.Entry) error {
	log.Trace("assign.ToStringMap")
	slice, err := AsMapSlice(v)

	if err != nil {
		return xlog.Error(log, err)
	}

	for i := range slice {
		l2 := xlog.AddPath(log, slice[i].Key)

		if key, err := cast.AsString(slice[i].Key); err != nil {
			return xlog.Error(l2, err)
		} else if val, err := cast.AsString(slice[i].Value); err != nil {
			return xlog.Error(l2, err)
		} else {
			ref.Put(key, val)
		}
	}

	return nil
}

func ToUntypedMap(v interface{}, ref raml.UntypedMap, log *logrus.Entry) error {
	log.Trace("assign.ToUntypedMap")
	slice, err := AsMapSlice(v)

	if err != nil {
		return xlog.Error(log, err)
	}

	for i := range slice {
		l2 := xlog.AddPath(log, slice[i].Key)

		if key, err := cast.AsString(slice[i].Key); err != nil {
			return xlog.Error(l2, err)
		} else {
			ref.Put(key, slice[i].Value)
		}
	}

	return nil
}
