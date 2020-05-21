package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"gopkg.in/yaml.v3"
)

// Helper function for uint values in assign functions
func ToUint(v *yaml.Node, ptr *uint) error {
	if val, err := xyml.ToUint(v); err != nil {
		return err
	} else {
		*ptr = val
	}

	return nil
}

// Helper function for uint values in assign functions
func AsUintPtr(v *yaml.Node, ptr **uint) error {
	if val, err := xyml.ToUint(v); err != nil {
		return err
	} else {
		*ptr = &val
	}

	return nil
}
