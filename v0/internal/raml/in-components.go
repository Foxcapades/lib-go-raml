package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

/******************************************************************************/

func makeUses() hasUses {
	return hasUses{NewStringMap()}
}

type hasUses struct {
	uses raml.StringMap
}

func (u *hasUses) Uses() raml.StringMap {
	return u.uses
}

func (u *hasUses) out(out raml.AnyMap) {
	if u.uses.Len() > 0 {
		out.Put(rmeta.KeyUses, u.uses)
	}
}

func (u *hasUses) in(value *yaml.Node) error {
	return assign.ToStringMap(value, u.uses)
}

/******************************************************************************/

func makeAnnTypes() hasAnnTypes {
	return hasAnnTypes{NewUntypedMap()}
}

type hasAnnTypes struct {
	annTypes raml.UntypedMap
}

func (h *hasAnnTypes) AnnotationTypes() raml.UntypedMap {
	return h.annTypes
}

func (h *hasAnnTypes) out(out raml.AnyMap) {
	if h.annTypes.Len() > 0 {
		out.Put(rmeta.KeyAnnotationTypes, h.annTypes)
	}
}

func (h *hasAnnTypes) in(value *yaml.Node) error {
	return h.annTypes.UnmarshalRAML(value)
}

/******************************************************************************/

func makeSecSchemes() hasSecSchemes {
	return hasSecSchemes{NewUntypedMap()}
}

type hasSecSchemes struct {
	secSchemes raml.UntypedMap
}

func (h *hasSecSchemes) SecuritySchemes() raml.UntypedMap {
	return h.secSchemes
}

func (h *hasSecSchemes) out(out raml.AnyMap) {
	if h.secSchemes.Len() > 0 {
		out.Put(rmeta.KeySecuritySchemes, h.secSchemes)
	}
}

func (h *hasSecSchemes) in(value *yaml.Node) error {
	return h.secSchemes.UnmarshalRAML(value)
}

/******************************************************************************/

func makeHasTypes() hasTypes {
	return hasTypes{NewDataTypeMap()}
}

type hasTypes struct {
	mp raml.DataTypeMap
}

func (h *hasTypes) Types() raml.DataTypeMap {
	return h.mp
}

func (h *hasTypes) Schemas() raml.DataTypeMap {
	return h.mp
}

func (h *hasTypes) out(out raml.AnyMap) {
	logrus.Trace("internal.hasTypes.out")
	if h.mp.Len() > 0 {
		logrus.Debug("appending types")
		out.Put(rmeta.KeyTypes, h.mp)
	}
}

func (h *hasTypes) in(v *yaml.Node) error {
	logrus.Trace("internal.hasTypes.in")
	return h.mp.UnmarshalRAML(v)
}

/******************************************************************************/

func makeHasTraits() hasTraits {
	return hasTraits{NewUntypedMap()}
}

type hasTraits struct {
	traits raml.UntypedMap
}

func (h *hasTraits) Traits() raml.UntypedMap {
	return h.traits
}

func (h *hasTraits) out(out raml.AnyMap) {
	if h.traits.Len() > 0 {
		out.Put(rmeta.KeyTraits, h.traits)
	}
}

func (h *hasTraits) in(v *yaml.Node) error {
	return h.traits.UnmarshalRAML(v)
}

/******************************************************************************/

func makeResTypes() hasResTypes {
	return hasResTypes{NewUntypedMap()}
}

type hasResTypes struct {
	mp raml.UntypedMap
}

func (h *hasResTypes) ResourceTypes() raml.UntypedMap {
	return h.mp
}

func (h *hasResTypes) out(out raml.AnyMap) {
	if h.mp.Len() > 0 {
		out.Put(rmeta.KeyResourceTypes, h.mp)
	}
}

func (h *hasResTypes) in(v *yaml.Node) error {
	return h.mp.UnmarshalRAML(v)
}

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

func makeUsage() hasUsage {
	return hasUsage{}
}

type hasUsage struct {
	val *string
}

func (h *hasUsage) Usage() option.String {
	return option.NewMaybeString(h.val)
}

func (h *hasUsage) out(out raml.AnyMap) {
	out.PutNonNil(rmeta.KeyUsage, h.val)
}

func (h *hasUsage) in(v *yaml.Node) error {
	return assign.AsStringPtr(v, &h.val)
}

/******************************************************************************/

func makeExtra() hasExtra {
	return hasExtra{NewAnyMap()}
}

type hasExtra struct {
	mp raml.AnyMap
}

func (h *hasExtra) ExtraFacets() raml.AnyMap {
	return h.mp
}

func (h *hasExtra) out(out raml.AnyMap) {
	h.mp.ForEach(func(k interface{}, v interface{}) { out.Put(k, v) })
}

func (h *hasExtra) in(k, v *yaml.Node) {
	h.mp.Put(k.Value, v)
}
