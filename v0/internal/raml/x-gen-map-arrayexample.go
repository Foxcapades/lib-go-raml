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

func NewArrayExampleMap() *ArrayExampleMap {
	return &ArrayExampleMap{
		index: make(map[string]*raml.ArrayExample),
	}
}

// ArrayExampleMap generated @ 2020-05-21T01:49:31.367162698-04:00
type ArrayExampleMap struct {
	slice []mapPair
	index map[string]*raml.ArrayExample
}

func (o *ArrayExampleMap) Len() uint {
	return uint(len(o.slice))
}

func (o *ArrayExampleMap) Empty() bool {
	return len(o.slice) == 0
}

func (o *ArrayExampleMap) Put(key string, value raml.ArrayExample) raml.ArrayExampleMap {
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *ArrayExampleMap) PutNonNil(key string, value raml.ArrayExample) raml.ArrayExampleMap {
	logrus.Trace("internal.ArrayExampleMap.PutNonNil")

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
	o.slice[ind.Get()].val = value
	return out
}

func (o *ArrayExampleMap) ReplaceOrPut(key string, value raml.ArrayExample) raml.ArrayExample {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, mapPair{key: key, val: value})

		return nil
	}

	out := *o.index[key]
	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *ArrayExampleMap) Get(key string) raml.ArrayExample {
	logrus.Trace("internal.ArrayExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *ArrayExampleMap) At(index uint) (key option.String, value raml.ArrayExample) {

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.ArrayExample)

	return
}

func (o *ArrayExampleMap) IndexOf(key string) option.Uint {
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

func (o *ArrayExampleMap) Has(key string) bool {
	logrus.Trace("internal.ArrayExampleMap.Has")

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
		if o.slice[i].key == key {
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

func (o *ArrayExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	return xyml.ForEachMap(val, func(key, val *yaml.Node) error {
		altKey := key.Value

		tmpVal := NewArrayExample()

		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)

		return nil
	})
}

func (o *ArrayExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(xyml.Indent)

	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
