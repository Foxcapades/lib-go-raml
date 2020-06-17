package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"gopkg.in/yaml.v3"
)

func PropertySortingHat(val *yaml.Node) (raml.Property, error) {
	if val, err := TypeSortingHat(val); err != nil {
		return nil, err
	} else {
		return val.(raml.Property), nil
	}
}
