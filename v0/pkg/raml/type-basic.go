package raml

type Base struct {
	Default     interface{}            `yaml:"default,omitempty"`
	Type        string                 `yaml:"type"`
	Example     interface{}            `yaml:"example,omitempty"`
	Examples    interface{}            `yaml:"examples,omitempty"`
	DisplayName *string                `yaml:"displayName,omitempty"`
	Description *string                `yaml:"description,omitempty"`
	Annotations map[string]interface{} `yaml:"-,inline"`
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
	return b.Default == nil &&
		b.Example == nil &&
		b.Examples == nil &&
		b.DisplayName == nil &&
		b.Description == nil &&
		(b.Annotations == nil || len(b.Annotations) == 0) &&
		b.Facets == nil &&
		b.Xml == nil &&
		b.Enum == nil &&
		(b.Required == nil || *b.Required)
}
