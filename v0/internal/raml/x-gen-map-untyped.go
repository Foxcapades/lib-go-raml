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

func NewUntypedMap(log *logrus.Entry) *UntypedMap {
	return &UntypedMap{
		log:   xlog.WithType(log, "internal.UntypedMap"),
		index: make(map[string]*interface{}),
	}
}

// UntypedMap generated @ 2020-05-20T01:05:35.571783841-04:00
type UntypedMap struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	index map[string]*interface{}
}

func (o *UntypedMap) Len() uint {
	return uint(len(o.slice))
}

func (o *UntypedMap) Put(key string, value interface{}) raml.UntypedMap {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

func (o *UntypedMap) PutNonNil(key string, value interface{}) raml.UntypedMap {
	if !util.IsNil(value) {
		return o.Put(key, value)
	}
	return o
}

func (o *UntypedMap) Replace(key string, value interface{}) option.Untyped {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return option.NewEmptyUntyped()
	}

	out := option.NewMaybeUntyped(o.index[key])

	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *UntypedMap) ReplaceOrPut(key string, value interface{}) option.Untyped {
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

func (o *UntypedMap) Get(key string) option.Untyped {
	if !o.Has(key) {
		return option.NewEmptyUntyped()
	}

	return option.NewMaybeUntyped(o.index[key])
}

func (o *UntypedMap) At(index uint) (key option.String, value option.Untyped) {
	tmp := &o.slice[index]
	key = option.NewString(tmp.Key.(string))
	
	if util.IsNil(tmp.Value) {
		value = option.NewEmptyUntyped()
	} else {
		value = option.NewUntyped(tmp.Value.(interface{}))
	}

	return
}

func (o *UntypedMap) IndexOf(key string) option.Uint {
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

func (o *UntypedMap) Has(key string) bool {
	_, ok := o.index[key]
	return ok
}

func (o *UntypedMap) Delete(key string) option.Untyped {
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

func (o UntypedMap) ForEach(fn func(string, interface{})) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o UntypedMap) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *UntypedMap) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.UntypedMap.UnmarshalRAML")
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

		tmpVal := tmp.Value

		o.Put(key, tmpVal)
	}

	return nil
}
