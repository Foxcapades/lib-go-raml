package collection

import (
	"gopkg.in/yaml.v2"
)

type StringAnyMap struct {
	serialize yaml.MapSlice
	lookup    map[string]interface{}
}

func (S StringAnyMap) MarshalYAML() (interface{}, error) {
	return S.serialize, nil
}

func (S *StringAnyMap) Put(key string, value interface{}) *StringAnyMap {
	EnsureStrAnyMap(&S.lookup)

	if _, ok := S.lookup[key]; ok {
		S.Delete(key)
	}

	S.serialize = append(S.serialize, yaml.MapItem{Key: key, Value: value})
	S.lookup[key] = value

	return S
}

func (S *StringAnyMap) Get(key string) interface{} {
	EnsureStrAnyMap(&S.lookup)

	return S.lookup[key]
}

func (S *StringAnyMap) Delete(key string) (out interface{}) {
	EnsureStrAnyMap(&S.lookup)

	out = S.lookup[key]
	delete(S.lookup, key)
	for i := range S.serialize {
		if S.serialize[i].Key == key {
			S.serialize = append(S.serialize[:i], S.serialize[i+1:]...)
			break
		}
	}

	return
}
