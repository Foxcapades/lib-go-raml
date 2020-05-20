package cast

import (
	"fmt"
	"reflect"
)

const (
	errUnexpectedType = "expected type %s, instead got %s"
)

func AsString(v interface{}) (string, error) {
	if str, ok := v.(string); ok {
		return str, nil
	} else {
		return "", fmt.Errorf(errUnexpectedType, "string", reflect.TypeOf(v))
	}
}
