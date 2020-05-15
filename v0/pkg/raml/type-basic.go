package raml

const (
	keyDefault     = "default"
	keyType        = "type"
	keyExample     = "example"
	keyExamples    = "examples"
	keyDisplayName = "displayName"
	keyDescription = "description"
	keyFacets      = "facets"
	keyXml         = "xml"
	keyEnum        = "enum"
	keyRequired    = "required"
)

type Base struct {
	Default     interface{}            `yaml:"default,omitempty"`
	Type        string                 `yaml:"type"`
	Example     interface{}            `yaml:"example,omitempty"`
	Examples    interface{}            `yaml:"examples,omitempty"`
	DisplayName *string                `yaml:"displayName,omitempty"`
	Description *string                `yaml:"description,omitempty"`
	Annotations map[string]interface{} `yaml:",inline"`
	Facets      interface{}            `yaml:"facets,omitempty"`
	Xml         interface{}            `yaml:"xml,omitempty"`
	Enum        interface{}            `yaml:"enum,omitempty"`
	Required    *bool                  `yaml:"required,omitempty"`
}

func (b *Base) ToRaml() (string, error) {
	return dataTypeRaml(b)
}

//
//func (b *Base) UnmarshalYAML(value *yaml.Node) error {
//	if value.Kind == yaml.ScalarNode {
//		b.Type = value.Value
//		return nil
//	}
//	if value.Kind != yaml.MappingNode {
//		return fmt.Errorf("Expected YAML map, instead got %s", value.Tag)
//	}
//
//	for i := 0; i < len(value.Content); i++ {
//		key := value.Content[i].Value
//		i++
//		if err := b.assign(key, value.Content[i]); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (b *Base) assign(key string, val *yaml.Node) error {
//	if key[0] == '(' {
//		b.Annotations[key] = val
//		return nil
//	}
//
//	// "type" is intentionally ignored
//	switch key {
//	case keyDefault:
//		b.Default = val
//	case keyExample:
//		b.Example = val
//	case keyExamples:
//		b.Examples = val
//	case keyDisplayName:
//		tmp := val.Value
//		b.DisplayName = &tmp
//	case keyDescription:
//		tmp := val.Value
//		b.Description = &tmp
//	case keyFacets:
//		b.Facets = val
//	case keyXml:
//		b.Xml = val
//	case keyEnum:
//		b.Enum = val
//	case keyRequired:
//		tmp := val.Value == "true"
//		b.Required = &tmp
//	}
//	return nil
//}

func (b *Base) canSimplify() bool {
	if b.Default != nil {
		return false
	}
	if b.Example != nil {
		return false
	}
	if b.Examples != nil {
		return false
	}
	if b.DisplayName != nil {
		return false
	}
	if b.Description != nil {
		return false
	}
	if b.Annotations != nil && len(b.Annotations) > 0 {
		return false
	}
	if b.Facets != nil {
		return false
	}
	if b.Xml != nil {
		return false
	}
	if b.Enum != nil {
		return false
	}
	if b.Required != nil {
		return *b.Required
	}
	return true
}
