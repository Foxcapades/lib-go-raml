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

func NewAnnotationMap() *AnnotationMap {
	return &AnnotationMap{
		index: make(map[string]*raml.Annotation),
	}
}

// AnnotationMap generated @ 2020-05-22T11:24:08.538522822-04:00
type AnnotationMap struct {
	slice []mapPair
	index map[string]*raml.Annotation
}

func (o *AnnotationMap) Len() uint {
	return uint(len(o.slice))
}

func (o *AnnotationMap) Empty() bool {
	return len(o.slice) == 0
}

func (o *AnnotationMap) Put(key string, value raml.Annotation) raml.AnnotationMap {
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *AnnotationMap) PutNonNil(key string, value raml.Annotation) raml.AnnotationMap {
	logrus.Trace("internal.AnnotationMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *AnnotationMap) Replace(key string, value raml.Annotation) raml.Annotation {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *AnnotationMap) ReplaceOrPut(key string, value raml.Annotation) raml.Annotation {
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

func (o *AnnotationMap) Get(key string) raml.Annotation {
	logrus.Trace("internal.AnnotationMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *AnnotationMap) At(index uint) (key option.String, value raml.Annotation) {

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.Annotation)

	return
}

func (o *AnnotationMap) IndexOf(key string) option.Uint {
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

func (o *AnnotationMap) Has(key string) bool {
	logrus.Trace("internal.AnnotationMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *AnnotationMap) Delete(key string) raml.Annotation {
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

func (o AnnotationMap) ForEach(fn func(string, raml.Annotation)) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o AnnotationMap) MarshalYAML() (interface{}, error) {
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

func (o *AnnotationMap) UnmarshalRAML(val *yaml.Node) (err error) {
	return xyml.ForEachMap(val, func(key, val *yaml.Node) error {
		altKey := key.Value

		tmpVal := NewAnnotation()

		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)

		return nil
	})
}

func (o *AnnotationMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(xyml.Indent)

	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
