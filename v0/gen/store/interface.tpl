{{define "map" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

// {{.Name}}Map generated @ {{.Time}}
type {{.Name}}Map interface {
	Unmarshaler
	yaml.Marshaler

	// Len returns the current number of elements in the
	// map.
	Len() uint

	// Put inserts a new element at the end of the map.
	//
	// If there was already an entry in the map with the given
	// key, it will be replaced.  The new key/value pair will
	// be appended to the end of the map regardless of whether
	// the key already existed.
	Put(key {{.KType}}, value {{.Type}}) {{.Name}}Map

	PutNonNil(key {{.KType}}, value {{if (and (not .Raml) (ne .Kind "Untyped"))}}*{{end}}{{.Type}}) {{.Name}}Map

	// Replace takes the given key and replaces the already
	// stored value with the new given value without changing
	// the map order.
	//
	// If the key did not already exist, this method will not
	// insert it.
	//{{if .Raml}}
	// This method returns the value previously stored at
	// `key` if such a value existed.{{else}}
	// This method returns an option which will contain the
	// value previously stored at `key` if such a value
	// existed.{{end}}
	Replace(key {{.KType}}, value {{.Type}}) {{if not .Raml}}option.{{end}}{{.Kind}}

	// ReplaceOrPut takes the given key/value pair and either
	// replaces an existing value if the key was already in
	// the map, or appends the key/value pair to the end of
	// the map if no such key previously existed.
	//{{if .Raml}}
	// This method returns the value previously stored at
	// `key` if such a value existed.{{else}}
	// This method returns an option which will contain the
	// value previously stored at `key` if such a value
	// existed.{{end}}
	ReplaceOrPut(key {{.KType}}, value {{.Type}}) {{if not .Raml}}option.{{end}}{{.Kind}}
{{if .Raml}}
	// Get returns the value stored at the given key (if it
	// exists) or nil if no such key was found.{{else}}
	// Get returns an option that either contains the value
	// stored at the given key (if it exists) or nothing if no
	// such key was found.{{end}}
	Get(key {{.KType}}) {{if not .Raml}}option.{{end}}{{.Kind}}

	// At returns the key/value pair stored at the given
	// position in the map.
	//
	// This method makes no attempt to verify that the given
	// index is valid, and will panic if attempting to fetch
	// a value that is greater than the length of the ordered
	// map.
	At(index uint) (key option.{{.KName}}, value {{if not .Raml}}option.{{end}}{{.Kind}})

	// IndexOf returns an option containing the position in
	// the map of the entry matching the given key, or an
	// empty option if the given key does not appear in the
	// map.
	IndexOf(key {{.KType}}) option.Uint

	// Has returns whether or not the given key exists in the
	// map.
	Has(key {{.KType}}) bool

	// Delete removes the value stored at the given key from
	// the map.
	//
	// If no such key existed, this method does nothing.
	//{{if .Raml}}
	// Returns the value previously stored at the given key if
	// such a value existed.{{else}}
	// Returns an option that will contain the value
	// previously stored at the given key if such a value
	// existed.{{end}}
	Delete(key {{.KType}}) {{if not .Raml}}option.{{end}}{{.Kind}}

	ForEach(fn func(k {{.KType}}, v {{.Type}}))
}
{{end}}