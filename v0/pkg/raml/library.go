package raml

type Library struct {
	Types         map[string]TypeDef `yaml:"types,omitempty"`
	ResourceTypes []byte             `yaml:"resourceTypes,omitempty"`
	Traits        []byte             `yaml:"traits,omitempty"`
	Uses          map[string]string  `yaml:"uses,omitempty"`
}
