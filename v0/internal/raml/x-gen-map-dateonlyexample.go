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

func NewDateOnlyExampleMap() *DateOnlyExampleMap {
	return &DateOnlyExampleMap{
		index: make(map[string]*raml.DateOnlyExample),
	}
}

// DateOnlyExampleMap generated @ 2020-05-22T11:24:08.538522822-04:00
type DateOnlyExampleMap struct {
	slice []mapPair
	index map[string]*raml.DateOnlyExample
}

func (o *DateOnlyExampleMap) Len() uint {
	return uint(len(o.slice))
}

func (o *DateOnlyExampleMap) Empty() bool {
	return len(o.slice) == 0
}

func (o *DateOnlyExampleMap) Put(key string, value raml.DateOnlyExample) raml.DateOnlyExampleMap {
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *DateOnlyExampleMap) PutNonNil(key string, value raml.DateOnlyExample) raml.DateOnlyExampleMap {
	logrus.Trace("internal.DateOnlyExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *DateOnlyExampleMap) Replace(key string, value raml.DateOnlyExample) raml.DateOnlyExample {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *DateOnlyExampleMap) ReplaceOrPut(key string, value raml.DateOnlyExample) raml.DateOnlyExample {
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

func (o *DateOnlyExampleMap) Get(key string) raml.DateOnlyExample {
	logrus.Trace("internal.DateOnlyExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *DateOnlyExampleMap) At(index uint) (key option.String, value raml.DateOnlyExample) {

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.DateOnlyExample)

	return
}

func (o *DateOnlyExampleMap) IndexOf(key string) option.Uint {
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

func (o *DateOnlyExampleMap) Has(key string) bool {
	logrus.Trace("internal.DateOnlyExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *DateOnlyExampleMap) Delete(key string) raml.DateOnlyExample {
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

func (o DateOnlyExampleMap) ForEach(fn func(string, raml.DateOnlyExample)) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o DateOnlyExampleMap) MarshalYAML() (interface{}, error) {
	out := xyml.MapNode(len(o.slice))

	for i := range o.slice {
		var val interface{}

		if v, ok := o.slice[i].val.(raml.Marshaler); ok {
			tmp := NewAnyMap()
			if s, err := v.MarshalRAML(tmp); err != nil {
				return nil, err
			} else if s {
				_, t2 := tmp.At(0)
				val = t2.Get()
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

func (o *DateOnlyExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	return xyml.ForEachMap(val, func(key, val *yaml.Node) error {
		altKey := key.Value

		tmpVal := NewDateOnlyExample()

		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)

		return nil
	})
}

func (o *DateOnlyExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(xyml.Indent)

	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
