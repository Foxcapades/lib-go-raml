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

func NewFileExampleMap() *FileExampleMap {
	return &FileExampleMap{
		index: make(map[string]*raml.FileExample),
	}
}

// FileExampleMap generated @ 2020-05-20T18:40:12.501365164-04:00
type FileExampleMap struct {
	slice []mapPair
	index map[string]*raml.FileExample
}

func (o *FileExampleMap) Len() uint {
	logrus.Trace("internal.FileExampleMap.Len")
	return uint(len(o.slice))
}

func (o *FileExampleMap) Put(key string, value raml.FileExample) raml.FileExampleMap {
	logrus.Trace("internal.FileExampleMap.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

func (o *FileExampleMap) PutNonNil(key string, value raml.FileExample) raml.FileExampleMap {
	logrus.Trace("internal.FileExampleMap.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, value)
	}

	return o
}

func (o *FileExampleMap) Replace(key string, value raml.FileExample) raml.FileExample {
	logrus.Trace("internal.FileExampleMap.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		return nil
	}

	out := *o.index[key]

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *FileExampleMap) ReplaceOrPut(key string, value raml.FileExample) raml.FileExample {
	logrus.Trace("internal.FileExampleMap.ReplaceOrPut")

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

func (o *FileExampleMap) Get(key string) raml.FileExample {
	logrus.Trace("internal.FileExampleMap.Get")

	if !o.Has(key) {
		return nil
	}

	return *o.index[key]
}

func (o *FileExampleMap) At(index uint) (key option.String, value raml.FileExample) {

	logrus.Trace("internal.FileExampleMap.At")

	tmp := &o.slice[index]
	key = option.NewString(tmp.key.(string))

	value = tmp.val.(raml.FileExample)

	return
}

func (o *FileExampleMap) IndexOf(key string) option.Uint {
	logrus.Trace("internal.FileExampleMap.IndexOf")
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

func (o *FileExampleMap) Has(key string) bool {
	logrus.Trace("internal.FileExampleMap.Has")

	_, ok := o.index[key]
	return ok
}

func (o *FileExampleMap) Delete(key string) raml.FileExample {
	logrus.Trace("internal.FileExampleMap.Delete")

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

func (o FileExampleMap) ForEach(fn func(string, raml.FileExample)) {
	logrus.Trace("internal.FileExampleMap.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o FileExampleMap) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.FileExampleMap.MarshalYAML")

	out := xyml.MapNode(len(o.slice) * 2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *FileExampleMap) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.FileExampleMap.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		altKey := key.Value

		tmpVal := NewFileExample()
		if err = tmpVal.UnmarshalRAML(val); err != nil {
			return err
		}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *FileExampleMap) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
