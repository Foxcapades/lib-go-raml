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

func NewTimeOnlyExampleMap() *TimeOnlyExampleMap {
	return &TimeOnlyExampleMap{
		index: make(map[string]*raml.TimeOnlyExample),
	}
}

// TimeOnlyExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type TimeOnlyExampleMap struct {
	slice []mapPair
	index map[string]*raml.TimeOnlyExample
}

func (o *TimeOnlyExampleMap) Len() uint {
	logrus.Trace("internal.TimeOnlyExampleMap.Len")
	return uint(len(o.slice))
}

func (o *TimeOnlyExampleMap) Put(key string, value raml.TimeOnlyExample) raml.TimeOnlyExampleMap {
	logrus.Trace("internal.TimeOnlyExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *TimeOnlyExampleMap) PutNonNil(key string, value raml.TimeOnlyExample) raml.TimeOnlyExampleMap {
	logrus.Trace("internal.TimeOnlyExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *TimeOnlyExampleMap) Replace(key string, value raml.TimeOnlyExample) raml.TimeOnlyExample {
	logrus.Trace("internal.TimeOnlyExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *TimeOnlyExampleMap) ReplaceOrPut(key string, value raml.TimeOnlyExample) raml.TimeOnlyExample {
	logrus.Trace("internal.TimeOnlyExampleMap.ReplaceOrPut")

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

func (o *TimeOnlyExampleMap) Get(key string) raml.TimeOnlyExample {
	logrus.Trace("internal.TimeOnlyExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *TimeOnlyExampleMap) At(index uint) (key option.String, value raml.TimeOnlyExample) {

	logrus.Trace("internal.TimeOnlyExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.TimeOnlyExample)

	return
}

func (o *TimeOnlyExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.TimeOnlyExampleMap.IndexOf")
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

func (o *TimeOnlyExampleMap) Has(key string) bool {
	logrus.Trace("internal.TimeOnlyExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *TimeOnlyExampleMap) Delete(key string) raml.TimeOnlyExample {
	logrus.Trace("internal.TimeOnlyExampleMap.Delete")

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

func (o TimeOnlyExampleMap) ForEach(fn func(string, raml.TimeOnlyExample)) {
	logrus.Trace("internal.TimeOnlyExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o TimeOnlyExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.TimeOnlyExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *TimeOnlyExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.TimeOnlyExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewTimeOnlyExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *TimeOnlyExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
