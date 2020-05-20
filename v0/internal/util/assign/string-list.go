package assign

import (
	"reflect"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/sirupsen/logrus"
)

func AsStringList(v interface{}, ptr *[]string, log *logrus.Entry) error {
	log.Trace("assign.AsStringList")

	arr, err := AsAnyList(v, log)

	if err != nil {
		return err
	}

	tmp := make([]string, len(arr))

	for i := range arr {
		l2 := xlog.AddPath(log, i)
		if str, ok := arr[i].(string); ok {
			tmp[i] = str
		} else {
			return xlog.Errorf(l2, errReqType, "string", reflect.TypeOf(arr[i]))
		}
	}

	*ptr = tmp
	return nil
}
