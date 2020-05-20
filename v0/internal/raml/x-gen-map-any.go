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

// AnyMap generated @ 2020-05-20T18:40:12.501365164-04:00
type AnyMap struct {
	slice []mapPair
	index map[interface{}]*interface{}
}

func (o *AnyMap) Len() uint {
	logrus.Trace("internal.AnyMap.Len")
	return uint(len(o.slice))
}

func (o *AnyMap) Put(key interface{}, value interface{}) raml.AnyMap {
	logrus.Trace("internal.AnyMap.Put")
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
	logrus.Trace("internal.AnyMap.Replace")

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
	logrus.Trace("internal.AnyMap.ReplaceOrPut")

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

	logrus.Trace("internal.AnyMap.At")

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
	logrus.Trace("internal.AnyMap.IndexOf")
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
	logrus.Trace("internal.AnyMap.Delete")

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
	logrus.Trace("internal.AnyMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o AnyMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.AnyMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *AnyMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.AnyMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey, err := xyml.CastScalarToYmlType(key)
		if err != nil {
			return err
		}

		tmpVal := val

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *AnyMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
