package raml

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

func (b *Base) GetType() string {
	return b.Type
}

func (b *Base) SetType(t string) {
	b.Type = t
}

func (b *Base) ToRAML() (string, error) {
	return dataTypeRaml(b)
}

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
