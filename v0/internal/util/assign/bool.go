package assign

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"gopkg.in/yaml.v3"
)

// Helper function for bool values in assign functions
func AsBool(v *yaml.Node, ptr *bool) error {
	if val, err := xyml.ToBool(v); err != nil {
		return err
	} else {
		*ptr = val
	}

	return nil
}

// Helper function for bool values in assign functions
func AsBoolPtr(v *yaml.Node, ptr **bool) error {
	if val, err := xyml.ToBool(v); err != nil {
		return err
	} else {
		*ptr = &val
	}

	return nil
}
