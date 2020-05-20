package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewApiSpec() *ApiSpec {

	return &ApiSpec{
		baseUriParameters: NewUntypedMap(),
		resources:         NewUntypedMap(),

		hasAnnotations: makeAnnotations(),
		hasAnnTypes:    makeAnnTypes(),
		hasExtra:       makeExtra(),
		hasResTypes:    makeResTypes(),
		hasSecSchemes:  makeSecSchemes(),
		hasTraits:      makeHasTraits(),
		hasTypes:       makeHasTypes(),
		hasUses:        makeUses(),
	}
}

type ApiSpec struct {
	hasAnnotations
	hasAnnTypes
	hasExtra
	hasResTypes
	hasSecSchemes
	hasTraits
	hasTypes
	hasUses

	title             string
	version           *string
	baseUri           *string
	description       *string
	protocols         []string
	mediaType         []string
	documentation     []interface{}
	securedBy         []string
	baseUriParameters raml.UntypedMap
	resources         raml.UntypedMap
}

func (a ApiSpec) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.ApiSpec.MarshalYAML")
	out := NewAnyMap().
		Put(rmeta.KeyTitle, a.title).
		PutNonNil(rmeta.KeyVersion, a.version).
		PutNonNil(rmeta.KeyBaseUri, a.baseUri).
		PutNonNil(rmeta.KeyDescription, a.description).
		PutNonNil(rmeta.KeyProtocols, a.protocols).
		PutNonNil(rmeta.KeyMediaType, a.mediaType).
		PutNonNil(rmeta.KeyDocumentation, a.documentation).
		PutNonNil(rmeta.KeySecuredBy, a.securedBy)

	a.hasUses.out(out)
	a.hasSecSchemes.out(out)

	if a.baseUriParameters.Len() > 0 {
		out.Put(rmeta.KeyBaseUriParams, a.baseUriParameters)
	}

	a.hasTraits.out(out)
	a.hasAnnTypes.out(out)
	a.hasResTypes.out(out)
	a.hasAnnotations.out(out)
	a.hasExtra.out(out)
	a.resources.ForEach(func(k string, v interface{}) { out.Put(k, v) })
	a.hasTypes.out(out)

	return out, nil
}

func (a *ApiSpec) UnmarshalYAML(raw *yaml.Node) error {
	logrus.Trace("internal.ApiSpec.UnmarshalYAML")
	return xyml.ForEachMap(raw, a.assign)
}

func (a *ApiSpec) Description() option.String {
	return option.NewMaybeString(a.description)
}

func (a *ApiSpec) Title() string {
	return a.title
}

func (a *ApiSpec) Version() option.String {
	return option.NewMaybeString(a.version)
}

func (a *ApiSpec) BaseUri() option.String {
	return option.NewMaybeString(a.baseUri)
}

func (a *ApiSpec) BaseUriParameters() raml.UntypedMap {
	return a.baseUriParameters
}

func (a *ApiSpec) Protocols() []string {
	return a.protocols
}

func (a *ApiSpec) MediaTypes() []string {
	return a.mediaType
}

func (a *ApiSpec) Documentation() []interface{} {
	return a.documentation
}

func (a *ApiSpec) SecuredBy() []string {
	return a.securedBy
}

func (a *ApiSpec) Resources() raml.UntypedMap {
	return a.resources
}

func (a *ApiSpec) assign(k, v *yaml.Node) error {
	logrus.Trace("internal.ApiSpec.assign")
	if k.Tag != xyml.String {
		a.hasExtra.in(k, v)
		return nil
	}

	if used, err := a.hasAnnotations.in(k.Value, v); err != nil {
		return err
	} else if used {
		return nil
	}

	if k.Value[0] == '/' {
		a.resources.Put(k.Value, v)
	}

	switch k.Value {
	case rmeta.KeyTitle:
		return assign.AsString(v, &a.title)
	case rmeta.KeyVersion:
		return assign.AsStringPtr(v, &a.version)
	case rmeta.KeyBaseUri:
		return assign.AsStringPtr(v, &a.baseUri)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(v, &a.description)
	case rmeta.KeyProtocols:
		return assign.AsStringList(v, &a.protocols)
	case rmeta.KeyMediaType:
		if v.Tag == xyml.String {
			a.mediaType = append(a.mediaType, v.Value)
			return nil
		}
		return assign.AsStringList(v, &a.mediaType)
	case rmeta.KeyDocumentation:
		return assign.AnyList(v, &a.documentation)
	case rmeta.KeySecuredBy:
		return assign.AsStringList(v, &a.securedBy)
	case rmeta.KeyUses:
		return a.hasUses.in(v)
	case rmeta.KeySecuritySchemes:
		return a.secSchemes.UnmarshalRAML(v)
	case rmeta.KeyBaseUriParams:
		return assign.ToUntypedMap(v, a.baseUriParameters)
	case rmeta.KeyTraits:
		return a.hasTraits.in(v)
	case rmeta.KeyAnnotationTypes:
		return a.hasAnnTypes.in(v)
	case rmeta.KeyResourceTypes:
		return a.hasResTypes.in(v)
	case rmeta.KeyTypes, rmeta.KeySchemas:
		return a.hasTypes.in(v)
	}

	a.hasExtra.in(k, v)
	return nil
}
