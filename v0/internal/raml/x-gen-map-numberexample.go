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

func NewNumberExampleMap() *NumberExampleMap {
	return &NumberExampleMap{
		index: make(map[string]*raml.NumberExample),
	}
}

// NumberExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type NumberExampleMap struct {
	slice []mapPair
	index map[string]*raml.NumberExample
}

func (o *NumberExampleMap) Len() uint {
	logrus.Trace("internal.NumberExampleMap.Len")
	return uint(len(o.slice))
}

func (o *NumberExampleMap) Put(key string, value raml.NumberExample) raml.NumberExampleMap {
	logrus.Trace("internal.NumberExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *NumberExampleMap) PutNonNil(key string, value raml.NumberExample) raml.NumberExampleMap {
	logrus.Trace("internal.NumberExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *NumberExampleMap) Replace(key string, value raml.NumberExample) raml.NumberExample {
	logrus.Trace("internal.NumberExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *NumberExampleMap) ReplaceOrPut(key string, value raml.NumberExample) raml.NumberExample {
	logrus.Trace("internal.NumberExampleMap.ReplaceOrPut")

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

func (o *NumberExampleMap) Get(key string) raml.NumberExample {
	logrus.Trace("internal.NumberExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *NumberExampleMap) At(index uint) (key option.String, value raml.NumberExample) {

	logrus.Trace("internal.NumberExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.NumberExample)

	return
}

func (o *NumberExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.NumberExampleMap.IndexOf")
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

func (o *NumberExampleMap) Has(key string) bool {
	logrus.Trace("internal.NumberExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *NumberExampleMap) Delete(key string) raml.NumberExample {
	logrus.Trace("internal.NumberExampleMap.Delete")

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

func (o NumberExampleMap) ForEach(fn func(string, raml.NumberExample)) {
	logrus.Trace("internal.NumberExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o NumberExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.NumberExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *NumberExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.NumberExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewNumberExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *NumberExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
