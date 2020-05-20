package raml

import (
	"fmt"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"strings"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewDataTypeMap() *DataTypeMap {
	return &DataTypeMap{
		index: make(map[string]*raml.DataType),
	}
}

// DataTypeMap generated @ 2020-05-20T18:40:12.501365164-04:00
type DataTypeMap struct {
	slice []mapPair
	index map[string]*raml.DataType
}

func (o *DataTypeMap) Len() uint {
	logrus.Trace("internal.DataTypeMap.Len")
	return uint(len(o.slice))
}

func (o *DataTypeMap) Put(key string, value raml.DataType) raml.DataTypeMap {
	logrus.Trace("internal.DataTypeMap.Put")
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
	logrus.Trace("internal.DataTypeMap.Replace")

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
	logrus.Trace("internal.DataTypeMap.ReplaceOrPut")

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

	logrus.Trace("internal.DataTypeMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.DataType)

	return
}

func (o *DataTypeMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.DataTypeMap.IndexOf")
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
	logrus.Trace("internal.DataTypeMap.Delete")

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
	logrus.Trace("internal.DataTypeMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o DataTypeMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.DataTypeMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *DataTypeMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.DataTypeMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal, err := TypeSortingHat(val)
		if err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *DataTypeMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
