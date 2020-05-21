package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

func NewDocumentation() *Documentation {
	return &Documentation{extra: NewAnyMap()}
}

type Documentation struct {
	title   string
	content string
	extra   raml.AnyMap
}

func (d *Documentation) Title() string {
	return d.title
}

func (d *Documentation) SetTitle(title string) raml.Documentation {
	d.title = title
	return d
}

func (d *Documentation) Content() string {
	return d.content
}

func (d *Documentation) SetContent(content string) raml.Documentation {
	d.content = content
	return d
}

func (d *Documentation) ExtraFacets() raml.AnyMap {
	return d.extra
}

func (d *Documentation) MarshalRAML(out raml.AnyMap) (simple bool, err error) {
	out.Put(rmeta.KeyTitle, d.title).
		Put(rmeta.KeyContent, d.content)
	d.extra.ForEach(func(k, v interface{}) { out.Put(k, v) })
	return false, nil
}

func (d *Documentation) UnmarshalRAML(val *yaml.Node) error {
	return xyml.ForEachMap(val, func(k, v *yaml.Node) error {
		switch k.Value {
		case rmeta.KeyTitle:
			d.title = v.Value
		case rmeta.KeyContent:
			d.content = v.Value
		default:
			if key, err := xyml.CastYmlTypeToScalar(k); err != nil {
				return err
			} else {
				d.extra.Put(key, v)
			}
		}
		return nil
	})
}
