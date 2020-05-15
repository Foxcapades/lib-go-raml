package raml

const (
	TypeDatetime = "datetime"

	DefaultDatetimeFormat = "rfc3339"
)

func NewDatetime() *Datetime {
	return &Datetime{Base: Base{Type: TypeDatetime}}
}

type Datetime struct {
	Base   `yaml:",inline"`
	Format *string `yaml:"format,omitempty"`
}

func (d *Datetime) ToRaml() (string, error) {
	return dataTypeRaml(d)
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

func (d *Datetime) MarshalYAML() (interface{}, error) {
	if d.canSimplify() {
		return d.Type, nil
	}
	return dtAlias(*d), nil
}

//func (o *Datetime) assign(key string, val *yaml.Node) error {
//	switch key {
//	case keyFormat:
//		tmp := val.Value
//		o.Format = &tmp
//	}
//	return nil
//}

func (d *Datetime) canSimplify() bool {
	if !d.Base.canSimplify() {
		return false
	}
	if d.Format != nil && *d.Format != DefaultDatetimeFormat {
		return false
	}

	return true
}

func (d *Datetime) GetType() string {
	return TypeDatetime
}
