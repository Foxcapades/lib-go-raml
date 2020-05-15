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
	Base

	UniqueItems *bool    `yaml:"uniqueItems"`
	Items       *TypeDef `yaml:"items"`
	MinItems    *uint    `yaml:"minItems"`
	MaxItems    *uint    `yaml:"maxItems"`
}

func (a *Array) ToRaml() (string, error) {
	return dataTypeRaml(a)
}

func (a *Array) GetType() string {
	return TypeArray
}

type arAlias Array

func (a *Array) MarshalYAML() (interface{}, error) {
	if a.canSimplify() {
		return a.Type, nil
	}
	return arAlias(*a), nil
}

//
//func (a *Array) UnmarshalYAML(f func(interface{}) error) error {
//	return f(a)
//}
//
//func (a *Array) UnmarshalYAML(value *yaml.Node) error {
//	if err := a.Base.UnmarshalYAML(value); err != nil {
//		return err
//	}
//
//	for i := 0; i < len(value.Content); i++ {
//		key := value.Content[i].Value
//		i++
//		if err := a.assign(key, value.Content[i]); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (a *Array) assign(key string, val *yaml.Node) error {
//	switch key {
//	case "items":
//		if tmp, err := handleItems(val); err != nil {
//			return err
//		} else {
//			a.Items = &tmp
//		}
//	case "uniqueItems":
//		tmp := val.Value == "true"
//		a.UniqueItems = &tmp
//	case "minItems":
//		tmp, err := ParseUint(val)
//		if err != nil {
//			return err
//		}
//		a.MinItems = &tmp
//	case "maxItems":
//		tmp, err := ParseUint(val)
//		if err != nil {
//			return err
//		}
//		a.MaxItems = &tmp
//	}
//	return nil
//}
//
//func handleItems(val *yaml.Node) (TypeDef, error) {
//	var out TypeDef
//	err := out.UnmarshalYAML(val)
//	return out, err
//}

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
