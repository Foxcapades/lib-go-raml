package collection

func EnsureAnyMap(ptr *map[interface{}]interface{}) {
	if *ptr == nil {
		*ptr = make(map[interface{}]interface{})
	}
}

func EnsureStrAnyMap(ptr *map[string]interface{}) {
	if *ptr == nil {
		*ptr = make(map[string]interface{})
	}
}
