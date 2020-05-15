package raml

const TypeString = "string"

func NewString() *String {
	return &String{Base: Base{Type: TypeString}}
}

type String struct {
	Base `yaml:",inline"`

	Pattern   *string `yaml:"pattern,omitempty"`
	MinLength *uint   `yaml:"minLength,omitempty"`
	MaxLength *uint   `yaml:"maxLength,omitempty"`
}

func (s *String) ToRAML() (string, error) {
	return dataTypeRaml(s)
}

type strAlias String

func (s String) MarshalYAML() (interface{}, error) {
	if s.canSimplify() {
		return s.Type, nil
	}
	return strAlias(s), nil
}

func (s *String) canSimplify() bool {
	return s.Base.canSimplify() &&
		s.Pattern == nil &&
		s.MinLength == nil &&
		s.MaxLength == nil
}
