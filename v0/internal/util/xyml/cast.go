package xyml

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// ToBool attempts to parse the given raw YAML value as a bool.
// If the value cannot be parsed as a bool, returns an error.
func ToBool(y *yaml.Node) (bool, error) {
	logrus.Trace("xyml.ToBool")

	if err := RequireBool(y); err != nil {
		return false, err
	}

	switch strings.ToLower(y.Value) {
	case "y", "yes", "true", "on":
		return true, nil
	case "n", "no", "false", "off":
		return false, nil
	}

	return false, fmt.Errorf("unrecognized boolean value '%s' at %d:%d",
		y.Value, y.Line, y.Column)
}

// ToFloat64 attempts to parse the given raw YAML value as a float.
// If the value cannot be parsed as a float, returns an error.
func ToFloat64(y *yaml.Node) (float64, error) {
	logrus.Trace("xyml.ToFloat64")

	if err := RequireFloat(y); err != nil {
		return 0, err
	}

	if val, err := strconv.ParseFloat(y.Value, 64); err != nil {
		return 0, err
	} else {
		return val, nil
	}
}

// ToFloat64 attempts to parse the given raw YAML value as an int.
// If the value cannot be parsed as an int, returns an error.
func ToInt64(y *yaml.Node) (int64, error) {
	logrus.Trace("xyml.ToInt64")

	if err := RequireInt(y); err != nil {
		return 0, err
	}

	if val, err := strconv.ParseInt(y.Value, 10, 64); err != nil {
		return 0, err
	} else {
		return val, nil
	}
}

// ToString attempts to parse the given raw YAML value as a string.
// If the value cannot be parsed as a string, returns an error.
func ToString(y *yaml.Node) (string, error) {
	logrus.Trace("xyml.ToString")

	if err := RequireString(y); err != nil {
		return "", err
	}

	return y.Value, nil
}

func CastYmlTypeToScalar(v *yaml.Node) (interface{}, error) {
	switch v.Tag {
	case String:
		return v.Value, nil
	case Int:
		return strconv.ParseInt(v.Value, 10, 64)
	case Float:
		return strconv.ParseFloat(v.Value, 64)
	case Nil:
		return nil, nil
	}
	return nil, errors.New("cannot cast type " + v.Tag)
}

func CastScalarToYmlType(v interface{}) (*yaml.Node, error) {
	k := reflect.TypeOf(v).Kind()
	switch k {
	case reflect.String:
		return StringNode(v.(string)), nil
	case reflect.Int:
		return IntNode(int64(v.(int))), nil
	case reflect.Int8:
		return IntNode(int64(v.(int8))), nil
	case reflect.Int16:
		return IntNode(int64(v.(int16))), nil
	case reflect.Int32:
		return IntNode(int64(v.(int32))), nil
	case reflect.Int64:
		return IntNode(v.(int64)), nil
	case reflect.Uint:
		return IntNode(int64(v.(uint))), nil
	case reflect.Uint8:
		return IntNode(int64(v.(uint8))), nil
	case reflect.Uint16:
		return IntNode(int64(v.(uint16))), nil
	case reflect.Uint32:
		return IntNode(int64(v.(uint32))), nil
	case reflect.Uint64:
		tmp := v.(uint64)
		if uint64(math.MaxInt64) < tmp {
			return nil, errors.New("uint value would overflow int64")
		}
		return IntNode(int64(tmp)), nil
	case reflect.Float32:
		return FloatNode(float64(v.(float32))), nil
	case reflect.Float64:
		return FloatNode(v.(float64)), nil
	}
	return nil, errors.New("unsupported type conversion for " + reflect.TypeOf(v).String())
}

func CastAnyToYmlType(v interface{}) (*yaml.Node, error) {
	k := reflect.TypeOf(v).Kind()
	switch k {
	case reflect.Map:
		r := reflect.ValueOf(v)
		if r.IsNil() {
			return MapNode(0), nil
		}
		tmp := MapNode(r.Len() * 2)
		it := r.MapRange()
		for it.Next() {
			if err := AppendToMap(tmp, it.Key().Interface(), it.Value().Interface()); err != nil {
				return nil, err
			}
		}
		return tmp, nil
	case reflect.Array, reflect.Slice:
		r := reflect.ValueOf(v)
		if r.IsNil() {
			return SequenceNode(0), nil
		}
		ln := r.Len()
		tmp := SequenceNode(ln)
		for i := 0; i < ln; i++ {
			if err := AppendToSlice(tmp, r.Index(i).Interface()); err != nil {
				return nil, err
			}
		}
		return tmp, nil
	}

	return CastScalarToYmlType(v)
}
