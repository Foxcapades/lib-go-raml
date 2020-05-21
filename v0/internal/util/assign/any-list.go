package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"gopkg.in/yaml.v3"
)

func AnyList(v *yaml.Node, list *[]interface{}) error {
	*list = make([]interface{}, 0, len(v.Content))

	return xyml.ForEachList(v, func(v *yaml.Node) error {
		*list = append(*list, v)
		return nil
	})
}
