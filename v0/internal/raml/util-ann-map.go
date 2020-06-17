package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// UnmarshalAnnotationMapRAML unmarshals the given YAML node into the given
// AnnotationMap.
func UnmarshalAnnotationMapRAML(aMap raml.AnnotationMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmpVal := NewAnnotation()

		if err := UnmarshalUntypedMapRAML(tmpVal, v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmpVal)

		return nil
	})
}
