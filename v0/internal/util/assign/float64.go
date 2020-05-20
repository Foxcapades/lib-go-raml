package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"gopkg.in/yaml.v3"
)

func AsFloat64Ptr(v *yaml.Node, ptr **float64) error {
	if val, err := xyml.ToFloat64(v); err != nil {
		return err
	} else {
		*ptr = &val
	}

	return nil
}
