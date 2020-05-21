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

func NewUntypedMap() *UntypedMap {
	return &UntypedMap{
		index: make(map[string]*interface{}),
	}
}

// UntypedMap generated @ 2020-05-20T21:46:00.242352937-04:00
type UntypedMap struct {
	slice []mapPair
	index map[string]*interface{}
}

func (o *UntypedMap) Len() uint {
	return uint(len(o.slice))
}

func (o *UntypedMap) Put(key string, value interface{}) raml.UntypedMap {
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *UntypedMap) PutNonNil(key string, value interface{}) raml.UntypedMap {
	logrus.Trace("internal.UntypedMap.PutNonNil")

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
	o.slice[ind.Get()].val = value
	return out
}

func (o *UntypedMap) ReplaceOrPut(key string, value interface{}) option.Untyped {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, mapPair{key: key, val: value})

		return option.NewEmptyUntyped()
	}

	out := option.NewMaybeUntyped(o.index[key])
	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *UntypedMap) Get(key string) option.Untyped {
	logrus.Trace("internal.UntypedMap.Get")

	if !o.Has(key) {
		return option.NewEmptyUntyped()
	}

	return option.NewMaybeUntyped(o.index[key])
}

func (o *UntypedMap) At(index uint) (key option.String, value option.Untyped) {

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	if util.IsNil(tmp.val) {
		value = option.NewEmptyUntyped()
	} else {
		value = option.NewUntyped(tmp.val.(interface{}))
	}

	return
}

func (o *UntypedMap) IndexOf(key string) option.Uint {
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

func (o *UntypedMap) Has(key string) bool {
	logrus.Trace("internal.UntypedMap.Has")

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
		if o.slice[i].key == key {
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

func (o *UntypedMap) UnmarshalRAML(val *yaml.Node) (err error) {
	return xyml.ForEachMap(val, func(key, val *yaml.Node) error {
		altKey := key.Value

		tmpVal := val

		o.Put(altKey, tmpVal)

		return nil
	})
}

func (o *UntypedMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(xyml.Indent)

	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
