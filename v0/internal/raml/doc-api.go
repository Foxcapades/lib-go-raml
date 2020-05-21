package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

func NewApiSpec() *APISpec {

	return &APISpec{
		baseURIParameters: NewUntypedMap(),
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

type APISpec struct {
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
	baseURI           *string
	description       *string
	protocols         []string
	mediaType         []string
	documentation     []interface{}
	securedBy         []string
	baseURIParameters raml.UntypedMap
	resources         raml.UntypedMap
}

func (a APISpec) MarshalYAML() (interface{}, error) {
	out := NewAnyMap().
		Put(rmeta.KeyTitle, a.title).
		PutNonNil(rmeta.KeyVersion, a.version).
		PutNonNil(rmeta.KeyBaseURI, a.baseURI).
		PutNonNil(rmeta.KeyDescription, a.description).
		PutNonNil(rmeta.KeyProtocols, a.protocols).
		PutNonNil(rmeta.KeyMediaType, a.mediaType).
		PutNonNil(rmeta.KeyDocumentation, a.documentation).
		PutNonNil(rmeta.KeySecuredBy, a.securedBy)

	a.hasUses.out(out)
	a.hasSecSchemes.out(out)

	if a.baseURIParameters.Len() > 0 {
		out.Put(rmeta.KeyBaseURIParams, a.baseURIParameters)
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

func (a *APISpec) UnmarshalYAML(raw *yaml.Node) error {
	return xyml.ForEachMap(raw, a.assign)
}

func (a *APISpec) Description() option.String {
	return option.NewMaybeString(a.description)
}

func (a *APISpec) Title() string {
	return a.title
}

func (a *APISpec) Version() option.String {
	return option.NewMaybeString(a.version)
}

func (a *APISpec) BaseURI() option.String {
	return option.NewMaybeString(a.baseURI)
}

func (a *APISpec) BaseURIParameters() raml.UntypedMap {
	return a.baseURIParameters
}

func (a *APISpec) Protocols() []string {
	return a.protocols
}

func (a *APISpec) MediaTypes() []string {
	return a.mediaType
}

func (a *APISpec) Documentation() []interface{} {
	return a.documentation
}

func (a *APISpec) SecuredBy() []string {
	return a.securedBy
}

func (a *APISpec) Resources() raml.UntypedMap {
	return a.resources
}

func (a *APISpec) assign(k, v *yaml.Node) error {
	if !xyml.IsString(k) {
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
	case rmeta.KeyBaseURI:
		return assign.AsStringPtr(v, &a.baseURI)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(v, &a.description)
	case rmeta.KeyProtocols:
		return assign.AsStringList(v, &a.protocols)
	case rmeta.KeyMediaType:
		if xyml.IsString(v) {
			a.mediaType = append(a.mediaType, v.Value)
		} else {
			return assign.AsStringList(v, &a.mediaType)
		}
	case rmeta.KeyDocumentation:
		return assign.AnyList(v, &a.documentation)
	case rmeta.KeySecuredBy:
		return assign.AsStringList(v, &a.securedBy)
	case rmeta.KeyUses:
		return a.hasUses.in(v)
	case rmeta.KeySecuritySchemes:
		return a.secSchemes.UnmarshalRAML(v)
	case rmeta.KeyBaseURIParams:
		return assign.ToUntypedMap(v, a.baseURIParameters)
	case rmeta.KeyTraits:
		return a.hasTraits.in(v)
	case rmeta.KeyAnnotationTypes:
		return a.hasAnnTypes.in(v)
	case rmeta.KeyResourceTypes:
		return a.hasResTypes.in(v)
	case rmeta.KeyTypes, rmeta.KeySchemas:
		return a.hasTypes.in(v)
	default:
		a.hasExtra.in(k, v)
	}

	return nil
}
