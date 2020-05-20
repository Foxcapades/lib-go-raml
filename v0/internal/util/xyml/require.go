package xyml

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	errReqNodeType = "expected node type %s, instead got %s at %d:%d"
)

// RequireBool verifies that the given YAML node is a boolean type.
// If it is not, returns an error.
func RequireBool(y *yaml.Node) error {
	logrus.Trace("xyml.RequireBool")

	if !IsBool(y) {
		return fmt.Errorf(errReqNodeType, Bool, y.Tag, y.Line, y.Column)
	}

	return nil
}

// RequireFloat verifies that the given YAML node is a floating point type.
// If it is not, returns an error.
func RequireFloat(y *yaml.Node) error {
	logrus.Trace("xyml.RequireFloat")

	if !IsFloat(y) {
		return fmt.Errorf(errReqNodeType, Float, y.Tag, y.Line, y.Column)
	}

	return nil
}

// RequireInt verifies that the given YAML node is an integral type.
// If it is not, returns an error.
func RequireInt(y *yaml.Node) error {
	logrus.Trace("xyml.RequireFloat")

	if !IsInt(y) {
		return fmt.Errorf(errReqNodeType, Int, y.Tag, y.Line, y.Column)
	}

	return nil
}

// RequireList verifies that the given YAML node is a sequence type.
// If it is not, returns an error.
func RequireList(y *yaml.Node) error {
	logrus.Trace("xyml.RequireList")

	if !IsList(y) {
		return fmt.Errorf(errReqNodeType, Sequence, y.Tag, y.Line, y.Column)
	}

	return nil
}

// RequireMapping verifies that the given YAML node is a mapping type.
// If it is not, returns an error.
func RequireMapping(y *yaml.Node) error {
	logrus.Trace("xyml.RequireMapping")

	if !IsMap(y) {
		return fmt.Errorf(errReqNodeType, Map, y.Tag, y.Line, y.Column)
	}

	return nil
}

// RequireString verifies that the given YAML node is a string type.
// If it is not, returns an error.
func RequireString(y *yaml.Node) error {
	logrus.Trace("xyml.RequireString")

	if !IsString(y) {
		return fmt.Errorf(errReqNodeType, String, y.Tag, y.Line, y.Column)
	}

	return nil
}
