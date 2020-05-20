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

func NewObjectExampleMap() *ObjectExampleMap {
	return &ObjectExampleMap{
		index: make(map[string]*raml.ObjectExample),
	}
}

// ObjectExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type ObjectExampleMap struct {
	slice []mapPair
	index map[string]*raml.ObjectExample
}

func (o *ObjectExampleMap) Len() uint {
	logrus.Trace("internal.ObjectExampleMap.Len")
	return uint(len(o.slice))
}

func (o *ObjectExampleMap) Put(key string, value raml.ObjectExample) raml.ObjectExampleMap {
	logrus.Trace("internal.ObjectExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *ObjectExampleMap) PutNonNil(key string, value raml.ObjectExample) raml.ObjectExampleMap {
	logrus.Trace("internal.ObjectExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *ObjectExampleMap) Replace(key string, value raml.ObjectExample) raml.ObjectExample {
	logrus.Trace("internal.ObjectExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *ObjectExampleMap) ReplaceOrPut(key string, value raml.ObjectExample) raml.ObjectExample {
	logrus.Trace("internal.ObjectExampleMap.ReplaceOrPut")

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

func (o *ObjectExampleMap) Get(key string) raml.ObjectExample {
	logrus.Trace("internal.ObjectExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *ObjectExampleMap) At(index uint) (key option.String, value raml.ObjectExample) {

	logrus.Trace("internal.ObjectExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.ObjectExample)

	return
}

func (o *ObjectExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.ObjectExampleMap.IndexOf")
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

func (o *ObjectExampleMap) Has(key string) bool {
	logrus.Trace("internal.ObjectExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *ObjectExampleMap) Delete(key string) raml.ObjectExample {
	logrus.Trace("internal.ObjectExampleMap.Delete")

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

func (o ObjectExampleMap) ForEach(fn func(string, raml.ObjectExample)) {
	logrus.Trace("internal.ObjectExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o ObjectExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.ObjectExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *ObjectExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.ObjectExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewObjectExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *ObjectExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
