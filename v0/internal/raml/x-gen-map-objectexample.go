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

func NewObjectExampleMap(log *logrus.Entry) *ObjectExampleMap {
	return &ObjectExampleMap{
		log:   xlog.WithType(log, "internal.ObjectExampleMap"),
		index: make(map[string]*raml.ObjectExample),
	}
}

// ObjectExampleMap generated @ 2020-05-20T01:05:35.571783841-04:00
type ObjectExampleMap struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	index map[string]*raml.ObjectExample
}

func (o *ObjectExampleMap) Len() uint {
	return uint(len(o.slice))
}

func (o *ObjectExampleMap) Put(key string, value raml.ObjectExample) raml.ObjectExampleMap {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

func (o *ObjectExampleMap) PutNonNil(key string, value raml.ObjectExample) raml.ObjectExampleMap {
	if !util.IsNil(value) {
		return o.Put(key, value)
	}
	return o
}

func (o *ObjectExampleMap) Replace(key string, value raml.ObjectExample) raml.ObjectExample {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *ObjectExampleMap) ReplaceOrPut(key string, value raml.ObjectExample) raml.ObjectExample {
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

func (o *ObjectExampleMap) Get(key string) raml.ObjectExample {
	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *ObjectExampleMap) At(index uint) (key option.String, value raml.ObjectExample) {
	tmp := &o.slice[index]
	key = option.NewString(tmp.Key.(string))
	value = tmp.Value.(raml.ObjectExample)

	return
}

func (o *ObjectExampleMap) IndexOf(key string) option.Uint {
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

func (o *ObjectExampleMap) Has(key string) bool {
	_, ok := o.index[key]
	return ok
}

func (o *ObjectExampleMap) Delete(key string) raml.ObjectExample {
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

func (o ObjectExampleMap) ForEach(fn func(string, raml.ObjectExample)) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o ObjectExampleMap) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *ObjectExampleMap) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.ObjectExampleMap.UnmarshalRAML")
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

		tmpVal := NewObjectExample(l2)
		if err = tmpVal.UnmarshalRAML(tmp.Value, l2); err != nil {
			return xlog.Error(l2, err)
		}

		o.Put(key, tmpVal)
	}

	return nil
}
