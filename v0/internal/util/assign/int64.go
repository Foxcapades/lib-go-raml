package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"gopkg.in/yaml.v3"
)

func AsInt64Ptr(v *yaml.Node, ptr **int64) error {
	if val, err := xyml.ToInt64(v); err != nil {
		return err
	} else {
		*ptr = &val
	}

	return nil
}
