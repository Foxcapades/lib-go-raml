package raml

import (
	"fmt"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"strings"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewDataTypeMap() *DataTypeMap {
	return &DataTypeMap{
		index: make(map[string]*raml.DataType),
	}
}

// DataTypeMap generated @ 2020-05-21T01:49:31.367162698-04:00
type DataTypeMap struct {
	slice []mapPair
	index map[string]*raml.DataType
}

func (o *DataTypeMap) Len() uint {
	return uint(len(o.slice))
}

func (o *DataTypeMap) Empty() bool {
	return len(o.slice) == 0
}

func (o *DataTypeMap) Put(key string, value raml.DataType) raml.DataTypeMap {
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *DataTypeMap) PutNonNil(key string, value raml.DataType) raml.DataTypeMap {
	logrus.Trace("internal.DataTypeMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *DataTypeMap) Replace(key string, value raml.DataType) raml.DataType {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *DataTypeMap) ReplaceOrPut(key string, value raml.DataType) raml.DataType {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, mapPair{key: key, val: value})

		return nil
	}

	out := *o.index[key]
	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *DataTypeMap) Get(key string) raml.DataType {
	logrus.Trace("internal.DataTypeMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *DataTypeMap) At(index uint) (key option.String, value raml.DataType) {

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.DataType)

	return
}

func (o *DataTypeMap) IndexOf(key string) option.Uint {
	if !o.Has(key) {
		return option.NewEmptyUint()
	}

	for i := range o.slice {
		if o.slice[i].key == key {
			return option.NewUint(uint(i))
		}
	}

	panic("invalid map state, index out of sync")
}

func (o *DataTypeMap) Has(key string) bool {
	logrus.Trace("internal.DataTypeMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *DataTypeMap) Delete(key string) raml.DataType {
	if !o.Has(key) {
		return nil
	}

	out := *o.index[key]
	delete(o.index, key)

	for i := range o.slice {
		if o.slice[i].key == key {
			o.slice = append(o.slice[:i], o.slice[i+1:]...)
			return out
		}
	}

	panic("invalid map state, index out of sync")
}

func (o DataTypeMap) ForEach(fn func(string, raml.DataType)) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o DataTypeMap) MarshalYAML() (interface{}, error) {
	out := xyml.MapNode(len(o.slice))

	for i := range o.slice {
		var val interface{}

		if v, ok := o.slice[i].val.(raml.Marshaler); ok {
			tmp := NewAnyMap()
			if s, err := v.MarshalRAML(tmp); err != nil {
				return nil, err
			} else if s {
				val = tmp.Get(rmeta.KeyType).Get()
			} else {
				val = tmp
			}
		} else {
			val = o.slice[i].val
		}

		if err := xyml.AppendToMap(out, o.slice[i].key, val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (o *DataTypeMap) UnmarshalRAML(val *yaml.Node) (err error) {
	return xyml.ForEachMap(val, func(key, val *yaml.Node) error {
		altKey := key.Value

		tmpVal, err := TypeSortingHat(val)

		if err != nil {
			return err
		}

		o.Put(altKey, tmpVal)

		return nil
	})
}

func (o *DataTypeMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(xyml.Indent)

	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
