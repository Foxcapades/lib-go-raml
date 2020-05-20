package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func AnyList(v *yaml.Node, list *[]interface{}) error {
	logrus.Trace("assign.AnyList")

	if err := xyml.RequireList(v); err != nil {
		return err
	}

	out := make([]interface{}, len(v.Content))

	for i, node := range v.Content {
		out[i] = node
	}

	*list = out

	return nil
}
