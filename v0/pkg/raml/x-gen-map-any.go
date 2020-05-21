package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

// AnyMap generated @ 2020-05-21T01:49:31.367162698-04:00
type AnyMap interface {
	Unmarshaler
	yaml.Marshaler

	// Len returns the current number of elements in the
	// map.
	Len() uint

	// Empty returns whether the current map has 0 elements.
	Empty() bool

	// Put inserts a new element at the end of the map.
	//
	// If there was already an entry in the map with the given
	// key, it will be replaced.  The new key/value pair will
	// be appended to the end of the map regardless of whether
	// the key already existed.
	Put(key interface{}, value interface{}) AnyMap

	PutNonNil(key interface{}, value interface{}) AnyMap

	// Replace takes the given key and replaces the already
	// stored value with the new given value without changing
	// the map order.
	//
	// If the key did not already exist, this method will not
	// insert it.
	//
	// This method returns an option which will contain the
	// value previously stored at `key` if such a value
	// existed.
	Replace(key interface{}, value interface{}) option.Untyped

	// ReplaceOrPut takes the given key/value pair and either
	// replaces an existing value if the key was already in
	// the map, or appends the key/value pair to the end of
	// the map if no such key previously existed.
	//
	// This method returns an option which will contain the
	// value previously stored at `key` if such a value
	// existed.
	ReplaceOrPut(key interface{}, value interface{}) option.Untyped

	// Get returns an option that either contains the value
	// stored at the given key (if it exists) or nothing if no
	// such key was found.
	Get(key interface{}) option.Untyped

	// At returns the key/value pair stored at the given
	// position in the map.
	//
	// This method makes no attempt to verify that the given
	// index is valid, and will panic if attempting to fetch
	// a value that is greater than the length of the ordered
	// map.
	At(index uint) (key option.Untyped, value option.Untyped)

	// IndexOf returns an option containing the position in
	// the map of the entry matching the given key, or an
	// empty option if the given key does not appear in the
	// map.
	IndexOf(key interface{}) option.Uint

	// Has returns whether or not the given key exists in the
	// map.
	Has(key interface{}) bool

	// Delete removes the value stored at the given key from
	// the map.
	//
	// If no such key existed, this method does nothing.
	//
	// Returns an option that will contain the value
	// previously stored at the given key if such a value
	// existed.
	Delete(key interface{}) option.Untyped

	ForEach(fn func(k interface{}, v interface{}))
}
