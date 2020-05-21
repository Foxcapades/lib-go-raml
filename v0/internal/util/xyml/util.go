package xyml

import (
	"gopkg.in/yaml.v3"
)

func ForEachMap(y *yaml.Node, fn func(k, v *yaml.Node) error) error {
	if err := RequireMapping(y); err != nil {
		return err
	}

	for i := 0; i < len(y.Content); i += 2 {
		if err := fn(y.Content[i], y.Content[i+1]); err != nil {
			return err
		}
	}

	return nil
}

func ForEachList(y *yaml.Node, fn func(v *yaml.Node) error) error {
	if err := RequireList(y); err != nil {
		return err
	}

	for _, val := range y.Content {
		if err := fn(val); err != nil {
			return err
		}
	}

	return nil
}
