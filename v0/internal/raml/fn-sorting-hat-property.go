package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
)

func PropertySortingHat(val interface{}, log *logrus.Entry) (raml.Property, error) {
	if val, err := TypeSortingHat(val, log); err != nil {
		return nil, err
	} else {
		return val.(raml.Property), nil
	}
}
