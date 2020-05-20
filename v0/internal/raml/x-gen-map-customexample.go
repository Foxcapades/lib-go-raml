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

func NewCustomExampleMap() *CustomExampleMap {
	return &CustomExampleMap{
		index: make(map[string]*raml.CustomExample),
	}
}

// CustomExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type CustomExampleMap struct {
	slice []mapPair
	index map[string]*raml.CustomExample
}

func (o *CustomExampleMap) Len() uint {
	logrus.Trace("internal.CustomExampleMap.Len")
	return uint(len(o.slice))
}

func (o *CustomExampleMap) Put(key string, value raml.CustomExample) raml.CustomExampleMap {
	logrus.Trace("internal.CustomExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *CustomExampleMap) PutNonNil(key string, value raml.CustomExample) raml.CustomExampleMap {
	logrus.Trace("internal.CustomExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *CustomExampleMap) Replace(key string, value raml.CustomExample) raml.CustomExample {
	logrus.Trace("internal.CustomExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *CustomExampleMap) ReplaceOrPut(key string, value raml.CustomExample) raml.CustomExample {
	logrus.Trace("internal.CustomExampleMap.ReplaceOrPut")

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

func (o *CustomExampleMap) Get(key string) raml.CustomExample {
	logrus.Trace("internal.CustomExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *CustomExampleMap) At(index uint) (key option.String, value raml.CustomExample) {

	logrus.Trace("internal.CustomExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.CustomExample)

	return
}

func (o *CustomExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.CustomExampleMap.IndexOf")
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

func (o *CustomExampleMap) Has(key string) bool {
	logrus.Trace("internal.CustomExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *CustomExampleMap) Delete(key string) raml.CustomExample {
	logrus.Trace("internal.CustomExampleMap.Delete")

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

func (o CustomExampleMap) ForEach(fn func(string, raml.CustomExample)) {
	logrus.Trace("internal.CustomExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o CustomExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.CustomExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *CustomExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.CustomExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewCustomExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *CustomExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
