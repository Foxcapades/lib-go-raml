package collection

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

const (
	errNotCollType = "OrderedMap.PutNonEmpty called with non-collection type %s"
)

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		index: make(map[interface{}]*interface{}),
	}
}

// OrderedMap is a builder style wrapper around a
// yaml.MapSlice allowing the MapSlice to be constructed
// in a chained manner.
//
// Convenience methods are provided to conditionally insert
// values into the map.
//
// Additionally OrderedMap verifies that the MapSlice are
// unique on entry.
type OrderedMap struct {
	slice yaml.MapSlice
	index map[interface{}]*interface{}
}

// Put inserts the given key/value pair into the map.
//
// If a value previously existed in the map with the same
// key, it will be removed and this value will be appended
// to the end of the map.
func (o *OrderedMap) Put(key, val interface{}) *OrderedMap {
	o.Delete(key)
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: val})
	o.index[key] = &val
	return o
}

// Delete removes the value from the map with the given key.
func (o *OrderedMap) Delete(key interface{}) {
	if _, ok := o.index[key]; ok {
		for i := range o.slice {
			if o.slice[i].Key == key {
				o.slice = append(o.slice[:i], o.slice[i+1:]...)
				break
			}
		}
		delete(o.index, key)
	}
}

func (o *OrderedMap) Get(key interface{}) (interface{}, bool) {
	if val, ok := o.index[key]; ok {
		return *val, ok
	}
	return nil, false
}

func (o *OrderedMap) At(index int) (key, val interface{}) {
	return o.slice[index].Key, o.slice[index].Value
}

func (o *OrderedMap) KeyAt(index int) interface{} {
	return o.slice[index].Key
}

func (o *OrderedMap) ValueAt(index int) interface{} {
	return o.slice[index].Value
}

func (o *OrderedMap) Len() int {
	return len(o.slice)
}

// PutNonNil adds the given value `val` to the map at the
// key `key` if `val` is not nil.
//
// If a key previously existed in the map with the given
// key, it will not be removed unless the given value is not
// nil.
//
// This method uses reflection, use one of the typed methods
// if available.
func (o *OrderedMap) PutNonNil(key, val interface{}) *OrderedMap {
	t := reflect.ValueOf(val)

	switch t.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if t.IsNil() {
			break
		}
		fallthrough
	default:
		if val != nil {
			o.Put(key, val)
		}
	}
	return o
}

// PutNonNilString appends the given string value to the map
// at the given key if the value is not nil.
func (o *OrderedMap) PutNonNilString(key interface{}, val *string) *OrderedMap {
	if val != nil {
		o.Put(key, val)
	}
	return o
}

// PutNonNilUint appends the given uint value to the map at
// the given key if the value is not nil.
func (o *OrderedMap) PutNonNilUint(key interface{}, val *uint) *OrderedMap {
	if val != nil {
		o.Put(key, val)
	}
	return o
}

// PutNonNilInt64 appends the given int64 value to the map
// at the given key if the value is not nil.
func (o *OrderedMap) PutNonNilInt64(key interface{}, val *int64) *OrderedMap {
	if val != nil {
		o.Put(key, val)
	}
	return o
}

// PutNonNilFloat64 appends the given float64 value to the
// map at the given key if the value is not nil.
func (o *OrderedMap) PutNonNilFloat64(key interface{}, val *float64) *OrderedMap {
	if val != nil {
		o.Put(key, val)
	}
	return o
}

// PutNonEqual appends the given value if it does not equal
// the given default value.
func (o *OrderedMap) PutNonEqual(key, val, def interface{}) *OrderedMap {
	if val != def {
		o.Put(key, val)
	}
	return o
}

// PutNonEmpty appends the given value to the map if the
// value is of type map or slice, and it's length is > 0.
//
// If the given value is neither a map or a slice, this
// method panics.
func (o *OrderedMap) PutNonEmpty(key, val interface{}) *OrderedMap {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Map, reflect.Slice:
		if v.Len() > 0 {
			o.Put(key, val)
		}
	default:
		panic(fmt.Sprintf(errNotCollType, v.Type()))
	}

	return o
}

// PutNonEmptyStrList appends the given slice of strings to
// the map if it's length is greater than 0.
func (o *OrderedMap) PutNonEmptyStrList(key interface{}, val []string) *OrderedMap {
	if len(val) > 0 {
		o.Put(key, val)
	}

	return o
}

// ApplyAnyMap inserts all the keys and values from the
// given map into this map.
func (o *OrderedMap) ApplyAnyMap(mp map[interface{}]interface{}) *OrderedMap {
	for k := range mp {
		o.Put(k, mp[k])
	}
	return o
}

// ApplyStrAnyMap inserts all the keys and values from the
// given map into this map.
func (o *OrderedMap) ApplyStrAnyMap(mp map[string]interface{}) *OrderedMap {
	for k := range mp {
		o.Put(k, mp[k])
	}
	return o
}

// ApplyStrStrMap inserts all the keys and values from the
// given map into this map.
func (o *OrderedMap) ApplyStrStrMap(mp map[string]string) *OrderedMap {
	for k := range mp {
		o.Put(k, mp[k])
	}
	return o
}

func (o OrderedMap) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *OrderedMap) UnmarshalYAML(fn func(interface{}) error) error {
	if err := fn(&o.slice); err != nil {
		return err
	}
	for i := range o.slice {
		o.Put(o.slice[i].Key, &o.slice[i].Value)
	}
	return nil
}

func (o *OrderedMap) ensureMap() {
	if o.index == nil {
		o.index = make(map[interface{}]*interface{})
	}
}
