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

func NewUnionExampleMap() *UnionExampleMap {
	return &UnionExampleMap{
		index: make(map[string]*raml.UnionExample),
	}
}

// UnionExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type UnionExampleMap struct {
	slice []mapPair
	index map[string]*raml.UnionExample
}

func (o *UnionExampleMap) Len() uint {
	logrus.Trace("internal.UnionExampleMap.Len")
	return uint(len(o.slice))
}

func (o *UnionExampleMap) Put(key string, value raml.UnionExample) raml.UnionExampleMap {
	logrus.Trace("internal.UnionExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *UnionExampleMap) PutNonNil(key string, value raml.UnionExample) raml.UnionExampleMap {
	logrus.Trace("internal.UnionExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *UnionExampleMap) Replace(key string, value raml.UnionExample) raml.UnionExample {
	logrus.Trace("internal.UnionExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *UnionExampleMap) ReplaceOrPut(key string, value raml.UnionExample) raml.UnionExample {
	logrus.Trace("internal.UnionExampleMap.ReplaceOrPut")

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

func (o *UnionExampleMap) Get(key string) raml.UnionExample {
	logrus.Trace("internal.UnionExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *UnionExampleMap) At(index uint) (key option.String, value raml.UnionExample) {

	logrus.Trace("internal.UnionExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.UnionExample)

	return
}

func (o *UnionExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.UnionExampleMap.IndexOf")
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

func (o *UnionExampleMap) Has(key string) bool {
	logrus.Trace("internal.UnionExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *UnionExampleMap) Delete(key string) raml.UnionExample {
	logrus.Trace("internal.UnionExampleMap.Delete")

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

func (o UnionExampleMap) ForEach(fn func(string, raml.UnionExample)) {
	logrus.Trace("internal.UnionExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o UnionExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.UnionExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *UnionExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.UnionExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewUnionExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *UnionExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
