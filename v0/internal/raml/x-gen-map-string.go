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

func NewStringMap(log *logrus.Entry) *StringMap {
	return &StringMap{
		log:   xlog.WithType(log, "internal.StringMap"),
		index: make(map[string]*string),
	}
}

// StringMap generated @ 2020-05-20T01:05:35.571783841-04:00
type StringMap struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	index map[string]*string
}

func (o *StringMap) Len() uint {
	return uint(len(o.slice))
}

func (o *StringMap) Put(key string, value string) raml.StringMap {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

func (o *StringMap) PutNonNil(key string, value *string) raml.StringMap {
	if !util.IsNil(value) {
		return o.Put(key, *value)
	}
	return o
}

func (o *StringMap) Replace(key string, value string) option.String {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return option.NewEmptyString()
	}

	out := option.NewMaybeString(o.index[key])

	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *StringMap) ReplaceOrPut(key string, value string) option.String {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
		return option.NewEmptyString()
	}

	out := option.NewMaybeString(o.index[key])
	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *StringMap) Get(key string) option.String {
	if !o.Has(key) {
		return option.NewEmptyString()
	}

	return option.NewMaybeString(o.index[key])
}

func (o *StringMap) At(index uint) (key option.String, value option.String) {
	tmp := &o.slice[index]
	key = option.NewString(tmp.Key.(string))
	
	if util.IsNil(tmp.Value) {
		value = option.NewEmptyString()
	} else {
		value = option.NewString(tmp.Value.(string))
	}

	return
}

func (o *StringMap) IndexOf(key string) option.Uint {
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

func (o *StringMap) Has(key string) bool {
	_, ok := o.index[key]
	return ok
}

func (o *StringMap) Delete(key string) option.String {
	if !o.Has(key) {
		return option.NewEmptyString()
	}

	out := option.NewMaybeString(o.index[key])
	delete(o.index, key)

	for i := range o.slice {
		if o.slice[i].Key == key {
			o.slice = append(o.slice[:i], o.slice[i+1:]...)
			return out
		}
	}
	panic("invalid map state, index out of sync")
}

func (o StringMap) ForEach(fn func(string, string)) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o StringMap) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *StringMap) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.StringMap.UnmarshalRAML")
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

		var tmpVal string
		if err = assign.AsString(tmp.Value, &tmpVal, l2); err != nil {
			return xlog.Error(l2, err)
		}

		o.Put(key, tmpVal)
	}

	return nil
}
