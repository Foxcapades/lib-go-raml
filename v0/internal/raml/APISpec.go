package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
	"io"
)

func NewApiSpec() *APISpec {

	return &APISpec{
		baseURIParameters: NewUntypedMap(),
		types:             NewDataTypeMap(),
		traits:            NewUntypedMap(),
		resourceTypes:     NewUntypedMap(),
		annotationTypes:   NewUntypedMap(),
		annotations:       NewAnnotationMap(),
		securitySchemes:   NewUntypedMap(),
		uses:              NewStringMap(),
		resources:         NewUntypedMap(),
		extra:             NewAnyMap(),
	}
}

type APISpec struct {
	title             string
	description       *string
	version           *string
	baseURI           *string
	baseURIParameters raml.UntypedMap
	protocols         []string
	mediaType         []string
	documentation     []raml.Documentation
	types             raml.DataTypeMap
	traits            raml.UntypedMap
	resourceTypes     raml.UntypedMap
	annotationTypes   raml.UntypedMap
	annotations       raml.AnnotationMap
	securitySchemes   raml.UntypedMap
	securedBy         []string
	uses              raml.StringMap
	resources         raml.UntypedMap
	extra             raml.AnyMap
}

func (a *APISpec) Title() string {
	return a.title
}

func (a *APISpec) SetTitle(t string) raml.APISpec {
	a.title = t
	return a
}

func (a *APISpec) Description() option.String {
	return option.NewMaybeString(a.description)
}

func (a *APISpec) SetDescription(s string) raml.APISpec {
	a.description = &s
	return a
}

func (a *APISpec) UnsetDescription() raml.APISpec {
	a.description = nil
	return a
}

func (a *APISpec) Version() option.String {
	return option.NewMaybeString(a.version)
}

func (a *APISpec) SetVersion(s string) raml.APISpec {
	a.version = &s
	return a
}

func (a *APISpec) UnsetVersion() raml.APISpec {
	a.version = nil
	return a
}

func (a *APISpec) BaseURI() option.String {
	return option.NewMaybeString(a.baseURI)
}

func (a *APISpec) SetBaseURI(s string) raml.APISpec {
	a.baseURI = &s
	return a
}

func (a *APISpec) UnsetBaseURI() raml.APISpec {
	a.baseURI = nil
	return a
}

func (a *APISpec) BaseURIParameters() raml.UntypedMap {
	return a.baseURIParameters
}

func (a *APISpec) Protocols() []string {
	return a.protocols
}

func (a *APISpec) SetProtocols(s []string) raml.APISpec {
	a.protocols = s
	return a
}

func (a *APISpec) UnsetProtocols() raml.APISpec {
	a.protocols = nil
	return a
}

func (a *APISpec) MediaTypes() []string {
	return a.mediaType
}

func (a *APISpec) SetMediaTypes(s []string) raml.APISpec {
	a.mediaType = s
	return a
}

func (a *APISpec) UnsetMediaTypes() raml.APISpec {
	a.mediaType = nil
	return a
}

func (a *APISpec) Documentation() []raml.Documentation {
	return a.documentation
}

func (a *APISpec) SetDocumentation(d []raml.Documentation) raml.APISpec {
	a.documentation = d
	return a
}

func (a *APISpec) UnsetDocumentation() raml.APISpec {
	a.documentation = nil
	return a
}

func (a *APISpec) Types() raml.DataTypeMap {
	return a.types
}

func (a *APISpec) SetTypes(t raml.DataTypeMap) raml.APISpec {
	if t == nil {
		return a.UnsetTypes()
	}
	a.types = t
	return a
}

func (a *APISpec) UnsetTypes() raml.APISpec {
	a.types = NewDataTypeMap()
	return a
}

func (a *APISpec) Schemas() raml.DataTypeMap {
	return a.types
}

func (a *APISpec) SetSchemas(t raml.DataTypeMap) raml.APISpec {
	if t == nil {
		return a.UnsetSchemas()
	}
	a.types = t
	return a
}

func (a *APISpec) UnsetSchemas() raml.APISpec {
	a.types = NewDataTypeMap()
	return a
}

func (a *APISpec) Traits() raml.UntypedMap {
	return a.traits
}

func (a *APISpec) SetTraits(t raml.UntypedMap) raml.APISpec {
	if t == nil {
		return a.UnsetTraits()
	}
	a.traits = t
	return a
}

func (a *APISpec) UnsetTraits() raml.APISpec {
	a.traits = NewUntypedMap()
	return a
}

func (a *APISpec) ResourceTypes() raml.UntypedMap {
	return a.resourceTypes
}

func (a *APISpec) SetResourceTypes(t raml.UntypedMap) raml.APISpec {
	if t == nil {
		return a.UnsetResourceTypes()
	}
	a.resourceTypes = t
	return a
}

func (a *APISpec) UnsetResourceTypes() raml.APISpec {
	a.resourceTypes = NewUntypedMap()
	return a
}

func (a *APISpec) AnnotationTypes() raml.UntypedMap {
	return a.annotationTypes
}

func (a *APISpec) SetAnnotationTypes(t raml.UntypedMap) raml.APISpec {
	if t == nil {
		return a.UnsetAnnotationTypes()
	}
	a.annotationTypes = t
	return a
}

