package raml

const TypeObject = "object"

func NewObject() *Object {
	return &Object{Base: Base{Type: TypeObject}}
}

type Object struct {
	Base `yaml:",inline"`

	Properties           map[string]TypeDef `yaml:"properties,omitempty"`
	MinProperties        *uint              `yaml:"minProperties,omitempty"`
	MaxProperties        *uint              `yaml:"maxProperties,omitempty"`
	AdditionalProperties *bool              `yaml:"additionalProperties,omitempty"`
	Discriminator        *string            `yaml:"discriminator,omitempty"`
	DiscriminatorValue   *string            `yaml:"discriminatorValue,omitempty"`
}

func (o *Object) GetType() string {
	return TypeObject
}

func (o *Object) ToRaml() (string, error) {
	return dataTypeRaml(o)
}

type objAlias Object

func (o *Object) MarshalYAML() (interface{}, error) {
	if o.canSimplify() {
		return o.Type, nil
	}
	return objAlias(*o), nil
}

//
//func (o *Object) UnmarshalYAML(value *yaml.Node) error {
//	if err := o.Base.UnmarshalYAML(value); err != nil {
//		return err
//	}
//
//	for i := 0; i < len(value.Content); i++ {
//		key := value.Content[i].Value
//		i++
//		if err := o.assign(key, value.Content[i]); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (o *Object) assign(key string, val *yaml.Node) error {
//	switch key {
//	case "properties":
//		tmp, err := handleProperties(val)
//		if err != nil {
//			return err
//		}
//		o.Properties = tmp
//	case "minProperties":
//		tmp, err := ParseUint(val)
//		if err != nil {
//			return err
//		}
//		o.MinProperties = &tmp
//	case "maxProperties":
//		tmp, err := ParseUint(val)
//		if err != nil {
//			return err
//		}
//		o.MaxProperties = &tmp
//	case "additionalProperties":
//		// Done this way because default value for this property is true
//		tmp := val.Value != "false"
//		o.AdditionalProperties = &tmp
//	case "discriminator":
//		tmp := val.Value
//		o.Discriminator = &tmp
//	case "discriminatorValue":
//		tmp := val.Value
//		o.DiscriminatorValue = &tmp
//	}
//	return nil
//}

func (o *Object) canSimplify() bool {
	if !o.Base.canSimplify() {
		return false
	}

	if o.Properties != nil && len(o.Properties) > 0 {
		return false
	}

	if o.MinProperties != nil {
		return false
	}

	if o.MaxProperties != nil {
		return false
	}

	if o.AdditionalProperties != nil {
		return *o.AdditionalProperties
	}

	if o.Discriminator != nil {
		return false
	}

	if o.DiscriminatorValue != nil {
		return false
	}

	return true
}

//
//func handleProperties(val *yaml.Node) (map[string]TypeDef, error) {
//	out := make(map[string]TypeDef, len(val.Content)/2)
//	for i := 0; i < len(val.Content); i++ {
//		key := val.Content[i].Value
//		tmp := TypeDef{}
//
//		i++
//		err := tmp.UnmarshalYAML(val.Content[i])
//		if err != nil {
//			return nil, err
//		}
//		out[key] = tmp
//	}
//	return out, nil
//}
