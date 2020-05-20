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

func NewBoolExampleMap() *BoolExampleMap {
	return &BoolExampleMap{
		index: make(map[string]*raml.BoolExample),
	}
}

// BoolExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type BoolExampleMap struct {
	slice []mapPair
	index map[string]*raml.BoolExample
}

func (o *BoolExampleMap) Len() uint {
	logrus.Trace("internal.BoolExampleMap.Len")
	return uint(len(o.slice))
}

func (o *BoolExampleMap) Put(key string, value raml.BoolExample) raml.BoolExampleMap {
	logrus.Trace("internal.BoolExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *BoolExampleMap) PutNonNil(key string, value raml.BoolExample) raml.BoolExampleMap {
	logrus.Trace("internal.BoolExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *BoolExampleMap) Replace(key string, value raml.BoolExample) raml.BoolExample {
	logrus.Trace("internal.BoolExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *BoolExampleMap) ReplaceOrPut(key string, value raml.BoolExample) raml.BoolExample {
	logrus.Trace("internal.BoolExampleMap.ReplaceOrPut")

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

func (o *BoolExampleMap) Get(key string) raml.BoolExample {
	logrus.Trace("internal.BoolExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *BoolExampleMap) At(index uint) (key option.String, value raml.BoolExample) {

	logrus.Trace("internal.BoolExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.BoolExample)

	return
}

func (o *BoolExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.BoolExampleMap.IndexOf")
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

func (o *BoolExampleMap) Has(key string) bool {
	logrus.Trace("internal.BoolExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *BoolExampleMap) Delete(key string) raml.BoolExample {
	logrus.Trace("internal.BoolExampleMap.Delete")

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

func (o BoolExampleMap) ForEach(fn func(string, raml.BoolExample)) {
	logrus.Trace("internal.BoolExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o BoolExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.BoolExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *BoolExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.BoolExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewBoolExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *BoolExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
