package assign

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func AsInt64Ptr(v *yaml.Node, ptr **int64) error {
	if val, err := xyml.ToInt(v, 10); err != nil {
		return err
	} else {
		*ptr = &val
	}

	return nil
}
