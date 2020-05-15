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

func (o *Object) ToRAML() (string, error) {
	return dataTypeRaml(o)
}

type objAlias Object

func (o *Object) MarshalYAML() (interface{}, error) {
	if o.canSimplify() {
		return o.Type, nil
	}
	return objAlias(*o), nil
}

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
