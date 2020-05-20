package raml

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type DateFormat interface {
	fmt.Stringer
	yaml.Unmarshaler
	yaml.Marshaler
}

