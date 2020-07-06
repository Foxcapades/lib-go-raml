package raml

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// FileExampleMap defines an ordered map of string to FileExample.
type FileExampleMap interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v FileExample) FileExampleMap

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *FileExample) FileExampleMap

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v FileExample) FileExampleMap

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v FileExample) FileExampleMap

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value FileExample, exists bool)

	

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) FileExampleMapEntry

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

	Delete(k string) FileExampleMap

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v FileExample)) FileExampleMap

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) FileExampleMap
}

// FileExampleMapEntry is a single entry in an instance of
// FileExampleMap.
type FileExampleMapEntry struct {
	Key string      `json:"key"`
	Val FileExample `json:"value"`
}

// NewFileExampleMap creates a new instance of FileExampleMap presized to the
// given size.
func NewFileExampleMap(size int) FileExampleMap {
	return &implFileExampleMap{
		ordered:  make([]FileExampleMapEntry, 0, size),
		index:    make(map[string]FileExample, size),
		outOrder: true,
	}
}

// FileExampleMap is an ordered map string to FileExample.
type implFileExampleMap struct {
	ordered  []FileExampleMapEntry
	index    map[string]FileExample
	outOrder bool
}

func (i implFileExampleMap) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implFileExampleMap) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i *implFileExampleMap) ToYAML() (*yaml.Node, error) {
	if i.outOrder {
		out := xyml.NewOrderedMapNode(i.Len())

		for j := range i.ordered {
			tmp := xyml.NewMapNode(1)
			_ = xyml.MapAppend(tmp, i.ordered[j].Key, i.ordered[j].Val)
			if err := xyml.SequenceAppend(out, tmp); err != nil {
				return nil, err
			}
		}

		return out, nil
	}

	out := xyml.NewMapNode(i.Len())

	for j := range i.ordered {
		xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val)
	}

	return out, nil
}

func (i *implFileExampleMap) Put(k string, v FileExample) FileExampleMap {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, FileExampleMapEntry{k, v})
	return i
}

func (i *implFileExampleMap) PutIfNotNil(k string, v *FileExample) FileExampleMap {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implFileExampleMap) ReplaceOrPut(k string, v FileExample) FileExampleMap {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implFileExampleMap) ReplaceIfExists(k string, v FileExample) FileExampleMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implFileExampleMap) Get(k string) (value FileExample, exists bool) {
	v, ok := i.index[k]
	return v, ok
}



func (i *implFileExampleMap) At(j int) FileExampleMapEntry {
	return i.ordered[j]
}

func (i *implFileExampleMap) Len() int {
	return len(i.ordered)
}

func (i *implFileExampleMap) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implFileExampleMap) IndexOf(k string) int {
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

func (i *implFileExampleMap) Delete(k string) FileExampleMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implFileExampleMap) ForEach(f func(k string, v FileExample)) FileExampleMap {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implFileExampleMap) SerializeOrdered(b bool) FileExampleMap {
	i.outOrder = b
	return i
}
