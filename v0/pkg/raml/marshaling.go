package raml

import "github.com/sirupsen/logrus"

type Unmarshaler interface {
	UnmarshalRAML(value interface{}, log *logrus.Entry) error
}

type Marshaler interface {
	MarshalRAML(out AnyMap) (simple bool, err error)
}
