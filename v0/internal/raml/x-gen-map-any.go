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

func NewAnyMap() *AnyMap {
	return &AnyMap{
		index: make(map[interface{}]*interface{}),
	}
}

// AnyMap generated @ 2020-05-22T11:24:08.538522822-04:00
type AnyMap struct {
	slice []mapPair
	index map[interface{}]*interface{}
}

func (o *AnyMap) Len() uint {
	return uint(len(o.slice))
}

func (o *AnyMap) Empty() bool {
	return len(o.slice) == 0
}

func (o *AnyMap) Put(key interface{}, value interface{}) raml.AnyMap {
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *AnyMap) PutNonNil(key interface{}, value interface{}) raml.AnyMap {
	logrus.Trace("internal.AnyMap.PutNonNil")

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
	o.slice[ind.Get()].val = value
	return out
}

func (o *AnyMap) ReplaceOrPut(key interface{}, value interface{}) option.Untyped {
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

func (o *AnyMap) Get(key interface{}) option.Untyped {
	logrus.Trace("internal.AnyMap.Get")

	if !o.Has(key) {
		return option.NewEmptyUntyped()
	}

	return option.NewMaybeUntyped(o.index[key])
}

func (o *AnyMap) At(index uint) (key option.Untyped, value option.Untyped) {

	tmp := &o.slice[index]
	key = option.NewUntyped(tmp.key.(string))

	if util.IsNil(tmp.val) {
		value = option.NewEmptyUntyped()
	} else {
		value = option.NewUntyped(tmp.val.(interface{}))
	}

	return
}

func (o *AnyMap) IndexOf(key interface{}) option.Uint {
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

func (o *AnyMap) Has(key interface{}) bool {
	logrus.Trace("internal.AnyMap.Has")

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
		if o.slice[i].key == key {
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
	out := xyml.MapNode(len(o.slice))

	for i := range o.slice {
		var val interface{}

		if v, ok := o.slice[i].val.(raml.Marshaler); ok {
			tmp := NewAnyMap()
			if s, err := v.MarshalRAML(tmp); err != nil {
				return nil, err
			} else if s {
				_, t2 := tmp.At(0)
				val = t2.Get()
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

func (o *AnyMap) UnmarshalRAML(val *yaml.Node) (err error) {
	return xyml.ForEachMap(val, func(key, val *yaml.Node) error {
		altKey, err := xyml.CastScalarToYmlType(key)
		if err != nil {
			return err
		}

		tmpVal := val

		o.Put(altKey, tmpVal)

		return nil
	})
}

func (o *AnyMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(xyml.Indent)

	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
