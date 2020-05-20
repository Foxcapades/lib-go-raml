package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func ToStringMap(v *yaml.Node, ref raml.StringMap) error {
	logrus.Trace("assign.ToStringMap")

	if err := xyml.RequireMapping(v); err != nil {
		return err
	}

	for i := 0; i < len(v.Content); i += 2 {
		key := v.Content[i]

		if err := xyml.RequireString(key); err != nil {
			return err
		}

		val := v.Content[i+1]

		if err := xyml.RequireString(val); err != nil {
			return err
		}

		ref.Put(key.Value, val.Value)
	}

	return nil
}

// TODO: decompose the yaml node further
func ToUntypedMap(v *yaml.Node, ref raml.UntypedMap) error {
	logrus.Trace("assign.ToUntypedMap")

	if err := xyml.RequireMapping(v); err != nil {
		return err
	}

	for i := 0; i < len(v.Content); i += 2 {
		if err := xyml.RequireString(v.Content[i]); err != nil {
			return err
		}

		ref.Put(v.Content[i].Value, v.Content[i+1])
	}

	return nil
}
