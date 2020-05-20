package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func NewApiSpec(log *logrus.Entry) *ApiSpec {
	log = xlog.WithType(log, "internal.ApiSpec")

	return &ApiSpec{
		log:               log,
		baseUriParameters: NewUntypedMap(log),
		resources:         NewUntypedMap(log),

		hasAnnotations: makeAnnotations(log),
		hasAnnTypes:    makeAnnTypes(log),
		hasExtra:       makeExtra(log),
		hasResTypes:    makeResTypes(log),
		hasSecSchemes:  makeSecSchemes(log),
		hasTraits:      makeHasTraits(log),
		hasTypes:       makeHasTypes(log),
		hasUses:        makeUses(log),
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

	log *logrus.Entry

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
	a.log.Trace("internal.ApiSpec.MarshalYAML")
	out := NewAnyMap(a.log).
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
	a.hasTypes.out(out, a.log)

	return out, nil
}

func (a *ApiSpec) UnmarshalYAML(fn func(interface{}) error) error {
	a.log.Trace("internal.ApiSpec.UnmarshalYAML")
	var raw yaml.MapSlice
	if err := fn(&raw); err != nil {
		return err
	}

	for i := range raw {
		l2 := xlog.AddPath(a.log, raw[i].Key)
		if err := a.assign(raw[i].Key, raw[i].Value, l2); err != nil {
			return err
		}
	}

	return nil
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

func (a *ApiSpec) assign(k, v interface{}, log *logrus.Entry) error {
	log.Trace("internal.ApiSpec.assign")
	key, ok := k.(string)

	if !ok {
		a.hasExtra.in(k, v)
		return nil
	}

	if used, err := a.hasAnnotations.in(key, v); err != nil {
		return err
	} else if used {
		return nil
	}

	if key[0] == '/' {
		a.resources.Put(key, v)
	}

	switch key {
	case rmeta.KeyTitle:
		return assign.AsString(v, &a.title, log)
	case rmeta.KeyVersion:
		return assign.AsStringPtr(v, &a.version, log)
	case rmeta.KeyBaseUri:
		return assign.AsStringPtr(v, &a.baseUri, log)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(v, &a.description, log)
	case rmeta.KeyProtocols:
		return assign.AsStringList(v, &a.protocols, log)
	case rmeta.KeyMediaType:
		if str, ok := v.(string); ok {
			a.mediaType = append(a.mediaType, str)
			return nil
		}
		return assign.AsStringList(v, &a.mediaType, log)
	case rmeta.KeyDocumentation:
		if tmp, err := assign.AsAnyList(v, log); err != nil {
			return xlog.Error(log, err)
		} else {
			a.documentation = tmp
		}
	case rmeta.KeySecuredBy:
		return assign.AsStringList(v, &a.securedBy, log)
	case rmeta.KeyUses:
		return a.hasUses.in(v, log)
	case rmeta.KeySecuritySchemes:
		return a.secSchemes.UnmarshalRAML(v, log)
	case rmeta.KeyBaseUriParams:
		return assign.ToUntypedMap(v, a.baseUriParameters, log)
	case rmeta.KeyTraits:
		return a.hasTraits.in(v, log)
	case rmeta.KeyAnnotationTypes:
		return a.hasAnnTypes.in(v, log)
	case rmeta.KeyResourceTypes:
		return a.hasResTypes.in(v, log)
	case rmeta.KeyTypes, rmeta.KeySchemas:
		return a.hasTypes.in(v, log)
	}

	a.hasExtra.in(k, v)
	return nil
}
