package raml

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type DateFormat interface {
	fmt.Stringer
	yaml.Unmarshaler
	yaml.Marshaler
}
