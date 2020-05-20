package xlog

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	pathKey = "path"
	KeyFunc = "func"
	KeyType = "type"
)

func WithType(log *logrus.Entry, t string) *logrus.Entry {
	return log.WithField(KeyType, t)
}

func Error(log *logrus.Entry, i ...interface{}) error {
	tmp := fmt.Sprint(i...)
	log.Error(tmp)
	return errors.New(tmp)
}

func Errorf(log *logrus.Entry, f string, i ...interface{}) error {
	tmp := fmt.Sprintf(f, i...)
	log.Error(tmp)
	return errors.New(tmp)
}

// OptError is a bad bad boy
//
// Deprecated: remove me
func OptError(err error) error {
	return err
}

func AddPath(log *logrus.Entry, nxt interface{}) *logrus.Entry {
	if tmp, ok := log.Data[pathKey]; ok {
		return log.WithField(pathKey, append(tmp.([]interface{}), nxt))
	}

	return log.WithField(pathKey, []interface{}{nxt})
}
