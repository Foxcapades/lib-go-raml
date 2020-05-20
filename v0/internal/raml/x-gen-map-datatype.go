package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func NewDataTypeMap(log *logrus.Entry) *DataTypeMap {
	return &DataTypeMap{
		log:   xlog.WithType(log, "internal.DataTypeMap"),
		index: make(map[string]*raml.DataType),
	}
}

// DataTypeMap generated @ 2020-05-20T01:05:35.571783841-04:00
type DataTypeMap struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	index map[string]*raml.DataType
}

func (o *DataTypeMap) Len() uint {
	return uint(len(o.slice))
}

func (o *DataTypeMap) Put(key string, value raml.DataType) raml.DataTypeMap {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

func (o *DataTypeMap) PutNonNil(key string, value raml.DataType) raml.DataTypeMap {
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
	o.slice[ind.Get()].Value = value
	return out
}

func (o *DataTypeMap) ReplaceOrPut(key string, value raml.DataType) raml.DataType {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
		return nil
	}

	out := *o.index[key]
	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *DataTypeMap) Get(key string) raml.DataType {
	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *DataTypeMap) At(index uint) (key option.String, value raml.DataType) {
	tmp := &o.slice[index]
	key = option.NewString(tmp.Key.(string))
	value = tmp.Value.(raml.DataType)

	return
}

func (o *DataTypeMap) IndexOf(key string) option.Uint {
	if !o.Has(key) {
		return option.NewEmptyUint()
	}
	for i := range o.slice {
		if o.slice[i].Key == key {
			return option.NewUint(uint(i))
		}
	}
	panic("invalid map state, index out of sync")
}

func (o *DataTypeMap) Has(key string) bool {
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
		if o.slice[i].Key == key {
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
	return o.slice, nil
}

func (o *DataTypeMap) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.DataTypeMap.UnmarshalRAML")
	yml, err := assign.AsMapSlice(val)

	if err != nil {
		return xlog.Error(log, err)
	}

	for i := range yml {
		tmp := &yml[i]
		l2 := xlog.AddPath(log, tmp.Key)

		key := ""

		if err = assign.AsString(tmp.Key, &key, l2); err != nil {
			return xlog.Error(l2, err)
		}

		tmpVal, err := TypeSortingHat(tmp.Value, l2)
		if err != nil {
			return xlog.Error(l2, err)
		}

		o.Put(key, tmpVal)
	}

	return nil
}
