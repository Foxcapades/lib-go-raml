package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// UnmarshalUntypedMapRAML umarshals the given YAML node into an UntypedMap.
func UnmarshalUntypedMapRAML(uMap raml.UntypedMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		uMap.Put(k.Value, v)

		return nil
	})
}
