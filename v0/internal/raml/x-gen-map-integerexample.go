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

func NewIntegerExampleMap() *IntegerExampleMap {
	return &IntegerExampleMap{
		index: make(map[string]*raml.IntegerExample),
	}
}

// IntegerExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type IntegerExampleMap struct {
	slice []mapPair
	index map[string]*raml.IntegerExample
}

func (o *IntegerExampleMap) Len() uint {
	logrus.Trace("internal.IntegerExampleMap.Len")
	return uint(len(o.slice))
}

func (o *IntegerExampleMap) Put(key string, value raml.IntegerExample) raml.IntegerExampleMap {
	logrus.Trace("internal.IntegerExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *IntegerExampleMap) PutNonNil(key string, value raml.IntegerExample) raml.IntegerExampleMap {
	logrus.Trace("internal.IntegerExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *IntegerExampleMap) Replace(key string, value raml.IntegerExample) raml.IntegerExample {
	logrus.Trace("internal.IntegerExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *IntegerExampleMap) ReplaceOrPut(key string, value raml.IntegerExample) raml.IntegerExample {
	logrus.Trace("internal.IntegerExampleMap.ReplaceOrPut")

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

func (o *IntegerExampleMap) Get(key string) raml.IntegerExample {
	logrus.Trace("internal.IntegerExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *IntegerExampleMap) At(index uint) (key option.String, value raml.IntegerExample) {

	logrus.Trace("internal.IntegerExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.IntegerExample)

	return
}

func (o *IntegerExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.IntegerExampleMap.IndexOf")
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

func (o *IntegerExampleMap) Has(key string) bool {
	logrus.Trace("internal.IntegerExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *IntegerExampleMap) Delete(key string) raml.IntegerExample {
	logrus.Trace("internal.IntegerExampleMap.Delete")

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

func (o IntegerExampleMap) ForEach(fn func(string, raml.IntegerExample)) {
	logrus.Trace("internal.IntegerExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o IntegerExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.IntegerExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *IntegerExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.IntegerExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewIntegerExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *IntegerExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
