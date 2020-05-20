package util

import "reflect"

func IsNil(val interface{}) bool {
	if val == nil {
		return true
	}

	ref := reflect.ValueOf(val)

	switch ref.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Array:
		return ref.IsNil()
	case reflect.Invalid:
		return true
	}

	return false
}
