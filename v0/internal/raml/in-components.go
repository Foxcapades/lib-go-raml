package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

/******************************************************************************/

func makeUses(log *logrus.Entry) hasUses {
	return hasUses{NewStringMap(log)}
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

func (u *hasUses) in(value interface{}, log *logrus.Entry) error {
	return assign.ToStringMap(value, u.uses, log)
}

/******************************************************************************/

func makeAnnTypes(log *logrus.Entry) hasAnnTypes {
	return hasAnnTypes{
		annTypes: NewUntypedMap(log),
	}
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

func (h *hasAnnTypes) in(value interface{}, log *logrus.Entry) error {
	return h.annTypes.UnmarshalRAML(value, log)
}

/******************************************************************************/

func makeSecSchemes(log *logrus.Entry) hasSecSchemes {
	return hasSecSchemes{NewUntypedMap(log)}
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

func (h *hasSecSchemes) in(value interface{}, log *logrus.Entry) error {
	return h.secSchemes.UnmarshalRAML(value, log)
}

/******************************************************************************/

func makeHasTypes(log *logrus.Entry) hasTypes {
	return hasTypes{NewDataTypeMap(log)}
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

func (h *hasTypes) out(out raml.AnyMap, l *logrus.Entry) {
	l.Trace("internal.hasTypes.out")
	if h.mp.Len() > 0 {
		l.Debug("appending types")
		out.Put(rmeta.KeyTypes, h.mp)
	}
}

func (h *hasTypes) in(v interface{}, l *logrus.Entry) error {
	l.Trace("internal.hasTypes.in")
	return xlog.OptError(l, h.mp.UnmarshalRAML(v, l))
}

/******************************************************************************/

func makeHasTraits(log *logrus.Entry) hasTraits {
	return hasTraits{NewUntypedMap(log)}
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

func (h *hasTraits) in(v interface{}, l *logrus.Entry) error {
	return h.traits.UnmarshalRAML(v, l)
}

/******************************************************************************/

func makeResTypes(log *logrus.Entry) hasResTypes {
	return hasResTypes{NewUntypedMap(log)}
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

func (h *hasResTypes) in(v interface{}, l *logrus.Entry) error {
	return h.mp.UnmarshalRAML(v, l)
}

/******************************************************************************/

func makeAnnotations(log *logrus.Entry) hasAnnotations {
	return hasAnnotations{log, NewAnnotationMap(log)}
}

type hasAnnotations struct {
	log *logrus.Entry
	mp raml.AnnotationMap
}

func (h *hasAnnotations) Annotations() raml.AnnotationMap {
	return h.mp
}

func (h *hasAnnotations) out(out raml.AnyMap) {
	h.mp.ForEach(func(k string, v raml.Annotation) { out.Put(k, v) })
}

func (h *hasAnnotations) in(k string, v interface{}) (bool, error) {
	if k[0] != '(' {
		return false, nil
	}

	tmp := NewUntypedMap(h.log)

	if err := tmp.UnmarshalRAML(v, h.log); err != nil {
		return false, err
	}

	h.mp.Put(k, tmp)

	return true, nil
}

/******************************************************************************/

func makeUsage(log *logrus.Entry) hasUsage {
	return hasUsage{log, nil}
}

type hasUsage struct {
	log *logrus.Entry
	val *string
}

func (h *hasUsage) Usage() option.String {
	return option.NewMaybeString(h.val)
}

func (h *hasUsage) out(out raml.AnyMap) {
	out.PutNonNil(rmeta.KeyUsage, h.val)
}

func (h *hasUsage) in(v interface{}) error {
	return assign.AsStringPtr(v, &h.val, h.log)
}

/******************************************************************************/

func makeExtra(log *logrus.Entry) hasExtra {
	return hasExtra{NewAnyMap(log)}
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

func (h *hasExtra) in(k, v interface{}) {
	h.mp.Put(k, v)
}