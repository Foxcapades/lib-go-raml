package raml

const TypeString = "string"

func NewString() *String {
	return &String{Base: Base{Type: TypeString},
	}
}

type String struct {
	Base  `yaml:",inline"`

	Pattern   *string `yaml:"pattern,omitempty"`
	MinLength *uint   `yaml:"minLength,omitempty"`
	MaxLength *uint   `yaml:"maxLength,omitempty"`
}

func (s *String) ToRaml() (string, error) {
	return dataTypeRaml(s)
}

type strAlias String
func (s *String) MarshalYAML() (interface{}, error) {
	if s.canSimplify() {
		return s.Type, nil
	}
	return strAlias(*s), nil
}
//
//func (s *String) UnmarshalYAML(value *yaml.Node) error {
//	if err := s.Base.UnmarshalYAML(value); err != nil {
//		return err
//	}
//
//	for i := 0; i < len(value.Content); i++ {
//		key := value.Content[i].Value
//		i++
//		if err := s.assign(key, value.Content[i]); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (s *String) assign(key string, val *yaml.Node) error {
//	switch key {
//	case "pattern":
//		tmp := val.Value
//		s.Pattern = &tmp
//	case "minLength":
//		tmp, err := strconv.ParseUint(val.Value, 10, strconv.IntSize)
//		if err != nil {
//			return err
//		}
//		t2 := uint(tmp)
//		s.MinLength = &t2
//	case "maxLength":
//		tmp, err := strconv.ParseUint(val.Value, 10, strconv.IntSize)
//		if err != nil {
//			return err
//		}
//		t2 := uint(tmp)
//		s.MaxLength = &t2
//	}
//	return nil
//}

func (s *String) canSimplify() bool {
	if !s.Base.canSimplify() {
		return false
	}

	if s.Pattern != nil {
		return false
	}

	if s.MinLength != nil {
		return false
	}

	if s.MaxLength != nil {
		return false
	}

	return true
}
