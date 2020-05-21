package xyml

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

const (
	errReqNodeType = "expected node type %s, instead got %s at %d:%d"
)

// RequireBool verifies that the given YAML node is a boolean type.
// If it is not, returns an error.
func RequireBool(y *yaml.Node) error {
	if !IsBool(y) {
		return mkErr(Bool, y)
	}

	return nil
}

// RequireFloat verifies that the given YAML node is a floating point type.
// If it is not, returns an error.
func RequireFloat(y *yaml.Node) error {
	if !IsFloat(y) {
		return mkErr(Float, y)
	}

	return nil
}

// RequireInt verifies that the given YAML node is an integral type.
// If it is not, returns an error.
func RequireInt(y *yaml.Node) error {
	if !IsInt(y) {
		return mkErr(Int, y)
	}

	return nil
}

// RequireList verifies that the given YAML node is a sequence type.
// If it is not, returns an error.
func RequireList(y *yaml.Node) error {
	if !IsList(y) {
		return mkErr(Sequence, y)
	}

	return nil
}

// RequireMapping verifies that the given YAML node is a mapping type.
// If it is not, returns an error.
func RequireMapping(y *yaml.Node) error {
	if !IsMap(y) {
		return mkErr(Map, y)
	}

	return nil
}

// RequireString verifies that the given YAML node is a string type.
// If it is not, returns an error.
func RequireString(y *yaml.Node) error {
	if !IsString(y) {
		return mkErr(String, y)
	}

	return nil
}

func mkErr(tag string, y *yaml.Node) error {
	return fmt.Errorf(errReqNodeType, tag, y.Tag, y.Line, y.Column)
}
