package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func UnmarshalFacetMapRAML(dMap raml.FacetMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewFacet()
		if err := UnmarshalDataTypeMapRAML(tmp, v); err != nil {
			return err
		}

		dMap.Put(k.Value, tmp)
		return nil
	})
}
