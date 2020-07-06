package raml

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// IntegerExampleMap defines an ordered map of string to IntegerExample.
type IntegerExampleMap interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v IntegerExample) IntegerExampleMap

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *IntegerExample) IntegerExampleMap

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v IntegerExample) IntegerExampleMap

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v IntegerExample) IntegerExampleMap

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value IntegerExample, exists bool)

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) IntegerExampleMapEntry

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

	Delete(k string) IntegerExampleMap

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v IntegerExample)) IntegerExampleMap

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) IntegerExampleMap
}

// IntegerExampleMapEntry is a single entry in an instance of
// IntegerExampleMap.
type IntegerExampleMapEntry struct {
	Key string         `json:"key"`
	Val IntegerExample `json:"value"`
}

// NewIntegerExampleMap creates a new instance of IntegerExampleMap presized to the
// given size.
func NewIntegerExampleMap(size int) IntegerExampleMap {
	return &implIntegerExampleMap{
		ordered:  make([]IntegerExampleMapEntry, 0, size),
		index:    make(map[string]IntegerExample, size),
		outOrder: true,
	}
}

// IntegerExampleMap is an ordered map string to IntegerExample.
type implIntegerExampleMap struct {
	ordered  []IntegerExampleMapEntry
	index    map[string]IntegerExample
	outOrder bool
}

func (i implIntegerExampleMap) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implIntegerExampleMap) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i *implIntegerExampleMap) ToYAML() (*yaml.Node, error) {
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

func (i *implIntegerExampleMap) Put(k string, v IntegerExample) IntegerExampleMap {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, IntegerExampleMapEntry{k, v})
	return i
}

func (i *implIntegerExampleMap) PutIfNotNil(k string, v *IntegerExample) IntegerExampleMap {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implIntegerExampleMap) ReplaceOrPut(k string, v IntegerExample) IntegerExampleMap {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implIntegerExampleMap) ReplaceIfExists(k string, v IntegerExample) IntegerExampleMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implIntegerExampleMap) Get(k string) (value IntegerExample, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implIntegerExampleMap) At(j int) IntegerExampleMapEntry {
	return i.ordered[j]
}

func (i *implIntegerExampleMap) Len() int {
	return len(i.ordered)
}

func (i *implIntegerExampleMap) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implIntegerExampleMap) IndexOf(k string) int {
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

func (i *implIntegerExampleMap) Delete(k string) IntegerExampleMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implIntegerExampleMap) ForEach(f func(k string, v IntegerExample)) IntegerExampleMap {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implIntegerExampleMap) SerializeOrdered(b bool) IntegerExampleMap {
	i.outOrder = b
	return i
}
