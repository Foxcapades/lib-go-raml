package raml

const (
	TypeDatetime = "datetime"

	DefaultDatetimeFormat = "rfc3339"
)

func NewDatetime() *Datetime {
	return &Datetime{Base: Base{Type: TypeDatetime}}
}

type Datetime struct {
	Base `yaml:",inline"`
	Format *string `yaml:"format,omitempty"`
}

func (o *Datetime) ToRaml() (string, error) {
	return dataTypeRaml(o)
}

//func (o *Datetime) UnmarshalYAML(value *yaml.Node) error {
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

type dtAlias Datetime
func (b *Datetime) MarshalYAML() (interface{}, error) {
	if b.canSimplify() {
		return b.Type, nil
	}
	return dtAlias(*b), nil
}

//func (o *Datetime) assign(key string, val *yaml.Node) error {
//	switch key {
//	case keyFormat:
//		tmp := val.Value
//		o.Format = &tmp
//	}
//	return nil
//}

func (o *Datetime) canSimplify() bool {
	if !o.Base.canSimplify() {
		return false
	}
	if o.Format != nil && *o.Format != DefaultDatetimeFormat {
		return false
	}

	return true
}
