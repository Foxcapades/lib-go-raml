package raml

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// StringMap defines an ordered map of string to string.
type StringMap interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v string) StringMap

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *string) StringMap

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v string) StringMap

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v string) StringMap

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value string, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of string if an entry was found
	// with the key `k`.
	GetOpt(k string) option.String

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) StringMapEntry

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

	Delete(k string) StringMap

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v string)) StringMap

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) StringMap
}

// StringMapEntry is a single entry in an instance of
// StringMap.
type StringMapEntry struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

// NewStringMap creates a new instance of StringMap presized to the
// given size.
func NewStringMap(size int) StringMap {
	return &implStringMap{
		ordered:  make([]StringMapEntry, 0, size),
		index:    make(map[string]string, size),
		outOrder: true,
	}
}

// StringMap is an ordered map string to string.
type implStringMap struct {
	ordered  []StringMapEntry
	index    map[string]string
	outOrder bool
}

func (i implStringMap) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implStringMap) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i *implStringMap) ToYAML() (*yaml.Node, error) {
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

func (i *implStringMap) Put(k string, v string) StringMap {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, StringMapEntry{k, v})
	return i
}

func (i *implStringMap) PutIfNotNil(k string, v *string) StringMap {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implStringMap) ReplaceOrPut(k string, v string) StringMap {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implStringMap) ReplaceIfExists(k string, v string) StringMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implStringMap) Get(k string) (value string, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implStringMap) GetOpt(k string) option.String {
	if v, ok := i.index[k]; ok {
		return option.NewString(v)
	}

	return option.NewEmptyString()
}

func (i *implStringMap) At(j int) StringMapEntry {
	return i.ordered[j]
}

func (i *implStringMap) Len() int {
	return len(i.ordered)
}

func (i *implStringMap) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implStringMap) IndexOf(k string) int {
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

func (i *implStringMap) Delete(k string) StringMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implStringMap) ForEach(f func(k string, v string)) StringMap {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implStringMap) SerializeOrdered(b bool) StringMap {
	i.outOrder = b
	return i
}
