package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func UnmarshalStringMapRAML(sMap raml.StringMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		if err := xyml.RequireString(v); err != nil {
			return err
		}

		sMap.Put(k.Value, v.Value)

		return nil
	})
}
