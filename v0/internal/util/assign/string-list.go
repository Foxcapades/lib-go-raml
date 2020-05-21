package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"gopkg.in/yaml.v3"
)

func AsStringList(v *yaml.Node, ptr *[]string) error {
	tmp := make([]string, 0, len(v.Content))

	if err := xyml.ForEachList(v, func(v *yaml.Node) error {
		if err := xyml.RequireString(v); err != nil {
			return err
		}

		tmp = append(tmp, v.Value)

		return nil
	}); err != nil {
		return err
	}

	*ptr = tmp

	return nil
}
