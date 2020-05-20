package xyml

import "gopkg.in/yaml.v3"

func AppendToMap(node *yaml.Node, key, val interface{}) error {
	sKey, err := CastScalarToYmlType(key)
	if err != nil {
		return err
	}

	yVal, err := CastAnyToYmlType(val)
	if err != nil {
		return err
	}

	node.Content = append(node.Content, sKey, yVal)
	return nil
}

func AppendToSlice(node *yaml.Node, val interface{}) error {
	if yVal, err := CastAnyToYmlType(val); err != nil {
		return err
	} else {
		node.Content = append(node.Content, yVal)
	}
	return nil
}
