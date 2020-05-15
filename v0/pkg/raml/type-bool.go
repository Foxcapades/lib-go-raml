package raml

const TypeBoolean = "boolean"

func NewBoolean() *Boolean {
	return &Boolean{Base{Type: TypeBoolean}}
}

type Boolean struct {
	Base
}

func (b *Boolean) GetType() string {
	return TypeBoolean
}

type boAlias Boolean

func (b *Boolean) MarshalYAML() (interface{}, error) {
	if b.canSimplify() {
		return b.Type, nil
	}
	return boAlias(*b), nil
}
