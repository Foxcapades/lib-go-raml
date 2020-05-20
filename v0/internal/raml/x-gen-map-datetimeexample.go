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

func NewDatetimeExampleMap() *DatetimeExampleMap {
	return &DatetimeExampleMap{
		index: make(map[string]*raml.DatetimeExample),
	}
}

// DatetimeExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type DatetimeExampleMap struct {
	slice []mapPair
	index map[string]*raml.DatetimeExample
}

func (o *DatetimeExampleMap) Len() uint {
	logrus.Trace("internal.DatetimeExampleMap.Len")
	return uint(len(o.slice))
}

func (o *DatetimeExampleMap) Put(key string, value raml.DatetimeExample) raml.DatetimeExampleMap {
	logrus.Trace("internal.DatetimeExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *DatetimeExampleMap) PutNonNil(key string, value raml.DatetimeExample) raml.DatetimeExampleMap {
	logrus.Trace("internal.DatetimeExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *DatetimeExampleMap) Replace(key string, value raml.DatetimeExample) raml.DatetimeExample {
	logrus.Trace("internal.DatetimeExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *DatetimeExampleMap) ReplaceOrPut(key string, value raml.DatetimeExample) raml.DatetimeExample {
	logrus.Trace("internal.DatetimeExampleMap.ReplaceOrPut")

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

func (o *DatetimeExampleMap) Get(key string) raml.DatetimeExample {
	logrus.Trace("internal.DatetimeExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *DatetimeExampleMap) At(index uint) (key option.String, value raml.DatetimeExample) {

	logrus.Trace("internal.DatetimeExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.DatetimeExample)

	return
}

func (o *DatetimeExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.DatetimeExampleMap.IndexOf")
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

func (o *DatetimeExampleMap) Has(key string) bool {
	logrus.Trace("internal.DatetimeExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *DatetimeExampleMap) Delete(key string) raml.DatetimeExample {
	logrus.Trace("internal.DatetimeExampleMap.Delete")

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

func (o DatetimeExampleMap) ForEach(fn func(string, raml.DatetimeExample)) {
	logrus.Trace("internal.DatetimeExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o DatetimeExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.DatetimeExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *DatetimeExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.DatetimeExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewDatetimeExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *DatetimeExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
