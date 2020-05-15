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

func (d *Datetime) ToRAML() (string, error) {
	return dataTypeRaml(d)
}

type dtAlias Datetime

func (d *Datetime) MarshalYAML() (interface{}, error) {
	if d.canSimplify() {
		return d.Type, nil
	}
	return dtAlias(*d), nil
}

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
