package raml

const (
	valAny = "any"
	valNil = "nil"
)

const (
	TypeUnion = "union"
	TypeAny   = "any"
	TypeNil   = "nil"
)

func NewAny() *Any {
	tmp := Any(valAny)
	return &tmp
}

type Any string

func (t *Any) GetType() string {
	return TypeAny
}

func (t *Any) UnmarshalYAML(func(interface{}) error) error {
	*t = valAny
	return nil
}

func (t *Any) MarshalYAML() (interface{}, error) {
	return valAny, nil
}

func (t *Any) ToRAML() (string, error) {
	return HeaderDataType + `

type: any`, nil
}

func NewNil() *Nil {
	tmp := Nil(valNil)
	return &tmp
}

type Nil string

func (t *Nil) GetType() string {
	return TypeNil
}

func (t *Nil) UnmarshalYAML(func(interface{}) error) error {
	*t = "nil"
	return nil
}

func (t *Nil) MarshalYAML() (interface{}, error) {
	return "nil", nil
}

func (t *Nil) ToRAML() (string, error) {
	return HeaderDataType + `

type: nil`, nil
}

func NewUnion() *Union {
	tmp := Union("")
	return &tmp
}

type Union string

func (t Union) GetType() string {
	return string(t)
}

func (t Union) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *Union) UnmarshalYAML(f func(interface{}) error) error {
	tmp := ""
	err := f(&tmp)
	*t = Union(tmp)
	return err
}

func (t Union) ToRAML() (string, error) {
	return HeaderDataType + `

type: ` + string(t), nil
}

func NewCustomType() *CustomType {
	tmp := CustomType("")
	return &tmp
}

type CustomType string

func (t CustomType) GetType() string {
	return string(t)
}

func (t *CustomType) UnmarshalYAML(f func(interface{}) error) error {
	tmp := ""
	err := f(&tmp)
	*t = CustomType(tmp)
	return err
}

func (t CustomType) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t CustomType) ToRAML() (string, error) {
	return HeaderDataType + `

type: ` + string(t), nil
}
