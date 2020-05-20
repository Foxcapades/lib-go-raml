package assign

import (
	"fmt"
	"reflect"
)

// Helper function for uint values in assign functions
func ToUint(v interface{}, ptr *uint) error {
	if val, ok := v.(uint); ok {
		*ptr = val
		return nil
	}
	if val, ok := v.(int); ok {
		*ptr = uint(val)
		return nil
	}
	return fmt.Errorf(errReqType, "uint", reflect.TypeOf(v))
}

// Helper function for uint values in assign functions
func AsUintPtr(v interface{}, ptr **uint) error {
	if val, ok := v.(uint); ok {
		*ptr = &val
		return nil
	}
	if val, ok := v.(int); ok {
		tmp := uint(val)
		*ptr = &tmp
		return nil
	}
	return fmt.Errorf(errReqType, "uint", reflect.TypeOf(v))
}
