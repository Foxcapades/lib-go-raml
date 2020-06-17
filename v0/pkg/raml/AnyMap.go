package raml

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// AnyMap defines an ordered map of interface{} to interface{}.
type AnyMap interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v interface{}) AnyMap

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v interface{}) AnyMap

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v interface{}) AnyMap

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v interface{}) AnyMap

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value interface{}, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of interface{} if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Untyped

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) AnyMapEntry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k interface{}) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k interface{}) int

	Delete(k interface{}) AnyMap

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v interface{})) AnyMap
}

// AnyMapEntry is a single entry in an instance of
// AnyMap.
type AnyMapEntry struct {
	Key interface{} `json:"key"`
	Val interface{} `json:"value"`
}

// NewAnyMap creates a new instance of AnyMap presized to the
// given size.
func NewAnyMap(size int) AnyMap {
	return &implAnyMap{
		ordered: make([]AnyMapEntry, 0, size),
		index:   make(map[interface{}]interface{}, size),
	}
}

// AnyMap is an ordered map interface{} to interface{}.
type implAnyMap struct {
	ordered []AnyMapEntry
	index   map[interface{}]interface{}
}

func (i implAnyMap) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implAnyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implAnyMap) ToYAML() (*yaml.Node, error) {
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

func (i *implAnyMap) Put(k interface{}, v interface{}) AnyMap {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, AnyMapEntry{k, v})
	return i
}

func (i *implAnyMap) PutIfNotNil(k interface{}, v interface{}) AnyMap {
	if !gomp.IsNil(v) {
		return i.Put(k, gomp.Deref(v))
	}

	return i
}

func (i *implAnyMap) ReplaceOrPut(k interface{}, v interface{}) AnyMap {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implAnyMap) ReplaceIfExists(k interface{}, v interface{}) AnyMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implAnyMap) Get(k interface{}) (value interface{}, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implAnyMap) GetOpt(k interface{}) option.Untyped {
	if v, ok := i.index[k]; ok {
		return option.NewUntyped(v)
	}

	return option.NewEmptyUntyped()
}

func (i *implAnyMap) At(j int) AnyMapEntry {
	return i.ordered[j]
}

func (i *implAnyMap) Len() int {
	return len(i.ordered)
}

func (i *implAnyMap) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implAnyMap) IndexOf(k interface{}) int {
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

func (i *implAnyMap) Delete(k interface{}) AnyMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implAnyMap) ForEach(f func(k interface{}, v interface{})) AnyMap {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
