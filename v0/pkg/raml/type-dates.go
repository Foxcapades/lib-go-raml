package raml

const (
	TypeDateOnly     = "date-only"
	TypeTimeOnly     = "time-only"
	TypeDatetimeOnly = "datetime-only"
)

func NewDateOnly() *GenericDatetime {
	return &GenericDatetime{Base: Base{Type: TypeDateOnly}}
}

func NewTimeOnly() *GenericDatetime {
	return &GenericDatetime{Base: Base{Type: TypeTimeOnly}}
}

func NewDatetimeOnly() *GenericDatetime {
	return &GenericDatetime{Base: Base{Type: TypeDatetimeOnly}}
}

type GenericDatetime struct {
	Base `yaml:",inline"`
}

type genAlias GenericDatetime

func (b *GenericDatetime) MarshalYAML() (interface{}, error) {
	if b.canSimplify() {
		return b.Type, nil
	}
	return genAlias(*b), nil
}
