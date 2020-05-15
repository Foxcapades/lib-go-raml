package raml

const TypeObject = "object"

func NewObject() *Object {
	return &Object{Base: Base{Type: TypeObject}}
}

type Object struct {
	Base `yaml:",inline"`

	Properties           map[string]*TypeDef `yaml:"properties,omitempty"`
	MinProperties        *uint               `yaml:"minProperties,omitempty"`
	MaxProperties        *uint               `yaml:"maxProperties,omitempty"`
	AdditionalProperties *bool               `yaml:"additionalProperties,omitempty"`
	Discriminator        *string             `yaml:"discriminator,omitempty"`
	DiscriminatorValue   *string             `yaml:"discriminatorValue,omitempty"`
}

func (o *Object) ToRAML() (string, error) {
	return dataTypeRaml(o)
}

type objAlias Object

func (o Object) MarshalYAML() (interface{}, error) {
	if o.canSimplify() {
		return o.Type, nil
	}
	return objAlias(o), nil
}

func (o *Object) canSimplify() bool {
	return o.Base.canSimplify() &&
		(o.Properties == nil || len(o.Properties) == 0) &&
		o.MinProperties == nil &&
		o.MaxProperties == nil &&
		o.AdditionalProperties == nil &&
		o.Discriminator == nil &&
		o.DiscriminatorValue == nil
}
