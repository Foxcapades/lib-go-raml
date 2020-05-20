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

func NewPropertyMap() *PropertyMap {
	return &PropertyMap{
		index: make(map[string]*raml.Property),
	}
}

// PropertyMap generated @ 2020-05-20T18:40:12.501365164-04:00
type PropertyMap struct {
	slice []mapPair
	index map[string]*raml.Property
}

func (o *PropertyMap) Len() uint {
	logrus.Trace("internal.PropertyMap.Len")
	return uint(len(o.slice))
}

func (o *PropertyMap) Put(key string, value raml.Property) raml.PropertyMap {
	logrus.Trace("internal.PropertyMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *PropertyMap) PutNonNil(key string, value raml.Property) raml.PropertyMap {
	logrus.Trace("internal.PropertyMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *PropertyMap) Replace(key string, value raml.Property) raml.Property {
	logrus.Trace("internal.PropertyMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *PropertyMap) ReplaceOrPut(key string, value raml.Property) raml.Property {
	logrus.Trace("internal.PropertyMap.ReplaceOrPut")

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

func (o *PropertyMap) Get(key string) raml.Property {
	logrus.Trace("internal.PropertyMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *PropertyMap) At(index uint) (key option.String, value raml.Property) {

	logrus.Trace("internal.PropertyMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.Property)

	return
}

func (o *PropertyMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.PropertyMap.IndexOf")
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

func (o *PropertyMap) Has(key string) bool {
	logrus.Trace("internal.PropertyMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *PropertyMap) Delete(key string) raml.Property {
	logrus.Trace("internal.PropertyMap.Delete")

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

func (o PropertyMap) ForEach(fn func(string, raml.Property)) {
	logrus.Trace("internal.PropertyMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o PropertyMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.PropertyMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *PropertyMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.PropertyMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal, err := PropertySortingHat(val)
		if err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *PropertyMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
