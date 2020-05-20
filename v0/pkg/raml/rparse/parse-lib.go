package rparse

import (
	"os"
	"path"

	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	impl "github.com/Foxcapades/lib-go-raml-types/v0/internal/raml"
)

func ParseLibraryFile(filePath string) (raml.Library, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(file)
	out := impl.NewLibrary(logrus.WithField("file", path.Join(
		path.Base(path.Dir(filePath)), path.Base(filePath))))
	if err = dec.Decode(out); err != nil {
		return nil, err
	}
	return out, nil
}
