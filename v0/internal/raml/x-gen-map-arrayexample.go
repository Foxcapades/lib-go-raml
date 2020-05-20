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

func NewArrayExampleMap(log *logrus.Entry) *ArrayExampleMap {
	return &ArrayExampleMap{
		log:   xlog.WithType(log, "internal.ArrayExampleMap"),
		index: make(map[string]*raml.ArrayExample),
	}
}

// ArrayExampleMap generated @ 2020-05-20T01:05:35.571783841-04:00
type ArrayExampleMap struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	index map[string]*raml.ArrayExample
}

func (o *ArrayExampleMap) Len() uint {
	return uint(len(o.slice))
}

func (o *ArrayExampleMap) Put(key string, value raml.ArrayExample) raml.ArrayExampleMap {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

func (o *ArrayExampleMap) PutNonNil(key string, value raml.ArrayExample) raml.ArrayExampleMap {
	if !util.IsNil(value) {
		return o.Put(key, value)
	}
	return o
}

func (o *ArrayExampleMap) Replace(key string, value raml.ArrayExample) raml.ArrayExample {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *ArrayExampleMap) ReplaceOrPut(key string, value raml.ArrayExample) raml.ArrayExample {
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

func (o *ArrayExampleMap) Get(key string) raml.ArrayExample {
	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *ArrayExampleMap) At(index uint) (key option.String, value raml.ArrayExample) {
	tmp := &o.slice[index]
	key = option.NewString(tmp.Key.(string))
	value = tmp.Value.(raml.ArrayExample)

	return
}

func (o *ArrayExampleMap) IndexOf(key string) option.Uint {
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

func (o *ArrayExampleMap) Has(key string) bool {
	_, ok := o.index[key]
	return ok
}

func (o *ArrayExampleMap) Delete(key string) raml.ArrayExample {
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

func (o ArrayExampleMap) ForEach(fn func(string, raml.ArrayExample)) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o ArrayExampleMap) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *ArrayExampleMap) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.ArrayExampleMap.UnmarshalRAML")
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

		tmpVal := NewArrayExample(l2)
		if err = tmpVal.UnmarshalRAML(tmp.Value, l2); err != nil {
			return xlog.Error(l2, err)
		}

		o.Put(key, tmpVal)
	}

	return nil
}
