package raml

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// PropertyMap defines an ordered map of string to Property.
type PropertyMap interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v Property) PropertyMap

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *Property) PropertyMap

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v Property) PropertyMap

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v Property) PropertyMap

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value Property, exists bool)

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) PropertyMapEntry

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

	Delete(k string) PropertyMap

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v Property)) PropertyMap

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) PropertyMap
}

// PropertyMapEntry is a single entry in an instance of
// PropertyMap.
type PropertyMapEntry struct {
	Key string   `json:"key"`
	Val Property `json:"value"`
}

// NewPropertyMap creates a new instance of PropertyMap presized to the
// given size.
func NewPropertyMap(size int) PropertyMap {
	return &implPropertyMap{
		ordered:  make([]PropertyMapEntry, 0, size),
		index:    make(map[string]Property, size),
		outOrder: true,
	}
}

// PropertyMap is an ordered map string to Property.
type implPropertyMap struct {
	ordered  []PropertyMapEntry
	index    map[string]Property
	outOrder bool
}

func (i implPropertyMap) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implPropertyMap) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i *implPropertyMap) ToYAML() (*yaml.Node, error) {
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

func (i *implPropertyMap) Put(k string, v Property) PropertyMap {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, PropertyMapEntry{k, v})
	return i
}

func (i *implPropertyMap) PutIfNotNil(k string, v *Property) PropertyMap {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implPropertyMap) ReplaceOrPut(k string, v Property) PropertyMap {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implPropertyMap) ReplaceIfExists(k string, v Property) PropertyMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implPropertyMap) Get(k string) (value Property, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implPropertyMap) At(j int) PropertyMapEntry {
	return i.ordered[j]
}

func (i *implPropertyMap) Len() int {
	return len(i.ordered)
}

func (i *implPropertyMap) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implPropertyMap) IndexOf(k string) int {
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

func (i *implPropertyMap) Delete(k string) PropertyMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implPropertyMap) ForEach(f func(k string, v Property)) PropertyMap {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implPropertyMap) SerializeOrdered(b bool) PropertyMap {
	i.outOrder = b
	return i
}
