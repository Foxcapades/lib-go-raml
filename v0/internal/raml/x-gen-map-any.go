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

func NewAnyMap(log *logrus.Entry) *AnyMap {
	return &AnyMap{
		log:   xlog.WithType(log, "internal.AnyMap"),
		index: make(map[interface{}]*interface{}),
	}
}

// AnyMap generated @ 2020-05-20T01:05:35.571783841-04:00
type AnyMap struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	index map[interface{}]*interface{}
}

func (o *AnyMap) Len() uint {
	return uint(len(o.slice))
}

func (o *AnyMap) Put(key interface{}, value interface{}) raml.AnyMap {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

func (o *AnyMap) PutNonNil(key interface{}, value interface{}) raml.AnyMap {
	if !util.IsNil(value) {
		return o.Put(key, value)
	}
	return o
}

func (o *AnyMap) Replace(key interface{}, value interface{}) option.Untyped {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return option.NewEmptyUntyped()
	}

	out := option.NewMaybeUntyped(o.index[key])

	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *AnyMap) ReplaceOrPut(key interface{}, value interface{}) option.Untyped {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
		return option.NewEmptyUntyped()
	}

	out := option.NewMaybeUntyped(o.index[key])
	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *AnyMap) Get(key interface{}) option.Untyped {
	if !o.Has(key) {
		return option.NewEmptyUntyped()
	}

	return option.NewMaybeUntyped(o.index[key])
}

func (o *AnyMap) At(index uint) (key option.Untyped, value option.Untyped) {
	tmp := &o.slice[index]
	key = option.NewUntyped(tmp.Key.(string))
	
	if util.IsNil(tmp.Value) {
		value = option.NewEmptyUntyped()
	} else {
		value = option.NewUntyped(tmp.Value.(interface{}))
	}

	return
}

func (o *AnyMap) IndexOf(key interface{}) option.Uint {
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

func (o *AnyMap) Has(key interface{}) bool {
	_, ok := o.index[key]
	return ok
}

func (o *AnyMap) Delete(key interface{}) option.Untyped {
	if !o.Has(key) {
		return option.NewEmptyUntyped()
	}

	out := option.NewMaybeUntyped(o.index[key])
	delete(o.index, key)

	for i := range o.slice {
		if o.slice[i].Key == key {
			o.slice = append(o.slice[:i], o.slice[i+1:]...)
			return out
		}
	}
	panic("invalid map state, index out of sync")
}

func (o AnyMap) ForEach(fn func(interface{}, interface{})) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o AnyMap) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *AnyMap) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.AnyMap.UnmarshalRAML")
	yml, err := assign.AsMapSlice(val)

	if err != nil {
		return xlog.Error(log, err)
	}

	for i := range yml {
		tmp := &yml[i]
		

		key := tmp.Key

		tmpVal := tmp.Value

		o.Put(key, tmpVal)
	}

	return nil
}
