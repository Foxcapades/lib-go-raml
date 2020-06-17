package assign

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func AsString(v *yaml.Node, ptr *string) error {
	if err := xyml.RequireString(v); err != nil {
		return err
	}

	*ptr = v.Value

	return nil
}

func AsStringPtr(v *yaml.Node, ptr **string) error {
	if err := xyml.RequireString(v); err != nil {
		return err
	}

	tmp := v.Value
	*ptr = &tmp

	return nil
}
