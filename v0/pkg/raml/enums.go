package raml

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type NumberFormat interface {
	fmt.Stringer
	yaml.Unmarshaler
	yaml.Marshaler
}

type IntegerFormat interface {
	fmt.Stringer
	yaml.Unmarshaler
	yaml.Marshaler
}

type DatetimeFormat interface {
	fmt.Stringer
}