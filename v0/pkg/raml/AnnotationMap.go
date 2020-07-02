package raml

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// AnnotationMap defines an ordered map of string to Annotation.
type AnnotationMap interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v Annotation) AnnotationMap

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *Annotation) AnnotationMap

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v Annotation) AnnotationMap

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v Annotation) AnnotationMap

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value Annotation, exists bool)

	

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) AnnotationMapEntry

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

	Delete(k string) AnnotationMap

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v Annotation)) AnnotationMap
}

// AnnotationMapEntry is a single entry in an instance of
// AnnotationMap.
type AnnotationMapEntry struct {
	Key string     `json:"key"`
	Val Annotation `json:"value"`
}

// NewAnnotationMap creates a new instance of AnnotationMap presized to the
// given size.
func NewAnnotationMap(size int) AnnotationMap {
	return &implAnnotationMap{
		ordered: make([]AnnotationMapEntry, 0, size),
		index:   make(map[string]Annotation, size),
	}
}

// AnnotationMap is an ordered map string to Annotation.
type implAnnotationMap struct {
	ordered []AnnotationMapEntry
	index   map[string]Annotation
}

func (i implAnnotationMap) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implAnnotationMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implAnnotationMap) ToYAML() (*yaml.Node, error) {
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

func (i *implAnnotationMap) Put(k string, v Annotation) AnnotationMap {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, AnnotationMapEntry{k, v})
	return i
}

func (i *implAnnotationMap) PutIfNotNil(k string, v *Annotation) AnnotationMap {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implAnnotationMap) ReplaceOrPut(k string, v Annotation) AnnotationMap {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implAnnotationMap) ReplaceIfExists(k string, v Annotation) AnnotationMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implAnnotationMap) Get(k string) (value Annotation, exists bool) {
	v, ok := i.index[k]
	return v, ok
}



func (i *implAnnotationMap) At(j int) AnnotationMapEntry {
	return i.ordered[j]
}

func (i *implAnnotationMap) Len() int {
	return len(i.ordered)
}

func (i *implAnnotationMap) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implAnnotationMap) IndexOf(k string) int {
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

func (i *implAnnotationMap) Delete(k string) AnnotationMap {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implAnnotationMap) ForEach(f func(k string, v Annotation)) AnnotationMap {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}