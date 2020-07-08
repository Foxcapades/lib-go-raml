package raml

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// ObjectExampleMap defines an ordered map of string to ObjectExample.
type ObjectExampleMap interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v ObjectExample) ObjectExampleMap

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *ObjectExample) ObjectExampleMap

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v ObjectExample) ObjectExampleMap

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v ObjectExample) ObjectExampleMap

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value ObjectExample, exists bool)

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) ObjectExampleMapEntry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k string) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k string) int

	Delete(k string) ObjectExampleMap

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v ObjectExample)) ObjectExampleMap

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) ObjectExampleMap
}

// ObjectExampleMapEntry is a single entry in an instance of
// ObjectExampleMap.
type ObjectExampleMapEntry struct {
	Key string        `json:"key"`
	Val ObjectExample `json:"value"`
}

// NewObjectExampleMap creates a new instance of ObjectExampleMap presized to the
// given size.
func NewObjectExampleMap(size int) ObjectExampleMap {
	return &implObjectExampleMap{
		ordered:  make([]ObjectExampleMapEntry, 0, size),
		index:    make(map[string]ObjectExample, size),
		outOrder: true,
	}
}

// ObjectExampleMap is an ordered map string to ObjectExample.
type implObjectExampleMap struct {
	ordered  []ObjectExampleMapEntry
	index    map[string]ObjectExample
	outOrder bool
}

func (i implObjectExampleMap) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implObjectExampleMap) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i implObjectExampleMap) ToYAML() (*yaml.Node, error) {
	if i.outOrder {
		out := xyml.NewOrderedMapNode(i.Len())

		for j := range i.ordered {
			tmp := xyml.NewMapNode(1)
			if e := xyml.MapAppend(tmp, i.ordered[j].Key, i.ordered[j].Val); e != nil {
				return nil, e
			}
			if err := xyml.SequenceAppend(out, tmp); err != nil {
				return nil, err
			}
		}

		return out, nil
	}

	out := xyml.NewMapNode(i.Len())

	for j := range i.ordered {
		if e := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); e != nil {
			return nil, e
		}
	}

	return out, nil
}

func (i *implObjectExampleMap) Put(k string, v ObjectExample) ObjectExampleMap {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, ObjectExampleMapEntry{k, v})
	return i
}

func (i *implObjectExampleMap) PutIfNotNil(k string, v *ObjectExample) ObjectExampleMap {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implObjectExampleMap) ReplaceOrPut(k string, v ObjectExample) ObjectExampleMap {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implObjectExampleMap) ReplaceIfExists(k string, v ObjectExample) ObjectExampleMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implObjectExampleMap) Get(k string) (value ObjectExample, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implObjectExampleMap) At(j int) ObjectExampleMapEntry {
	return i.ordered[j]
}

func (i *implObjectExampleMap) Len() int {
	return len(i.ordered)
}

func (i *implObjectExampleMap) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implObjectExampleMap) IndexOf(k string) int {
	if _, ok := i.index[k]; !ok {
		return -1
	}
	for j := range i.ordered {
		if i.ordered[j].Key == k {
			return j
		}
	}
	panic("invalid map state, out of sync")
}

func (i *implObjectExampleMap) Delete(k string) ObjectExampleMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implObjectExampleMap) ForEach(f func(k string, v ObjectExample)) ObjectExampleMap {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implObjectExampleMap) SerializeOrdered(b bool) ObjectExampleMap {
	i.outOrder = b
	return i
}
