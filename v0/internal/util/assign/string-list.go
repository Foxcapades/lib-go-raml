package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func AsStringList(v *yaml.Node, ptr *[]string) error {
	logrus.Trace("assign.AsStringList")

	if err := xyml.RequireList(v); err != nil {
		return err
	}

	tmp := make([]string, len(v.Content))

	for i, node := range v.Content {
		if err := xyml.RequireString(node); err != nil {
			return err
		}
		tmp[i] = node.Value
	}

	*ptr = tmp
	return nil
}
