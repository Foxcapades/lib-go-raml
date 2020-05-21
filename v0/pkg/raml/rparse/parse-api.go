package rparse

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"gopkg.in/yaml.v3"
	"io"

	impl "github.com/Foxcapades/lib-go-raml-types/v0/internal/raml"
)

func ParseApiDoc(red io.Reader) (raml.APISpec, error) {
	out := impl.NewApiSpec()
	dec := yaml.NewDecoder(red)

	if err := dec.Decode(out); err != nil {
		return nil, err
	}

	return out, nil
}
