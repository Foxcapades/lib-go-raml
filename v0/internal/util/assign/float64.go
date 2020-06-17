package assign

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func AsFloat64Ptr(v *yaml.Node, ptr **float64) error {
	if val, err := xyml.ToFloat(v); err != nil {
		return err
	} else {
		*ptr = &val
	}

	return nil
}
