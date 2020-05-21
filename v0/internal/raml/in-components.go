package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"gopkg.in/yaml.v3"
)

/******************************************************************************/

func makeAnnotations() hasAnnotations {
	return hasAnnotations{NewAnnotationMap()}
}

type hasAnnotations struct {
	mp raml.AnnotationMap
}

func (h *hasAnnotations) Annotations() raml.AnnotationMap {
	return h.mp
}

func (h *hasAnnotations) out(out raml.AnyMap) {
	h.mp.ForEach(func(k string, v raml.Annotation) { out.Put(k, v) })
}

func (h *hasAnnotations) in(k string, v *yaml.Node) (bool, error) {
	if k[0] != '(' {
		return false, nil
	}

	tmp := NewUntypedMap()

	if err := tmp.UnmarshalRAML(v); err != nil {
		return false, err
	}

	h.mp.Put(k, tmp)

	return true, nil
}

/******************************************************************************/

// Deprecated: don't use this
func makeExtra() hasExtra {
	return hasExtra{NewAnyMap()}
}

// Deprecated: don't use this
type hasExtra struct {
	mp raml.AnyMap
}

func (h *hasExtra) ExtraFacets() raml.AnyMap {
	return h.mp
}

func (h *hasExtra) out(out raml.AnyMap) {
	h.mp.ForEach(func(k, v interface{}) { out.Put(k, v) })
}

func (h *hasExtra) in(k, v *yaml.Node) {
	h.mp.Put(k.Value, v)
}
