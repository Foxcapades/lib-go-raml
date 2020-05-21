package rparse

import (
	impl "github.com/Foxcapades/lib-go-raml-types/v0/internal/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"gopkg.in/yaml.v3"
	"io"
)

func ParseLibrary(red io.Reader) (raml.Library, error) {
	out := impl.NewLibrary()
	dec := yaml.NewDecoder(red)

	if err := dec.Decode(out); err != nil {
		return nil, err
	}

	return out, nil
}
