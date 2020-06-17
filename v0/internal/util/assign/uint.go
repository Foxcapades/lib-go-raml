package assign

import (
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// Helper function for uint values in assign functions
func ToUint(v *yaml.Node, ptr *uint) error {
	if val, err := xyml.ToInt(v, 10); err != nil {
		return err
	} else {
		*ptr = uint(val)
	}

	return nil
}

// Helper function for uint values in assign functions
func AsUintPtr(v *yaml.Node, ptr **uint) error {
	if val, err := xyml.ToInt(v, 10); err != nil {
		return err
	} else {
		tmp := uint(val)
		*ptr = &tmp
	}

	return nil
}
