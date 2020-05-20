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

// AnnotationMap generated @ 2020-05-20T18:40:12.501365164-04:00
type AnnotationMap struct {
	slice []mapPair
	index map[string]*raml.Annotation
}

func (o *AnnotationMap) Len() uint {
	logrus.Trace("internal.AnnotationMap.Len")
	return uint(len(o.slice))
}

func (o *AnnotationMap) Put(key string, value raml.Annotation) raml.AnnotationMap {
	logrus.Trace("internal.AnnotationMap.Put")
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
	logrus.Trace("internal.AnnotationMap.Replace")

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
	logrus.Trace("internal.AnnotationMap.ReplaceOrPut")

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

	logrus.Trace("internal.AnnotationMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.Annotation)

	return
}

func (o *AnnotationMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.AnnotationMap.IndexOf")
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
	logrus.Trace("internal.AnnotationMap.Delete")

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
	logrus.Trace("internal.AnnotationMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o AnnotationMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.AnnotationMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *AnnotationMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.AnnotationMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewAnnotation()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *AnnotationMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
