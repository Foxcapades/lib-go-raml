package raml

import (
	"fmt"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"strings"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewStringMap() *StringMap {
	return &StringMap{
		index: make(map[string]*string),
	}
}

// StringMap generated @ 2020-05-21T01:49:31.367162698-04:00
type StringMap struct {
	slice []mapPair
	index map[string]*string
}

func (o *StringMap) Len() uint {
	return uint(len(o.slice))
}

func (o *StringMap) Empty() bool {
	return len(o.slice) == 0
}

func (o *StringMap) Put(key string, value string) raml.StringMap {
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *StringMap) PutNonNil(key string, value *string) raml.StringMap {
	logrus.Trace("internal.StringMap.PutNonNil")

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
	o.slice[ind.Get()].val = value
	return out
}

func (o *StringMap) ReplaceOrPut(key string, value string) option.String {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, mapPair{key: key, val: value})

		return option.NewEmptyString()
	}

	out := option.NewMaybeString(o.index[key])
	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *StringMap) Get(key string) option.String {
	logrus.Trace("internal.StringMap.Get")

	if !o.Has(key) {
		return option.NewEmptyString()
	}

	return option.NewMaybeString(o.index[key])
}

func (o *StringMap) At(index uint) (key option.String, value option.String) {

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	if util.IsNil(tmp.val) {
		value = option.NewEmptyString()
	} else {
		value = option.NewString(tmp.val.(string))
	}

	return
}

func (o *StringMap) IndexOf(key string) option.Uint {
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

func (o *StringMap) Has(key string) bool {
	logrus.Trace("internal.StringMap.Has")

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
		if o.slice[i].key == key {
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

func (o *StringMap) UnmarshalRAML(val *yaml.Node) (err error) {
	return xyml.ForEachMap(val, func(key, val *yaml.Node) error {
		altKey := key.Value

		var tmpVal string

		if err = assign.AsString(val, &tmpVal); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)

		return nil
	})
}

func (o *StringMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(xyml.Indent)

	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
