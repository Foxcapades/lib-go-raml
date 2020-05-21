package xyml

import "gopkg.in/yaml.v3"

// AppendToMap takes the given key and value parameters, converts them to YAML
// nodes and pushes them into the given YAML mapping.
func AppendToMap(node *yaml.Node, key, val interface{}) error {
	if sKey, err := CastScalarToYmlType(key); err != nil {
		return err
	} else if yVal, err := CastAnyToYmlType(val); err != nil {
		return err
	} else {
		node.Content = append(node.Content, sKey, yVal)
	}

	return nil
}

// AppendToSlice takes the given value parameter, converts it to YAML and
// appends it to the given YAML sequence.
func AppendToSlice(node *yaml.Node, val interface{}) error {
	if yVal, err := CastAnyToYmlType(val); err != nil {
		return err
	} else {
		node.Content = append(node.Content, yVal)
	}

	return nil
}
