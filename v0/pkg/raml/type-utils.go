package raml

import (
	"gopkg.in/yaml.v2"
	"strings"
)

func dataTypeRaml(o interface{}) (string, error) {
	out := strings.Builder{}
	out.WriteString(HeaderDataType)
	out.WriteString("\n\n")

	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(o); err != nil {
		return "", err
	}
	return out.String(), nil
}