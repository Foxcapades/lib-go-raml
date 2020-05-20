package rparse

import (
	"os"
	"path"

	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	impl "github.com/Foxcapades/lib-go-raml-types/v0/internal/raml"
)

func ParseApiDoc(filePath string) (raml.ApiSpec, error) {
	log := logrus.WithField("file", path.Join(path.Base(path.Dir(filePath)), path.Base(filePath)))
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	dec := yaml.NewDecoder(file)
	out := impl.NewApiSpec(log)

	if err = dec.Decode(out); err != nil {
		return nil, xlog.Error(log, err)
	}

	return out, nil
}