func (a *APISpec) UnsetAnnotationTypes() raml.APISpec {
	a.annotationTypes = NewUntypedMap()
	return a
}

func (a *APISpec) Annotations() raml.AnnotationMap {
	return a.annotations
}

func (a *APISpec) SetAnnotations(t raml.AnnotationMap) raml.APISpec {
	if t == nil {
		return a.UnsetAnnotations()
	}
	a.annotations = t
	return a
}

func (a *APISpec) UnsetAnnotations() raml.APISpec {
	a.annotations = NewAnnotationMap()
	return a
}

func (a *APISpec) SecuritySchemes() raml.UntypedMap {
	return a.securitySchemes
}

func (a *APISpec) SetSecuritySchemes(t raml.UntypedMap) raml.APISpec {
	if t == nil {
		return a.UnsetSecuritySchemes()
	}
	a.securitySchemes = t
	return a
}

func (a *APISpec) UnsetSecuritySchemes() raml.APISpec {
	a.securitySchemes = NewUntypedMap()
	return a
}

func (a *APISpec) SecuredBy() []string {
	return a.securedBy
}

func (a *APISpec) SetSecuredBy(s []string) raml.APISpec {
	a.securedBy = s
	return a
}

func (a *APISpec) UnsetSecuredBy() raml.APISpec {
	a.securedBy = nil
	return a
}

func (a *APISpec) Uses() raml.StringMap {
	return a.uses
}

func (a *APISpec) SetUses(s raml.StringMap) raml.APISpec {
	if s == nil {
		return a.UnsetUses()
	}
	a.uses = s
	return a
}

func (a *APISpec) UnsetUses() raml.APISpec {
	a.uses = NewStringMap()
	return a
}

func (a *APISpec) Resources() raml.UntypedMap {
	return a.resources
}

func (a *APISpec) ExtraFacets() raml.AnyMap {
	return a.extra
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

	if !a.uses.Empty() {
		out.Put(rmeta.KeyUses, a.uses)
	}

	if !a.securitySchemes.Empty() {
		out.Put(rmeta.KeySecuritySchemes, a.securitySchemes)
	}

	if !a.baseURIParameters.Empty() {
		out.Put(rmeta.KeyBaseURIParams, a.baseURIParameters)
	}

	if !a.traits.Empty() {
		out.Put(rmeta.KeyTraits, a.traits)
	}

	if !a.annotationTypes.Empty() {
		out.Put(rmeta.KeyAnnotationTypes, a.annotationTypes)
	}

	if !a.resourceTypes.Empty() {
		out.Put(rmeta.KeyResourceTypes, a.resourceTypes)
	}

	a.annotations.ForEach(func(k string, v raml.Annotation) { out.Put(k, v) })
	a.extra.ForEach(func(k, v interface{}) { out.Put(k, v) })
	a.resources.ForEach(func(k string, v interface{}) { out.Put(k, v) })

	if !a.types.Empty() {
		out.Put(rmeta.KeyTypes, a.types)
	}

	return out, nil
}

func (a *APISpec) UnmarshalYAML(raw *yaml.Node) error {
	return xyml.ForEachMap(raw, a.assign)
}

func (a *APISpec) WriteRAML(w io.Writer) error {
	if _, err := w.Write([]byte(rmeta.HeaderRoot)); err != nil {
		return err
	}
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	return enc.Encode(a)
}

func (a *APISpec) assign(k, v *yaml.Node) error {
	if !xyml.IsString(k) {
		if val, err := xyml.CastYmlTypeToScalar(k); err != nil {
			return err
		} else {
			a.extra.Put(val, v)
		}
		return nil
	}

	if k.Value[0] == '(' {
		tmp := NewUntypedMap()
		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}
		a.annotations.Put(k.Value, tmp)
		return nil
	}

	if k.Value[0] == '/' {
		a.resources.Put(k.Value, v)
		return nil
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
		return xyml.ForEachList(v, func(v *yaml.Node) error {
			tmp := NewDocumentation()
			if err := tmp.UnmarshalRAML(v); err != nil {
				return err
			}
			a.documentation = append(a.documentation, tmp)
			return nil
		})
	case rmeta.KeySecuredBy:
		return assign.AsStringList(v, &a.securedBy)
	case rmeta.KeyUses:
		return a.uses.UnmarshalRAML(v)
	case rmeta.KeySecuritySchemes:
		return a.securitySchemes.UnmarshalRAML(v)
	case rmeta.KeyBaseURIParams:
		return assign.ToUntypedMap(v, a.baseURIParameters)
	case rmeta.KeyTraits:
		return a.traits.UnmarshalRAML(v)
	case rmeta.KeyAnnotationTypes:
		return a.annotationTypes.UnmarshalRAML(v)
	case rmeta.KeyResourceTypes:
		return a.resourceTypes.UnmarshalRAML(v)
	case rmeta.KeyTypes, rmeta.KeySchemas:
		return a.types.UnmarshalRAML(v)
	default:
		if key, err := xyml.CastYmlTypeToScalar(k); err != nil {
			return err
		} else {
			a.extra.Put(key, v)
		}
	}

	return nil
}
