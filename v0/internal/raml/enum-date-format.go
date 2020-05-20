package raml

type DateFormat string

const (
	DateFormatRfc3339 DateFormat = "rfc3339"
	DateFormatRfc2616 DateFormat = "rfc2616"
)

func (d DateFormat) String() string {
	return string(d)
}

func (d *DateFormat) UnmarshalYAML(fn func(interface{}) error) error {
	return fn(d)
}

func (d DateFormat) MarshalYAML() (interface{}, error) {
	return string(d), nil
}

func (d DateFormat) render() bool {
	return d != DateFormatRfc3339
}
