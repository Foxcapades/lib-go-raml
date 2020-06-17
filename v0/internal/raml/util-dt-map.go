package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func UnmarshalDataTypeMapRAML(dMap raml.DataTypeMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		val, err := TypeSortingHat(v)
		if err != nil {
			return err
		}

		dMap.Put(k.Value, val)
		return nil
	})
}

func UnmarshalPropertyMapRAML(dMap raml.PropertyMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		val, err := TypeSortingHat(v)
		if err != nil {
			return err
		}

		dMap.Put(k.Value, val.(raml.Property))
		return nil
	})
}
