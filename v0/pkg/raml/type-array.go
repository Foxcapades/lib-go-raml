package raml

const (
	TypeArray             = "array"
	DefaultArrayMinLength = uint(0)
	DefaultArrayMaxLength = uint(2_147_483_647)
)

func NewArray() *Array {
	return &Array{Base: Base{Type: TypeArray}}
}

type Array struct {
	Base `yaml:",inline"`

	UniqueItems *bool    `yaml:"uniqueItems"`
	Items       *TypeDef `yaml:"items"`
	MinItems    *uint    `yaml:"minItems"`
	MaxItems    *uint    `yaml:"maxItems"`
}

func (a *Array) ToRAML() (string, error) {
	return dataTypeRaml(a)
}

type arAlias Array

func (a *Array) MarshalYAML() (interface{}, error) {
	if a.canSimplify() {
		return a.Type, nil
	}
	return arAlias(*a), nil
}

func (a *Array) canSimplify() bool {
	if !a.Base.canSimplify() {
		return false
	}
	if a.UniqueItems != nil {
		return !*a.UniqueItems
	}
	if a.Items != nil {
		return false
	}
	if a.MinItems != nil && *a.MinItems != DefaultArrayMinLength {
		return false
	}
	if a.MaxItems != nil && *a.MaxItems != DefaultArrayMaxLength {
		return false
	}

	return true
}
