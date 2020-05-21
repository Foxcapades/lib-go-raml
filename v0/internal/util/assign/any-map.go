package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"gopkg.in/yaml.v3"
)

// ToStringMap appends the values of the given YAML mapping to the given
// StringMap.
func ToStringMap(v *yaml.Node, ref raml.StringMap) error {
	return xyml.ForEachMap(v, func(key, val *yaml.Node) error {
		if err := xyml.RequireString(key); err != nil {
			return err
		} else if err := xyml.RequireString(val); err != nil {
			return err
		}

		ref.Put(key.Value, val.Value)

		return nil
	})
}

// TODO: decompose the yaml node further
func ToUntypedMap(y *yaml.Node, ref raml.UntypedMap) error {
	return xyml.ForEachMap(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		ref.Put(k.Value, v)

		return nil
	})
}
