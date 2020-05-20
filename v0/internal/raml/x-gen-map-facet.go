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

func NewFacetMap() *FacetMap {
	return &FacetMap{
		index: make(map[string]*raml.Facet),
	}
}

// FacetMap generated @ 2020-05-20T18:40:12.501365164-04:00
type FacetMap struct {
	slice []mapPair
	index map[string]*raml.Facet
}

func (o *FacetMap) Len() uint {
	logrus.Trace("internal.FacetMap.Len")
	return uint(len(o.slice))
}

func (o *FacetMap) Put(key string, value raml.Facet) raml.FacetMap {
	logrus.Trace("internal.FacetMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *FacetMap) PutNonNil(key string, value raml.Facet) raml.FacetMap {
	logrus.Trace("internal.FacetMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *FacetMap) Replace(key string, value raml.Facet) raml.Facet {
	logrus.Trace("internal.FacetMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *FacetMap) ReplaceOrPut(key string, value raml.Facet) raml.Facet {
	logrus.Trace("internal.FacetMap.ReplaceOrPut")

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

func (o *FacetMap) Get(key string) raml.Facet {
	logrus.Trace("internal.FacetMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *FacetMap) At(index uint) (key option.String, value raml.Facet) {

	logrus.Trace("internal.FacetMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.Facet)

	return
}

func (o *FacetMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.FacetMap.IndexOf")
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

func (o *FacetMap) Has(key string) bool {
	logrus.Trace("internal.FacetMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *FacetMap) Delete(key string) raml.Facet {
	logrus.Trace("internal.FacetMap.Delete")

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

func (o FacetMap) ForEach(fn func(string, raml.Facet)) {
	logrus.Trace("internal.FacetMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o FacetMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.FacetMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *FacetMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.FacetMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewFacet()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *FacetMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
